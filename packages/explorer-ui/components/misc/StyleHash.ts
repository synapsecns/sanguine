import {getExplorerTxUrl} from '@urls'
import {getNetworkTextHoverColor} from '@utils/styles/networks'

export function StyleHash({ sourceInfo, limiter = 6 }) {
  if (sourceInfo.hash) {
    return (
      // @ts-expect-error TS(2304): Cannot find name 'a'.
      <a
        // @ts-expect-error TS(2304): Cannot find name 'className'.
        className={`${getNetworkTextHoverColor(
          sourceInfo.chainId
        )} hover:underline `}
        // @ts-expect-error TS(2304): Cannot find name 'href'.
        href={getExplorerTxUrl({
          hash: sourceInfo.hash,
          chainId: sourceInfo.chainId,
        })}
        // @ts-expect-error TS(2552): Cannot find name 'onClick'. Did you mean 'onclick'... Remove this comment to see the full error message
        onClick={(e) => e.stopPropagation()}
        // @ts-expect-error TS(2304): Cannot find name 'target'.
        target="_blank"
        // @ts-expect-error TS(2304): Cannot find name 'rel'.
        rel="noreferrer"
      >
        // @ts-expect-error TS(18004): No value exists in scope for the shorthand propert... Remove this comment to see the full error message
        {sourceInfo.hash.toLowerCase().slice(0, limiter)}...
        // @ts-expect-error TS(2304): Cannot find name 'sourceInfo'.
        {sourceInfo.hash.toLowerCase().slice(-limiter, sourceInfo.hash.length)}
      </a>
    )
  } else {
    return '--'
  }
}
