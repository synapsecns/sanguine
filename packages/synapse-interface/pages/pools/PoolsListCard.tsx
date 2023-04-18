import { Link, useHistory } from 'react-router-dom'

import { getPoolUrl } from '@urls'
import { POOL_INVERTED_ROUTER_INDEX } from '@constants/poolRouter'

import { POOLS_MAP } from '@hooks/pools/usePools'

import { useGenericPoolData } from '@hooks/pools/useGenericPoolData'
import { useChainSwitcher } from '@hooks/wallet/useChainSwitcher'
import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'

import Card from '@tw/Card'
import Grid from '@tw/Grid'

import ApyTooltip from '@components/ApyTooltip'

import { getPoolStats } from './getPoolStats'

import { CHAIN_INFO_MAP } from '@constants/networks'

export default function PoolsListCard({ poolName, chainId }) {
  const history = useHistory()
  const { chainId: activeChainId } = useActiveWeb3React()
  const [poolData] = useGenericPoolData(chainId, poolName)

  const triggerChainSwitch = useChainSwitcher()

  const poolTokens = POOLS_MAP[chainId][poolName]
  const poolRouterIndex = POOL_INVERTED_ROUTER_INDEX[chainId][poolName]

  const { apy, fullCompoundedApyStr, totalLockedUSDStr } =
    getPoolStats(poolData)

  const { chainName, chainImg } = CHAIN_INFO_MAP[chainId]

  return (
    <div>
      <Link
        onClick={() => {
          if (chainId != activeChainId) {
            triggerChainSwitch(chainId)
            history.push(getPoolUrl({ poolRouterIndex }))
          }
        }}
        to={getPoolUrl({ poolRouterIndex })}
      >
        <Card
          title={
            <PoolsCardTitle
              chainImg={chainImg}
              poolName={poolName}
              chainName={chainName}
            />
          }
          titleClassName="text-white font-light text-xl"
          className={`
            bg-bgBase transition-all rounded-xl items-center
            hover:bg-bgLight
            py-6 mt-4 pr-2
            border border-transparent
          `}
          divider={false}
        >
          <Grid gap={3} cols={{ xs: 3 }} className="mt-8">
            <div>
              <h3 className="text-sm text-opacity-50 text-secondaryTextColor">
                Assets
              </h3>
              <CoinLabels coins={poolTokens} />
            </div>
            <div>
              <h3 className="text-sm text-opacity-50 text-secondaryTextColor">
                TVL
              </h3>
              <div className="mt-2 text-white">
                ${totalLockedUSDStr ?? <i className="opacity-50"> - </i>}
              </div>
            </div>
            <div>
              <h3 className="text-sm text-opacity-50 text-secondaryTextColor">
                APY{' '}
                {/* {fullCompoundedApyStr && (
                <ApyTooltip className="-m-8" apyData={apy} />
              )} */}
              </h3>
              <div className="mt-2 text-green-400">
                <>
                  {fullCompoundedApyStr ?? <i className="opacity-50"> - </i>}%
                </>
              </div>
            </div>
          </Grid>
        </Card>
      </Link>
    </div>
  )
}

function PoolsCardTitle({ chainName, poolName, chainImg }) {
  let displayPoolName = poolName.replace(chainName, `<b>${chainName}</b>`)

  return (
    <div className="flex items-center">
      <img src={chainImg} className="w-6 h-6 mr-2 rounded-full" />
      <div dangerouslySetInnerHTML={{ __html: displayPoolName }} />
    </div>
  )
}

function CoinLabels({ coins }) {
  return (
    <div className="flex mt-3">
      {coins.map((coin, i) => (
        <img alt="" className="w-5 mr-1 rounded-full" src={coin.icon} key={i} />
      ))}
    </div>
  )
}
