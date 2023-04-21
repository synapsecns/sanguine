import { Zero } from '@ethersproject/constants'
// import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'
import { useGenericPoolApyData } from '@hooks/pools/useGenericPoolApyData'
import { useEthPrice, useAvaxPrice } from '@hooks/usePrices'
import {
  calcBnSum,
  calcIfZero,
  getTokenBalanceInfo,
  getPoolTokenInfoArr,
  MAX_BN_POW,
} from '@utils/poolDataFuncs'
import { STAKING_MAP_TOKENS } from '@constants/tokens'
import { useSynapseContext } from '@utils/SynapseProvider'
import { fetchBalance } from '@wagmi/core'
import { PoolTokenObject } from '@types'

export const useGenericPoolData = async (chainId, poolName, address) => {
  const SynapseSDK = useSynapseContext()
  const poolToken = STAKING_MAP_TOKENS[chainId][poolName]
  const { poolType, addresses, poolTokens, swapAddresses } = poolToken
  const poolTokenAddress = addresses[chainId]
  const poolAddress = swapAddresses[chainId]
  if (!poolTokenAddress || !poolAddress) {
    return null
  }

  // get LP token
  const lpTokenAddress = (await SynapseSDK.getPoolInfo(chainId, poolAddress))
    ?.lpToken

  // get balances of all tokens in the pool
  const tokenBalances: PoolTokenObject[] = []
  const userTokenBalances: PoolTokenObject[] = []
  let poolTokenSum = Zero
  let userLpTokenBalance = Zero
  let lpTokenBalance = Zero
  for (const token of poolTokens.length) {
    const isLP = token.addresses[chainId] === lpTokenAddress

    const rawBalance =
      (
        await fetchBalance({
          address: poolAddress,
          chainId,
          token: token.addresses[chainId],
        })
      )?.value ?? Zero
    const rawUserBalance =
      (
        await fetchBalance({
          address,
          chainId,
          token: token.addresses[chainId],
        })
      )?.value ?? Zero
    const balance = rawBalance?.pow(18 - token.decimals[chainId]) ?? Zero
    const userBalance =
      rawUserBalance?.pow(18 - token.decimals[chainId]) ?? Zero

    // add to balances
    tokenBalances.push({
      rawBalance,
      balance,
      token,
      isLP,
    })

    // add to user balances
    userTokenBalances.push({
      rawBalance: rawUserBalance,
      balance: userBalance,
      token,
      isLP,
    })

    // set lp variables
    if (isLP) {
      userLpTokenBalance = userBalance
      lpTokenBalance = balance
    }

    // running sum of all tokens in the pool
    poolTokenSum = poolTokenSum.add(balance)
  }

  const virtualPriceResult = poolTokenSum.div(
    tokenBalances.filter((t) => t.isLP)[0].balance
  )

  console.log('lp token', `0x${lpTokenAddress.slice(2)}`)

  console.log('tokenBalances', tokenBalances)
  // const swapContract = useGenericSwapContract(chainId, poolName)
  // get
  // const poolApyData = useGenericPoolApyData(chainId, poolToken)

  // const lpTokenContract = useGenericContract(
  //   chainId,
  //   poolToken.addresses[chainId],
  //   LPTOKEN_ABI
  // )

  const ethPrice = useEthPrice()
  const avaxPrice = useAvaxPrice()

  let virtualPrice
  if (lpTokenBalance.isZero()) {
    virtualPrice = MAX_BN_POW
  } else {
    virtualPrice = virtualPriceResult?.[0]
  }

  const { tokenBalancesSum, tokenBalancesUSD } = getTokenBalanceInfo({
    tokenBalances,
    prices: {
      ethPrice,
      avaxPrice,
    },
    poolType,
  })

  // const { adminFee, swapFee } = await swapStorageRequest
  // (weeksPerYear * KEEPPerWeek * KEEPPrice) / (BTCPrice * BTCInPool)

  // User share data
  const userShare = userLpTokenBalance
    .mul(MAX_BN_POW)
    .div(calcIfZero(lpTokenBalance))

  const userPoolTokenBalances = tokenBalances.map((token) =>
    userShare.mul(token.balance).div(MAX_BN_POW)
  )

  const userPoolTokenBalancesSum = calcBnSum(userPoolTokenBalances)

  const sharedArgs = {
    lpTokenBalance: tokenBalances.filter((t) => !t.isLP)[0].balance,
    tokenBalancesSum,
  }

  const generalPoolTokens = getPoolTokenInfoArr({
    tokenBalances,
    ...sharedArgs,
  })

  const userPoolTokens = getPoolTokenInfoArr({
    tokenBalances: userTokenBalances,
    ...sharedArgs,
  })

  const poolDataObj = {
    name: poolName,
    tokens: generalPoolTokens,
    totalLocked: tokenBalancesSum,
    totalLockedUSD: tokenBalancesUSD,
    virtualPrice,
    volume: 'XXX', // TODO
    utilization: 'XXX', // TODO
    apy: poolApyData, //? DIFF
  }

  let userShareData
  if (address) {
    userShareData = {
      name: poolName,
      share: userShare,
      value: userPoolTokenBalancesSum,
      avgBalance: userPoolTokenBalancesSum,
      tokens: userPoolTokens,
      lpTokenBalance: userLpTokenBalance,
      // the code was always doing this, i could not find out why
      lpTokenMinted: userLpTokenBalance,
    }
  } else {
    userShareData = null
  }
  return [poolDataObj, userShareData]

  // return [poolData, userPoolData]
}
