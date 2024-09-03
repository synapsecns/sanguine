import { useAccount } from 'wagmi'
import { useRouter } from 'next/router'
import Link from 'next/link'
import { useTranslations } from 'next-intl'

import StandardPageContainer from '@layouts/StandardPageContainer'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import StakeCard from './StakeCard'
import { POOL_BY_ROUTER_INDEX } from '@/constants/tokens'
import { POOL_PATH } from '@/constants/urls'
import { ChevronLeftIcon } from '@heroicons/react/outline'

export async function getStaticPaths({ locales }) {
  const poolRouterIndices = Object.keys(POOL_BY_ROUTER_INDEX)

  const paths = poolRouterIndices.flatMap((routerIndex) =>
    locales.map((locale) => ({
      params: { routerIndex },
      locale,
    }))
  )

  return {
    paths,
    fallback: false,
  }
}

export async function getStaticProps({ params, locale }) {
  return {
    props: {
      messages: (await import(`../../messages/${locale}.json`)).default,
      routerIndex: params.routerIndex,
    },
  }
}

const SingleStakePage = () => {
  const router = useRouter()
  const { routerIndex } = router.query
  const { address } = useAccount()
  const { chain } = useAccount()

  const t = useTranslations('Pools')

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
                {t('Back to Pool')}
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
