import { getCoinTextColor } from '@styles/coins'
import { formatAmount } from '@utils/formatAmount'
import { AssetImage } from '@components/misc/AssetImage'
import { addressToSymbol } from '@utils/addressToSymbol'
import { TOKEN_HASH_MAP } from '@constants/tokens/basic'

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
  const t = chainId && tokenAddress && TOKEN_HASH_MAP[chainId]?.[tokenAddress]

  let styledCoinClass
  if (styledCoin === true) {
    styledCoinClass =
      'bg-gray-700 rounded-xl text-[10px] pl-[5px] pr-[5px] pt-[2px] pb-[2px] text-white  items-center  h-fit ml-2'
  } else {
    styledCoinClass = t && `${getCoinTextColor(t)} ${textSize}`
  }

  let showToken
  if (tokenSymbol) {
    const displaySymbol = addressToSymbol({ tokenAddress, chainId })
    showToken = <div className={styledCoinClass}>{displaySymbol}</div>
  } else {
    showToken = <span className={`${textSize} text-slate-400`}>--</span>
  }
  return (
    <div className="flex justify-between items-center ">
      <div className="flex flex-row items-center ">
        <AssetImage
          tokenAddress={tokenAddress}
          chainId={chainId}
          className={`${iconSize} inline mr-1 rounded-lg hover:opacity-[0.8] transition-all ease-in-out`}
        />
        <div className={`${textSize} pl-1 whitespace-nowrap text-white`}>
          {formatAmount(formattedValue)}
        </div>
      </div>
      {showToken}
    </div>
  )
}
