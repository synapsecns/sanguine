import TOKEN_DISTRIBUTOR from '@abis/tokenDistributor.json'
import { NRV_AIRDROP_ADDRESSES, SYN_AIRDROP_ADDRESSES } from "@constants/airdrop"
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'


import { useContract } from "@hooks/contracts/useContract"

export function useAirdropContracts() {
  const { chainId, account } = useActiveWeb3React()
  const nrvAirdropContract = useContract(NRV_AIRDROP_ADDRESSES[chainId], TOKEN_DISTRIBUTOR)
  const synAirdropContract = useContract(SYN_AIRDROP_ADDRESSES[chainId], TOKEN_DISTRIBUTOR)

  return {
    nrvAirdropContract, synAirdropContract
  }
}


