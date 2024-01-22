import _ from 'lodash'
import { CHAINS } from 'synapse-constants'

const ChainId = CHAINS.ChainId

/**
 * Represents an ERC20-like token with a unique address, chainId, and some metadata.
 */
export class Token {
  addresses?: any
  decimals?: any
  depositTokens?: any
  description?: any
  docUrl?: any
  icon?: any
  isEthSwap?: any
  isMeta?: any
  isNative?: any
  name?: any
  nativeTokens?: any
  poolId?: any
  poolName?: any
  poolTokens?: any
  poolType?: any
  routerIndex?: any
  swapAddresses?: any
  swapDepositAddresses?: any
  swapEthAddresses?: any
  swapWrapperAddresses?: any
  swapableType?: any
  symbol?: any
  wrapperAddresses?: any
  /**
   * @param {Object} obj - An object.
   * @param {Object.<number,string>} [obj.addresses] - addresses of the actual token
   * @param {Object.<number,string>} [obj.wrapperAddresses] - addresses of the wrapper token in cases like gmx
   * @param {(number|Object.<number,number>)} [obj.decimals] - decimals of the currency
   * @param {string} [obj.symbol] - symbol of the currency
   * @param {string} [obj.name] - human readable name of the currency
   * @param {string} [obj.logo] - logo path
   * @param {string} [obj.poolName] - poolName for pool
   * @param {Object.<number,string>} [obj.swapAddresses] - standard
   * @param {Object.<number,string>} [obj.swapDepositAddresses] - for (metaswap)
   * @param {Object.<number,string>} [obj.swapEthAddresses] - for (eth/native)
   * @param {string} [obj.routerIndex] - for internal url path routing
   * @param {(number|Object.<number,number>)} [obj.poolId] - poolId used in smart contract
   * @param {Token[]} [obj.poolTokens] - poolTokens (basic tokens)
   * @param {string} [obj.description] - description of the token
   * @param {string} [obj.docUrl] - url link to the documentation
   * @param {boolean} [obj.forceMeta] - force the token to be treated as a
   * @param {boolean} [obj.isNative] - whether the token is a native token
   * @param {string} [obj.swapableType] - the type of swap a la `SYN`, `ETH`, etc
   *  metaswap even if swapDepositAddresses isnt present
   */
  constructor({
    addresses,
    wrapperAddresses,
    decimals,
    symbol,
    name,
    logo,
    poolName,
    swapAddresses,
    swapWrapperAddresses,
    swapDepositAddresses,
    swapEthAddresses,
    routerIndex,
    poolId,
    poolType,
    poolTokens,
    depositTokens,
    nativeTokens,
    description,
    docUrl,
    forceMeta,
    isNative,
    swapableType,
  }: {
    addresses?: any
    wrapperAddresses?: any
    decimals?: any
    symbol?: any
    name?: any
    logo?: any
    poolName?: any
    swapAddresses?: any
    swapWrapperAddresses?: any
    swapDepositAddresses?: any
    swapEthAddresses?: any
    routerIndex?: any
    poolId?: any
    poolType?: any
    poolTokens?: any
    depositTokens?: any
    nativeTokens?: any
    description?: any
    docUrl?: any
    forceMeta?: any
    isNative?: any
    swapableType?: any
  }) {
    let isMeta
    if (swapDepositAddresses || forceMeta) {
      isMeta = true
    } else {
      isMeta = false
    }
    this.addresses = addresses
    this.wrapperAddresses = wrapperAddresses
    // this.decimals             = decimals
    this.decimals = makeMultiChainObj(decimals)
    this.symbol = symbol
    this.name = name
    this.icon = logo
    this.poolName = poolName
    this.swapAddresses = swapAddresses
    this.swapWrapperAddresses = swapWrapperAddresses
    this.swapDepositAddresses = swapDepositAddresses
    this.swapEthAddresses = swapEthAddresses
    this.routerIndex = routerIndex

    this.poolTokens = poolTokens
    this.nativeTokens = nativeTokens ?? poolTokens
    this.depositTokens = depositTokens ?? this.nativeTokens
    this.description = description
    this.docUrl = docUrl ?? ''

    this.poolId = makeMultiChainObj(poolId)
    this.poolType = poolType

    this.isMeta = isMeta
    this.isEthSwap = swapEthAddresses ? true : false
    this.isNative = isNative ?? false
    this.swapableType = swapableType
  }
  /**
   * Returns true if the two tokens are equivalent, i.e. have the same chainId and address.
   *
   * @param other other token to compare
   */
  // equals(otherToken) {
  //   // short circuit on reference equality
  //   if (this === otherToken) {
  //     return true
  //   } else {
  //     return (this.address === otherToken.address)
  //   }

  // }
}

function makeMultiChainObj(valOrObj) {
  if (_.isObject(valOrObj)) {
    return valOrObj
  } else {
    const obj = {}
    for (const [chainName, chainId] of _.entries(ChainId)) {
      obj[chainId] = valOrObj
    }
    return obj
  }
}
