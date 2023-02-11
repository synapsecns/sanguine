import {getCoinTextColor} from '@styles/coins'
import {formatAmount} from '@utils/formatAmount'
import {AssetImage} from '@components/misc/AssetImage'
import {addressToSymbol} from '@utils/addressToSymbol'
import {TOKEN_HASH_MAP} from '@constants/tokens/basic'

export function IconAndAmount({
  formattedValue,
  tokenAddress,
  chainId,
  tokenSymbol,
  textSize = 'text-2xl',
  iconSize = 'w-6 h-6',
  styledCoin = false,
}) {
  tokenAddress = tokenAddress && tokenAddress.toLowerCase()
  const t = chainId && tokenAddress && TOKEN_HASH_MAP[chainId][tokenAddress]

  let styledCoinClass
  if (styledCoin === true) {
    styledCoinClass = 'bg-gray-700 rounded-xl text-[10px] pl-[5px] pr-[5px]'
  } else {
    styledCoinClass = t && `${getCoinTextColor(t)} ${textSize}`
  }

  let showToken
  if (tokenSymbol) {
    let displaySymbol = addressToSymbol({ tokenAddress, chainId })
    showToken = <div className={styledCoinClass}>{displaySymbol}</div>
  } else {
    showToken = <span className={`${textSize} text-slate-400`}>--</span>
  }

  return (
    <div className="flex items-center">
      <AssetImage
        tokenAddress={tokenAddress}
        chainId={chainId}
        className={`${iconSize} mr-1 text-slate-400 self-center`}
      />
      <div className={`${textSize} ml-1 mr-2 text-white self-center`}>
        {formatAmount(formattedValue)}
      </div>
      {showToken}
    </div>
  )
}
