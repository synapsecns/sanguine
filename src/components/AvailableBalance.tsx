import { useMemo, useCallback } from 'react'
import { useAppDispatch } from '@/state/hooks'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { TokenBalance } from '@/utils/actions/fetchTokenBalances'
import { setInputAmount } from '@/state/slices/bridge/reducer'
import { BridgeableToken, Address } from 'types'

export const AvailableBalance = ({
  originChainId,
  originToken,
  inputAmount,
  connectedAddress,
  balances,
}: {
  originChainId: number
  originToken: BridgeableToken
  inputAmount: string
  connectedAddress: Address
  balances: TokenBalance[]
}) => {
  const dispatch = useAppDispatch()

  const currentTokenBalance = useMemo(() => {
    if (!balances) {
      return {
        rawBalance: null,
        parsedBalance: null,
        decimals: null,
      }
    } else {
      const matchedTokenBalance = balances?.find(
        (token: TokenBalance) => token?.token?.addresses[originChainId]
      )
      const decimals: number =
        typeof matchedTokenBalance?.token?.decimals === 'number'
          ? matchedTokenBalance?.token?.decimals
          : matchedTokenBalance?.token?.decimals[originChainId]

      return {
        rawBalance: matchedTokenBalance?.balance,
        parsedBalance: matchedTokenBalance?.parsedBalance,
        decimals: decimals,
      }
    }
  }, [balances])

  const userInputGreaterThanCurrentBalance: boolean = useMemo(() => {
    const formattedInput = stringToBigInt(
      inputAmount,
      currentTokenBalance.decimals
    )
    if (
      formattedInput === undefined ||
      formattedInput === null ||
      currentTokenBalance.rawBalance === undefined ||
      currentTokenBalance.rawBalance === null
    ) {
      return false
    } else {
      return Boolean(formattedInput > BigInt(currentTokenBalance.rawBalance))
    }
  }, [inputAmount, originToken, originChainId, currentTokenBalance])

  const handleAvailableBalanceClick = useCallback(() => {
    const maxAmount: string =
      formatBigIntToString(
        BigInt(currentTokenBalance.rawBalance ?? 0),
        currentTokenBalance.decimals ?? 0,
        18
      ) ?? '0.0'
    dispatch(setInputAmount(maxAmount))
  }, [dispatch, balances, currentTokenBalance])

  if (!connectedAddress) return

  if (userInputGreaterThanCurrentBalance) {
    return (
      <div
        onClick={handleAvailableBalanceClick}
        className="ml-px text-xs text-[--synapse-accent] cursor-pointer hover:underline active:opacity-40"
      >
        {currentTokenBalance.parsedBalance} available
      </div>
    )
  }
  if (connectedAddress) {
    return (
      <div
        onClick={handleAvailableBalanceClick}
        className="ml-px text-xs cursor-pointer hover:underline active:opacity-40 text-[--synapse-text-secondary]"
      >
        {currentTokenBalance.parsedBalance} available
      </div>
    )
  }
}
