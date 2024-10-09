import { Address } from 'viem'
import Link from 'next/link'
import { useTranslations } from 'next-intl'

import { Chain } from '@/utils/types'
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
  const t = useTranslations('Wallet')

  return (
    <div id="empty-portfolio-content" className="p-4">
      <p className="text-[#C2C2D6] mb-4">
        {t('No bridgeable assets found for {address} on {chainName}', {
          address: connectedAddress && shortened,
          chainName: connectedChain?.name,
        })}
      </p>
      <p className="text-[#C2C2D6] mb-4">
        {t("Don't see a chain or token you want to bridge?")}
      </p>
      <a className="text-[#C2C2D6]">
        {t('Let us know on')}
        <Link
          className="text-[#99E6FF] underline px-1"
          href={TWITTER_URL}
          target="_blank"
        >
          {t('Twitter')}
        </Link>
        {t('or')}
        <Link
          className="text-[#99E6FF] underline pl-1"
          href={DISCORD_URL}
          target="_blank"
        >
          {t('Discord')}
        </Link>
        .
      </a>
    </div>
  )
}
