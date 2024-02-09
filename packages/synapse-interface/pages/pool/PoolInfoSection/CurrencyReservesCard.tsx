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
              <div className="absolute inset-0 w-full h-full bg-bgDark">
                <div
                  className="h-full rounded-sm bg-bgBase/20 ring-1 ring-white/20"
                  style={{ width: token.percent }}
                ></div>
              </div>

              <div className="relative z-10 flex items-center h-full ml-2 space-x-2">
                <img
                  alt={`Icon for ${token.token.symbol}`}
                  className="w-[24px] h-[24px] rounded-full"
                  src={token.token.icon.src}
                />
                <div className="text-white text-md">
                  {numeral(token.balanceStr).format('0,0.00a')}
                </div>
                <div className="text-sm text-secondaryTextColor">
                  {token.token.symbol}
                </div>
              </div>

              <div className="relative z-10 flex items-center h-full text-white text-md">
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
