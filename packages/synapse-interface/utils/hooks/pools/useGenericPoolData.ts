import _ from 'lodash'
import { useEffect, useState, useContext } from 'react'
import { BigNumber } from '@ethersproject/bignumber'
import { Zero, One, AddressZero } from '@ethersproject/constants'
import LPTOKEN_ABI from '@abis/lpToken.json'
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
import { STAKING_MAP_TOKENS } from '@constants/tokens/staking'
import { useSynapseContext } from '@utils/SynapseProvider'
import { fetchBalance } from '@wagmi/core'

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
    .lpToken

  // get balances of all tokens in the pool
  const tokenBalances = {}
  for (const poolToken of poolTokens.length) {
    const balance = await fetchBalance({
      address: poolAddress,
      chainId,
      token: poolToken.addresses[chainId],
    })
    tokenBalances[poolToken.addresses[chainId]] = {
      rawBalance: balance,
      balance: Zero,
      token: poolToken,
    }
  }
  const balance = await fetchBalance({
    address: poolAddress,
    chainId,
    token: `0x${lpTokenAddress.slice(2)}`,
  })
  tokenBalances[lpTokenAddress] = {
    rawBalance: balance,
    balance: Zero,
    token: poolToken,
  }
  console.log('lp token', `0x${lpTokenAddress.slice(2)}`)

  console.log('tokenBalances', tokenBalances)
  // const swapContract = useGenericSwapContract(chainId, poolName)

  // get
  const poolApyData = useGenericPoolApyData(chainId, poolToken)

  // const lpTokenContract = useGenericContract(
  //   chainId,
  //   poolToken.addresses[chainId],
  //   LPTOKEN_ABI
  // )

  const ethPrice = useEthPrice()
  const avaxPrice = useAvaxPrice()

  // const virtualPriceResult = await SynapseSDK.getVirtualPrice()
  const virtualPriceResult = Zero

  // Pool token data
  // const tokenBalancesResults = useSingleContractMultipleData(
  //   chainId,
  //   swapContract,
  //   'getTokenBalance',
  //   poolTokens?.map((token, i) => [i]),
  //   { resultOnly: true }
  // )

  // const rawTokenBalances = tokenBalancesResults.map((item) => item?.[0] ?? One) //BigNumber.from(1))

  // const [swapStorageResult, virtualPriceResult] =
  //   useSingleContractMultipleMethods(
  //     chainId,
  //     swapContract,
  //     {
  //       swapStorage: [],
  //       getVirtualPrice: [],
  //     },
  //     { resultOnly: true }
  //   )

  // const [lpTokenBalanceOfResult, totalLpTokenSupplyResult] =
  //   useSingleContractMultipleMethods(
  //     chainId,
  //     lpTokenContract,
  //     {
  //       balanceOf: [address || AddressZero],
  //       totalSupply: [],
  //     },
  //     { resultOnly: true }
  //   )
  // THIS IS THE FRESHLY INTRODUCED CANCER

  // bahahahahhahah

  for (const tokenAddress of Object.keys(tokenBalances)) {
    const token = tokenBalances[tokenAddress]
    tokenBalances[tokenAddress].balance = BigNumber.from(10)
      .pow(18 - token.decimals[chainId]) // cast all to 18 decimals
      .mul(token.rawBalance)
  }
  // const userLpTokenBalance =
  const userLpTokenBalance = Zero
  const totalLpTokenBalance = tokenBalances[lpTokenAddress].balance

  let virtualPrice
  if (totalLpTokenBalance?.isZero()) {
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
      .div(calcIfZero(totalLpTokenBalance))

    const userPoolTokenBalances = tokenBalances.map((balance) =>
      userShare.mul(balance).div(MAX_BN_POW)
    )

    const userPoolTokenBalancesSum = calcBnSum(userPoolTokenBalances)

    const sharedArgs = {
      totalLpTokenBalance,
      tokenBalancesSum,
    }

    const generalPoolTokens = getPoolTokenInfoArr({
      poolTokenBalances: tokenBalances,
      ...sharedArgs,
      poolTokens,
      tokenBalances,
    })

    const userPoolTokens = getPoolTokenInfoArr({
      poolTokenBalances: userPoolTokenBalances,
      ...sharedArgs,
      poolTokens,
      tokenBalances,
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
