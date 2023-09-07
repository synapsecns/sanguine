import Card from '@tw/Card'
import Grid from '@tw/Grid'
import { getNetworkTextColor } from '@styles/chains'
import { CHAINS_BY_ID } from '@constants/chains'
import { Token } from '@types'

const NoPoolBody = ({
  pool,
  poolChainId,
}: {
  pool: Token
  poolChainId: number
}) => {
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
          Switch to{' '}
          <span className={`${getNetworkTextColor(pool.color)} font-medium`}>
            {CHAINS_BY_ID[poolChainId].name}
          </span>{' '}
          to interact with the <u>{pool.name}</u> pool.
        </div>
      </Card>
    </Grid>
  ) : (
    <Grid cols={{ xs: 1 }} gap={2}>
      <div className="w-full pt-4 text-center text-gray-400">
        Invalid Pool ID
      </div>
    </Grid>
  )
}

export default NoPoolBody
