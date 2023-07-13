import { useMemo } from 'react'
import { formatBigIntToString } from '@/utils/bigint/format'

function removeLeadingZeros(inputValue: number): number {
  const numberString = inputValue.toString()
  const integerPart = parseInt(numberString)
  const decimalPart = parseFloat(
    numberString.substring(integerPart.toString().length)
  )

  return integerPart + decimalPart
}

const PriceImpactDisplay = ({ priceImpact }: { priceImpact: bigint }) => {
  let colorClassName: string
  let labelText: string
  let content: any

  const priceImpactValue: number = useMemo(() => {
    if (!priceImpact) return 0

    let formattedPriceImpact = Number(
      formatBigIntToString(priceImpact * 100n, 18, 18)
    )

    if (priceImpact > 0n && formattedPriceImpact === 0) {
      formattedPriceImpact = removeLeadingZeros(
        Number(formatBigIntToString(priceImpact * 100n, 18, 10))
      )
    }

    return formattedPriceImpact
  }, [priceImpact])

  const priceImpactDisplayValue: string = useMemo(() => {
    if (Math.abs(priceImpactValue) < 0.01) {
      return '<0.01'
    } else {
      return priceImpactValue.toFixed(2)
    }
  }, [priceImpactValue])

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
          {priceImpactDisplayValue}%
        </span>
      </div>
    )
  }
  return content
}
export default PriceImpactDisplay
