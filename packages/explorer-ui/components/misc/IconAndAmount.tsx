import { getCoinTextColor } from '@styles/coins'
import { formatAmount } from '@utils/formatAmount'
import { AssetImage } from '@components/misc/AssetImage'
import { addressToSymbol } from '@utils/addressToSymbol'
import { TOKEN_HASH_MAP, tokenAddressToToken } from 'synapse-constants'
import { addressToDecimals } from '@utils/addressToDecimals'
import { Tooltip as ReactTooltip } from 'react-tooltip'

export function IconAndAmount({
  formattedValue,
  tokenAddress,
  chainId,
  tokenSymbol,
  textSize = 'text-2xl',
  iconSize = 'w-4 h-4 rounded-full',
  styledCoin = false,
}) {
  const t = chainId && tokenAddress && tokenAddressToToken( chainId, tokenAddress )

  let styledCoinClass
  if (styledCoin === true) {
    styledCoinClass =
      'bg-gray-700 rounded-xl text-[10px] text-white items-center ml-2 pl-[5px] pr-[5px] pt-[2px] pb-[2px]'
  } else {
    styledCoinClass = t && `${getCoinTextColor(t)} ${textSize}`
  }
2
  let amount
  let showToken
  if (tokenSymbol) {
    const displaySymbol = addressToSymbol({ tokenAddress, chainId }) || tokenSymbol
    showToken = <div className={styledCoinClass}>{displaySymbol}</div>
    const dec = 10 ** addressToDecimals({ tokenAddress, chainId })
    // Need a cleaner way of doing this.
    if (formattedValue > 10000000) {
      amount = formattedValue / (dec / 10 ** 6)
    }
    else {
      amount = formattedValue
    }
  } else {
    const displaySymbol = addressToSymbol({ tokenAddress, chainId })
    showToken = displaySymbol ? (
      <div className={styledCoinClass}>{displaySymbol}</div>
    ) : (
      <span className={`${textSize} text-slate-400`}>--</span>
    )
    const dec = 10 ** addressToDecimals({ tokenAddress, chainId })
    amount = formattedValue / (dec / 10 ** 6)
  }

  return (
    <div className="flex items-center">
      <div className="flex flex-row items-center">
        <AssetImage
          tokenAddress={tokenAddress}
          tokenSymbol={tokenSymbol}
          chainId={chainId}
          className={`${iconSize} min-w-[1.5rem] min-h-[1.5rem] inline-block mr-[.5rem]`}
        />
        <div
          data-tooltip-content={amount}
          data-tooltip-id="amount"
          className='flex-1 text-white'
        >
          {formatAmount(amount)}
        </div>
      </div>
      {showToken}
      <ReactTooltip id="amount" className="z-50 rounded-xl" />
    </div>
  )
}
