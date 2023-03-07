import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'

import { ChainId } from '@constants/networks'
import { SYN, NUSD, NETH, ETH, WETH, SYN_FRAX, FRAX, ONEETH, WMOVR, WAVAX, WJEWEL, SYNJEWEL, SYNAVAX } from '@constants/tokens/basic'
import { BRIDGE_SWAPABLE_TYPE_POOLS_BY_CHAIN, MINT_BURN_TOKENS } from '@constants/tokens/tokenGroups'

import { MAX_GAS_THRESHOLD } from '@utils/gas'

import { useBridgeZapContract, useGenericBridgeZapContract } from '@hooks/contracts/useBridgeZapContract'
import { useBridgeConfigContract } from '@hooks/contracts/useBridgeConfigContract'
import { matchSymbolWithinPool } from '@utils/matchSymbolWithinPool'


const OTHER_CHAINS_WITH_ETH = [
  ChainId.ARBITRUM,
  ChainId.BOBA,
  ChainId.FANTOM,
  ChainId.OPTIMISM,
  ChainId.AVALANCHE,
  ChainId.HARMONY,
  ChainId.METIS,
  ChainId.MOONBEAM,
  ChainId.KLAYTN, 
  ChainId.DOGECHAIN,
  ChainId.CANTO
]

export function useCalculateBridgeRate({fromChainId, toChainId}) {
  const bridgeConfigContract = useBridgeConfigContract()
  const fromChainZapContract = useBridgeZapContract()
  const toChainZapContract = useGenericBridgeZapContract(toChainId)
  // console.log({ fromChainZapContract, toChainZapContract, toChainId})

  return async function calculateBridgeRate({
    fromCoin,
    toCoin,
    amountToGive,
  }) {
    if ((fromChainId != ChainId.TERRA) && !fromChainZapContract) throw new Error('Bridge Zap contract is not loaded')
    if ((toChainId != ChainId.TERRA) && !toChainZapContract) throw new Error('Bridge Zap contract is not loaded')

    if (fromCoin.symbol == ETH.symbol) {
      fromCoin = WETH
    }

    if (toCoin.symbol == ETH.symbol) {
      toCoin = WETH
    }

    const fromChainTokens = BRIDGE_SWAPABLE_TYPE_POOLS_BY_CHAIN[fromChainId][fromCoin.swapableType].poolTokens
    const toChainTokens   = BRIDGE_SWAPABLE_TYPE_POOLS_BY_CHAIN[toChainId][toCoin.swapableType].poolTokens

    const tokenIndexFrom = fromChainTokens.findIndex(i => matchSymbolWithinPool(i, fromCoin) )
    const tokenIndexTo   = toChainTokens.findIndex(i => matchSymbolWithinPool(i, toCoin) )

    // console.log({ fromChainTokens, toChainTokens, tokenIndexFrom, tokenIndexTo })
    let intermedieteToken
    let bridgeConfigIntermedieteToken
    if (fromCoin.swapableType === "SYN") {
      intermedieteToken = SYN
    } else if (["LINK", "HIGHSTREET", "DOG", "JUMP", "SFI", "NFD", "OHM", "SOLAR", "GMX", "NEWO", "VSTA", "SDT", "UNIDX", "USDB", "UST", "XJEWEL", "H2O", "L2DAO", "PLS", "AGEUR", "WBTC", "KLAYTN_USDC", "KLAYTN_USDT", "KLAYTN_DAI", "DOGECHAIN_BUSD"].includes(fromCoin.swapableType)) {
      intermedieteToken = fromCoin
    } else if (fromCoin.swapableType === "JEWEL") {
      intermedieteToken = WJEWEL
      if ([ChainId.HARMONY, ChainId.AVALANCHE].includes(toChainId)) {
        bridgeConfigIntermedieteToken = SYNJEWEL
      }

    } else if (fromCoin.swapableType === "AVAX") {
      // if (
      //   (fromChainId == ChainId.HARMONY) &&
      //   (toChainId == ChainId.AVALANCHE)
      // ) {
      if (
          (fromChainId == ChainId.HARMONY) &&
          (toChainId == ChainId.AVALANCHE)
      ) {
        bridgeConfigIntermedieteToken = WAVAX
        intermedieteToken = WAVAX //SYNAVAX
      }
      else if (
        (fromChainId == ChainId.AVALANCHE) &&
        (toChainId == ChainId.HARMONY)
      ) {
        bridgeConfigIntermedieteToken = SYNAVAX
        intermedieteToken = SYNAVAX
      }
      else {
        intermedieteToken = WAVAX
      }
    } else if (fromCoin.swapableType === "MOVR") {
      intermedieteToken = WMOVR
    } else if (fromCoin.swapableType == "FRAX") {
      if (toChainId == ChainId.ETH) {
        bridgeConfigIntermedieteToken = FRAX
      } else {
        bridgeConfigIntermedieteToken = SYN_FRAX
      }
    } else if (fromCoin.swapableType === "ETH") {
      intermedieteToken = NETH
      if (toChainId == ChainId.ETH) {
        bridgeConfigIntermedieteToken = WETH
      } else {
        bridgeConfigIntermedieteToken = NETH
      }
    } else {
      intermedieteToken = NUSD
    }
    bridgeConfigIntermedieteToken = bridgeConfigIntermedieteToken ?? intermedieteToken
    // console.log({ bridgeConfigIntermedieteToken })
    /**
    /**
     * FYI Bridge fee done in decimals of NUSD/SYN (18 decimals)
     */

    let atg
    let multiplier
    if ([fromCoin.swapableType, toCoin.swapableType].includes("UST")) {
      multiplier = BigNumber.from(10).pow(12)
      atg = amountToGive
      // .div(
      //   BigNumber.from(10).pow(6)
      // )

    } else {
      multiplier = 1
      if ((fromChainId === ChainId.KLAYTN || toChainId === ChainId.KLAYTN) || (fromChainId === ChainId.DOGECHAIN || toChainId === ChainId.DOGECHAIN)) {
        atg = amountToGive
      } else {
        atg = amountToGive
        .mul(
          BigNumber.from(10).pow(18 - fromCoin.decimals[fromChainId])
        )
      }
    }

    const bridgeFeeRequest = bridgeConfigContract["calculateSwapFee(string,uint256,uint256)"](
      _.toLower(
        bridgeConfigIntermedieteToken.wrapperAddresses?.[toChainId]
          ?? bridgeConfigIntermedieteToken.addresses[toChainId]
      ),
      toChainId,
      atg
    )
    // console.log("address: ", bridgeConfigIntermedieteToken.addresses[toChainId])
    // console.log("chainId: ",  toChainId)
    // console.log("atg: ", atg.toString())
    // console.log({fromCoin, toCoin})
    // console.log(await bridgeFeeRequest)

    let amountToReceiveFromChain
    if (amountToGive.isZero()) {
      amountToReceiveFromChain = Zero
    } else if (MINT_BURN_TOKENS.map(t => t.symbol).includes(fromCoin.symbol)) {
      amountToReceiveFromChain = amountToGive
    } else if (fromChainId == ChainId.ETH) {
      if ( OTHER_CHAINS_WITH_ETH.includes(toChainId) && (toCoin.swapableType == "ETH")) {
        amountToReceiveFromChain = amountToGive
      } else {
        const liquidityAmounts = fromChainTokens.map(t => {
          if (matchSymbolWithinPool(t, fromCoin)) {
            return amountToGive
          } else {
            return 0
          }
        })
        amountToReceiveFromChain = await fromChainZapContract.calculateTokenAmount(
          liquidityAmounts,
          true,
          { gasLimit: MAX_GAS_THRESHOLD[fromChainId] }
        )
      }
    } else {
      // console.log([intermedieteToken.addresses[fromChainId],
      //   tokenIndexFrom,
      //   0,
      //   amountToGive,])
      amountToReceiveFromChain = await fromChainZapContract.calculateSwap(
        intermedieteToken.addresses[fromChainId],
        tokenIndexFrom,
        0,
        amountToGive,
      )
    }

    // const bridgeFee = await bridgeConfigContract.calculateSwapFee(
    //   intermedieteToken.addresses[toChainId],
    //   toChainId,
    //   amountToReceiveFromChain
    // )

    const bridgeFee = await bridgeFeeRequest

    // console.log({ amountToReceiveFromChain, amountToGive, bridgeFee, toCoin })
    // console.log(bridgeFee)
    amountToReceiveFromChain = safeBnSubtract(
      amountToReceiveFromChain,
      bridgeFee
    )


    let amountToReceiveToChain
    if (amountToReceiveFromChain.isZero()) {
      amountToReceiveToChain = Zero
    } else if (MINT_BURN_TOKENS.map(t => t.symbol).includes(toCoin.symbol)) {
      amountToReceiveToChain = amountToReceiveFromChain
    } else if (toChainId == ChainId.ETH) {
      if ( OTHER_CHAINS_WITH_ETH.includes(fromChainId) && (fromCoin.swapableType == "ETH") ) {
        amountToReceiveToChain = amountToReceiveFromChain
      } else {
        amountToReceiveToChain = await toChainZapContract.calculateRemoveLiquidityOneToken(
          amountToReceiveFromChain,
          tokenIndexTo,
        )
      }

    } else {
      amountToReceiveToChain = await toChainZapContract.calculateSwap(
        intermedieteToken.addresses[toChainId],
        0,
        tokenIndexTo,
        amountToReceiveFromChain
      )
    }

    return {
      amountToReceive: amountToReceiveToChain,
      bridgeFee: bridgeFee?.mul(multiplier) ?? Zero
      // bridgeFee: bridgeFee ?? Zero
    }
    // return Promise.all([amountToReceiveToChain, bridgeFee ?? Zero])
  }
}


function safeBnSubtract(a, b) {
  if (a.gt(b)) {
    return a.sub(b)
  } else {
    return Zero
  }
}

