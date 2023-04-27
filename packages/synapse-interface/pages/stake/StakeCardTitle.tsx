import { Token } from '@/utils/types'

interface StakeCardTitleProps {
  token: Token
  poolTokens: Token[]
  poolLabel: string
}

const StakeCardTitle = ({
  token,
  poolTokens,
  poolLabel,
}: StakeCardTitleProps) => {
  return (
    <div className="px-2 mb-5">
      <div className="inline-flex items-center mt-2">
        {poolTokens && (
          <div className="items-center hidden mr-4 md:flex lg:flex">
            {poolTokens.map((coin) => (
              <img
                key={coin.symbol}
                className="relative inline-block w-6 -mr-1 text-white shadow-solid"
                src={coin.icon}
              />
            ))}
          </div>
        )}
        <h3 className="mr-2 text-xl font-medium text-white">{poolLabel}</h3>
      </div>

      <div className="text-lg font-normal text-white text-opacity-70">
        <span className="text-green-400">
          {/* {fullyCompoundedApyLabel}%  */}- %
        </span>
        APY
        {/* <ApyTooltip
          apyData={apyData}
          baseApyData={baseApyData}
          className="ml-1"
        /> */}
      </div>
    </div>
  )
}

export default StakeCardTitle
