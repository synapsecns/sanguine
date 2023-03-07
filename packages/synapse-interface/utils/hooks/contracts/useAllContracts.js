
import { POOLS_BY_CHAIN, FAKE_POOLS_BY_CHAIN } from '@constants/tokens/poolsByChain'
import { BASIC_TOKENS_BY_CHAIN } from '@constants/tokens/basic'

import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { useTokenContracts } from '@hooks/contracts/useContract'
import { useMemo } from 'react'


export function useAllContracts() {
  const { chainId } = useActiveWeb3React()

  const tokenArr = [
    ...POOLS_BY_CHAIN[chainId],
    ...BASIC_TOKENS_BY_CHAIN[chainId],
    ...FAKE_POOLS_BY_CHAIN[chainId],
  ]

  const allTokenContracts = useTokenContracts(tokenArr)

  return useMemo(() => allTokenContracts, [chainId]) // previously wasnt doing memo here
}
