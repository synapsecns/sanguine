import { formatBNToPercentString, formatBNToString } from '@bignumber/format'
import { BigNumber } from '@ethersproject/bignumber'

const PriceImpactDisplay = ({ priceImpact }: { priceImpact: BigNumber }) => {
  let colorClassName
  let labelText
  // TODO: Check this
  let priceImpactBP = Number(
    formatBNToString(priceImpact.mul(BigNumber.from(-100)), 18, 2)
  )

  console.log('priceImpactBP', priceImpactBP, priceImpact.toString())
  if (priceImpactBP > 0) {
    colorClassName = 'text-green-500'
    labelText = 'Bonus'
  } else {
    colorClassName = 'text-red-500'
    labelText = 'Price Impact'
  }

  let content
  if (priceImpactBP == 0) {
    content = ''
  } else {
    content = (
      <div className="text-center cursor-pointer sm:text-right">
        <p className=" text-sm font-medium opacity-70 pb-0.5 text-gray-300">
          {labelText}
        </p>
        <span
          className={`
          pl-2 text-md font-medium ml-auto
          ${colorClassName}
        `}
        >
          {priceImpactBP}%
        </span>
      </div>
    )
  }
  return content
}
export default PriceImpactDisplay
