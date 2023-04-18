import PoolsListCard from './PoolsListCard'

const PoolsOnChain = ({
  chainId,
  poolsArr,
  connectedChainId,
}: {
  chainId: number
  poolsArr: any
  connectedChainId: number
}) => {
  if (poolsArr.length == 0 && connectedChainId != chainId) {
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
export default PoolsOnChain
