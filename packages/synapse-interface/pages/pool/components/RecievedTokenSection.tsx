import _ from 'lodash'

import { commify } from '@ethersproject/units'
import { cleanNumberInput } from '@utils/cleanNumberInput'
import { Token } from '@types'
const RecievedTokenSection = ({
  inputState,
  poolTokens,
  label,
}: {
  inputState: any
  poolTokens: Token[]
  label?: string
}) => {
  return (
    <div className="text-center sm:text-left">
      <p className="text-sm font-medium bg-opacity-70 pb-0.5 text-secondaryTextColor text-opacity-50 mb-2">
        {label ?? 'You will receive '}
      </p>
      {poolTokens
        .filter(
          (token) => Number(cleanNumberInput(inputState[token.symbol])) > 0
        )
        .map((token) => {
          return (
            <div className="flex items-center">
              <img
                alt=""
                className="inline-block -mt-0.5 w-4 mr-1"
                src={token.icon}
              ></img>
              <span className="text-sm text-white">
                {commify(_.round(inputState[token.symbol], 2))}
              </span>
              <span className="text-sm px-1 font-medium text-secondaryTextColor text-opacity-50">
                {token.symbol}
              </span>
            </div>
          )
        })}
    </div>
  )
}
export default RecievedTokenSection
