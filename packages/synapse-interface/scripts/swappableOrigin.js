const fs = require('fs')
const { execSync } = require('child_process')

const { ethers } = require('ethers')

// Provider URLs
const providers = require('./providers.json')
// Contract ABIs
const SynapseRouterABI = require('./abi/SynapseRouter.json')
const SynapseCCTPRouterABI = require('./abi/SynapseCCTPRouter.json')
const SwapQuoterABI = require('./abi/SwapQuoter.json')
// ETH address
const ETH = '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'

// Format is { chainId: providerUrl }
// Replace providerUrl with new ethers.providers.JsonRpcProvider(providerUrl)
Object.keys(providers).forEach((chainId) => {
  providers[chainId] = new ethers.providers.JsonRpcProvider(providers[chainId])
})

// Contract addresses
const SynapseRouterAddress = '0x7e7a0e201fd38d3adaa9523da6c109a07118c96a'
const SynapseCCTPRouterAddress = '0xd359bc471554504f683fbd4f6e36848612349ddf'

// Chain IDs where SynapseCCTPRouter is allowed
const allowedChainIdsForSynapseCCTPRouter = [1, 42161, 43114]

// Get SynapseRouter contract instances for each chain
const SynapseRouters = {}
const SwapQuoters = {}
Object.keys(providers).forEach((chainId) => {
  SynapseRouters[chainId] = new ethers.Contract(
    SynapseRouterAddress,
    SynapseRouterABI,
    providers[chainId]
  )
  SwapQuoters[chainId] = new ethers.Contract(
    SynapseRouters[chainId].swapQuoter(),
    SwapQuoterABI,
    providers[chainId]
  )
})

// Get SynapseCCTPRouter contract instances for each chain
// Only include chains where SynapseCCTPRouter is allowed
const SynapseCCTPRouters = {}
Object.keys(allowedChainIdsForSynapseCCTPRouter).forEach((chainId) => {
  SynapseCCTPRouters[chainId] = new ethers.Contract(
    SynapseCCTPRouterAddress,
    SynapseCCTPRouterABI,
    providers[chainId]
  )
})

// Function to get list of tokens that could be swapped into SynapseBridge tokens for a given chain
const getBridgeSwappableOrigin = async (chainId) => {
  // Get WETH address
  const weth = await SwapQuoters[chainId].weth()
  // Get list of supported tokens
  const bridgeTokens = await SynapseRouters[chainId].bridgeTokens()
  const pools = await SynapseRouters[chainId].allPools()

  // Collect map from bridge token to symbols by doing tokenToSymbol for each bridge token
  const bridgeTokenSymbols = {}
  await Promise.all(
    bridgeTokens.map(async (bridgeToken) => {
      const symbol = await SynapseRouters[chainId].tokenToSymbol(bridgeToken)
      bridgeTokenSymbols[bridgeToken] = symbol
    })
  )
  // Collect map from supported tokens into set of bridge token symbols
  const tokenBridgeSymbols = {}
  // Add all bridge tokens to tokenBridgeSymbols
  bridgeTokens.forEach((bridgeToken) => {
    tokenBridgeSymbols[bridgeToken] = new Set([bridgeTokenSymbols[bridgeToken]])
  })
  pools.forEach((pool) => {
    // Collect set of bridge token symbols that are present in the pool
    const bridgeSymbols = getBridgeSymbolsSet(
      bridgeTokens,
      bridgeTokenSymbols,
      pool.tokens.map((token) => token.token).concat(pool.lpToken)
    )
    // Add collected set to tokenBridgeSymbols for each token in the pool
    pool.tokens.forEach((token) => {
      addSetToMap(tokenBridgeSymbols, token.token, bridgeSymbols)
    })
    // Add collected set to tokenBridgeSymbols for the pool lpToken (if it is a bridge token)
    if (bridgeTokens.includes(pool.lpToken)) {
      addSetToMap(tokenBridgeSymbols, pool.lpToken, bridgeSymbols)
    }
  })
  // If WETH is present in the map, add ETH with the same set of bridge token symbols
  if (Object.keys(tokenBridgeSymbols).includes(weth)) {
    addSetToMap(tokenBridgeSymbols, ETH, tokenBridgeSymbols[weth])
  }
  return tokenBridgeSymbols
}

// Gets a set of bridge token symbols from a list of tokens
const getBridgeSymbolsSet = (
  bridgeTokens,
  bridgeTokenSymbols,
  tokenAddresses
) => {
  const bridgeSymbols = new Set()
  tokenAddresses.forEach((token) => {
    if (bridgeTokens.includes(token)) {
      bridgeSymbols.add(bridgeTokenSymbols[token])
    }
  })
  return bridgeSymbols
}

// Adds values from set to map[key]
const addSetToMap = (map, key, set) => {
  if (!map[key]) {
    map[key] = new Set()
  }
  set.forEach((value) => {
    map[key].add(value)
  })
}

// Transforms and sorts map {token: set} into {token: array}
const transformMap = (map) => {
  const result = {}
  Object.keys(map)
    .sort()
    .forEach((token) => {
      result[token] = Array.from(map[token]).sort()
    })
  return result
}

const printOriginTokens = async () => {
  // Collect swappableOrigin for each chain
  const swappableOrigin = {}
  await Promise.all(
    Object.keys(providers)
      .sort()
      .map(async (chainId) => {
        swappableOrigin[chainId] = transformMap(
          await getBridgeSwappableOrigin(chainId)
        )
      })
  )
  prettyPrint(swappableOrigin, './data/swappableOrigin.json')
}

// Writes obj to fn as a pretty printed JSON file
const prettyPrint = (obj, fn) => {
  fs.writeFileSync(fn, JSON.stringify(obj, null, 2))
  // Run prettier on the file using terminal command:
  execSync(`npx prettier --write ${fn}`)
}

printOriginTokens()
