import { BigNumber } from '@ethersproject/bignumber'
import { MaxUint256, Zero } from '@ethersproject/constants'

import { useSingleContractMultipleMethods } from '@hooks/multicall'
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'

export function useCheckTokenApproval({ srcTokenContract, swapAddress, spendingValue }) {
  const { chainId, account } = useActiveWeb3React()
  const [ tokenNameArr, allowanceArr ] = useSingleContractMultipleMethods(
    chainId,
    srcTokenContract,
    {
      'name': [],
      'allowance': [account, swapAddress]
    },
    { resultOnly: true }
  )
  const existingAllowance = allowanceArr?.[0]
  const tokenName = tokenNameArr?.[0]


  if (existingAllowance?.gte(spendingValue)) {
    return {
      status: true,
      existingAllowance,
      tokenName
    }
  } else {
    return {
      status: false,
      existingAllowance,
      tokenName
    }
  }
}




