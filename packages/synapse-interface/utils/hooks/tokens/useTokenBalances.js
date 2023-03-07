import _ from 'lodash'
import { Zero } from '@ethersproject/constants'

import { CHAIN_PARAMS } from '@constants/networks'
import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'
import { useMiniChefContract } from '@hooks/contracts/useMiniChefContract'

import {
  useTokenContract,
  useGenericTokenContract,
} from '@hooks/contracts/useContract'

import { usePool } from '@hooks/pools/usePools'

import {
  useSingleCallResult,
  useMultipleContractSingleData,
  useSingleContractMultipleMethods,
} from '@hooks/multicall'
import { ERC20_INTERFACE } from '@constants/interfaces'
import {
  useGenericMulticall2Contract,
  useMulticall2Contract,
} from '@hooks/contracts/useMulticallContract'
import { formatCommifyBn } from '@bignumber/format'

/**
 * @param {Token} t
 * @return {BigNumber}
 */
export function useTokenBalance(t) {
  const { account, chainId } = useActiveWeb3React()

  const multicallContract = useMulticall2Contract()

  const tokenContract = useTokenContract(t)

  let sharedArgs
  if (t?.isNative) {
    sharedArgs = [
      chainId,
      multicallContract,
      'getEthBalance',
      [account],
      { resultOnly: true },
    ]
  } else {
    sharedArgs = [
      chainId,
      tokenContract,
      'balanceOf',
      [account],
      { resultOnly: true },
    ]
  }
  const singleCallResult = useSingleCallResult(...sharedArgs)

  return singleCallResult?.balance ?? Zero // balance
}

/**
 * @param {number} chainId
 * @param {Token} t
 */
export function useGenericTokenBalance(chainId, t) {
  const { account } = useActiveWeb3React()
  const tokenContract = useGenericTokenContract(chainId, t)
  const multicallContract = useGenericMulticall2Contract(chainId)

  let sharedArgs
  if (t?.isNative) {
    sharedArgs = [
      chainId,
      multicallContract,
      'getEthBalance',
      [account],
      { resultOnly: true },
    ]
  } else {
    sharedArgs = [
      chainId,
      tokenContract,
      'balanceOf',
      [account],
      { resultOnly: true },
    ]
  }

  const singleCallResult = useSingleCallResult(...sharedArgs)

  return singleCallResult?.balance ?? Zero // balance
}

/**
 * @param {Token[]} tokens
 */
export function useTokenBalances(tokens) {
  const { account, chainId } = useActiveWeb3React()
  const multipleCallResults = useMultipleContractSingleData(
    chainId,
    tokens?.map((t) => t.addresses[chainId]),
    ERC20_INTERFACE,
    'balanceOf',
    [account],
    { resultOnly: true }
  )

  const balanceObj = _.fromPairs(
    _.zip(
      tokens.map((t) => t.symbol),
      multipleCallResults?.map((item) => item?.balance ?? Zero)
    )
  )

  return balanceObj
}

export function useStakedBalance(poolId) {
  const { account, chainId } = useActiveWeb3React()
  const miniChefContract = useMiniChefContract()

  const [userInfoObject, pendingSynapseObj] = useSingleContractMultipleMethods(
    chainId,
    miniChefContract,
    {
      userInfo: [poolId, account],
      pendingSynapse: [poolId, account],
    }
  )
  const amount = userInfoObject.result?.amount ?? Zero
  const reward = pendingSynapseObj.result?.[0] ?? Zero

  return { amount, reward }
}

export function usePoolTokenBalances(poolName) {
  const poolTokens = usePool(poolName)
  const poolTokenBalances = useTokenBalances(poolTokens)

  return poolTokenBalances
}

export function getNativeBalance() {
  const { account, chainId } = useActiveWeb3React()

  const multicallContract = useMulticall2Contract()

  const nativeCurrency = CHAIN_PARAMS[chainId].nativeCurrency

  let sharedArgs = [
    chainId,
    multicallContract,
    'getEthBalance',
    [account],
    { resultOnly: true },
  ]

  const singleCallResult = useSingleCallResult(...sharedArgs)

  const balance = singleCallResult?.balance ?? Zero
  return formatCommifyBn(balance, nativeCurrency, 4)
}
