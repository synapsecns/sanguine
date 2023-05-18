import { QuestionMarkCircleIcon } from '@heroicons/react/outline'
import { TOKEN_HASH_MAP } from '@constants/tokens/basic'
import { getTokenAddressUrl } from '@urls'
import Image from 'next/image'

export function AssetImage({ tokenAddress, chainId, className }) {
  tokenAddress = tokenAddress && tokenAddress.toLowerCase()
  if (hasRequiredData({ tokenAddress, chainId })) {
    const t = TOKEN_HASH_MAP[chainId][tokenAddress]
    return (
      <a href={getTokenAddressUrl({ tokenAddress, chainId })}>
        <Image
          className={`inline w-5 h-5 mr-2 rounded-md ${className}`}
          src={t?.icon}
          alt=""
        />
      </a>
    )
  } else {
    return (
      <QuestionMarkCircleIcon
        className={`inline w-5 h-5 mr-2 rounded-md ${className}`}
        strokeWidth={2}
      />
    )
  }
}

function hasRequiredData({ tokenAddress, chainId }) {
  return tokenAddress && chainId && TOKEN_HASH_MAP[chainId][tokenAddress]
}
