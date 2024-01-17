import Link from 'next/link'
import Button from '@/components/ui/tailwind/Button'
import { DOCS_URL, BRIDGE_PATH } from '@/constants/urls'
import { SectionContainer } from '../../../components/landing/shared'

export default function HeroSection() {
  return (
    <SectionContainer dataTestId="landing-hero-section">
      <h1
        className={`
          text-center text-4xl
          font-medium
          md:text-[46px]
        `}
      >
        Secure cross-chain<br />communication
      </h1>
      <p className="hidden text-center opacity-80 sm:block">
        Synapse is the most widely used, extensible, secure cross-<br />
        chain communications network. Build truly cross-chain<br />
        applications using the Synapse Protocol.
      </p>
      <p className="text-center opacity-80 sm:hidden">
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
            border hover:opacity-70 active:opacity-40
            bg-zinc-100 border-zinc-400
            dark:bg-zinc-900 dark:border-zinc-500
            rounded-md
            text-center transition-all duration-75
          `}
        >
          <Button
            onClick={() => null}
          >
            Build on Synapse
          </Button>
        </Link>
        <Link
          href={BRIDGE_PATH}
          className={`
            h-12 border-[#AC8FFF] flex items-center border
            text-base px-4 py-3 hover:opacity-70 active:opacity-40
            rounded-md
            transition-all duration-75
          `}
          style={{
            background:
              'linear-gradient(310.65deg, rgba(255, 0, 255, 0.2) -17.9%, rgba(172, 143, 255, 0.2) 86.48%)',
          }}
        >
          <Button onClick={() => null}>
            Enter Bridge
          </Button>
        </Link>
      </div>
    </SectionContainer>
  )
}
