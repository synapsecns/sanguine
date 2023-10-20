import React, { memo } from 'react'

import PoolCard from './PoolCard'
import { LoaderIcon } from 'react-hot-toast'

const PoolCards = memo(
  ({ pools, address }: { pools: any; address: string }) => {
    const poolChainIds = pools ? Object.keys(pools) : []

    return (
      <>
        {pools ? (
          poolChainIds.map((chainId) => {
            return (
              <React.Fragment key={chainId}>
                {pools[chainId] &&
                  pools[chainId]?.length > 0 &&
                  pools[chainId].map((pool, i) => {
                    return (
                      <div key={i}>
                        <PoolCard
                          key={pool?.poolName}
                          pool={pool}
                          chainId={Number(chainId)}
                          address={address}
                        />
                      </div>
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
