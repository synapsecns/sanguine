const fs = require('fs')
const { execSync } = require('child_process')

const { ethers } = require('ethers')

// Provider URLs
const providers = require('../data/providers.json')
// List of ignored bridge symbols
const ignoredBridgeSymbols = require('../data/ignoredBridgeSymbols.json')
// Symbol overrides (for tokens with incorrect on-chain symbols)
const symbolOverrides = require('../data/symbolOverrides.json')
// Contract ABIs
const SynapseRouterABI = require('../abi/SynapseRouter.json')
const SynapseCCTPABI = require('../abi/SynapseCCTP.json')
const SynapseCCTPRouterABI = require('../abi/SynapseCCTPRouter.json')
const SwapQuoterABI = require('../abi/SwapQuoter.json')
const ERC20ABI = require('../abi/IERC20Metadata.json')
const DefaultPoolABI = require('../abi/IDefaultPool.json')
// ETH address
const ETH = '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'

// Format is { chainId: providerUrl }
// Replace providerUrl with new ethers.providers.JsonRpcProvider(providerUrl)
Object.keys(providers).forEach((chainId) => {
  providers[chainId] = new ethers.providers.JsonRpcProvider(providers[chainId])
})

// Contract addresses
const SynapseRouterAddress = '0x7e7a0e201fd38d3adaa9523da6c109a07118c96a'
const SynapseCCTPRouterAddress = '0xd5a597d6e7ddf373a92C8f477DAAA673b0902F48'
const SynapseCCTPAddress = '0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E'

// Chain IDs where SynapseCCTPRouter is allowed
const allowedChainIdsForSynapseCCTPRouter = [1, 10, 137, 8453, 42161, 43114]

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
// into SynapseBridge tokens for a given chain.
const getBridgeOriginMap = async (chainId) => {
  // Get WETH address
  const weth = await SwapQuoters[chainId].weth()
  // Get list of supported tokens
  let bridgeTokens = await SynapseRouters[chainId].bridgeTokens()
  const pools = await SynapseRouters[chainId].allPools()

  // Collect map from bridge token to symbols by doing tokenToSymbol for each bridge token
  const allTokenSymbols = {}
  await Promise.all(
    bridgeTokens.map(async (bridgeToken) => {
      const symbol = await SynapseRouters[chainId].tokenToSymbol(bridgeToken)
      // Skip if symbol is in ignoredBridgeSymbols
      if (!ignoredBridgeSymbols.includes(symbol)) {
        allTokenSymbols[bridgeToken] = new Set([symbol])
      }
    })
  )
  // List of bridge tokens without ignored symbols
  bridgeTokens = Object.keys(allTokenSymbols)
  // Collect map from supported tokens into set of bridge token symbols
  const tokensToSymbols = {}
  // Add all bridge tokens to tokensToSymbols
  Object.keys(allTokenSymbols).forEach((bridgeToken) => {
    tokensToSymbols[bridgeToken] = new Set(allTokenSymbols[bridgeToken])
  })
  // List of sets of tokens in each pool
  const poolSets = []
  pools.forEach((pool) => {
    // Collect set of bridge token symbols that are present in the pool
    const bridgeSymbols = getBridgeSymbolsSet(
      allTokenSymbols,
      pool.tokens.map((token) => token.token).concat(pool.lpToken)
    )
    const poolSet = new Set()
    // Add collected set to tokensToSymbols for each token in the pool
    pool.tokens.forEach((token) => {
      addSetToMap(tokensToSymbols, token.token, bridgeSymbols)
      poolSet.add(token.token)
    })
    // Add collected set to tokensToSymbols for the pool lpToken (if it is a bridge token)
    if (bridgeTokens.includes(pool.lpToken)) {
      addSetToMap(tokensToSymbols, pool.lpToken, bridgeSymbols)
      poolSet.add(pool.lpToken)
    }
    // Save set of tokens in the pool
    poolSets.push(poolSet)
  })
  // If WETH is present in the map, add ETH with the same set of bridge token symbols
  if (weth in tokensToSymbols) {
    addSetToMap(tokensToSymbols, ETH, tokensToSymbols[weth])
    poolSets.forEach((poolSet) => {
      if (poolSet.has(weth)) {
        poolSet.add(ETH)
      }
    })
  }
  return {
    originMap: tokensToSymbols,
    poolSets,
  }
}

