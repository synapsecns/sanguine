import { useMemo, useState, useEffect, useContext, useCallback } from 'react'
import { Address } from 'viem'

import { SynapseSDK } from '@synapsecns/sdk-router'
import { Web3Context } from 'providers/Web3Provider'

import { fetchBridgeQuote } from '@/utils/actions/fetchBridgeQuote'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import { Receipt } from '@/components/Receipt'

import { BridgeableToken, Chain, WidgetProps } from 'types'
import { ChainSelect } from '@/components/ui/ChainSelect'
import { TokenSelect } from '@/components/ui/TokenSelect'
import { useBridgeQuote, QuoteCallbackState } from '@/hooks/useBridgeQuote'
import { useAllowance } from '@/hooks/useAllowance'
import {
  ApproveCallbackState,
  UseApproveCallbackProps,
  useApproveCallback,
} from '@/hooks/useApproveCallback'
import {
  UseBridgeCallbackArgs,
  useBridgeCallback,
  BridgeCallbackState,
} from '@/hooks/useBridgeCallback'

import { useAppDispatch } from '@/state/hooks'
import {
  setDestinationChain,
  setOriginChain,
  setOriginToken,
  setDestinationToken,
  setTokens,
} from '@/state/slices/bridge/reducer'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { BridgeButton } from './BridgeButton'
import { isOnlyZeroes } from '@/utils/isOnlyZeroes'
import { toHexStr } from '@/utils/toHexStr'

import { generateTheme } from '@/utils/generateTheme'

