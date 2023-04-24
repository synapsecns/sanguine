// import { Link, useHistory } from 'react-router-dom'
import Link from 'next/link'
import { getPoolUrl } from '@urls'
// import { POOL_INVERTED_ROUTER_INDEX } from '@constants/poolRouter'
import { switchNetwork } from '@wagmi/core'
import { useEffect, useState } from 'react'
// import { POOLS_MAP } from '@hooks/pools/usePools'

import { useGenericPoolData } from '@hooks/pools/useGenericPoolData'
// import { useChainSwitcher } from '@hooks/wallet/useChainSwitcher'
// import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'

import Card from '@tw/Card'
import Grid from '@tw/Grid'

import { getPoolStats } from './getPoolStats'
import { memo } from 'react'
import { CHAINS_BY_ID } from '@constants/chains'
import { STAKING_MAP_TOKENS } from '@constants/tokens'
import { useSynapseContext } from '@/utils/SynapseProvider'
const PoolsListCard = memo(
  ({
    poolName,
    chainId,
    address,
    connectedChainId,
  }: {
    poolName: string
    chainId: number
    address: string
    connectedChainId: number
  }) => {
    const [poolData, setPoolData] = useState(undefined)
    const SynapseSDK = useSynapseContext()

    console.log('PoolsListCard RERENDER')
    useEffect(() => {
      if (
        connectedChainId === undefined ||
        chainId === undefined ||
        address === undefined ||
        poolName === undefined
      ) {
        return
      }
      // TODO - separate the apy and tvl so they load async.
      useGenericPoolData(chainId, poolName, address, SynapseSDK)
        .then((res) => {
          setPoolData(res.poolDataObj)
        })
        .catch((err) => {
          console.log('ERROR useGenericPoolData: ', err)
        })
    }, [])
    // const [poolData] =
    const chain = CHAINS_BY_ID[chainId]
    const poolTokens = STAKING_MAP_TOKENS[chainId][poolName]
    // const poolRouterIndex = POOL_INVERTED_ROUTER_INDEX[chainId][poolName]
    const { apy, fullCompoundedApyStr, totalLockedUSDStr } =
      getPoolStats(poolData)

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
          // href={getPoolUrl({ poolRouterIndex })}
          href="{getPoolUrl({ poolRouterIndex })}" // TODO: fix this
        >
          <Card
            title={
              <PoolsCardTitle
                chainImg={chain?.chainImg?.src}
                poolName={poolName}
                chainName={chain?.chainName}
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
                {poolTokens?.poolTokens && (
                  <CoinLabels coins={poolTokens?.poolTokens} />
                )}
              </div>
              <div>
                <h3 className="text-sm text-opacity-50 text-secondaryTextColor">
                  TVL
                </h3>
                <div className={'mt-2 text-white '}>
                  {totalLockedUSDStr ? (
                    '$' + totalLockedUSDStr
                  ) : (
                    <div className="animate-pulse rounded bg-slate-700 h-6 w-12" />
                  )}
                </div>
              </div>
              <div>
                <h3 className="text-sm text-opacity-50 text-secondaryTextColor">
                  APY{' '}
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
)

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
      {coins
        ? coins?.map((coin, i) => (
            <img
              alt=""
              className="w-5 mr-1 rounded-full"
              src={coin.icon.src}
              key={i}
            />
          ))
        : null}
    </div>
  )
}
export default PoolsListCard
