import { useAccount, useNetwork } from 'wagmi'
import { useSelector } from 'react-redux'
import { RootState } from '../../store/store'
import toast from 'react-hot-toast'
import { animated } from 'react-spring'
import BridgeWatcher from '@/pages/bridge/BridgeWatcher'
import { useRouter } from 'next/router'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import {
  updateFromValue,
  setBridgeQuote,
  setIsLoading,
  setDestinationAddress,
  addBridgeTxHash,
} from '@/slices/bridge/reducer'

import {
  setShowDestinationAddress,
  setShowSettingsSlideOver,
} from '@/slices/bridgeDisplaySlice'

import { EMPTY_BRIDGE_QUOTE_ZERO } from '@/constants/bridge'

import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { getErc20TokenAllowance } from '@/actions/getErc20TokenAllowance'
import { subtractSlippage } from '@/utils/slippage'
import { commify } from '@ethersproject/units'
import { formatBigIntToString, powBigInt } from '@/utils/bigint/format'
import { calculateExchangeRate } from '@/utils/calculateExchangeRate'
import { useEffect, useRef, useState } from 'react'
import { Token } from '@/utils/types'
import { getWalletClient } from '@wagmi/core'
import { txErrorHandler } from '@/utils/txErrorHandler'
import { AcceptedChainId, CHAINS_ARR, CHAINS_BY_ID } from '@/constants/chains'
import { approveToken } from '@/utils/approveToken'
import { PageHeader } from '@/components/PageHeader'
import Card from '@/components/ui/tailwind/Card'
import BridgeExchangeRateInfo from '@/components/StateManagedBridge/BridgeExchangeRateInfo'
import { Transition } from '@headlessui/react'
import {
  SECTION_TRANSITION_PROPS,
  TRANSITION_PROPS,
} from '@/styles/transitions'
import { InputContainer } from '@/components/StateManagedBridge/InputContainer'
import { OutputContainer } from '@/components/StateManagedBridge/OutputContainer'
import SettingsSlideOver from '@/components/StateManagedBridge/SettingsSlideOver'
import Button from '@/components/ui/tailwind/Button'
import { SettingsIcon } from '@/components/icons/SettingsIcon'
import { DestinationAddressInput } from '@/components/StateManagedBridge/DestinationAddressInput'
import { isAddress } from '@ethersproject/address'
import { BridgeTransactionButton } from '@/components/StateManagedBridge/BridgeTransactionButton'
import ExplorerToastLink from '@/components/ExplorerToastLink'
import { Address, zeroAddress } from 'viem'
import { stringToBigInt } from '@/utils/bigint/format'
import { Warning } from '@/components/Warning'
import { useAppDispatch } from '@/store/hooks'
import { NetworkTokenBalancesAndAllowances } from '@/utils/actions/fetchPortfolioBalances'
import {
  fetchAndStoreSingleTokenAllowance,
  fetchAndStoreSingleTokenBalance,
} from '@/slices/portfolio/hooks'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { updateSingleTokenAllowance } from '@/slices/portfolio/actions'

