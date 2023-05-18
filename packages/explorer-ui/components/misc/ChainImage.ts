import {QuestionMarkCircleIcon} from '@heroicons/react/outline'

import {CHAIN_INFO_MAP} from '@constants/networks'
import Image from 'next/image'

export function ChainImage({ chainId, imgSize = 'w-4 h-4', className }) {
  if (chainId) {
    const chainImg = CHAIN_INFO_MAP[chainId]
    return (
      // @ts-expect-error TS(2749): 'Image' refers to a value, but is being used as a ... Remove this comment to see the full error message
      <Image
        // @ts-expect-error TS(2304): Cannot find name 'src'.
        src={chainImg.chainImg}
        // @ts-expect-error TS(2349): This expression is not callable.
        className={`${imgSize} rounded-full mr-2 inline ${className}`}
        // @ts-expect-error TS(2304): Cannot find name 'alt'.
        alt={chainImg.chainName}
      />
    )
  } else {
    return (
      <QuestionMarkCircleIcon
        className={`${imgSize} rounded-full mr-2 inline`}
        // @ts-expect-error TS(2304): Cannot find name 'strokeWidth'.
        strokeWidth={1}
      />
    )
  }
}
