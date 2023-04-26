import { commifyBnToString } from '@bignumber/format'
import AugmentWithUnits from '../components/AugmentWithUnits'
import InfoSectionCard from './InfoSectionCard'
import { displaySymbol } from '@utils/displaySymbol'

export default function CurrencyReservesCard({
  chainId,
  title,
  tokens,
  poolToken,
}) {
  return (
    <InfoSectionCard title={title}>
      {tokens.map((token, idx) => {
        return (
          <div key={idx}>
            <CurrencyInfoListItem
              chainId={chainId}
              key={token.symbol}
              {...token}
              token={poolToken.poolTokens[idx]}
            />
          </div>
        )
      })}
    </InfoSectionCard>
  )
}

function CurrencyInfoListItem({ chainId, percent, value, token }) {
  const { icon, swapableType } = token
  const symbol = displaySymbol(chainId, token)

  let decimalsToDisplay
  if (swapableType == 'USD') {
    decimalsToDisplay = 0
  } else {
    decimalsToDisplay = 2
  }

  return (
    <div className="flex items-center justify-between my-2 text-sm font-medium text-white">
      <div className="flex items-center">
        <img className="relative mr-2 w-7" src={icon} />
        <div>{symbol}</div>
      </div>
      <div>{percent}</div>

      <AugmentWithUnits
        content={commifyBnToString(value, decimalsToDisplay)}
        label={symbol}
      />
    </div>
  )
}