// Function to get list of tokens that could be swapped
// into SynapseCCTP tokens for a given chain.
const getCCTPOriginMap = async (chainId) => {
  // Return empty map if CCTP is not supported on the chain
  if (!SynapseCCTPs[chainId]) {
    return {}
  }
  // Get map from token to set of bridge token symbols
  const cctpTokenSymbols = await getCCTPBridgeSymbols(chainId)
  const tokensToSymbols = {}
  // Add all bridge tokens to tokensToSymbols
  Object.keys(cctpTokenSymbols).forEach((token) => {
    tokensToSymbols[token] = new Set([cctpTokenSymbols[token]])
  })
  // Add tokens from whitelisted pools
  await Promise.all(
    Object.keys(cctpTokenSymbols).map(async (cctpToken) => {
      const pool = await SynapseCCTPs[chainId].circleTokenPool(cctpToken)
      const tokens = await getPoolTokens(chainId, pool)
      tokens.forEach((token) => {
        addSetToMap(
          tokensToSymbols,
          token,
          new Set([cctpTokenSymbols[cctpToken]])
        )
      })
    })
  )
  return tokensToSymbols
}

// Function to get a list of bridge token symbols that could be swapped
// into a token on a destination chain.
const getDestinationBridgeSymbols = async (chainId, token) => {
  // Get list of connected bridge tokens: (symbol, token) pairs
  const connectedBridgeTokens = await SynapseRouters[
    chainId
  ].getConnectedBridgeTokens(token)
  const symbolSet = new Set()
  connectedBridgeTokens.forEach((bridgeToken) => {
    symbolSet.add(bridgeToken.symbol)
  })
  // Get a list of bridge token symbols from CCTP if CCTP is supported on the chain
  if (SynapseCCTPRouters[chainId]) {
    const connectedCctpTokens = await SynapseCCTPRouters[
      chainId
    ].getConnectedBridgeTokens(token)
    connectedCctpTokens.forEach((bridgeToken) => {
      symbolSet.add(bridgeToken.symbol)
    })
  }
  return Array.from(symbolSet).sort()
}

// Function to get a map from CCTP tokens to bridge token symbols
const getCCTPBridgeSymbols = async (chainId) => {
  // Return empty map if CCTP is not supported on the chain
  if (!SynapseCCTPs[chainId]) {
    return {}
  }
  // Get a list of bridge tokens
  // List entries are (symbol, token) pairs
  const cctpTokens = await SynapseCCTPs[chainId].getBridgeTokens()
  // Return map from token to bridge symbol
  const cctpTokenToSymbol = {}
  cctpTokens.forEach((bridgeToken) => {
    cctpTokenToSymbol[bridgeToken.token] = bridgeToken.symbol
  })
  return cctpTokenToSymbol
}

