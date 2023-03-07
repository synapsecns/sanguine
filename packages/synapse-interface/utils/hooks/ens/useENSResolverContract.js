import ENS_PUBLIC_RESOLVER_ABI from '@abis/ensPublicResolver.json'
import { ChainId } from '@constants/networks'
import { useGenericContract } from '@hooks/contracts/useContract'

export function useENSResolverContract(address, withSignerIfPossible) {

  return useGenericContract(ChainId.ETH, address, ENS_PUBLIC_RESOLVER_ABI, withSignerIfPossible)
}