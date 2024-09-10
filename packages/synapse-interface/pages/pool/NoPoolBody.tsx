import { useTranslations } from 'next-intl'

import Card from '@tw/Card'
import Grid from '@tw/Grid'
import { Token } from '@types'
import { CHAINS_BY_ID } from '@constants/chains'
import { getNetworkTextColor } from '@styles/chains'

const NoPoolBody = ({
  pool,
  poolChainId,
}: {
  pool: Token
  poolChainId: number
}) => {
  const t = useTranslations('Pools')

  return pool && poolChainId ? (
    <Grid cols={{ xs: 1 }} gap={2}>
      <Card
        title="Pool Info "
        className={`
          bg-bgBase
          my-8 transform transition-all duration-100 rounded-lg place-self-center
          min-w-4/5 sm:min-w-3/4 md:min-w-3/5 lg:min-w-1/2
        `}
        divider={false}
      >
        <div className="w-full pt-4 text-center text-gray-400">
          {t('Switch to')}{' '}
          <span className={`${getNetworkTextColor(pool.color)} font-medium`}>
            {CHAINS_BY_ID[poolChainId].name}
          </span>{' '}
          {t('to interact with the')} <u>{pool.name}</u> {t('pool')}.
        </div>
      </Card>
    </Grid>
  ) : (
    <Grid cols={{ xs: 1 }} gap={2}>
      <div className="w-full pt-4 text-center text-gray-400">
        {t('Invalid Pool ID')}
      </div>
    </Grid>
  )
}

export default NoPoolBody
