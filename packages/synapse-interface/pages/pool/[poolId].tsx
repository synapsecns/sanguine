import _ from 'lodash'
import { useEffect, useState } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import { useDispatch, useSelector } from 'react-redux'
import { useRouter } from 'next/router'
import StandardPageContainer from '@layouts/StandardPageContainer'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import { DEFAULT_FROM_CHAIN } from '@/constants/swap'
import PoolBody from './PoolBody'
import NoPoolBody from './NoPoolBody'
import { fetchPoolData, resetPoolData } from '@/slices/poolDataSlice'
import { RootState } from '@/store/store'
import LoadingSpinner from '@/components/ui/tailwind/LoadingSpinner'

const PoolPage = () => {
  const router = useRouter()
  const { poolId } = router.query
  const { address } = useAccount()
  const { chain } = useNetwork()
  const [connectedChainId, setConnectedChainId] = useState(0)

  const { pool, isLoading } = useSelector((state: RootState) => state.poolData)

  const dispatch: any = useDispatch()

  useEffect(() => {
    setConnectedChainId(chain?.id ?? DEFAULT_FROM_CHAIN)
  }, [chain])

  useEffect(() => {
    if (poolId) {
      dispatch(resetPoolData())
      dispatch(fetchPoolData({ poolName: String(poolId) }))
    }
  }, [poolId, address])

  return (
    <LandingPageWrapper>
      <StandardPageContainer
        connectedChainId={connectedChainId}
        address={address}
      >
        {!pool || isLoading ? (
          <div className="flex items-center justify-center">
            <LoadingSpinner />
          </div>
        ) : pool && connectedChainId === pool.chainId ? (
          <PoolBody address={address} connectedChainId={connectedChainId} />
        ) : (
          <NoPoolBody pool={pool} poolChainId={pool.chainId} />
        )}
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

export default PoolPage
