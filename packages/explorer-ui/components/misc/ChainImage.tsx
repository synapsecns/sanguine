import Image from 'next/image'
import { QuestionMarkCircleIcon } from '@heroicons/react/outline'

//@ts-ignore
import { CHAINS } from '@synapsecns/synapse-constants'

const CHAINS_BY_ID = CHAINS.CHAINS_BY_ID

export const ChainImage = ({ chainId, imgSize = 'w-4 h-4', className }) => {
  if (chainId) {
    const chain = CHAINS_BY_ID[chainId]
    return (
      <Image
        src={chain.chainImg}
        className={`${imgSize} rounded-full mr-2 inline ${className}`}
        alt={chain.name}
        height={16}
        width={16}
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
