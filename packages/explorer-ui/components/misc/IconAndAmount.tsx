import { Tooltip as ReactTooltip } from 'react-tooltip'
import { AssetImage } from '@components/misc/AssetImage'
import { formatAmount } from '@utils/formatAmount'
import { addressToDecimals } from '@utils/addressToDecimals'

export const IconAndAmount = ({
  formattedValue,
  tokenAddress,
  chainId,
  tokenSymbol,
  iconSize = 'w-4 h-4 rounded-full',
  className = '',
}) => {
  2
  let amount
  if (tokenSymbol) {
    const dec = 10 ** addressToDecimals({ tokenAddress, chainId })
    if (formattedValue > 10000000) {
      amount = formattedValue / (dec / 10 ** 6)
    } else {
      amount = formattedValue
    }
  } else {
    const dec = 10 ** addressToDecimals({ tokenAddress, chainId })
    amount = formattedValue / (dec / 10 ** 6)
  }

  const displayAmount = amount ? formatAmount(amount) : '< 0.001'

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
