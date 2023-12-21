import { useContext, useMemo } from 'react'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { checkExists } from '@/utils/checkExists'
import { useCurrentTokenBalance } from './useCurrentTokenBalance'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { Web3Context } from 'providers/Web3Provider'
import { isOnlyZeroes } from '@/utils/isOnlyZeroes'
import { useWalletState } from '@/state/slices/wallet/hooks'
import { ZeroAddress } from 'ethers'

export const useValidations = (): {
  hasValidSelections: boolean
  hasEnoughBalance: boolean
  isInputValid: boolean
  isApproved: boolean
  onSelectedChain: boolean
} => {
  const {
    debouncedInputAmount,
    originChainId,
    originToken,
    destinationChainId,
    destinationToken,
  } = useBridgeState()

  const { balances, allowance } = useWalletState()

  const currentTokenBalance = useCurrentTokenBalance()

  const web3Context = useContext(Web3Context)
  const { networkId } = web3Context.web3Provider

  const hasValidSelections: boolean = useMemo(() => {
    return Boolean(
      originChainId && originToken && destinationChainId && destinationToken
    )
  }, [originChainId, originToken, destinationChainId, destinationToken])

  const hasEnoughBalance: boolean = useMemo(() => {
    if (
      !checkExists(debouncedInputAmount) ||
      !checkExists(currentTokenBalance.rawBalance)
    ) {
      return false
    } else {
      const formattedInput = stringToBigInt(
        debouncedInputAmount,
        currentTokenBalance.decimals
      )
      return Boolean(BigInt(currentTokenBalance.rawBalance) >= formattedInput)
    }
  }, [balances, debouncedInputAmount, originToken, destinationToken])

  const isInputValid: boolean = useMemo(() => {
    if (debouncedInputAmount === '') return false
    if (isOnlyZeroes(debouncedInputAmount)) return false
    return true
  }, [debouncedInputAmount])

  const onSelectedChain: boolean = useMemo(() => {
    return networkId === originChainId
  }, [originChainId, networkId])

  const formattedInputAmount: bigint = useMemo(() => {
    return stringToBigInt(
      debouncedInputAmount ?? '0',
      originToken?.decimals[originChainId]
    )
  }, [debouncedInputAmount, originToken])

  const isApproved: boolean = useMemo(() => {
    if (originToken?.addresses[originChainId] === ZeroAddress) {
      return true
    }
    if (!checkExists(allowance)) return true
    if (!formattedInputAmount) return true
    return (
      formattedInputAmount <=
      stringToBigInt(allowance, originToken?.decimals[originChainId])
    )
  }, [formattedInputAmount, allowance, originToken, originChainId])

  return {
    hasValidSelections,
    hasEnoughBalance,
    isInputValid,
    isApproved,
    onSelectedChain,
  }
}
