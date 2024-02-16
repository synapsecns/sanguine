import React, { memo } from 'react'

import { LoaderIcon } from 'react-hot-toast'

import PoolCard from './PoolCard'


const PoolCards = memo(
  ({ pools, address }: { pools: any; address: string }) => {
    const poolChainIds = pools ? Object.keys(pools) : []

    return (
      <>
        {pools ? (
          poolChainIds.map((chainId) => {
            return (
              <React.Fragment key={chainId}>
                {pools[chainId]?.map((pool) =>
                  <PoolCard
                    key={pool?.poolName}
                    pool={pool}
                    address={address}
                  />
                )}
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
