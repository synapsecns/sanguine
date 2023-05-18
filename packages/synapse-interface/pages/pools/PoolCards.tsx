import Grid from '@tw/Grid'
import { Tab } from '@headlessui/react'
import _ from 'lodash'
import LoadingPoolCard from '@components/loading/LoadingPoolCard'

import PoolsListCard from './PoolCard'
import { PageHeader } from '@components/PageHeader'
import { memo, useEffect, useState } from 'react'
import {
  getSynPrices,
  getEthPrice,
  getAvaxPrice,
} from '@/utils/actions/getPrices'
import { PoolCardInfo } from '@types'
const PoolCards = memo(
  ({
    arr,
    connectedChainId,
    address,
  }: {
    arr: PoolCardInfo[]
    connectedChainId: number
    address: string
  }) => {
    const [synPrices, setSynPrices] = useState(undefined)
    const [ethPrice, setEthPrice] = useState(undefined)
    const [avaxPrice, setAvaxPrice] = useState(undefined)

    // Prices to reduce number of calls
    useEffect(() => {
      getSynPrices()
        .then((res) => {
          setSynPrices(res)
        })
        .catch((err) => console.log('Could not get syn prices', err))
      getEthPrice()
        .then((res) => {
          setEthPrice(res)
        })
        .catch((err) => console.log('Could not get eth prices', err))
      getAvaxPrice()
        .then((res) => {
          setAvaxPrice(res)
        })
        .catch((err) => console.log('Could not get avax prices', err))
    }, [])
    return (
      <Tab.Group>
        <div className="flex-wrap justify-between mb-8 px-36 md:flex">
          <PageHeader title="Pools" subtitle="Provide liquidity." />
          <Grid
            cols={{ xs: 1, sm: 1, md: 1, lg: 1 }}
            className="justify-center md:float-right place-items-center"
          >
            <Tab.List className="flex min-w-[360px] p-1.5 space-x-2">
              {synPrices &&
                ethPrice &&
                avaxPrice &&
                arr &&
                arr.map(({ label }, index) => {
                  return (
                    <Tab
                      key={index}
                      className={({ selected }) => {
                        return `
                        bg-bgLight
                          px-4 py-2 rounded-lg
                          text-sm text-white
                          transform-gpu transition-all duration-75
                          hover:bg-bgLighter
                          border-transparent
                          ${
                            selected
                              ? 'border-gradient-br-magenta-melrose-bgDarker border-solid border'
                              : 'bg:bg-bgLight'
                          }
                        `
                      }}
                    >
                      {label}
                    </Tab>
                  )
                })}
            </Tab.List>
          </Grid>
        </div>
        <Tab.Panels className="flex justify-center">
          {synPrices && ethPrice && avaxPrice && arr ? (
            arr.map(({ poolsByChain }, index) => {
              // DOUBLE CHECK HERE
              return (
                <Tab.Panel key={index}>
                  <Grid cols={{ xs: 1, sm: 1, md: 2, lg: 3 }} gap={4}>
                    {/* Render the pools for the selected chain first */}
                    {poolsByChain[connectedChainId] &&
                      poolsByChain[connectedChainId]?.length > 0 &&
                      poolsByChain[connectedChainId].map((pool) => {
                        return (
                          <PoolsListCard
                            key={pool?.poolName}
                            pool={pool}
                            chainId={connectedChainId}
                            connectedChainId={connectedChainId}
                            address={address}
                            prices={{ synPrices, ethPrice, avaxPrice }}
                          />
                        )
                      })}

                    {/* Render all the other pools */}
                    {synPrices &&
                      ethPrice &&
                      avaxPrice &&
                      _.entries(poolsByChain)
                        .filter(
                          ([otherChainId, poolsArr]) =>
                            Number(otherChainId) != connectedChainId
                        )
                        .map(
                          ([otherChainId, poolsArr], index) =>
                            poolsByChain[otherChainId] &&
                            poolsByChain[otherChainId]?.length > 0 &&
                            poolsByChain[otherChainId].map((pool) => {
                              return (
                                <PoolsListCard
                                  key={pool?.poolName}
                                  pool={pool}
                                  chainId={Number(otherChainId)}
                                  connectedChainId={connectedChainId}
                                  address={address}
                                  prices={{ synPrices, ethPrice, avaxPrice }}
                                />
                              )
                            })
                        )}
                  </Grid>
                </Tab.Panel>
              )
            })
          ) : (
            <Grid
              cols={{ xs: 1, sm: 1, md: 2, lg: 3 }}
              gap={4}
              className="w-[90%]"
            >
              <LoadingPoolCard />
              <LoadingPoolCard />
              <LoadingPoolCard />
            </Grid>
          )}
        </Tab.Panels>
      </Tab.Group>
    )
  }
)

export default PoolCards
