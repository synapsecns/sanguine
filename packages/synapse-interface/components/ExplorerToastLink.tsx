import { ExternalLinkIcon } from '@heroicons/react/outline'
import { getExplorerTxUrl } from '@urls'

const ExplorerToastLink = ({
  transactionHash,
  chainId,
}: {
  transactionHash: string
  chainId: number
}) => {
  const explorerTxUrl = getExplorerTxUrl({ hash: transactionHash, chainId })
  const len = transactionHash.length
  return (
    <a target="_blank" href={explorerTxUrl} className="hover:text-blue-500">
      {transactionHash.slice(0, 6)}...{transactionHash.slice(len - 4, len)}
      <ExternalLinkIcon className="inline w-4 h-4 ml-2" />
    </a>
  )
}
export default ExplorerToastLink
