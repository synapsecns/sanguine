import { useEffect, useMemo, useRef, useState } from 'react'
import { ZeroAddress } from 'ethers'
import { type BridgeableToken } from 'types'

import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { isValidBridgeQuote } from '@/utils/isValidBridgeQuote'
import { parseBigIntValue } from '@/utils/parseBigIntValue'
import { selectBridgeQuote } from '@/utils/selectBridgeQuote'

type BigIntLike = bigint | { toString(): string }

type BridgeModulePauseLike = {
  chainId?: number
  toChainId?: number
  bridgeModuleName: string
}

type BridgeQuoteTransactionLike = {
  data?: string
  to?: string
  value?: string | null
}

type ExecutableBridgeQuoteTransactionLike = Required<
  Pick<BridgeQuoteTransactionLike, 'data' | 'to'>
> &
  BridgeQuoteTransactionLike

type BridgeQuoteLike = {
  moduleNames: string[]
  nativeFee?: unknown
  tx?: BridgeQuoteTransactionLike | null
}

type ExecutableBridgeQuoteLike = BridgeQuoteLike & {
  tx: ExecutableBridgeQuoteTransactionLike
}

type FeeDataLike = {
  gasPrice?: BigIntLike | null
  maxFeePerGas?: BigIntLike | null
}

type ProviderLike = {
  estimateGas: (transaction: {
    data: string
    from: string
    to: string
    value?: string
  }) => Promise<BigIntLike>
  getFeeData: () => Promise<FeeDataLike>
}

type NativeSafeMaxState = {
  amountWei: bigint | null
  fillAmount: string | null
  isClickable: boolean
  labelAmount: string | null
  status: 'idle' | 'loading' | 'ready' | 'unavailable'
}

const DISPLAY_DECIMALS = 4
const FILL_DECIMALS = 18
const LOADING_NATIVE_SAFE_MAX_STATE: NativeSafeMaxState = {
  amountWei: null,
  fillAmount: null,
  isClickable: false,
  labelAmount: null,
  status: 'loading',
}
const IDLE_NATIVE_SAFE_MAX_STATE: NativeSafeMaxState = {
  amountWei: null,
  fillAmount: null,
  isClickable: false,
  labelAmount: null,
  status: 'idle',
}
const UNAVAILABLE_NATIVE_SAFE_MAX_STATE: NativeSafeMaxState = {
  amountWei: null,
  fillAmount: null,
  isClickable: false,
  labelAmount: null,
  status: 'unavailable',
}

const hasExecutableQuoteTx = (
  transaction: BridgeQuoteTransactionLike | null | undefined
): transaction is ExecutableBridgeQuoteTransactionLike => {
  return Boolean(transaction?.to && transaction?.data)
}

const toBigIntValue = (value: BigIntLike | null | undefined): bigint | null => {
  if (typeof value === 'bigint') {
    return value
  }

  if (!value) {
    return null
  }

  try {
    return BigInt(value.toString())
  } catch {
    return null
  }
}

const calculateBufferedGasFeeWei = (
  gasLimitWei: bigint,
  feePerGasWei: bigint
): bigint => {
  const estimatedGasFeeWei = gasLimitWei * feePerGasWei

  return (estimatedGasFeeWei * 3n + 1n) / 2n
}

const calculateSafeMaxWei = (
  rawBalanceWei: bigint,
  nativeFeeWei: bigint,
  bufferedGasFeeWei: bigint
): bigint => {
  const reservedWei = nativeFeeWei + bufferedGasFeeWei

  return rawBalanceWei > reservedWei ? rawBalanceWei - reservedWei : 0n
}

const getBootstrapQuoteAmountWei = (rawBalanceWei: bigint): bigint => {
  if (rawBalanceWei <= 1n) {
    return rawBalanceWei
  }

  return rawBalanceWei / 2n
}

const getDisplayAmount = (amountWei: bigint, decimals: number): string => {
  return formatBigIntToString(amountWei, decimals, DISPLAY_DECIMALS) ?? '0.0'
}

const getFillAmount = (amountWei: bigint, decimals: number): string => {
  return formatBigIntToString(amountWei, decimals, FILL_DECIMALS) ?? '0.0'
}

const getSafeMaxState = (
  amountWei: bigint,
  decimals: number
): NativeSafeMaxState => {
  return {
    amountWei,
    fillAmount: getFillAmount(amountWei, decimals),
    isClickable: amountWei > 0n,
    labelAmount: getDisplayAmount(amountWei, decimals),
    status: 'ready',
  }
}

