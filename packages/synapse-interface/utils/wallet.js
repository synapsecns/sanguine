

import { CHAIN_PARAMS, ChainId } from '@constants/networks'
import { getCompleteUrl } from '@urls'


export const setupNetwork = async (chainId) => {
  const provider = window.ethereum

  if (provider) {
    try {
      console.log({chainId})
      let chainIdKey
      console.log([CHAIN_PARAMS[chainId], CHAIN_PARAMS[chainIdKey]])
      if (CHAIN_PARAMS[chainId]) {
        chainIdKey = chainId
      } else {
        chainIdKey = ChainId.BSC
      }

      await provider.request({
        method: 'wallet_addEthereumChain',
        params: [
          CHAIN_PARAMS[chainIdKey]
        ],
      })
      return true
    } catch (error) {
      console.error(error)
      return false
    }
  } else {
    console.error(
      "Can't setup the BSC network on metamask because window.ethereum is undefined"
    )
    return false
  }
}

/** Add target token to MetaMask */
export async function addTokenToWallet({ token, chainId, image }) {
  const provider = window.ethereum

  if (provider) {
    try {
      const address = token.addresses[chainId]
      // wasAdded is a boolean. Like any RPC method, an error may be thrown.
      const wasAdded = await provider.request({
        method: 'wallet_watchAsset',
        params: {
          type: 'ERC20', // Initially only supports ERC20, but eventually more!
          options: {
            address: address, // The address that the token is at.
            symbol: getMetamaskCompatibleSymbol(token.symbol), // A ticker symbol or shorthand, up to 5 chars.
            decimals: token.decimals[chainId], // The number of decimals in the token
            image: image ?? (token.icon ? getCompleteUrl(token.icon) : ''), // A string url of the token logo
          },
        },
      })

      if (wasAdded) {
        // console.log(`
        //   One thing I don't know why
        //   It doesn't even matter how hard you try
        //   Keep that in mind, ${token.symbol} designed this line
        //   To explain in due time
        // `)
      } else {
        // console.log(`You can't change the world without getting your hands dirty.`)
      }
    } catch (error) {
      console.log(error)
    }
  }
}

/**
 * the below is some wack fix for rpc issue
 * on char length im too lazy to fix correctly
 */
function getMetamaskCompatibleSymbol(rawSymbol) {
  let symbol
  if (rawSymbol?.length > 6) {
    symbol = rawSymbol.split('-').join('')
    if (symbol?.length > 6) {
      symbol = symbol.slice(0, 6)
    }
    return symbol
  } else {
    return rawSymbol
  }

}


