import 'wagmi/window'
import METAMASK_ICON from '@assets/icons/metamask.svg'
import { getCompleteUrl } from '@urls'

import Button from '@tw/Button'
import { getSwapBorderHoverStyleForCoin } from '@styles/tokens'

export default function AddToWalletButton({ token, icon, className, chainId }) {
  return (
    <Button
      onClick={() => addTokenToWallet({ token, icon, chainId })}
      className={`
        flex items-center group
        px-2 py-1
        border border-[#88818C] hover:border-opacity-70 active:border-opacity-70
        !rounded-full focus:ring-0 active:ring-0 outline-none
        transform-gpu transition duration-500 ease-in-out
        ${className}
      `}
    >
      <small
        className={`
          hidden group-hover:inline-block transition duration-500 ease-in-out mr-1
        `}
      >
        Add to Wallet
      </small>
      <img alt="metamask icon" src={METAMASK_ICON} className="inline w-5 h-5" />
    </Button>
  )
}

export function AddToWalletMiniButton({
  token,
  icon,
  className,
  iconFirst = false,
  chainId,
}) {
  return (
    <Button
      onClick={() => addTokenToWallet({ token, icon, chainId })}
      className={`
        px-2 !pt-0 !pb-0.5 group
        !border-transparent hover:border ${getSwapBorderHoverStyleForCoin(
          token
        )}
        !rounded-full focus:ring-0 active:ring-0 outline-none
        transform-gpu transition duration-500 ease-in-out ${className}
      `}
    >
      {iconFirst && (
        <img
          alt="metamask icon"
          src={METAMASK_ICON}
          className="inline w-5 h-5 mr-1 transition-all duration-200 ease-in-out opacity-50 group-hover:opacity-95"
        />
      )}
      <small
        className={`
          hidden group-hover:inline-block transition duration-200 ease-in-out mr-1 -mt-0.5
        `}
      >
        Add to Wallet
      </small>
      {!iconFirst && (
        <img
          alt="metamask icon"
          src={METAMASK_ICON?.src}
          className="inline w-5 h-5 transition-all duration-200 ease-in-out opacity-50 group-hover:opacity-95"
        />
      )}
    </Button>
  )
}



export const addTokenToWallet = async ({ token, chainId, icon }) => {
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
            image: icon ?? (token.icon ? getCompleteUrl(token.icon) : ''), // A string url of the token logo
          },
        },
      })
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
