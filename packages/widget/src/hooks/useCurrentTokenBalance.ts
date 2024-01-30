import { useBridgeState } from '@/state/slices/bridge/hooks'
import { useContext, useMemo } from 'react'
import { Web3Context } from 'providers/Web3Provider'
import { TokenBalance } from '@/utils/actions/fetchTokenBalances'
import { useWalletState } from '@/state/slices/wallet/hooks'

export const useCurrentTokenBalance = () => {
  const web3Context = useContext(Web3Context)
  const { connectedAddress } = web3Context.web3Provider
  const { originChainId, originToken } = useBridgeState()
  const { balances } = useWalletState()

  return useMemo(() => {
    if (!Array.isArray(balances) || !originToken) {
      return {
        rawBalance: null,
        parsedBalance: null,
        decimals: null,
      }
    } else {
      const matchedTokenBalance = balances?.find(
        (token: TokenBalance) =>
          token?.token?.addresses[originChainId] ===
          originToken?.addresses[originChainId]
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
  }, [balances, originToken, originChainId, connectedAddress])
}
