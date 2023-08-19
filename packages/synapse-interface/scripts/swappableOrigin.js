const fs = require('fs')
const { execSync } = require('child_process')

const { ethers } = require('ethers')

// Provider URLs
const providers = require('./providers.json')
// Contract ABIs
const SynapseRouterABI = require('./abi/SynapseRouter.json')
const SynapseCCTPABI = require('./abi/SynapseCCTP.json')
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
const SynapseCCTPAddress = '0xfB2Bfc368a7edfD51aa2cbEC513ad50edEa74E84'

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
const SynapseCCTPs = {}
allowedChainIdsForSynapseCCTPRouter.forEach((chainId) => {
  SynapseCCTPRouters[chainId] = new ethers.Contract(
    SynapseCCTPRouterAddress,
    SynapseCCTPRouterABI,
    providers[chainId]
  )
  SynapseCCTPs[chainId] = new ethers.Contract(
    SynapseCCTPAddress,
    SynapseCCTPABI,
    providers[chainId]
  )
})

// Function to get list of tokens that could be swapped
// into SynapseBridge or SynapseCCTP tokens for a given chain.
const getSwappableOrigin = async (chainId) => {
  // Get WETH address
  const weth = await SwapQuoters[chainId].weth()
  // Get list of supported tokens
  const bridgeTokens = await SynapseRouters[chainId].bridgeTokens()
  const pools = await SynapseRouters[chainId].allPools()

  // Collect map from bridge token to symbols by doing tokenToSymbol for each bridge token
  const allTokenSymbols = {}
  await Promise.all(
    bridgeTokens.map(async (bridgeToken) => {
      const symbol = await SynapseRouters[chainId].tokenToSymbol(bridgeToken)
      allTokenSymbols[bridgeToken] = new Set([symbol])
    })
  )
  // Add entries from CCTP bridge tokens
  const cctpTokenSymbols = await getCCTPBridgeSymbols(chainId)
  Object.keys(cctpTokenSymbols).forEach((token) => {
    addSetToMap(allTokenSymbols, token, cctpTokenSymbols[token])
  })
  // Collect map from supported tokens into set of bridge token symbols
  const swappableToSymbols = {}
  // Add all bridge tokens to swappableToSymbols
  Object.keys(allTokenSymbols).forEach((bridgeToken) => {
    swappableToSymbols[bridgeToken] = new Set(allTokenSymbols[bridgeToken])
  })
  pools.forEach((pool) => {
    // Collect set of bridge token symbols that are present in the pool
    const bridgeSymbols = getBridgeSymbolsSet(
      allTokenSymbols,
      pool.tokens.map((token) => token.token).concat(pool.lpToken)
    )
    // Add collected set to swappableToSymbols for each token in the pool
    pool.tokens.forEach((token) => {
      addSetToMap(swappableToSymbols, token.token, bridgeSymbols)
    })
    // Add collected set to swappableToSymbols for the pool lpToken (if it is a bridge token)
    if (bridgeTokens.includes(pool.lpToken)) {
      addSetToMap(swappableToSymbols, pool.lpToken, bridgeSymbols)
    }
  })
  // If WETH is present in the map, add ETH with the same set of bridge token symbols
  if (weth in swappableToSymbols) {
    addSetToMap(swappableToSymbols, ETH, swappableToSymbols[weth])
  }
  return swappableToSymbols
}

const getCCTPBridgeSymbols = async (chainId) => {
  // Return empty map if CCTP is not supported on the chain
  if (!SynapseCCTPs[chainId]) {
    return {}
  }
  // Get a list of bridge tokens
  // List entries are (symbol, token) pairs
  const cctpTokens = await SynapseCCTPs[chainId].getBridgeTokens()
  // Return map from token to set of bridge token symbols
  const cctpTokenSymbols = {}
  cctpTokens.forEach((bridgeToken) => {
    addSetToMap(
      cctpTokenSymbols,
      bridgeToken.token,
      new Set([bridgeToken.symbol])
    )
  })
  return cctpTokenSymbols
}

// Gets a set of bridge token symbols from a list of tokens
const getBridgeSymbolsSet = (allTokenSymbols, tokenAddresses) => {
  const bridgeSymbols = new Set()
  tokenAddresses.forEach((token) => {
    if (token in allTokenSymbols) {
      // Add all values from allTokenSymbols[token] to bridgeSymbols
      allTokenSymbols[token].forEach((value) => {
        bridgeSymbols.add(value)
      })
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
          await getSwappableOrigin(chainId)
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
