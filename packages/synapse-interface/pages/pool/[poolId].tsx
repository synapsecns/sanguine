import _ from 'lodash'
import { useEffect, useState, useMemo } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import { useRouter } from 'next/router'
import StandardPageContainer from '@layouts/StandardPageContainer'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import { DEFAULT_FROM_CHAIN } from '@/constants/swap'
import { POOL_BY_ROUTER_INDEX, POOL_CHAINS_BY_NAME } from '@constants/tokens'
import PoolBody from './PoolBody'
import NoPoolBody from './NoPoolBody'

const PoolPage = () => {
  const router = useRouter()
  const { poolId } = router.query
  const { address: currentAddress } = useAccount()
  const { chain } = useNetwork()
  const [connectedChainId, setConnectedChainId] = useState(0)
  const [address, setAddress] = useState(undefined)
  const [pool, setPool] = useState(undefined)

  useEffect(() => {
    setConnectedChainId(chain?.id ?? DEFAULT_FROM_CHAIN)
  }, [chain])

  useEffect(() => {
    setAddress(currentAddress)
  }, [currentAddress])

  useEffect(() => {
    const poolFromName = POOL_BY_ROUTER_INDEX[String(poolId)]
    setPool(poolFromName)
  }, [poolId])

  const poolChainId = useMemo(
    () => (pool?.addresses ? Number(Object.keys(pool?.addresses)[0]) : 0),
    [pool]
  )
  return (
    <LandingPageWrapper>
      <StandardPageContainer
        connectedChainId={connectedChainId}
        address={address}
      >
        {pool && connectedChainId === poolChainId ? (
          <PoolBody
            pool={pool}
            address={address}
            poolChainId={poolChainId}
            connectedChainId={connectedChainId}
          />
        ) : (
          <NoPoolBody pool={pool} poolChainId={poolChainId} />
        )}
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

export default PoolPage
