import Link from 'next/link'
import { Address } from 'viem'
import { ANALYTICS_KAPPA, ANALYTICS_PATH } from '@/constants/urls'

export const getTransactionExplorerLink = ({
  kappa,
  fromChainId,
  toChainId,
}: {
  kappa: string
  fromChainId: number
  toChainId?: number
}): string => {
  if (typeof toChainId === 'number') {
    return `${ANALYTICS_KAPPA}${kappa}?chainIdFrom=${fromChainId}&chainIdTo=${toChainId}`
  } else {
    return `${ANALYTICS_KAPPA}${kappa}?chainIdFrom=${fromChainId}`
  }
}

export const UserExplorerLink = ({
  connectedAddress,
}: {
  connectedAddress?: Address | string
}) => {
  const explorerLink: string = connectedAddress
    ? `${ANALYTICS_PATH}address/${connectedAddress}`
    : ANALYTICS_PATH
  return (
    <div data-test-id="explorer-link" className="text-[#99E6FF] my-3">
      <Link href={explorerLink} target="_blank">
        <span className="hover:underline">Explorer</span> →
      </Link>
    </div>
  )
}
