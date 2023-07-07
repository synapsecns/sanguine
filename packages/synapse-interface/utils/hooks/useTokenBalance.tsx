import { useBalance, useNetwork, useAccount, Address } from 'wagmi'
import { Token } from '../types'

export const useTokenBalance = (token: Token) => {
  const { chain } = useNetwork()
  const { address } = useAccount()

  const tokenAddress = token && chain && token.addresses[chain.id]

  const balance = useBalance({
    address: address,
    token: tokenAddress as Address,
    chainId: chain && chain.id,
    watch: true,
  })

  return balance
}
