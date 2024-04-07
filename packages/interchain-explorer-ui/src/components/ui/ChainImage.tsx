import { CHAINS } from '@/constants/chains'
import Image from 'next/image'

export function ChainImage({
  chainId,
  width = 16,
  height = 16,
}: {
  chainId: number
  width?: number
  height?: number
}) {
  return (
    <Image
      src={CHAINS[chainId].imgUrl}
      alt={`${CHAINS[chainId].name} img`}
      width={width}
      height={height}
    />
  )
}
