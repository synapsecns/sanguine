import { useMemo } from 'react'
import { formatBNToString } from '@bignumber/format'
import { BigNumber } from '@ethersproject/bignumber'
import { WeiPerEther } from '@ethersproject/constants'

function removeLeadingZeros(inputValue: number): number {
  const numberString = inputValue.toString()
  const integerPart = parseInt(numberString)
  const decimalPart = parseFloat(
    numberString.substring(integerPart.toString().length)
  )

  return integerPart + decimalPart
}

const PriceImpactDisplay = ({ priceImpact }: { priceImpact: BigNumber }) => {
  let colorClassName: string
  let labelText: string
  let content: string | React.ReactNode

  const priceImpactValue: number = useMemo(() => {
    let formattedPriceImpact = Number(
      formatBNToString(priceImpact.mul(100), 18, 2)
    )

    if (priceImpact.gt(0) && formattedPriceImpact === 0) {
      formattedPriceImpact = removeLeadingZeros(
        Number(formatBNToString(priceImpact.mul(100), 18, 5))
      )
    }

    return formattedPriceImpact
  }, [priceImpact])

  if (priceImpactValue > 0) {
    colorClassName = 'text-green-500'
    labelText = 'Bonus'
  } else {
    colorClassName = 'text-red-500'
    labelText = 'Price Impact'
  }

  if (priceImpactValue == 0) {
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
          {priceImpactValue}%
        </span>
      </div>
    )
  }
  return content
}
export default PriceImpactDisplay