export const Widget = ({
  chainIds,
  web3Provider,
  networkProviders,
  theme,
  customTheme,
  tokens,
}: WidgetProps) => {
  const dispatch = useAppDispatch()
  const synapseSDK = new SynapseSDK(chainIds, networkProviders)
  const web3Context = useContext(Web3Context)
  const { connectedAddress, signer, provider, networkId } =
    web3Context.web3Provider

  const [inputAmount, setInputAmount] = useState<string>('')

  const { originChain, destinationChain, originToken, destinationToken } =
    useBridgeState()

  const themeVariables = (() => {
    if (theme === 'dark') return generateTheme({ bgColor: 'dark' })
    if (customTheme) return generateTheme(customTheme)
    return generateTheme()
  })()

  /** Fetch Web3 Provider Data */
  useEffect(() => {
    async function fetchData() {
      try {
        const signer = await web3Provider.getSigner()
        const address = await signer.getAddress()
        const network = await web3Provider.getNetwork()

        web3Context.setWeb3Provider({
          connectedAddress: address,
          networkId: Number(network?.chainId),
          signer,
          provider: web3Provider,
        })
      } catch (e) {
        console.log('Error', e)
      }
    }

    web3Provider && fetchData()
  }, [web3Provider])

  useEffect(() => {
    dispatch(setTokens(tokens))
  }, [tokens])

  const originTokenDecimals = useMemo(() => {
    if (typeof originToken.decimals === 'number') return originToken.decimals

    return originToken.decimals[originChain.id]
  }, [originToken])

  const destinationTokenDecimals = useMemo(() => {
    if (typeof destinationToken.decimals === 'number')
      return destinationToken.decimals

    return destinationToken.decimals[destinationChain.id]
  }, [destinationToken])

  const {
    state: quoteState,
    callback: fetchQuoteCallback,
    reset: resetQuote,
    quote,
    error: quoteError,
  } = useBridgeQuote({
    originChainId: originChain.id,
    originTokenAddress: originToken.addresses[originChain.id],
    destinationChainId: destinationChain.id,
    destinationTokenAddress: destinationToken.addresses[destinationChain.id],
    amount: stringToBigInt(inputAmount, originTokenDecimals),
    synapseSDK: synapseSDK,
  })

  const routerAddress: Address = quote?.routerAddress as Address

  const { allowance, checkAllowanceCallback } = useAllowance({
    spenderAddress: routerAddress as Address,
    tokenAddress: originToken.addresses[originChain.id] as Address,
    ownerAddress: connectedAddress as Address,
    chainId: originToken.addresses[originChain.id],
  })

  const useApproveCallbackArgs: UseApproveCallbackProps = {
    spenderAddress: routerAddress as Address,
    tokenAddress: originToken.addresses[originChain.id] as Address,
    ownerAddress: connectedAddress as Address,
    amount: stringToBigInt(inputAmount, originTokenDecimals),
    chainId: originChain.id,
    onSuccess: checkAllowanceCallback,
  }
  const {
    state: approveState,
    callback: approveCallback,
    error: approveError,
  } = useApproveCallback(useApproveCallbackArgs)

  console.log('approveState:', approveState)

  const useBridgeCallbackArgs: UseBridgeCallbackArgs = {
    destinationAddress: connectedAddress as Address,
    originRouterAddress: routerAddress,
    originChainId: originChain.id,
    destinationChainId: destinationChain.id,
    tokenAddress: originToken.addresses[originChain.id] as Address,
    amount: stringToBigInt(inputAmount, originTokenDecimals),
    originQuery: quote?.originQuery,
    destinationQuery: quote?.destQuery,
    synapseSDK,
    signer,
  }
  const {
    state: bridgeState,
    callback: bridgeCallback,
    error: bridgeError,
  } = useBridgeCallback(useBridgeCallbackArgs)

  const isConnectedNetworkSelectedOriginNetwork: boolean = useMemo(() => {
    return networkId === originChain?.id
  }, [originChain?.id, networkId])

  const formattedInputAmount: bigint = useMemo(
    () =>
      stringToBigInt(
        inputAmount ?? '0',
        originToken.decimals ?? originToken.decimals[originChain.id]
      ),
    [inputAmount, originToken]
  )

  const isInputValid: boolean = useMemo(() => {
    if (inputAmount === '') return false
    if (isOnlyZeroes(inputAmount)) return false
    return true
  }, [inputAmount])

  const isApproved: boolean = useMemo(() => {
    if (allowance === null) return true
    if (!formattedInputAmount) return true
    return formattedInputAmount <= allowance
  }, [formattedInputAmount, allowance])

  /** Handle refreshing quotes */
  useEffect(() => {
    if (isInputValid) {
      fetchQuoteCallback()
    } else {
      resetQuote()
    }
  }, [
    inputAmount,
    originToken,
    destinationToken,
    originChain,
    destinationChain,
    isInputValid,
  ])

  const handleSwitchNetwork = useCallback(async () => {
    try {
      const hexChainId: string = toHexStr(originChain?.id)
      await provider.send('wallet_switchEthereumChain', [
        { chainId: hexChainId },
      ])
    } catch (error) {
      console.error('handleSwitchNetwork ', error)
    }
  }, [originChain?.id, provider])

  const handleInputAmountChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const cleanedInput: string = cleanNumberInput(event.target.value)
    setInputAmount(cleanedInput)
  }

  const maxAmountOut = useMemo(() => {
    if (!quote || !quote.maxAmountOut) {
      return 0
    }

    const max = BigInt(quote.maxAmountOut.toString())

    return formatBigIntToString(max, destinationTokenDecimals, 4)
  }, [quote])

  const handleOriginChainSelection = useCallback(
    (newOriginChain: Chain) => {
      dispatch(setOriginChain(newOriginChain))
    },
    [dispatch]
  )

  const handleDestinationChainSelection = useCallback(
    (newDestinationChain: Chain) => {
      dispatch(setDestinationChain(newDestinationChain))
    },
    [dispatch]
  )

  const handleOriginTokenSelection = useCallback(
    (newOriginToken: BridgeableToken) => {
      dispatch(setOriginToken(newOriginToken))
    },
    [dispatch]
  )

  const handleDestinationTokenSelection = useCallback(
    (newDestinationToken: BridgeableToken) => {
      dispatch(setDestinationToken(newDestinationToken))
    },
    [dispatch]
  )

  return (
    <div
      style={themeVariables}
      className="w-[374px] bg-[--synapse-bg-root] p-2 text-[--synapse-text-primary] rounded-lg font-medium flex flex-col gap-2"
    >
      <div className="border rounded-md bg-[--synapse-bg-surface] border-[--synapse-border] p-2 flex flex-col gap-2">
        <ChainSelect
          label="From"
          chain={originChain}
          onChange={handleOriginChainSelection}
        />
        <div className="flex pb-2">
          <input
            className="text-3xl w-full font-semibold bg-[--synapse-bg-surface] placeholder:text-[--synapse-border-hover] focus:outline-none"
            placeholder="0"
            value={inputAmount}
            onChange={handleInputAmountChange}
          />
          <TokenSelect
            label="In"
            token={originToken}
            onChange={handleOriginTokenSelection}
          />
        </div>
      </div>
      <div className="border rounded-md bg-[--synapse-bg-surface] border-[--synapse-border] p-2 flex flex-col gap-2">
        <ChainSelect
          label="To"
          chain={destinationChain}
          onChange={handleDestinationChainSelection}
        />
        <div className="flex items-center justify-between pb-1">
          <input
            className="text-3xl w-full font-semibold bg-[--synapse-bg-surface] placeholder:text-[--synapse-border-hover] focus:outline-none cursor-not-allowed"
            disabled={true}
            placeholder=""
            value={
              quoteState === QuoteCallbackState.LOADING ? '...' : maxAmountOut
            }
          />
          <TokenSelect
            label="Out"
            token={destinationToken}
            onChange={handleDestinationTokenSelection}
          />
        </div>
      </div>
      <Receipt
        quote={quote ?? null}
        send={formatBigIntToString(
          stringToBigInt(inputAmount, originTokenDecimals),
          originTokenDecimals,
          4
        )}
        receive={maxAmountOut}
      />
      <BridgeButton
        originChain={originChain}
        isApproved={isApproved}
        isDisabled={!isInputValid}
        isWrongNetwork={!isConnectedNetworkSelectedOriginNetwork}
        handleApprove={approveCallback}
        handleBridge={bridgeCallback}
        handleSwitchNetwork={handleSwitchNetwork}
        isApprovalPending={approveState === ApproveCallbackState.PENDING}
        isBridgePending={bridgeState === BridgeCallbackState.PENDING}
        approveError={approveError}
        bridgeError={bridgeError}
      />
    </div>
  )
}
