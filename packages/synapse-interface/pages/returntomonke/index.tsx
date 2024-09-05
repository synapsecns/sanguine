import { useEffect, useState } from 'react'
import { useAccount } from 'wagmi'
import { useTranslations } from 'next-intl'
import deepmerge from 'deepmerge'

import Grid from '@tw/Grid'
import { DEFAULT_FROM_CHAIN } from '@/constants/swap'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import StandardPageContainer from '@layouts/StandardPageContainer'

import PfpGeneratorCard from './PfpGeneratorCard'

export async function getStaticProps({ locale }) {
  const userMessages = (await import(`../../messages/${locale}.json`)).default
  const defaultMessages = (await import(`../../messages/en-US.json`)).default
  const messages = deepmerge(defaultMessages, userMessages)

  return {
    props: {
      messages,
    },
  }
}

const ReturnToMonkePage = () => {
  const { address: currentAddress, chain } = useAccount()
  const [connectedChainId, setConnectedChainId] = useState(0)
  const [address, setAddress] = useState(undefined)

  const t = useTranslations('ReturnToMonke')

  useEffect(() => {
    setConnectedChainId(chain?.id ?? DEFAULT_FROM_CHAIN)
  }, [chain])

  useEffect(() => {
    setAddress(currentAddress)
  }, [currentAddress])
  return (
    <LandingPageWrapper>
      <StandardPageContainer
        connectedChainId={connectedChainId}
        address={address}
      >
        <div className="flex justify-between">
          <div>
            <div className="text-2xl text-white">
              {t('Generate Synaptic Profile Picture')}
            </div>
          </div>
        </div>
        <div className="py-6">
          <Grid
            cols={{ xs: 1 }}
            gap={6}
            className="justify-center px-2 py-16 sm:px-6 md:px-8"
          >
            <div className="pb-3 place-self-center">
              <PfpGeneratorCard />
            </div>
          </Grid>
        </div>
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

export default ReturnToMonkePage
