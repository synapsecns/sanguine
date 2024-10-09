import { useEffect, useMemo, useRef, useState } from 'react'
import { Address } from 'viem'
import { useAccount, useSwitchChain } from 'wagmi'
import { useRouter } from 'next/router'
import { getWalletClient, waitForTransactionReceipt } from '@wagmi/core'
import { useTranslations } from 'next-intl'

import { CHAINS_BY_ID } from '@/constants/chains'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'
import { setIsWalletPending } from '@/slices/wallet/reducer'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { Token } from '@/utils/types'
import { txErrorHandler } from '@/utils/txErrorHandler'
import { approveToken } from '@/utils/approveToken'
import { stringToBigInt } from '@/utils/bigint/format'
import {
  fetchAndStoreSingleNetworkPortfolioBalances,
  usePortfolioState,
} from '@/slices/portfolio/hooks'
import {
  updatePendingBridgeTransaction,
  addPendingBridgeTransaction,
  removePendingBridgeTransaction,
} from '@/slices/transactions/actions'
import { useAppDispatch } from '@/store/hooks'
import { getUnixTimeMinutesFromNow } from '@/utils/time'
import { isTransactionReceiptError } from '@/utils/isTransactionReceiptError'
import { wagmiConfig } from '@/wagmiConfig'
import { useWalletState } from '@/slices/wallet/hooks'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { resetBridgeQuote } from '@/slices/bridgeQuote/reducer'
import { fetchBridgeQuote } from '@/slices/bridgeQuote/thunks'
import { useIsBridgeApproved } from '@/utils/hooks/useIsBridgeApproved'
import { isTransactionUserRejectedError } from '@/utils/isTransactionUserRejectedError'
import { useBridgeValidations } from '@/components/StateManagedBridge/hooks/useBridgeValidations'
import { useStaleQuoteUpdater } from '@/components/StateManagedBridge/hooks/useStaleQuoteUpdater'
import { ChainSelect } from './components/ChainSelect'
import { OPTIMISM } from '@/constants/chains/master'
import { CustomAmountInput } from './components/CustomAmountInput'
import { USDC } from '@/constants/tokens/bridgeable'
import { useMaintenance } from '@/components/Maintenance/Maintenance'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { TransactionSummary } from './components/TransactionSummary'

