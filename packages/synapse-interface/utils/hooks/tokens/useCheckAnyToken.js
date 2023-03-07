import { BigNumber } from '@ethersproject/bignumber'
import { MaxUint256, Zero } from '@ethersproject/constants'
import { useGenericTokenContract, useTokenContract } from '@hooks/contracts/useContract'

import { useSingleCallResult } from '@hooks/multicall'
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'

export function useCheckAnyToken({ targetChainId, token, dappAddr }) {
  const { account } = useActiveWeb3React()
  // const account = "0x157b1c8C3De55B87be545B3DBc9D17f06F0112F8"
  const tokenContract = useGenericTokenContract(targetChainId, token)
  const allowanceArr = useSingleCallResult(
    targetChainId,
    tokenContract,
    'allowance',
    [account, dappAddr],
    { resultOnly: true }
  )

  const existingAllowance = allowanceArr?.[0]

  return existingAllowance?.gt(0)
}




