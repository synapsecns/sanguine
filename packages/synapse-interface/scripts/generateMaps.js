require('dotenv').config()
// TODO: handle HYPE-ETH, HYPE-USDC pairs from the RFQ API
const { ethers } = require('ethers')

const { prettyPrintTS } = require('./utils/prettyPrintTs')
const { fetchGasZipData } = require('./utils/fetchGasZipData')
const { fetchRfqData } = require('./utils/fetchRfqData')
// List of ignored bridge symbols
const ignoredBridgeSymbols = require('./data/ignoredBridgeSymbols.json')
// Symbol overrides (for tokens with incorrect on-chain symbols)
const symbolOverrides = require('./data/symbolOverrides.json')
const providerOverrides = require('./data/providerOverrides.json')
// Contract ABIs
const SynapseRouterABI = require('./abi/SynapseRouter.json')
const SynapseCCTPABI = require('./abi/SynapseCCTP.json')
const SynapseCCTPRouterABI = require('./abi/SynapseCCTPRouter.json')
const SwapQuoterABI = require('./abi/SwapQuoter.json')
const FastBridgeRouterABI = require('./abi/FastBridgeRouter.json')
const ERC20ABI = require('./abi/IERC20Metadata.json')
const DefaultPoolABI = require('./abi/IDefaultPool.json')
// const rfqResponse = require('./data/rfqResponse.json')
// ETH address
const ETH = '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'

// Contract addresses
const SynapseRouterAddress = '0x7e7a0e201fd38d3adaa9523da6c109a07118c96a'
const SynapseCCTPRouterAddress = '0xd5a597d6e7ddf373a92C8f477DAAA673b0902F48'
const SynapseCCTPAddress = '0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E'
const FastBridgeRouterAddress = '0x00cD000000003f7F682BE4813200893d4e690000'

// Chain IDs where SynapseBridge is allowed
const allowedChainIdsForSynapseBridge = [
  1, 10, 25, 56, 137, 250, 288, 1088, 1284, 1285, 2000, 7700, 8217, 8453, 81457,
  42161, 43114, 53935, 1313161554, 1666600000,
]

// Chain IDs where SynapseCCTPRouter is allowed
const allowedChainIdsForSynapseCCTPRouter = [1, 10, 137, 8453, 42161, 43114]

// Chain IDs where RFQ is allowed
const allowedChainIdsForRfq = [
  1, 10, 56, 130, 480, 8453, 42161, 59144, 80094, 81457, 534352,
]

const allChainIds = Array.from(
  new Set([
    ...allowedChainIdsForSynapseBridge,
    ...allowedChainIdsForSynapseCCTPRouter,
    ...allowedChainIdsForRfq,
  ])
)
if (!process.env.RPC_URL) {
  throw new Error('RPC_URL is not defined in the environment variables')
}
// Format is { chainId: provider }
const providers = allChainIds.reduce((acc, chainId) => {
  acc[chainId] = new ethers.providers.JsonRpcProvider(
    providerOverrides[chainId] || `${process.env.RPC_URL}/${chainId}`
  )
  return acc
}, {})