export const CustomBridge = () => {
  const dispatch = useAppDispatch()
  const { address, isConnected } = useAccount()
  const { synapseSDK } = useSynapseContext()
  const router = useRouter()
  const { query, pathname } = router

  const currentSDKRequestID = useRef(0)
  const quoteTimeout = 15000

  const t = useTranslations('Bridge')

  const [isTyping, setIsTyping] = useState(false)

  const [fromChainId, setFromChainId] = useState(10)
  const [fromToken, setFromToken] = useState<Token>(USDC)
  const [toChainId, setToChainId] = useState(534352)
  const [toToken, setToToken] = useState<Token>(USDC)
  const [fromValue, setFromValue] = useState('0')

  const { balances } = usePortfolioState()

  const fromChainBalances = balances[fromChainId]
  const fromTokenBalance = fromChainBalances?.find(
    (token) => token.tokenAddress === fromToken.addresses[fromChainId]
  ).parsedBalance

  const { destinationAddress }: BridgeState = useBridgeState()

  const { bridgeQuote, isLoading } = useBridgeQuoteState()

  const isApproved = useIsBridgeApproved()

  const { hasValidQuote, hasSufficientBalance } = useBridgeValidations()

  const { isWalletPending } = useWalletState()

  const {
    isBridgePaused,
    pausedModulesList,
    BridgeMaintenanceProgressBar,
    BridgeMaintenanceWarningMessage,
  } = useMaintenance()

  useEffect(() => {
    segmentAnalyticsEvent(
      `[Custom Bridge page] arrives`,
      {
        fromChainId,
        query,
        pathname,
      },
      true
    )
  }, [query])

  useEffect(() => {
    if (
      fromToken &&
      toToken &&
      fromToken?.decimals[fromChainId] &&
      stringToBigInt(fromValue, fromToken?.decimals[fromChainId]) > 0n
    ) {
      console.log('trying to set bridge quote')
      getAndSetBridgeQuote()
    } else {
      dispatch(resetBridgeQuote())
    }
  }, [fromChainId, toChainId, fromToken, toToken, fromValue])

  const getAndSetBridgeQuote = async () => {
    currentSDKRequestID.current += 1
    const thisRequestId = currentSDKRequestID.current

    const currentTimestamp: number = getUnixTimeMinutesFromNow(0)

    try {
      if (thisRequestId === currentSDKRequestID.current) {
        await dispatch(
          fetchBridgeQuote({
            synapseSDK,
            fromChainId,
            toChainId,
            fromToken,
            toToken,
            debouncedFromValue: fromValue,
            requestId: thisRequestId,
            currentTimestamp,
            address,
            pausedModulesList,
          })
        )
      }
    } catch (err) {
      console.log(err)
      if (thisRequestId === currentSDKRequestID.current) {
        dispatch(resetBridgeQuote())

        return
      }
    }
  }

  const isUpdaterEnabled =
    isConnected &&
    hasValidQuote &&
    hasSufficientBalance &&
    isApproved &&
    !isWalletPending

  const isQuoteStale = useStaleQuoteUpdater(
    bridgeQuote,
    getAndSetBridgeQuote,
    isUpdaterEnabled,
    quoteTimeout
  )

  const approveTxn = async () => {
    try {
      dispatch(setIsWalletPending(true))
      const tx = approveToken(
        bridgeQuote?.routerAddress,
        fromChainId,
        fromToken?.addresses[fromChainId],
        stringToBigInt(fromValue, fromToken?.decimals[fromChainId])
      )
      await tx
      /** Re-fetch bridge quote to re-check approval state */
      getAndSetBridgeQuote()
    } catch (error) {
      return txErrorHandler(error)
    } finally {
      dispatch(setIsWalletPending(false))
    }
  }

  const executeBridge = async () => {
    const currentTimestamp: number = getUnixTimeMinutesFromNow(0)

    segmentAnalyticsEvent(
      `[Custom Bridge] initiates bridge`,
      {
        id: bridgeQuote.id,
        originChainId: fromChainId,
        destinationChainId: toChainId,
        inputAmount: fromValue,
        expectedReceivedAmount: bridgeQuote.outputAmountString,
        slippage: bridgeQuote.exchangeRate,
        originToken: fromToken?.routeSymbol,
        destinationToken: toToken?.routeSymbol,
        exchangeRate: BigInt(bridgeQuote.exchangeRate.toString()),
        routerAddress: bridgeQuote.routerAddress,
        bridgeQuote,
      },
      true
    )

    dispatch(
      addPendingBridgeTransaction({
        id: currentTimestamp,
        originChain: CHAINS_BY_ID[fromChainId],
        originToken: fromToken,
        originValue: fromValue,
        destinationChain: CHAINS_BY_ID[toChainId],
        destinationToken: toToken,
        transactionHash: undefined,
        timestamp: undefined,
        isSubmitted: false,
        estimatedTime: bridgeQuote.estimatedTime,
        bridgeModuleName: bridgeQuote.bridgeModuleName,
        destinationAddress: destinationAddress,
        routerAddress: bridgeQuote.routerAddress,
      })
    )
    try {
      dispatch(setIsWalletPending(true))
      const wallet = await getWalletClient(wagmiConfig, {
        chainId: fromChainId,
      })

      const payload = await synapseSDK.bridge(
        address,
        bridgeQuote.routerAddress,
        fromChainId,
        toChainId,
        fromToken?.addresses[fromChainId as keyof Token['addresses']],
        stringToBigInt(fromValue, fromToken?.decimals[fromChainId]),
        bridgeQuote.originQuery,
        bridgeQuote.destQuery
      )

      const tx = await wallet.sendTransaction({
        ...payload,
      })

      segmentAnalyticsEvent(`[Custom Bridge] bridges successfully`, {
        id: bridgeQuote.id,
        originChainId: fromChainId,
        destinationChainId: toChainId,
        inputAmount: fromValue,
        expectedReceivedAmount: bridgeQuote.outputAmountString,
        slippage: bridgeQuote.exchangeRate,
        originToken: fromToken?.routeSymbol,
        destinationToken: toToken?.routeSymbol,
        exchangeRate: BigInt(bridgeQuote.exchangeRate.toString()),
        routerAddress: bridgeQuote.routerAddress,
        bridgeQuote,
      })
      dispatch(
        updatePendingBridgeTransaction({
          id: currentTimestamp,
          timestamp: undefined,
          transactionHash: tx,
          isSubmitted: false,
        })
      )
      dispatch(resetBridgeQuote())
      setFromValue('')

      await waitForTransactionReceipt(wagmiConfig, {
        hash: tx as Address,
        timeout: 60_000,
      })

      /** Update Origin Chain token balances after resolved tx or timeout reached */
      /** Assume tx has been actually resolved if above times out */
      dispatch(
        fetchAndStoreSingleNetworkPortfolioBalances({
          address,
          chainId: fromChainId,
        })
      )

      return tx
    } catch (error) {
      segmentAnalyticsEvent(`[Custom Bridge]  error bridging`, {
        errorCode: error.code,
      })
      dispatch(removePendingBridgeTransaction(currentTimestamp))
      console.error('Error executing bridge: ', error)

      /** Fetch balances if await transaction receipt times out */
      if (isTransactionReceiptError(error)) {
        dispatch(
          fetchAndStoreSingleNetworkPortfolioBalances({
            address,
            chainId: fromChainId,
          })
        )
      }

      if (isTransactionUserRejectedError(error)) {
        getAndSetBridgeQuote()
      }

      return txErrorHandler(error)
    } finally {
      dispatch(setIsWalletPending(false))
    }
  }

  const handleFromValueChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const swapFromValueString: string = cleanNumberInput(event.target.value)
    try {
      setFromValue(swapFromValueString)
    } catch (error) {
      console.error('Invalid value for conversion to BigInteger')
      const inputValue = event.target.value
      const regex = /^[0-9]*[.,]?[0-9]*$/

      if (regex.test(inputValue) || inputValue === '') {
        setFromValue(inputValue)
      }
    }
  }

  return (
    <div className="flex flex-col w-full max-w-lg mx-auto lg:mx-0">
      <div className="flex flex-col space-y-3">
        <TransactionSummary />
        {isConnected && (
          <div className="rounded-md bg-zinc-100 dark:bg-bgBase">
            <div className="flex items-center p-3 space-x-2 text-lg">
              <img
                src={fromToken?.icon.src}
                alt={fromToken?.symbol}
                className="w-6 h-6"
              />
              <div>To Bridge</div>
            </div>
            <div className="border-b border-zinc-200 dark:border-zinc-700"></div>
            <div className="flex items-center p-3 space-x-2 text-lg">
              <img
                src={fromToken?.icon.src}
                alt={fromToken?.symbol}
                className="w-5 h-5"
              />
              <div className="text-sm opacity-75">
                {fromTokenBalance}/{fromToken.symbol} from{' '}
                {CHAINS_BY_ID[fromChainId].name} to{' '}
                {CHAINS_BY_ID[toChainId].name}
              </div>
            </div>
          </div>
        )}
        <div className="rounded-md bg-zinc-100 dark:bg-bgBase">
          <div className="p-3 text-lg">
            Bridge to {CHAINS_BY_ID[toChainId].name}
          </div>
          <div className="border-b border-zinc-200 dark:border-zinc-700"></div>
          <div className="p-3">
            <ChainSelect
              label="From"
              isOrigin={true}
              onChange={() => {}}
              chain={OPTIMISM}
            />
            <div className="flex items-center justify-between pb-1 mt-5 border-b border-zinc-200 dark:border-zinc-700">
              <CustomAmountInput
                showValue={fromValue}
                handleFromValueChange={handleFromValueChange}
              />
              <div className="flex items-center flex-shrink-0 space-x-2">
                <img
                  src={fromToken?.icon.src}
                  alt={fromToken?.symbol}
                  className="w-6 h-6"
                />
                <div className="text-lg">{fromToken?.symbol}</div>
              </div>
            </div>
            <div className="flex justify-end mt-2 text-sm opacity-75">
              {isConnected ? (
                <button
                  className="hover:underline"
                  onClick={() => setFromValue(fromTokenBalance)}
                >
                  {fromTokenBalance} available
                </button>
              ) : (
                'Not connected'
              )}
            </div>
            {bridgeQuote?.outputAmountString !== '' && (
              <div className="flex justify-end mt-3 text-sm">
                <div>
                  <div className="opacity-75">
                    {bridgeQuote?.estimatedTime} seconds via{' '}
                    {bridgeQuote?.bridgeModuleName}
                  </div>
                  <div className="flex justify-end space-x-2">
                    <div className="opacity-75">Receive:</div>
                    <div>
                      {bridgeQuote?.outputAmountString} {toToken?.symbol}
                    </div>
                  </div>
                </div>
              </div>
            )}
            <div className="flex justify-end mt-3 text-sm">
              Powered by Synapse
            </div>
            <div className="mt-5">
              <TransactionButton
                fromChainId={fromChainId}
                fromToken={fromToken}
                toChainId={toChainId}
                bridgeQuote={bridgeQuote}
                fromValue={fromValue}
                fromTokenBalance={fromTokenBalance}
                isLoading={isLoading}
                approveTxn={approveTxn}
                executeBridge={executeBridge}
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

const TransactionButton = ({
  fromChainId,
  toChainId,
  fromToken,
  bridgeQuote,
  fromValue,
  fromTokenBalance,
  isLoading,
  approveTxn,
  executeBridge,
}) => {
  const { chain, isConnected } = useAccount()
  const { openConnectModal } = useConnectModal()
  const { switchChain } = useSwitchChain()

  const [isApproving, setIsApproving] = useState(false)
  const [isBridging, setIsBridging] = useState(false)

  const buttonClassName = `
    p-2 mb-2
    text-lg font-sans font-medium tracking-wide w-full
    shadow-[0_0_0_2px_#00C185,0_0_0_4px_#FF8736,0_0_0_6px_#FFC100] 
  `

  const comparableFromTokenBalance = stringToBigInt(
    fromTokenBalance,
    fromToken.decimals[fromChainId]
  )
  const comparableFromValue = stringToBigInt(
    fromValue,
    fromToken.decimals[fromChainId]
  )

  const isApproved = useMemo(() => {
    return (
      fromToken &&
      bridgeQuote?.allowance &&
      stringToBigInt(fromValue, fromToken.decimals[fromChainId]) <=
        bridgeQuote.allowance
    )
  }, [bridgeQuote, fromToken, fromValue, fromChainId])

  const handleApproveTxn = async () => {
    setIsApproving(true)
    try {
      await approveTxn()
    } catch (error) {
      console.error('Approval failed', error)
    } finally {
      setIsApproving(false)
    }
  }

  const handleBridgeTxn = async () => {
    setIsBridging(true)
    try {
      await executeBridge()
    } catch (error) {
      console.error('Bridge failed', error)
    } finally {
      setIsBridging(false)
    }
  }

  if (isLoading) {
    return <button className={buttonClassName}>Loading quote...</button>
  }

  if (!isConnected) {
    return (
      <button className={buttonClassName} onClick={openConnectModal}>
        Connect Wallet
      </button>
    )
  }

  if (isConnected && chain.id !== fromChainId) {
    return (
      <button
        className={buttonClassName}
        onClick={() => switchChain({ chainId: fromChainId })}
      >
        Switch to {CHAINS_BY_ID[fromChainId].name}
      </button>
    )
  }

  if (fromValue === '' || fromValue === '0') {
    return <button className={buttonClassName}>Enter an amount</button>
  }

  if (comparableFromValue > comparableFromTokenBalance) {
    return <button className={buttonClassName}>Insufficient balance</button>
  }

  if (!isApproved) {
    return (
      <button
        className={buttonClassName}
        onClick={handleApproveTxn}
        disabled={isApproving}
      >
        {isApproving ? 'Approving...' : 'Approve'}
      </button>
    )
  }

  return (
    <button
      className={buttonClassName}
      onClick={handleBridgeTxn}
      disabled={isBridging}
    >
      {isBridging ? 'Bridging...' : 'Bridge'}
    </button>
  )
}
