import { useBalance, useNetwork, useAccount, Address } from 'wagmi'
import { Token } from '../types'

export const useTokenBalance = ({ token }: { token: Token }) => {
  const { chain } = useNetwork()
  const { address } = useAccount()

  const tokenAddress = token.addresses[chain.id]
  const addressWithoutHex = tokenAddress.substring(2)
  const formattedAddress: Address = `0x${addressWithoutHex}`

  const balance = useBalance({
    address: address,
    token: formattedAddress,
    chainId: chain.id,
    watch: true,
  })

  return balance
}
