import { useAccount, useNetwork } from 'wagmi'
import { useRouter } from 'next/router'
import StandardPageContainer from '@layouts/StandardPageContainer'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import StakeCard from './StakeCard'
import { POOL_BY_ROUTER_INDEX } from '@/constants/tokens'
import Link from 'next/link'
import { POOL_PATH } from '@/constants/urls'
import { ChevronLeftIcon } from '@heroicons/react/outline'

const SingleStakePage = () => {
  const router = useRouter()
  const { routerIndex } = router.query
  const { address } = useAccount()
  const { chain } = useNetwork()

  const pool = POOL_BY_ROUTER_INDEX[routerIndex as string]

  if (!pool) return null

  return (
    <LandingPageWrapper>
      <StandardPageContainer
        connectedChainId={chain ? chain.id : 0}
        address={address}
      >
        <div className="justify-center py-8 mx-auto lg:flex lg:px-8">
          <div>
            <Link href={`${POOL_PATH}/${pool.routerIndex}`}>
              <div className="inline-flex items-center mb-3 text-sm font-light text-white hover:text-opacity-100">
                <ChevronLeftIcon className="w-4 h-4" />
                Back to Pool
              </div>
            </Link>
            <StakeCard address={address} chainId={pool.chainId} pool={pool} />
          </div>
        </div>
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

export default SingleStakePage
