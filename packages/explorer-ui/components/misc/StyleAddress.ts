import {ellipsizeString} from '@utils/ellipsizeString'
import {getAddressesUrl} from '@urls'
export function StyleAddress({ sourceInfo, limiter = 4 }) {
  if (sourceInfo.address) {
    return (
      // @ts-expect-error TS(2304): Cannot find name 'a'.
      <a
        // @ts-expect-error TS(2304): Cannot find name 'className'.
        className="underline hover:text-[#8FEBFF] transition ease-out hover:"
        // @ts-expect-error TS(2304): Cannot find name 'href'.
        href={getAddressesUrl({
          address: sourceInfo.address,
          chainIdTo: sourceInfo.chainId,
        })}
        // @ts-expect-error TS(2552): Cannot find name 'onClick'. Did you mean 'onclick'... Remove this comment to see the full error message
        onClick={(e) => e.stopPropagation()}
      >
        // @ts-expect-error TS(18004): No value exists in scope for the shorthand propert... Remove this comment to see the full error message
        {ellipsizeString({ string: sourceInfo.address, limiter, isZeroX: true })}
      </a>
    )
  } else {
    return '--'
  }
}
