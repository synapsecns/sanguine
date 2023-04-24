import PoolsListCard from './PoolsListCard'

const PoolsOnChain = ({
  chainId,
  poolsArr,
  connectedChainId,
  address,
}: {
  chainId: number
  poolsArr: any
  connectedChainId: number
  address: string
}) => {
  console.log()
  if (poolsArr?.length == 0 && connectedChainId != chainId) {
    return <></>
  } else {
    return (
      <>
        {poolsArr &&
          !(poolsArr?.length == 0 && connectedChainId != chainId) &&
          poolsArr.map((pt) => {
            return (
              <PoolsListCard
                key={pt.poolName}
                poolName={pt.poolName}
                chainId={chainId}
                connectedChainId={connectedChainId}
                address={address}
              />
            )
          })}
      </>
    )
  }
}
export default PoolsOnChain
