import { Tooltip as ReactTooltip } from 'react-tooltip'
import { AssetImage } from '@components/misc/AssetImage'
import { formatAmount } from '@utils/formatAmount'
import { addressToDecimals } from '@utils/addressToDecimals'
import { formatBigIntToString } from '@utils/formatBigIntToString'

export const IconAndAmount = ({
  value,
  tokenAddress,
  chainId,
  tokenSymbol,
  iconSize = 'w-4 h-4 rounded-full',
  className = '',
}) => {
  let amount
  const decimals = addressToDecimals({ tokenAddress, chainId })
  const formattedValue = formatBigIntToString(value, decimals)
  // const displayValue = amount ? formatAmount(amount) : '< 0.001'
  const displayAmount = formatAmount(formattedValue)

  if (tokenSymbol) {
    const dec = 10 ** addressToDecimals({ tokenAddress, chainId })
    if (value > 10000000) {
      amount = value / (dec / 10 ** 6)
    } else {
      amount = value
    }
  } else {
    const dec = 10 ** addressToDecimals({ tokenAddress, chainId })
    amount = value / (dec / 10 ** 6)
  }

  return (
    <div className={`flex items-center ${className}`}>
      <div className="flex flex-row items-center text-white">
        <AssetImage
          tokenAddress={tokenAddress}
          // tokenSymbol={tokenSymbol}
          chainId={chainId}
          className={`${iconSize} min-w-[1rem] min-h-[1rem] inline rounded-full`}
        />
        <div
          data-tooltip-content={displayAmount}
          data-tooltip-id="amount"
          className="flex-1 pl-1 mr-1 text-white"
        >
          {displayAmount}
        </div>
      </div>
      <span className="text-white">{tokenSymbol}</span>
      <ReactTooltip id="amount" className="z-50 rounded-xl" />
    </div>
  )
}