// Function to get a list of tokens in a pool
const getPoolTokens = async (chainId, poolAddress) => {
  const pool = new ethers.Contract(
    poolAddress,
    DefaultPoolABI,
    providers[chainId]
  )
  // To get a list of tokens we do getToken(i) calls until we get a revert
  const tokens = []
  for (let i = 0; ; i++) {
    try {
      tokens.push(await pool.getToken(i))
    } catch (e) {
      break
    }
  }
  return tokens
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

const sortMapByKeys = (map) => {
  const sortedMap = {}
  Object.keys(map)
    .sort()
    .forEach((key) => {
      sortedMap[key] = map[key]
    })
  return sortedMap
}

const printMaps = async () => {
  const bridgeMap = {}
  const bridgeSymbolsMap = {}
  console.log('Starting on chains: ', Object.keys(providers))
  await Promise.all(
    Object.keys(providers).map(async (chainId) => {
      // Get map from token to set of bridge token symbols
      const { originMap, poolSets } = await getBridgeOriginMap(chainId)
      // Add tokens from CCTP originMap to global originMap
      const cctpOriginMap = await getCCTPOriginMap(chainId)
      Object.keys(cctpOriginMap).forEach((token) => {
        addSetToMap(originMap, token, cctpOriginMap[token])
      })
      const tokens = {}
      await Promise.all(
        Object.keys(originMap).map(async (token) => {
          tokens[token] = {
            decimals: await getTokenDecimals(chainId, token),
            symbol: await getTokenSymbol(chainId, token),
            origin: Array.from(originMap[token]).sort(),
            destination: await getDestinationBridgeSymbols(chainId, token),
            swappable: extractSwappable(poolSets, token),
          }
        })
      )
      bridgeMap[chainId] = sortMapByKeys(tokens)
      bridgeSymbolsMap[chainId] = sortMapByKeys(extractBridgeSymbolsMap(tokens))
      console.log('Finished chain: ', chainId)
    })
  )
  prettyPrintTS(bridgeMap, 'BRIDGE_MAP', './constants/bridgeMap.ts')
}

// Extracts the list of tokens that can be swapped into a token via single on-chain swap
const extractSwappable = (poolSets, token) => {
  const tokenSet = new Set()
  poolSets.forEach((poolSet) => {
    // If token is in the pool, add all tokens except token to tokenSet
    if (poolSet.has(token)) {
      poolSet.forEach((poolToken) => {
        if (poolToken !== token) {
          tokenSet.add(poolToken)
        }
      })
    }
  })
  return Array.from(tokenSet).sort()
}

const extractBridgeSymbolsMap = (tokens) => {
  const bridgeSymbolsOriginSets = {}
  const bridgeSymbolsDestinationSets = {}
  // Add all bridge symbols that be swapped to/from each token
  Object.keys(tokens).forEach((token) => {
    tokens[token].origin.forEach((symbol) => {
      addSetToMap(bridgeSymbolsOriginSets, symbol, new Set([token]))
    })
    tokens[token].destination.forEach((symbol) => {
      addSetToMap(bridgeSymbolsDestinationSets, symbol, new Set([token]))
    })
  })
  // Bridge symbols are keys that are present in either map
  const bridgeSymbols = new Set([
    ...Object.keys(bridgeSymbolsOriginSets),
    ...Object.keys(bridgeSymbolsDestinationSets),
  ])
  const bridgeSymbolsMap = {}
  bridgeSymbols.forEach((symbol) => {
    bridgeSymbolsMap[symbol] = {
      origin: Array.from(bridgeSymbolsOriginSets[symbol]).sort(),
      destination: Array.from(bridgeSymbolsDestinationSets[symbol]).sort(),
    }
  })
  return bridgeSymbolsMap
}

const getTokenSymbol = async (chainId, token) => {
  // Check if token is ETH
  if (token === ETH) {
    // Get WETH address from SwapQuoter
    const weth = await SwapQuoters[chainId].weth()
    // Return "WETH" symbol without first character
    return getTokenSymbol(chainId, weth).then((symbol) => symbol.slice(1))
  }
  // Check if {chainId: {token: {symbol}}} is in symbolOverrides
  if (chainId in symbolOverrides && token in symbolOverrides[chainId]) {
    return symbolOverrides[chainId][token]
  }
  // Otherwise return symbol from ERC20 contract
  const symbol = await new ethers.Contract(
    token,
    ERC20ABI,
    providers[chainId]
  ).symbol()
  return symbol
}

const getTokenDecimals = async (chainId, token) => {
  // Check if token is ETH
  if (token === ETH) {
    return 18
  }
  // Otherwise return decimals from ERC20 contract
  const decimals = await new ethers.Contract(
    token,
    ERC20ABI,
    providers[chainId]
  ).decimals()
  return decimals
}

// Writes map export to a TypeScript file, then runs prettier on the file
const prettyPrintTS = (map, mapName, fn) => {
  console.log(`Writing ${mapName} to ${fn}`)
  const json = JSON.stringify(map)
  fs.writeFileSync(fn, `export const ${mapName} = ${json}`)
  // Run prettier on the file using terminal command:
  execSync(`npx prettier --write ${fn}`)
}

printMaps()
