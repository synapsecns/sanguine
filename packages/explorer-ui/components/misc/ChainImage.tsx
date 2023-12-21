import { QuestionMarkCircleIcon } from '@heroicons/react/outline'
import { CHAINS_BY_ID } from '../../constants/chains'
import Image from 'next/image'

export function ChainImage({ chainId, imgSize = 'w-4 h-4', className }) {
  if (chainId) {
    const chain = CHAINS_BY_ID[chainId]
    return (
      <Image
        src={chain.chainImg}
        className={`${imgSize} rounded-full mr-2 inline ${className}`}
        alt={chain.name}
      />
    )
  } else {
    return (
      <QuestionMarkCircleIcon
        className={`${imgSize} rounded-full mr-2 inline`}
        strokeWidth={1}
      />
    )
  }
}
