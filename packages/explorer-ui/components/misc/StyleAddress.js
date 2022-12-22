import { getAddressesUrl } from '@urls'
import { getNetworkTextHoverColor } from '@utils/styles/networks'
import { ellipsizeString } from '@utils/ellipsizeString'

export function StyleAddress({ sourceInfo, limiter = 6 }) {
  if (sourceInfo.address) {
    return (
      <a
        className={`${getNetworkTextHoverColor(
          sourceInfo.chainId
        )} hover:underline `}
        href={getAddressesUrl({
          address: sourceInfo.address,
          chainIdTo: sourceInfo.chainId,
        })}
        onClick={(e) => e.stopPropagation()}
      >
        {ellipsizeString({ string: sourceInfo.address, limiter })}
      </a>
    )
  } else {
    return '--'
  }
}
