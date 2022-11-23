import { getExplorerTxUrl } from '@urls'
import { getNetworkTextHoverColor } from '@utils/styles/networks'

export function StyleHash({ sourceInfo, limiter = 6 }) {
  if (sourceInfo.hash) {
    return (
      <a
        className={`${getNetworkTextHoverColor(
          sourceInfo.chainId
        )} hover:underline `}
        href={getExplorerTxUrl({
          hash: sourceInfo.hash,
          chainId: sourceInfo.chainId,
        })}
        onClick={(e) => e.stopPropagation()}
        target="_blank"
        rel="noreferrer"
      >
        {sourceInfo.hash.toLowerCase().slice(0, limiter)}...
        {sourceInfo.hash.toLowerCase().slice(-limiter, sourceInfo.hash.length)}
      </a>
    )
  } else {
    return '--'
  }
}
