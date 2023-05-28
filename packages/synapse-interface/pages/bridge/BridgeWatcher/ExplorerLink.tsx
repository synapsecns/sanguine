import { ExternalLinkIcon } from '@heroicons/react/outline'
import { getExplorerTxUrl } from '@urls'

const ExplorerLink = ({
  transactionHash,
  chainId,
  className,
  overrideExistingClassname = false,
  showIcon = false,
}) => {
  const explorerTxUrl = getExplorerTxUrl({ hash: transactionHash, chainId })
  const len = transactionHash?.length

  return (
    <a
      target="_blank"
      href={explorerTxUrl}
      className={`
        ${!overrideExistingClassname && 'hover:text-blue-500'}
        ${className}
      `}
    >
      {transactionHash?.slice(0, 6)}...{transactionHash?.slice(len - 4, len)}
      {showIcon && <ExternalLinkIcon className="inline w-4 h-4 ml-2" />}
    </a>
  )
}
export default ExplorerLink
