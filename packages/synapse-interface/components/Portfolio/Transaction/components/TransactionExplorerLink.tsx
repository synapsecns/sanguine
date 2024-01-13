import Link from 'next/link'
import { Address } from 'viem'
import { EXPLORER_KAPPA, EXPLORER_PATH } from '@/constants/urls'

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
    return `${EXPLORER_KAPPA}${kappa}?chainIdFrom=${fromChainId}&chainIdTo=${toChainId}`
  } else {
    return `${EXPLORER_KAPPA}${kappa}?chainIdFrom=${fromChainId}`
  }
}

export const UserExplorerLink = ({
  connectedAddress,
}: {
  connectedAddress?: Address | string
}) => {
  const explorerLink: string = connectedAddress
    ? `${EXPLORER_PATH}address/${connectedAddress}`
    : EXPLORER_PATH
  return (
    <div data-test-id="explorer-link" className="text-sky-500 my-3">
      <Link href={explorerLink} target="_blank">
        <span className="hover:underline">Explorer</span> →
      </Link>
    </div>
  )
}

export const getTransactionHashExplorerLink = ({
  transactionHash,
}: {
  transactionHash: string
}) => {
  return `${EXPLORER_PATH}txs?hash=${transactionHash}`
}
