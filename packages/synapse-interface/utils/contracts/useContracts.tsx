import { useNetwork, erc20ABI, getContract } from 'wagmi'
import { Token } from '../types'

/**
 * @param {Token} token token the token used
 */
export function useTokenContract({ token }: { token: Token }) {
  const { chain } = useNetwork()

  const contract = getContract({
    address: token ? token.addresses[chain.id] : '',
    abi: erc20ABI,
  })

  return contract
}
