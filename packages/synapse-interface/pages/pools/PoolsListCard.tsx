// import { Link, useHistory } from 'react-router-dom'
import Link from 'next/link'
import { getPoolUrl } from '@urls'
import { POOL_INVERTED_ROUTER_INDEX } from '@constants/poolRouter'
import { fetchSigner, getNetwork, switchNetwork } from '@wagmi/core'

// import { POOLS_MAP } from '@hooks/pools/usePools'

// import { useGenericPoolData } from '@hooks/pools/useGenericPoolData'
// import { useChainSwitcher } from '@hooks/wallet/useChainSwitcher'
// import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'

import Card from '@tw/Card'
import Grid from '@tw/Grid'

import ApyTooltip from '@components/ApyTooltip'

import { getPoolStats } from './getPoolStats'

import { CHAIN_INFO_MAP } from '@constants/chains'
import { POOLS_BY_CHAIN } from '@constants/tokens'
const PoolsListCard = ({ poolName, chainId, connectedChainId }) => {
  const [poolData] = useGenericPoolData(chainId, poolName)

  const poolTokens = POOLS_BY_CHAIN[chainId][poolName]
  const poolRouterIndex = POOL_INVERTED_ROUTER_INDEX[chainId][poolName]

  const { apy, fullCompoundedApyStr, totalLockedUSDStr } =
    getPoolStats(poolData)

  const { chainName, chainImg } = CHAIN_INFO_MAP[chainId]

  return (
    <div>
      <Link
        onClick={() => {
          if (address === undefined) {
            return alert('Please connect your wallet')
          }
          if (chainId != connectedChainId) {
            const res = switchNetwork({ chainId: chainId })
              .then((res) => {
                return res
              })
              .catch(() => {
                return undefined
              })
            if (res === undefined) {
              console.log("can't switch chain, chainId: ", chainId)
              return
            }
            // history.push(getPoolUrl({ poolRouterIndex }))
          }
        }}
        href={getPoolUrl({ poolRouterIndex })}
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

const PoolsCardTitle = ({ chainName, poolName, chainImg }) => {
  let displayPoolName = poolName.replace(chainName, `<b>${chainName}</b>`)

  return (
    <div className="flex items-center">
      <img src={chainImg} className="w-6 h-6 mr-2 rounded-full" />
      <div dangerouslySetInnerHTML={{ __html: displayPoolName }} />
    </div>
  )
}

const CoinLabels = ({ coins }) => {
  return (
    <div className="flex mt-3">
      {coins.map((coin, i) => (
        <img alt="" className="w-5 mr-1 rounded-full" src={coin.icon} key={i} />
      ))}
    </div>
  )
}
export default PoolsListCard
