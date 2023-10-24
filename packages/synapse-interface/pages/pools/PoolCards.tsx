import React, { memo } from 'react'

import PoolCard from './PoolCard'
import { LoaderIcon } from 'react-hot-toast'
import { useAppSelector } from '@/store/hooks'

const PoolCards = memo(
  ({ pools, address }: { pools: any; address: string }) => {
    const poolChainIds = pools ? Object.keys(pools) : []
    const { synPrices, ethPrice, avaxPrice } = useAppSelector(
      (state) => state.priceData
    )

    if (!ethPrice) {
      return null
    }

    return (
      <>
        {pools ? (
          poolChainIds.map((chainId) => {
            return (
              <React.Fragment key={chainId}>
                {synPrices.synPrice &&
                  ethPrice &&
                  avaxPrice &&
                  pools[chainId] &&
                  pools[chainId]?.length > 0 &&
                  pools[chainId].map((pool) => {
                    return (
                      <PoolCard
                        key={pool?.poolName}
                        pool={pool}
                        address={address}
                        ethPrice={ethPrice}
                        avaxPrice={avaxPrice}
                        synPrices={synPrices}
                      />
                    )
                  })}
              </React.Fragment>
            )
          })
        ) : (
          <div className="flex justify-center">
            <LoaderIcon />
          </div>
        )}
      </>
    )
  }
)

export default PoolCards
