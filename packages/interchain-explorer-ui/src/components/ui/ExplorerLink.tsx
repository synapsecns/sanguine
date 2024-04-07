import { CHAINS } from '@/constants/chains'
import { shortenHash } from '@/utils/shortenHash'

export function ExplorerLink({
  chainId,
  transactionHash,
}: {
  chainId: number
  transactionHash: string
}) {
  const { explorerUrl } = CHAINS[chainId]

  return (
    <a
      href={`${explorerUrl}tx/${transactionHash}`}
      className="hover:underline hover:text-blue-500"
      target="_blank"
      rel="noopener noreferrer"
    >
      {shortenHash(transactionHash)}
    </a>
  )
}
