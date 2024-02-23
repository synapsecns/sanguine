import Link from 'next/link'
import type { Address } from 'viem'
import type { Chain } from '@/utils/types'
import { shortenAddress } from '@/utils/shortenAddress'
import { DISCORD_URL, TWITTER_URL } from '@/constants/urls'

export const EmptyPortfolioContent = ({
  connectedAddress,
  connectedChain,
}: {
  connectedAddress: Address
  connectedChain: Chain
}) => {
  const shortened: string = shortenAddress(connectedAddress)
  return (
    <div id="empty-portfolio-content" className="p-4 text-sm">
      <p className="text-[#C2C2D6] mb-4">
        No bridgeable assets found {connectedAddress && `for ${shortened}`} on{' '}
        {connectedChain?.name}.
      </p>
      <p className="text-[#C2C2D6] mb-4">
        Don't see a chain or token you want to bridge?
      </p>
      <a className="text-[#C2C2D6]">
        Let us know on
        <Link
          className="text-[#99E6FF] underline px-1"
          href={TWITTER_URL}
          target="_blank"
        >
          Twitter
        </Link>
        or
        <Link
          className="text-[#99E6FF] underline pl-1"
          href={DISCORD_URL}
          target="_blank"
        >
          Discord
        </Link>
        .
      </a>
    </div>
  )
}