// Get SynapseRouter contract instances for each chain
const SynapseRouters = {}
const SwapQuoters = {}
allowedChainIdsForSynapseBridge.forEach((chainId) => {
  SynapseRouters[chainId] = new ethers.Contract(
    SynapseRouterAddress,
    SynapseRouterABI,
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

const FastBridgeRouters = {}
allowedChainIdsForRfq.forEach((chainId) => {
  FastBridgeRouters[chainId] = new ethers.Contract(
    FastBridgeRouterAddress,
    FastBridgeRouterABI,
    providers[chainId]
  )
})

const getSwapQuoter = async (chainId) => {
  if (SwapQuoters[chainId]) {
    return SwapQuoters[chainId]
  }
  const router = SynapseRouters[chainId] || FastBridgeRouters[chainId]
  const swapQuoterAddr = router ? await router.swapQuoter() : null
  if (!swapQuoterAddr) {
    return null
  }
  SwapQuoters[chainId] = new ethers.Contract(
    swapQuoterAddr,
    SwapQuoterABI,
    providers[chainId]
  )
  return SwapQuoters[chainId]
}

// Function to get list of tokens that could be swapped
// into SynapseBridge tokens for a given chain.
const getBridgeOriginMap = async (chainId) => {
  const swapQuoter = await getSwapQuoter(chainId)
  if (!swapQuoter) {
    return {
      originMap: {},
      poolSets: [],
    }
  }

  // Get WETH address
  const weth = await swapQuoter.weth()
  // Get list of supported tokens
  let bridgeTokens = (await SynapseRouters[chainId]?.bridgeTokens()) || []
  const pools = await swapQuoter.allPools()

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

const getFastBridgeOriginMap = async (chainId, rfqResponse) => {
  // Return empty map if FastBridge is not supported on the chain
  if (!FastBridgeRouters[chainId]) {
    return {}
  }
  const rfqTokens = getRFQBridgeTokens(chainId, rfqResponse)
  // Create map from token to symbol
  const rfqTokenSymbols = {}
  await Promise.all(
    rfqTokens.map(async (token) => {
      rfqTokenSymbols[token] = getRFQSymbol(
        await getTokenSymbol(chainId, token)
      )
    })
  )
  const tokensToSymbols = {}
  // Add all bridge tokens to tokensToSymbols
  Object.keys(rfqTokenSymbols).forEach((token) => {
    tokensToSymbols[token] = new Set([rfqTokenSymbols[token]])
  })
  const swapQuoter = await getSwapQuoter(chainId)
  if (!swapQuoter) {
    return tokensToSymbols
  }
  const weth = await swapQuoter.weth()
  if (rfqTokenSymbols[ETH]) {
    rfqTokenSymbols[weth] = rfqTokenSymbols[ETH]
    tokensToSymbols[weth] = new Set([rfqTokenSymbols[ETH]])
  }
  const pools = await swapQuoter.allPools()
  pools.forEach((pool) => {
    // Get the symbols of supported RFQ tokens in the pool
    const poolRFQSymbols = pool.tokens
      .filter((token) => rfqTokenSymbols[token.token])
      .map((token) => rfqTokenSymbols[token.token])
    if (poolRFQSymbols.length === 0) {
      return
    }
    // Every token in the pools is swappable into these symbols
    pool.tokens.forEach((token) => {
      addSetToMap(tokensToSymbols, token.token, new Set(poolRFQSymbols))
    })
  })
  return tokensToSymbols
}

const getGasZipOriginMap = async (chainId, gasZipChains) => {
  if (!gasZipChains.includes(Number(chainId))) {
    return {}
  }
  const tokensToSymbols = {}
  tokensToSymbols[ETH] = new Set(['Gas.zip'])
  return tokensToSymbols
}

// Function to get a list of bridge token symbols that could be swapped
// into a token on a destination chain.
const getDestinationBridgeSymbols = async (chainId, token) => {
  const symbolSet = new Set()
  if (SynapseRouters[chainId]) {
    // Get list of connected bridge tokens: (symbol, token) pairs
    const connectedBridgeTokens = await SynapseRouters[
      chainId
    ].getConnectedBridgeTokens(token)
    connectedBridgeTokens.forEach((bridgeToken) => {
      symbolSet.add(bridgeToken.symbol)
    })
  }
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

const getRFQBridgeTokens = (chainId, rfqResponse) => {
  return [
    ...new Set(
      rfqResponse
        .filter((quote) => quote.origin_chain_id === Number(chainId))
        .map((quote) => ethers.utils.getAddress(quote.origin_token_addr))
    ),
  ]
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
  console.log('Starting on chains: ', Object.keys(providers))

  const rfqResponse = await fetchRfqData()
  const gasZipChains = await fetchGasZipData()
  await Promise.all(
    Object.keys(providers).map(async (chainId) => {
      // Get map from token to set of bridge token symbols
      const { originMap, poolSets } = await getBridgeOriginMap(chainId)
      // Add tokens from CCTP originMap to global originMap
      const cctpOriginMap = await getCCTPOriginMap(chainId)
      Object.keys(cctpOriginMap).forEach((token) => {
        addSetToMap(originMap, token, cctpOriginMap[token])
      })
      // Add tokens from RFQ originMap to global originMap
      const rfqOriginMap = await getFastBridgeOriginMap(chainId, rfqResponse)
      Object.keys(rfqOriginMap).forEach((token) => {
        addSetToMap(originMap, token, rfqOriginMap[token])
      })
      // Add tokens from Gas.zip originMap to global originMap
      const gasZipOriginMap = await getGasZipOriginMap(chainId, gasZipChains)
      Object.keys(gasZipOriginMap).forEach((token) => {
        addSetToMap(originMap, token, gasZipOriginMap[token])
      })
      const tokens = {}
      await Promise.all(
        Object.keys(originMap).map(async (token) => {
          const decimals = await getTokenDecimals(chainId, token)
          const symbol = await getTokenSymbol(chainId, token)
          const origin = Array.from(originMap[token])
            .map((t) => (t === 'RFQ.USDC.e' ? 'RFQ.USDC' : t))
            .sort()
          const destination = await getDestinationBridgeSymbols(chainId, token)
          const swappable = extractSwappable(poolSets, token)

          tokens[token] = {
            decimals,
            symbol,
            origin,
            destination,
            swappable,
          }
          // Check if token is supported as destination asset in RFQ
          if (
            rfqResponse.some(
              (quote) =>
                ethers.utils.getAddress(quote.dest_token_addr) === token &&
                quote.dest_chain_id === Number(chainId)
            )
          ) {
            tokens[token].destination.push(getRFQSymbol(tokens[token].symbol))
          }
          // Check if token is a native asset on a GasZip supported chain
          if (token === ETH && gasZipChains.includes(Number(chainId))) {
            tokens[token].destination.push('Gas.zip')
          }
          tokens[token].destination = tokens[token].destination.sort()
        })
      )
      bridgeMap[chainId] = sortMapByKeys(tokens)
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

const getTokenSymbol = async (chainId, token) => {
  // Check if token is ETH
  if (token === ETH) {
    const swapQuoter = await getSwapQuoter(chainId)
    if (!swapQuoter) {
      return 'ETH'
    }
    // Get WETH address from SwapQuoter
    const weth = await swapQuoter.weth()
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

const getRFQSymbol = (symbol) => {
  if (symbol === 'USDC.e') {
    return 'RFQ.USDC'
  } else if (symbol === 'WETH') {
    return 'RFQ.ETH'
  } else {
    return `RFQ.${symbol}`
  }
}

printMaps()
