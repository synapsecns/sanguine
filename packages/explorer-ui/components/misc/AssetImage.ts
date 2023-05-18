import {QuestionMarkCircleIcon} from '@heroicons/react/outline'
import {TOKEN_HASH_MAP} from '@constants/tokens/basic'
import {getTokenAddressUrl} from '@urls'
import Image from 'next/image'

export function AssetImage({ tokenAddress, chainId, className }) {
  tokenAddress = tokenAddress && tokenAddress.toLowerCase()
  if (hasRequiredData({ tokenAddress, chainId })) {
    const t = TOKEN_HASH_MAP[chainId][tokenAddress]
    return (
      // @ts-expect-error TS(2304): Cannot find name 'a'.
      <a href={getTokenAddressUrl({tokenAddress, chainId})}>
      // @ts-expect-error TS(2749): 'Image' refers to a value, but is being used as a ... Remove this comment to see the full error message
      <Image
        className={`inline w-5 h-5 mr-2 rounded-md ${className}`}
        // @ts-expect-error TS(2304): Cannot find name 'src'.
        src={t?.icon}
        // @ts-expect-error TS(2304): Cannot find name 'alt'.
        alt=""
      /></a>
    )
  } else {
    return (
      // @ts-expect-error TS(2749): 'QuestionMarkCircleIcon' refers to a value, but is... Remove this comment to see the full error message
      <QuestionMarkCircleIcon
        // @ts-expect-error TS(2349): This expression is not callable.
        className={`inline w-5 h-5 mr-2 rounded-md ${className}`}
        // @ts-expect-error TS(2304): Cannot find name 'strokeWidth'.
        strokeWidth={2}
      />
    )
  }
}

function hasRequiredData({ tokenAddress, chainId }) {
  return tokenAddress && chainId && TOKEN_HASH_MAP[chainId][tokenAddress]
}
