import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'

import Grid from '@tw/Grid'
import { Tab } from '@headlessui/react'

import PoolsOnChain from './PoolsOnChain'

import { PageHeader } from '@components/PageHeader'

export default function PoolTabs({ arr }) {
  const { chainId } = useActiveWeb3React()
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
          return (
            <Tab.Panel key={index}>
              <Grid cols={{ xs: 1, sm: 1, md: 2, lg: 3 }} gap={4}>
                <PoolsOnChain
                  chainId={chainId}
                  poolsArr={poolsByChain[chainId]}
                  textLabel={textLabel ?? label}
                />
                {_.entries(poolsByChain)
                  .filter(([otherChainId, poolsArr]) => otherChainId != chainId)
                  .map(([otherChainId, poolsArr], index) => (
                    <PoolsOnChain
                      key={index}
                      chainId={otherChainId}
                      poolsArr={poolsArr}
                      textLabel={textLabel ?? label}
                    />
                  ))}
              </Grid>
            </Tab.Panel>
          )
        })}
      </Tab.Panels>
    </Tab.Group>
  )
}
