import { QuestionMarkCircleIcon } from '@heroicons/react/outline'
import Image from 'next/image'

type PortfolioNetworkProps = {
  displayName: string
  chainIcon: string
  isUnsupportedChain: boolean
}

export const PortfolioNetwork = ({
  displayName,
  chainIcon,
  isUnsupportedChain,
}: PortfolioNetworkProps) => {
  return (
    <div
      id="portfolio-network"
      className="flex flex-row justify-between flex-1 py-4 cursor-pointer"
    >
      <div className="flex flex-row items-center px-4">
        {isUnsupportedChain ? (
          <QuestionMarkCircleIcon className="w-6 h-6 mr-2 text-white rounded-md" />
        ) : (
          <Image
            className="w-6 h-6 mr-2 rounded-md"
            alt={`${displayName} img`}
            src={chainIcon}
          />
        )}
        <div className="text-lg font-medium text-white">
          {isUnsupportedChain ? 'Unsupported Network' : displayName}
        </div>
      </div>
    </div>
  )
}
