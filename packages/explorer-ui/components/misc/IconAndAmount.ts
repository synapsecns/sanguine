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
    styledCoinClass = 'bg-gray-700 rounded-xl text-[10px] pl-[5px] pr-[5px] pt-[2px] pb-[2px] text-white  items-center  h-fit ml-2'
  } else {
    styledCoinClass = t && `${getCoinTextColor(t)} ${textSize}`
  }

  let showToken
  if (tokenSymbol) {
    let displaySymbol = addressToSymbol({ tokenAddress, chainId })
    // @ts-expect-error TS(2304): Cannot find name 'div'.
    showToken = <div className={styledCoinClass}>{displaySymbol}</div>
  } else {
    // @ts-expect-error TS(2304): Cannot find name 'span'.
    showToken = <span className={`${textSize} text-slate-400`}>--</span>
  }
  return (
    // @ts-expect-error TS(2304): Cannot find name 'div'.
    <div className="flex justify-between items-center ">
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="flex flex-row items-center ">
        // @ts-expect-error TS(2749): 'AssetImage' refers to a value, but is being used ... Remove this comment to see the full error message
        <AssetImage
          tokenAddress={tokenAddress}
          chainId={chainId}
          // @ts-expect-error TS(2304): Cannot find name 'className'.
          className={`${iconSize} inline mr-1 rounded-lg hover:opacity-[0.8] transition-all ease-in-out`}
        />
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className={`${textSize} pl-1 whitespace-nowrap text-white`}>
          {formatAmount(formattedValue)}
        </div>
      </div>
      // @ts-expect-error TS(2304): Cannot find name 'showToken'.
      {showToken}
    </div>
  )
}
