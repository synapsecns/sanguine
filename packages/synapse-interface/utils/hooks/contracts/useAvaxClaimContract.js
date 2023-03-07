import AVAX_CLAIM_ABI from '@abis/tokenDistributor.json'
import { AVAX_CLAIM_ADDRESSES } from "@constants/avaxClaim"
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'


import { useContract } from "@hooks/contracts/useContract"

export function useAvaxClaimContract() {
  const { chainId, account } = useActiveWeb3React()
  const avaxClaimContract = useContract(AVAX_CLAIM_ADDRESSES[chainId], AVAX_CLAIM_ABI)

  return avaxClaimContract
}


