import AugmentWithUnits from '../components/AugmentWithUnits'
import InfoSectionCard from './InfoSectionCard'
import { displaySymbol } from '@utils/displaySymbol'
import LoadingRow from '@/components/loading/LoadingRow'
import { commify, formatBigIntToString } from '@utils/bigint/format'
import { stringToBigInt } from '@/utils/stringToBigNum'
import { useSelector } from 'react-redux'
import { RootState } from '@/store/store'

const CurrencyReservesCard = ({ chainId }: { chainId: number }) => {
  const { poolData } = useSelector((state: RootState) => state.poolData)

  return (
    <InfoSectionCard title="Currency Reserves">
      {poolData ? (
        poolData.tokens.map((tokenObj, idx) => {
          return (
            <div key={idx}>
              <CurrencyInfoListItem
                chainId={chainId}
                key={tokenObj.symbol}
                balance={tokenObj.balance}
                token={tokenObj.token}
                percent={tokenObj.percent}
              />
            </div>
          )
        })
      ) : (
        <>
          <LoadingRow />
          <LoadingRow />
        </>
      )}
    </InfoSectionCard>
  )
}

function CurrencyInfoListItem({ chainId, percent, balance, token }) {
  const symbol = displaySymbol(chainId, token)

  return (
    <div className="flex items-center justify-between my-2 text-sm font-medium text-white">
      <div className="flex items-center w-30">
        <img className="relative mr-2 w-7" src={token.icon.src} />
        <div>{symbol}</div>
      </div>
      <div className="text-right">{percent}</div>
      {balance && (
        <AugmentWithUnits
          content={commify(
            formatBigIntToString(
              stringToBigInt(`${balance}`, token.decimals[chainId]),
              token.decimals[chainId],
              -1
            )
          )}
          label={symbol}
        />
      )}
    </div>
  )
}

export default CurrencyReservesCard
