import { commifyBnToString } from '@bignumber/format'
import AugmentWithUnits from '../components/AugmentWithUnits'
import InfoSectionCard from './InfoSectionCard'
import { displaySymbol } from '@utils/displaySymbol'
import { PoolTokenObject, Token } from '@types'
import LoadingRow from '@/components/loading/LoadingRow'

const CurrencyReservesCard = ({
  chainId,
  title,
  poolData,
}: {
  chainId: number
  title: string
  poolData: any
}) => {
  return (
    <InfoSectionCard title={title}>
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
  let decimalsToDisplay = token.swapableType === 'USD' ? 0 : 2
  return (
    <div className="flex items-center justify-between my-2 text-sm font-medium text-white">
      <div className="flex items-center">
        <img className="relative mr-2 w-7" src={token.icon.src} />
        <div>{symbol}</div>
      </div>
      <div>{percent}</div>
      {balance && (
        <AugmentWithUnits
          content={commifyBnToString(balance, decimalsToDisplay)}
          label={symbol}
        />
      )}
    </div>
  )
}

export default CurrencyReservesCard
