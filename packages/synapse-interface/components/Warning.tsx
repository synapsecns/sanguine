import { Token } from '@/utils/types'

interface WarningProps {
  originChainId: number
  destinationChainId: number
  originToken: Token
  destinationToken: Token
}

export const Warning = ({
  originChainId,
  destinationChainId,
  originToken,
  destinationToken,
}: WarningProps) => {
  return <></>
}
