import Link from 'next/link'
import Button from '@tw/Button'
import { DOCS_URL, BRIDGE_PATH } from '@/constants/urls'
import { SectionContainer } from '../../../components/landing/shared'

export default function HeroSection() {
  return (
    <SectionContainer dataTestId="landing-hero-section">
      <h1
        className={`
          text-center text-4xl
          font-medium text-white
          md:text-[46px]
        `}
      >
        Secure cross-chain<br />communication
      </h1>
      <p className="hidden text-center text-secondaryTextColor sm:block">
        Synapse is the most widely used, extensible, secure cross-<br />
        chain communications network. Build truly cross-chain<br />
        applications using the Synapse Protocol.
      </p>
      <p className="text-center text-secondaryTextColor sm:hidden">
        Synapse is the most widely used, extensible, secure cross-chain
        communications network. Build truly cross-chain applications using the
        Synapse Protocol.
      </p>
      <div className="flex justify-center py-4 space-x-2 ">
        <Link
          href={DOCS_URL}
          target="_blank"
          className={`
            h-12 mr-2 px-4 py-3 flex items-center
            text-base border border-white hover:opacity-75
            bg-[#2f2f2f] hover:bg-[#2f2f2f] rounded-md
            text-center transform-gpu transition-all duration-75
          `}
        >
          <Button
            className="flex items-center justify-center font-medium"
            onClick={() => null}
          >
            Build on Synapse
          </Button>
        </Link>
        <Link
          href={BRIDGE_PATH}
          className={`
            h-12 border-[#AC8FFF] flex items-center border
            text-base px-4 py-3 hover:opacity-75 rounded-md
            text-center transform-gpu transition-all duration-75
          `}
          style={{
            background:
              'linear-gradient(310.65deg, rgba(255, 0, 255, 0.2) -17.9%, rgba(172, 143, 255, 0.2) 86.48%)',
            borderRadius: '10px',
          }}
        >
          <Button className="font-medium" onClick={() => null}>
            Enter Bridge
          </Button>
        </Link>
      </div>
    </SectionContainer>
  )
}
