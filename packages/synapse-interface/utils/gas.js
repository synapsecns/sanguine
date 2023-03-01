import _ from 'lodash'
import { GasPriceOracle } from 'gas-price-oracle'

import { ChainId, CHAIN_PARAMS } from '@constants/networks'


export const GAS_PRICES_NUMERIC = {
  DEFAULT:  5,
  STANDARD: 6,
  FAST:     7,
  INSTANT:  15,
  CUSTOM:   20,
}



export const GAS_PRICES_AVALANCHE = {
  DEFAULT:  225,
  STANDARD: 225,
  FAST:     225,
  INSTANT:  225,
  CUSTOM:   225,
}


const GAS_PRICE_ORACLE_OPTIONS = {
  chainId: ChainId.ETH,
  defaultRpc: CHAIN_PARAMS[ChainId.ETH].rpcUrls[0],
  timeout: 10000,
  defaultFallbackGasPrices: {
    instant:  69,
    fast:     42,
    standard: 21,
    low:      11,
  },
}


const DEFAULT_MAX_GAS = 10000000

export const MAX_GAS_THRESHOLD = {
  [ChainId.ETH]:       DEFAULT_MAX_GAS,
  [ChainId.BSC]:       DEFAULT_MAX_GAS,
  [ChainId.POLYGON]:   DEFAULT_MAX_GAS,
  [ChainId.FANTOM]:    DEFAULT_MAX_GAS,
  [ChainId.ARBITRUM]:  DEFAULT_MAX_GAS,
  [ChainId.OPTIMISM]:  DEFAULT_MAX_GAS,
  [ChainId.AVALANCHE]: DEFAULT_MAX_GAS

}


/**
 * @param {string} gasPriceSelected (DEFAULT|STANDARD|FAST|INSTANT|CUSTOM)
 * @param {number} chainId
 */
export async function getGasPrice(gasPriceSelected, chainId) {
  if (chainId == ChainId.ETH) {
    const GAS_PRICE_ORACLE = new GasPriceOracle(GAS_PRICE_ORACLE_OPTIONS)
    const gasPrices = await GAS_PRICE_ORACLE.gasPrices()
    let fmtKey
    if (!gasPriceSelected || ['DEFAULT', ''].includes(gasPriceSelected)) {
      fmtKey = 'fast'
    } else {
      fmtKey = _.toLower(gasPriceSelected)
    }
    return gasPrices[fmtKey]

  } else if (chainId == ChainId.AVALANCHE){
    return GAS_PRICES_AVALANCHE[gasPriceSelected] ?? GAS_PRICES_AVALANCHE.STANDARD
  } else {
    return GAS_PRICES_NUMERIC[gasPriceSelected] ?? GAS_PRICES_NUMERIC.STANDARD
  }

}



/**
 * Need to add gas price oracle on matic & non eth/bsc chains
 */


// optional fallbackGasPrices

