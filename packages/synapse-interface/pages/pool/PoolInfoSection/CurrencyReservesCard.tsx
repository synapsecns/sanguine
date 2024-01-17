import numeral from 'numeral'
import { useSelector } from 'react-redux'

import InfoSectionCard from './InfoSectionCard'
import { RootState } from '@/store/store'

const TokenLabels = ({ tokens }) => {
  return (
    <div className="mt-2 space-y-2">
      {tokens
        ? tokens?.map((token) => (
            <div
              className="relative flex items-center justify-between h-10"
              key={token.token.symbol}
            >
              <div className="absolute inset-0 w-full h-full">
                <div
                  className="h-full rounded-sm bg-fuchsia-500 bg-opacity-30"
                  style={{ width: token.percent }}
                />
              </div>

              <div className="relative z-10 flex items-center gap-2 p-2">
                <img
                  alt={`Icon for ${token.token.symbol}`}
                  className="w-6 h-6 rounded-full"
                  src={token.token.icon.src}
                />
                <div>
                  {numeral(token.balanceStr).format('0,0.00a')}
                </div>
                <span className="text-sm mt-px">
                  {token.token.symbol}
                </span>
              </div>

              <div className="relative z-10 flex items-center h-full">
                {numeral(token.percent).format('0,0%')}
              </div>
            </div>
          ))
        : null}
    </div>
  )
}

const CurrencyReservesCard = () => {
  const { poolData } = useSelector((state: RootState) => state.poolData)

  return (
    <InfoSectionCard title="Currency Reserves">
      {poolData ? <TokenLabels tokens={poolData.tokens} /> : null}
    </InfoSectionCard>
  )
}

export default CurrencyReservesCard
