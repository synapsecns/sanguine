import { ellipsizeString } from '@utils/ellipsizeString'
import { getAddressesUrl } from '@urls'
export function StyleAddress({ sourceInfo, limiter = 4 }) {
  if (sourceInfo.address) {
    return (
      <a
        className="underline hover:text-[#8FEBFF] transition ease-out hover:"
        href={getAddressesUrl({
          address: sourceInfo.address,
          chainIdTo: sourceInfo.chainId,
        })}
        onClick={(e) => e.stopPropagation()}
      >
        {ellipsizeString({
          string: sourceInfo.address,
          limiter,
          isZeroX: true,
        })}
      </a>
    )
  } else {
    return <div>'--'</div>
  }
}
