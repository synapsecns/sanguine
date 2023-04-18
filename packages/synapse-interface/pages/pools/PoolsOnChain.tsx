import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'

import PoolsListCard from './PoolsListCard'

export default function PoolsOnChain({ chainId, poolsArr }) {
  const { chainId: currentChainId } = useActiveWeb3React()

  if (poolsArr.length == 0 && currentChainId != chainId) {
    return <></>
  } else {
    return (
      <>
        {poolsArr.map((pt) => {
          return (
            <PoolsListCard
              key={pt.poolName}
              poolName={pt.poolName}
              chainId={chainId}
            />
          )
        })}
      </>
    )
  }
}
