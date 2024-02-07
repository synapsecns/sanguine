import { Address, useNetwork, erc20ABI } from 'wagmi'
import { getContract } from 'wagmi/actions'
import { Token } from '../types'

/**
 * @param {Token} token token the token used
 */
export function useTokenContract({ token }: { token: Token }) {
  const { chain } = useNetwork()

  const contract = getContract({
    address: token ? (token.addresses[chain.id] as Address) : `0x`,
    abi: erc20ABI,
  })

  return contract
}
