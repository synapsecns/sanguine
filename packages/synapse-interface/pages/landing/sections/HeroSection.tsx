import Link from 'next/link'
import Button from '@/components/ui/tailwind/Button'
import { DOCS_URL, BRIDGE_PATH } from '@/constants/urls'
import { SectionContainer } from '../../../components/landing/shared'
import { useTranslation } from 'react-i18next'

export default function HeroSection() {
  const { t } = useTranslation()
  return (
    <SectionContainer dataTestId="landing-hero-section">
      <div
        className={`
          mt-6 pb-4 text-left text-4xl
          font-medium text-white
          md:text-[46px] md:text-center
        max-w-[450px]
        mx-auto
        `}
      >
        <p>{t('secure-cross-chain-communication')}</p>
      </div>

      <div
        className={`
         text-secondaryTextColor
          max-w-[350px] flex items-center justify-center
          mx-auto
        `}
      >
        <p>{t('synapse-is-the-most')}</p>
      </div>

      <div className="flex justify-center py-4 space-x-2 ">
        <Link
          href={DOCS_URL}
          target="_blank"
          className={`
            h-12 mr-2 px-4 py-3 flex items-center
            text-base border border-white hover:opacity-75
            bg-[#2f2f2f] hover:bg-[#2f2f2f] rounded-lg
            text-center transform-gpu transition-all duration-75
          `}
        >
          <Button
            className="flex items-center justify-center font-medium"
            onClick={() => null}
          >
            {t('build-on-synapse')}
          </Button>
        </Link>
        <Link
          href={BRIDGE_PATH}
          className={`
            h-12 border-[#AC8FFF] flex items-center border
            text-base px-4 py-3 hover:opacity-75 rounded-lg
            text-center transform-gpu transition-all duration-75
          `}
          style={{
            background:
              'linear-gradient(310.65deg, rgba(255, 0, 255, 0.2) -17.9%, rgba(172, 143, 255, 0.2) 86.48%)',
            borderRadius: '10px',
          }}
        >
          <Button className="font-medium" onClick={() => null}>
            {t('Enter Bridge')}
          </Button>
        </Link>
      </div>
    </SectionContainer>
  )
}
