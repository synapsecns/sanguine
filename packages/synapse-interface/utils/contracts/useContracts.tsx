import { useNetwork, erc20ABI, useContract } from 'wagmi'
import { Token } from '../types'

/**
 * @param {Token} token token the token used
 */
export function useTokenContract({ token }: { token: Token }) {
  const { chain } = useNetwork()

  const contract = useContract({
    address: token ? token.addresses[chain.id] : '',
    abi: erc20ABI,
  })

  return contract
}