export const useNativeSafeMax = ({
  amountKey,
  connectedAddress,
  destinationChainId,
  destinationToken,
  isWalletPending,
  originChainId,
  originChainProvider,
  originToken,
  pausedModules,
  rawBalance,
  synapseSDK,
}: {
  amountKey: string
  connectedAddress?: string
  destinationChainId?: number
  destinationToken?: BridgeableToken
  isWalletPending: boolean
  originChainId?: number
  originChainProvider?: ProviderLike | null
  originToken?: BridgeableToken
  pausedModules: BridgeModulePauseLike[]
  rawBalance: bigint | string | null
  synapseSDK: any
}) => {
  const [state, setState] = useState<NativeSafeMaxState>(
    IDLE_NATIVE_SAFE_MAX_STATE
  )
  const requestIdRef = useRef(0)
  const resolvedContextKeyRef = useRef<string | null>(null)
  const contextVersionRef = useRef(0)
  const lastContextInputsRef = useRef<{
    coreInputsKey: string
    originChainProvider: ProviderLike | null | undefined
    synapseSDK: any
  } | null>(null)

  const pausedModulesKey = useMemo(
    () => JSON.stringify(pausedModules),
    [pausedModules]
  )
  const stablePausedModules = useMemo(() => pausedModules, [pausedModulesKey])
  const originTokenAddress =
    originToken && originChainId ? originToken.addresses[originChainId] : null
  const destinationTokenAddress =
    destinationToken && destinationChainId
      ? destinationToken.addresses[destinationChainId]
      : null
  const originTokenDecimals =
    originToken && originChainId ? originToken.decimals[originChainId] : null
  const rawBalanceWei = useMemo(
    () => parseBigIntValue(rawBalance),
    [rawBalance]
  )
  const isNativeOriginToken = originTokenAddress === ZeroAddress
  const coreInputsKey = JSON.stringify({
    connectedAddress: connectedAddress ?? null,
    destinationChainId: destinationChainId ?? null,
    destinationTokenAddress,
    isWalletPending,
    originChainId: originChainId ?? null,
    originTokenAddress,
    originTokenDecimals,
    pausedModulesKey,
    rawBalance: rawBalanceWei?.toString() ?? null,
  })

  if (
    lastContextInputsRef.current?.coreInputsKey !== coreInputsKey ||
    lastContextInputsRef.current?.originChainProvider !== originChainProvider ||
    lastContextInputsRef.current?.synapseSDK !== synapseSDK
  ) {
    lastContextInputsRef.current = {
      coreInputsKey,
      originChainProvider,
      synapseSDK,
    }
    contextVersionRef.current += 1
  }

  const contextKey = `${contextVersionRef.current}:${coreInputsKey}`

  useEffect(() => {
    let isActive = true
    requestIdRef.current += 1
    const currentRequestId = requestIdRef.current

    const isCurrentRequest = () => {
      return isActive && currentRequestId === requestIdRef.current
    }

    const finalizeState = (nextState: NativeSafeMaxState) => {
      if (!isCurrentRequest()) {
        return
      }

      resolvedContextKeyRef.current = contextKey
      setState(nextState)
    }

    if (!isNativeOriginToken) {
      resolvedContextKeyRef.current = null
      setState(IDLE_NATIVE_SAFE_MAX_STATE)

      return () => {
        isActive = false
      }
    }

    if (resolvedContextKeyRef.current === contextKey) {
      return () => {
        isActive = false
      }
    }

    if (
      !connectedAddress ||
      !originChainId ||
      !destinationChainId ||
      !originToken ||
      !destinationToken ||
      !originChainProvider ||
      !synapseSDK ||
      typeof originTokenDecimals !== 'number' ||
      rawBalanceWei === null
    ) {
      finalizeState(UNAVAILABLE_NATIVE_SAFE_MAX_STATE)

      return () => {
        isActive = false
      }
    }

    if (isWalletPending) {
      finalizeState(LOADING_NATIVE_SAFE_MAX_STATE)

      return () => {
        isActive = false
      }
    }

    if (rawBalanceWei === 0n) {
      finalizeState(getSafeMaxState(0n, originTokenDecimals))

      return () => {
        isActive = false
      }
    }

    setState(LOADING_NATIVE_SAFE_MAX_STATE)

    const estimateBufferedGasFeeWei = async (
      transaction: ExecutableBridgeQuoteTransactionLike
    ): Promise<bigint | null> => {
      let feeData: FeeDataLike

      try {
        feeData = await originChainProvider.getFeeData()
      } catch {
        return null
      }

      if (!isCurrentRequest()) {
        return null
      }

      const feePerGasWei =
        toBigIntValue(feeData.maxFeePerGas) ?? toBigIntValue(feeData.gasPrice)

      if (feePerGasWei === null) {
        return null
      }

      let gasLimit: BigIntLike

      try {
        gasLimit = await originChainProvider.estimateGas({
          data: transaction.data,
          from: connectedAddress,
          to: transaction.to,
          value: transaction.value
            ? BigInt(transaction.value).toString()
            : undefined,
        })
      } catch {
        return null
      }

      if (!isCurrentRequest()) {
        return null
      }

      const gasLimitWei = toBigIntValue(gasLimit)

      if (gasLimitWei === null) {
        return null
      }

      return calculateBufferedGasFeeWei(gasLimitWei, feePerGasWei)
    }

    const fetchSelectedQuote = async (
      amountWei: bigint
    ): Promise<ExecutableBridgeQuoteLike | null> => {
      let quotes: BridgeQuoteLike[]

      try {
        quotes = await synapseSDK.bridgeV2({
          fromAmount: amountWei.toString(),
          fromChainId: originChainId,
          fromSender: connectedAddress,
          fromToken: originToken.addresses[originChainId],
          slippagePercentage: 0.1,
          toChainId: destinationChainId,
          toRecipient: connectedAddress,
          toToken: destinationToken.addresses[destinationChainId],
        })
        quotes = quotes.filter(isValidBridgeQuote)
      } catch {
        return null
      }

      if (!isCurrentRequest()) {
        return null
      }

      const selectedQuote = selectBridgeQuote<BridgeQuoteLike>({
        quotes,
        originChainId,
        destinationChainId,
        pausedModules: stablePausedModules,
      })

      if (!selectedQuote || !hasExecutableQuoteTx(selectedQuote.tx)) {
        return null
      }

      const executableQuoteTx = selectedQuote.tx

      return {
        ...selectedQuote,
        tx: executableQuoteTx,
      }
    }

    void (async () => {
      try {
        const bootstrapQuoteAmountWei =
          getBootstrapQuoteAmountWei(rawBalanceWei)
        const bootstrapQuote = await fetchSelectedQuote(bootstrapQuoteAmountWei)

        if (!bootstrapQuote) {
          if (!isCurrentRequest()) {
            return
          }

          finalizeState(UNAVAILABLE_NATIVE_SAFE_MAX_STATE)
          return
        }

        const bootstrapBufferedGasFeeWei = await estimateBufferedGasFeeWei(
          bootstrapQuote.tx
        )

        if (bootstrapBufferedGasFeeWei === null) {
          if (!isCurrentRequest()) {
            return
          }

          finalizeState(UNAVAILABLE_NATIVE_SAFE_MAX_STATE)
          return
        }

        const bootstrapNativeFeeWei =
          parseBigIntValue(bootstrapQuote.nativeFee) ?? 0n
        const bootstrapSafeMaxWei = calculateSafeMaxWei(
          rawBalanceWei,
          bootstrapNativeFeeWei,
          bootstrapBufferedGasFeeWei
        )

        if (bootstrapSafeMaxWei === 0n) {
          finalizeState(getSafeMaxState(0n, originTokenDecimals))
          return
        }

        const refinedQuote = await fetchSelectedQuote(bootstrapSafeMaxWei)

        if (!refinedQuote) {
          if (!isCurrentRequest()) {
            return
          }

          finalizeState(UNAVAILABLE_NATIVE_SAFE_MAX_STATE)
          return
        }

        const refinedBufferedGasFeeWei = await estimateBufferedGasFeeWei(
          refinedQuote.tx
        )

        if (refinedBufferedGasFeeWei === null) {
          if (!isCurrentRequest()) {
            return
          }

          finalizeState(UNAVAILABLE_NATIVE_SAFE_MAX_STATE)
          return
        }

        const refinedNativeFeeWei =
          parseBigIntValue(refinedQuote.nativeFee) ?? 0n
        const refinedSafeMaxWei = calculateSafeMaxWei(
          rawBalanceWei,
          refinedNativeFeeWei,
          refinedBufferedGasFeeWei
        )

        finalizeState(getSafeMaxState(refinedSafeMaxWei, originTokenDecimals))
      } catch {
        finalizeState(UNAVAILABLE_NATIVE_SAFE_MAX_STATE)
      }
    })()

    return () => {
      isActive = false
    }
  }, [
    amountKey,
    connectedAddress,
    contextKey,
    destinationChainId,
    destinationToken,
    isNativeOriginToken,
    isWalletPending,
    originChainId,
    originChainProvider,
    originToken,
    originTokenDecimals,
    stablePausedModules,
    rawBalanceWei,
    synapseSDK,
  ])

  return {
    ...state,
    isNativeOriginToken,
  }
}
