import { useContext, useMemo } from 'react'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { checkExists } from '@/utils/checkExists'
import { useCurrentTokenBalance } from './useCurrentTokenBalance'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { Web3Context } from 'providers/Web3Provider'
import { isOnlyZeroes } from '@/utils/isOnlyZeroes'
import { useWalletState } from '@/state/slices/wallet/hooks'

export const useValidations = (): {
  hasValidSelections: boolean
  hasEnoughBalance: boolean
  isInputValid: boolean
  // hasEnoughApproved: boolean
  onSelectedChain: boolean
} => {
  const {
    inputAmount,
    debouncedInputAmount,
    originChainId,
    originToken,
    destinationChainId,
    destinationToken,
  } = useBridgeState()

  const { balances } = useWalletState()

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
      !checkExists(inputAmount) ||
      !checkExists(currentTokenBalance.rawBalance)
    ) {
      return false
    } else {
      const formattedInput = stringToBigInt(
        inputAmount,
        currentTokenBalance.decimals
      )
      return Boolean(BigInt(currentTokenBalance.rawBalance) >= formattedInput)
    }
  }, [balances, inputAmount, originToken, destinationToken])

  const isInputValid: boolean = useMemo(() => {
    if (debouncedInputAmount === '') return false
    if (isOnlyZeroes(debouncedInputAmount)) return false
    return true
  }, [debouncedInputAmount])

  const onSelectedChain: boolean = useMemo(() => {
    return networkId === originChainId
  }, [originChainId, networkId])

  return {
    hasValidSelections,
    hasEnoughBalance,
    isInputValid,
    // hasEnoughApproved,
    onSelectedChain,
  }
}
