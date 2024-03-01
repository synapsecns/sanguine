import _ from 'lodash'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { useAccount, useNetwork } from 'wagmi'
import { useAppDispatch } from '@/store/hooks'
import { fetchPoolData, resetPoolData } from '@/slices/poolDataSlice'
import { resetPoolDeposit } from '@/slices/poolDepositSlice'
import { resetPoolWithdraw } from '@/slices/poolWithdrawSlice'
import { usePoolDataState } from '@/slices/pools/hooks'
import { DEFAULT_FROM_CHAIN } from '@/constants/swap'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import { POOL_BY_ROUTER_INDEX } from '@constants/tokens'
import PoolBody from './PoolBody'
import NoPoolBody from './NoPoolBody'
import LoadingDots from '@/components/ui/tailwind/LoadingDots'
import StandardPageContainer from '@layouts/StandardPageContainer'

export const getStaticPaths = async () => {
  const paths = Object.keys(POOL_BY_ROUTER_INDEX).map((key) => ({
    params: { poolId: key },
  }))

  return {
    paths,
    fallback: false, // false or "blocking"
  }
}

export const getStaticProps = async (context) => {
  return { props: {} }
}

const PoolPage = () => {
  const dispatch = useAppDispatch()
  const router = useRouter()
  const { poolId } = router.query
  const { address } = useAccount()
  const { chain } = useNetwork()
  const [connectedChainId, setConnectedChainId] = useState(0)
  const [isClient, setIsClient] = useState(false)

  const { pool } = usePoolDataState()

  useEffect(() => {
    setIsClient(true)
  }, [])

  useEffect(() => {
    const handleRouteChange = () => {
      dispatch(resetPoolData())
      dispatch(resetPoolDeposit())
      dispatch(resetPoolWithdraw())
    }

    router.events.on('routeChangeStart', handleRouteChange)

    return () => {
      router.events.off('routeChangeStart', handleRouteChange)
    }
  }, [dispatch, router.events])

  useEffect(() => {
    setConnectedChainId(chain?.id ?? DEFAULT_FROM_CHAIN)
  }, [chain])

  useEffect(() => {
    if (poolId && isClient) {
      dispatch(resetPoolData())
      dispatch(fetchPoolData({ poolName: String(poolId) }))
    }
  }, [poolId, address, isClient])

  return (
    <LandingPageWrapper>
      <StandardPageContainer
        connectedChainId={connectedChainId}
        address={address}
      >
        {!pool || !poolId ? (
          <div className="flex items-center justify-center">
            <LoadingDots />
          </div>
        ) : pool ? (
          <PoolBody address={address} connectedChainId={connectedChainId} />
        ) : (
          <NoPoolBody pool={pool} poolChainId={pool.chainId} />
        )}
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

export default PoolPage