const StateManagedBridge = () => {
  const { address } = useAccount()
  const { chain } = useNetwork()
  const { synapseSDK } = useSynapseContext()
  const bridgeDisplayRef = useRef(null)
  const currentSDKRequestID = useRef(0)
  const router = useRouter()
  const { query, pathname } = router

  const portfolioBalances: NetworkTokenBalancesAndAllowances =
    usePortfolioBalances()

  const {
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    bridgeQuote,
    fromValue,
    destinationAddress,
  } = useSelector((state: RootState) => state.bridge)

  const { showSettingsSlideOver, showDestinationAddress } = useSelector(
    (state: RootState) => state.bridgeDisplay
  )

  let pendingPopup
  let successPopup

  const [isApproved, setIsApproved] = useState(false)

  const dispatch = useAppDispatch()

  useEffect(() => {
    segmentAnalyticsEvent(`[Bridge page] arrives`, {
      fromChainId,
      query,
      pathname,
    })
  }, [query])

  useEffect(() => {
    console.log(`[useEffect] fromChainId`, fromChainId)
    console.log(`[useEffect] fromToken`, fromToken?.symbol)
    console.log(`[useEffect] toChainId`, toChainId)
    console.log(`[useEffect] toToken`, toToken?.symbol)

    if (
      fromToken &&
      toToken &&
      fromToken?.decimals[fromChainId] &&
      stringToBigInt(fromValue, fromToken.decimals[fromChainId]) > 0n
    ) {
      console.log('trying to set bridge quote')
      getAndSetBridgeQuote()
    } else {
      dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
      dispatch(setIsLoading(false))
    }
  }, [
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    fromValue,
    address,
    portfolioBalances,
  ])

  // don't like this, rewrite: could be custom hook
  useEffect(() => {
    if (fromToken && fromToken?.addresses[fromChainId] === zeroAddress) {
      setIsApproved(true)
    } else {
      if (
        fromToken &&
        bridgeQuote?.allowance &&
        stringToBigInt(fromValue, fromToken.decimals[fromChainId]) <=
          bridgeQuote.allowance
      ) {
        setIsApproved(true)
      } else {
        setIsApproved(false)
      }
    }
  }, [bridgeQuote, fromToken, fromValue, fromChainId, toChainId])

  let quoteToast

  const getAndSetBridgeQuote = async () => {
    currentSDKRequestID.current += 1
    const thisRequestId = currentSDKRequestID.current
    // will have to handle deadlineMinutes here at later time, gets passed as optional last arg in .bridgeQuote()
    try {
      dispatch(setIsLoading(true))

      const { feeAmount, routerAddress, maxAmountOut, originQuery, destQuery } =
        await synapseSDK.bridgeQuote(
          fromChainId,
          toChainId,
          fromToken.addresses[fromChainId],
          toToken.addresses[toChainId],
          stringToBigInt(fromValue, fromToken.decimals[fromChainId])
        )

      console.log(`[getAndSetQuote] fromChainId`, fromChainId)
      console.log(`[getAndSetQuote] toChainId`, toChainId)
      console.log(`[getAndSetQuote] fromToken.symbol`, fromToken.symbol)
      console.log(`[getAndSetQuote] toToken.symbol`, toToken.symbol)
      console.log(`[getAndSetQuote] fromValue`, fromValue)
      console.log('feeAmount', feeAmount)
      console.log(`[getAndSetQuote] maxAmountOut`, maxAmountOut)

      if (!(originQuery && maxAmountOut && destQuery && feeAmount)) {
        dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
        dispatch(setIsLoading(false))
        return
      }

      const toValueBigInt = BigInt(maxAmountOut.toString()) ?? 0n

      const originTokenDecimals = fromToken.decimals[fromChainId]
      const adjustedFeeAmount =
        BigInt(feeAmount) <
        stringToBigInt(`${fromValue}`, fromToken.decimals[fromChainId])
          ? BigInt(feeAmount)
          : BigInt(feeAmount) / powBigInt(10n, BigInt(18 - originTokenDecimals))

      const isUnsupported = AcceptedChainId[fromChainId] ? false : true

      const allowance =
        fromToken.addresses[fromChainId] === zeroAddress ||
        address === undefined ||
        isUnsupported
          ? 0n
          : await getErc20TokenAllowance({
              address,
              chainId: fromChainId,
              tokenAddress: fromToken.addresses[fromChainId] as Address,
              spender: routerAddress,
            })

      if (fromToken.addresses[fromChainId] !== zeroAddress) {
        dispatch(
          updateSingleTokenAllowance({
            chainId: fromChainId,
            allowance: allowance,
            spender: routerAddress,
            token: fromToken,
          })
        )
      }

      const originMinWithSlippage = subtractSlippage(
        originQuery?.minAmountOut ?? 0n,
        'ONE_TENTH',
        null
      )
      const destMinWithSlippage = subtractSlippage(
        destQuery?.minAmountOut ?? 0n,
        'ONE_TENTH',
        null
      )

      let newOriginQuery = { ...originQuery }
      newOriginQuery.minAmountOut = originMinWithSlippage

      let newDestQuery = { ...destQuery }
      newDestQuery.minAmountOut = destMinWithSlippage
      if (thisRequestId === currentSDKRequestID.current) {
        dispatch(
          setBridgeQuote({
            outputAmount: toValueBigInt,
            outputAmountString: commify(
              formatBigIntToString(
                toValueBigInt,
                toToken.decimals[toChainId],
                8
              )
            ),
            routerAddress,
            allowance,
            exchangeRate: calculateExchangeRate(
              stringToBigInt(fromValue, fromToken.decimals[fromChainId]) -
                BigInt(adjustedFeeAmount),
              fromToken.decimals[fromChainId],
              toValueBigInt,
              toToken.decimals[toChainId]
            ),
            feeAmount,
            delta: BigInt(maxAmountOut.toString()),
            quotes: {
              originQuery: newOriginQuery,
              destQuery: newDestQuery,
            },
          })
        )

        toast.dismiss(quoteToast)
        const message = `Route found for bridging ${fromValue} ${fromToken.symbol} on ${CHAINS_BY_ID[fromChainId]?.name} to ${toToken.symbol} on ${CHAINS_BY_ID[toChainId]?.name}`
        console.log(message)
        quoteToast = toast(message, { duration: 3000 })
      }
    } catch (err) {
      console.log(err)
      if (thisRequestId === currentSDKRequestID.current) {
        toast.dismiss(quoteToast)
        const message = `No route found for bridging ${fromValue} ${fromToken.symbol} on ${CHAINS_BY_ID[fromChainId]?.name} to ${toToken.symbol} on ${CHAINS_BY_ID[toChainId]?.name}`
        console.log(message)
        quoteToast = toast(message, { duration: 3000 })

        dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
        return
      }
    } finally {
      if (thisRequestId === currentSDKRequestID.current) {
        dispatch(setIsLoading(false))
      }
    }
  }

  const approveTxn = async () => {
    try {
      const tx = approveToken(
        bridgeQuote?.routerAddress,
        fromChainId,
        fromToken?.addresses[fromChainId]
      ).then(() => {
        dispatch(
          fetchAndStoreSingleTokenAllowance({
            routerAddress: bridgeQuote?.routerAddress as Address,
            tokenAddress: fromToken?.addresses[fromChainId] as Address,
            address: address,
            chainId: fromChainId,
          })
        )
      })

      try {
        await tx
        setIsApproved(true)
      } catch (error) {
        return txErrorHandler(error)
      }
    } catch (error) {
      return txErrorHandler(error)
    }
  }

  const executeBridge = async () => {
    segmentAnalyticsEvent(`[Bridge] initiates bridge`, {
      address,
      originChainId: fromChainId,
      destinationChainId: toChainId,
      inputAmount: fromValue,
      expectedReceivedAmount: bridgeQuote.outputAmountString,
      slippage: bridgeQuote.exchangeRate,
    })
    try {
      const wallet = await getWalletClient({
        chainId: fromChainId,
      })

      const toAddress =
        destinationAddress && isAddress(destinationAddress)
          ? destinationAddress
          : address

      const data = await synapseSDK.bridge(
        toAddress,
        bridgeQuote.routerAddress,
        fromChainId,
        toChainId,
        fromToken.addresses[fromChainId as keyof Token['addresses']],
        stringToBigInt(fromValue, fromToken.decimals[fromChainId]),
        bridgeQuote.quotes.originQuery,
        bridgeQuote.quotes.destQuery
      )

      const payload =
        fromToken.addresses[fromChainId as keyof Token['addresses']] ===
          zeroAddress ||
        fromToken.addresses[fromChainId as keyof Token['addresses']] === ''
          ? {
              data: data.data,
              to: data.to,
              value: stringToBigInt(fromValue, fromToken.decimals[fromChainId]),
            }
          : data

      const tx = await wallet.sendTransaction(payload)

      const originChainName = CHAINS_BY_ID[fromChainId]?.name
      const destinationChainName = CHAINS_BY_ID[toChainId]?.name
      pendingPopup = toast(
        `Bridging from ${fromToken.symbol} on ${originChainName} to ${toToken.symbol} on ${destinationChainName}`,
        { id: 'bridge-in-progress-popup', duration: Infinity }
      )

      try {
        segmentAnalyticsEvent(`[Bridge] bridges successfully`, {
          address,
          originChainId: fromChainId,
          destinationChainId: toChainId,
          inputAmount: fromValue,
          expectedReceivedAmount: bridgeQuote.outputAmountString,
          slippage: bridgeQuote.exchangeRate,
        })
        dispatch(addBridgeTxHash(tx))
        dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
        dispatch(setDestinationAddress(null))
        dispatch(setShowDestinationAddress(false))
        dispatch(updateFromValue(''))

        const successToastContent = (
          <div>
            <div>
              Successfully initiated bridge from {fromToken.symbol} on{' '}
              {originChainName} to {toToken.symbol} on {destinationChainName}
            </div>
            <ExplorerToastLink
              transactionHash={tx ?? zeroAddress}
              chainId={fromChainId}
            />
          </div>
        )

        successPopup = toast.success(successToastContent, {
          id: 'bridge-success-popup',
          duration: 10000,
        })

        toast.dismiss(pendingPopup)

        setTimeout(async () => {
          await dispatch(
            fetchAndStoreSingleTokenBalance({
              token: fromToken,
              routerAddress: bridgeQuote?.routerAddress as Address,
              address: address,
              chainId: fromChainId,
            })
          )
        }, 2000)

        return tx
      } catch (error) {
        segmentAnalyticsEvent(`[Bridge] error bridging`, {
          address,
          errorCode: error.code,
        })
        console.log(`Transaction failed with error: ${error}`)
        toast.dismiss(pendingPopup)
      }
    } catch (error) {
      segmentAnalyticsEvent(`[Bridge]  error bridging`, {
        address,
        errorCode: error.code,
      })
      console.log('Error executing bridge', error)
      toast.dismiss(pendingPopup)
      return txErrorHandler(error)
    }
  }

  return (
    <div className="flex flex-col w-full max-w-lg mx-auto lg:mx-0">
      <div className="flex flex-col">
        <div className="flex items-center justify-between">
          <PageHeader
            title="Bridge"
            subtitle="Send your assets across chains."
          />
          <div>
            <Button
              className="flex items-center p-3 text-opacity-75 bg-bgLight hover:bg-bgLighter text-secondaryTextColor hover:text-white"
              onClick={() => {
                if (showSettingsSlideOver === true) {
                  dispatch(setShowSettingsSlideOver(false))
                } else {
                  dispatch(setShowSettingsSlideOver(true))
                }
              }}
            >
              {!showSettingsSlideOver ? (
                <>
                  <SettingsIcon className="w-5 h-5 mr-2" />
                  <span>Settings</span>
                </>
              ) : (
                <span>Close</span>
              )}
            </Button>
          </div>
        </div>
        <Card
          divider={false}
          className={`
                pt-5 pb-3 mt-5 overflow-hidden
                transition-all duration-100 transform rounded-xl
                bg-bgBase
              `}
        >
          <div ref={bridgeDisplayRef}>
            <Transition show={showSettingsSlideOver} {...TRANSITION_PROPS}>
              <animated.div>
                <SettingsSlideOver key="settings" />
              </animated.div>
            </Transition>
            <InputContainer />
            <OutputContainer />
            <Warning
              originChainId={fromChainId}
              destinationChainId={toChainId}
              originToken={fromToken}
              destinationToken={toToken}
            />
            <Transition
              appear={true}
              unmount={false}
              show={true}
              {...SECTION_TRANSITION_PROPS}
            >
              <BridgeExchangeRateInfo showGasDrop={true} />
            </Transition>
            {showDestinationAddress && (
              <DestinationAddressInput
                toChainId={toChainId}
                destinationAddress={destinationAddress}
              />
            )}
            <div className="mt-3 mb-3">
              <BridgeTransactionButton
                isApproved={isApproved}
                approveTxn={approveTxn}
                executeBridge={executeBridge}
              />
            </div>
          </div>
        </Card>
        {/* <ActionCardFooter link={HOW_TO_BRIDGE_URL} /> */}
      </div>
      <div className="mt-8">
        <BridgeWatcher
          fromChainId={fromChainId}
          toChainId={toChainId}
          address={address}
          destinationAddress={destinationAddress}
        />
      </div>
    </div>
  )
}

export default StateManagedBridge
