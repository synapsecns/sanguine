import { QuestionMarkCircleIcon } from '@heroicons/react/outline'
import {
  TOKEN_HASH_MAP,
  AVWETH,
  tokenAddressToToken,
  ETH,
  USDC,
  tokenSymbolToToken,
} from 'synapse-constants'
import { getTokenAddressUrl } from '@urls'
import Image from 'next/image'

export function AssetImage({
  tokenAddress,
  chainId,
  className,
  tokenSymbol = '',
}) {
  tokenAddress = tokenAddress
  if (hasRequiredData({ tokenAddress, chainId })) {
    const t = chainId && tokenAddress && TOKEN_HASH_MAP[chainId]?.[tokenAddress]
    return (
      <a href={getTokenAddressUrl({ tokenAddress, chainId })}>
        <div className="w-full relative">
          <div className="flex justify-between ">
            <div className="flex flex-row w-[90%] items-center">
              <Image className={`${className}`} src={t?.icon} alt="" />
            </div>
          </div>
        </div>
      </a>
    )
  }
  if (
    chainId === 43114 &&
    tokenAddress === '0x53f7c5869a859F0AeC3D334ee8B4Cf01E3492f21'
  ) {
    const t = AVWETH
    return (
      <a href={getTokenAddressUrl({ tokenAddress, chainId })}>
        <Image
          className={`inline mr-[.5rem] rounded-full ${className}`}
          src={t?.icon}
          alt=""
        />
      </a>
    )
  } else {
    return (
      // temporary fix until either symbolToToken works better as a function or explorer indexer has the right token addresses
      <Image className={`${className}`} src={USDC?.icon} alt="" />
      // <QuestionMarkCircleIcon
      //   className={`inline w-5 h-5 mr-2 rounded-md ${className}`}
      //   strokeWidth={2}
      // />
    )
  }
}

function hasRequiredData({ tokenAddress, chainId }) {
  return (
    tokenAddress &&
    chainId &&
    TOKEN_HASH_MAP[chainId] &&
    TOKEN_HASH_MAP[chainId][tokenAddress]
  )
}
