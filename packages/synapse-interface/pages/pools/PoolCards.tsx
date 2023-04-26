import Grid from '@tw/Grid'
import { Tab } from '@headlessui/react'
import _ from 'lodash'
import PoolsListCard from './PoolCard'

import { PageHeader } from '@components/PageHeader'
import { memo } from 'react'
const PoolCards = memo(
  ({
    arr,
    connectedChainId,
    address,
  }: {
    arr: any
    connectedChainId: number
    address: string
  }) => {
    console.log('PoolTabs RERENDER')
    return (
      <Tab.Group>
        <div className="flex-wrap justify-between mb-8 px-36 md:flex">
          <PageHeader title="Pools" subtitle="Provide liquidity." />
          <Grid
            cols={{ xs: 1, sm: 1, md: 1, lg: 1 }}
            className="justify-center mt-3 mb-5 md:float-right place-items-center"
          >
            <Tab.List className="flex min-w-[360px] p-1.5 space-x-2">
              {arr.map(({ label, val }, index) => {
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
          {arr.map(({ poolsByChain, textLabel, label }, index) => {
            // DOUBLE CHECK HERE
            return (
              <Tab.Panel key={index}>
                <Grid cols={{ xs: 1, sm: 1, md: 2, lg: 3 }} gap={4}>
                  {/* Render the pools for the selected chain first */}
                  {poolsByChain[connectedChainId] &&
                    poolsByChain[connectedChainId]?.length > 0 &&
                    poolsByChain[connectedChainId].map((pt) => {
                      return (
                        <PoolsListCard
                          key={pt.poolName}
                          poolName={pt.poolName}
                          chainId={connectedChainId}
                          connectedChainId={connectedChainId}
                          address={address}
                        />
                      )
                    })}

                  {/* Render all the other pools */}
                  {_.entries(poolsByChain)
                    .filter(
                      ([otherChainId, poolsArr]) =>
                        Number(otherChainId) != connectedChainId
                    )
                    .map(
                      ([otherChainId, poolsArr], index) =>
                        poolsByChain[otherChainId] &&
                        poolsByChain[otherChainId]?.length > 0 &&
                        poolsByChain[otherChainId].map((pt) => {
                          return (
                            <PoolsListCard
                              key={pt.poolName}
                              poolName={pt.poolName}
                              chainId={Number(otherChainId)}
                              connectedChainId={connectedChainId}
                              address={address}
                            />
                          )
                        })
                    )}
                </Grid>
              </Tab.Panel>
            )
          })}
        </Tab.Panels>
      </Tab.Group>
    )
  }
)

export default PoolCards
