import { Chain, Token } from '@/utils/types'
import { CHAINS_BY_ID } from '@/constants/chains'

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
  const originChain: Chain = CHAINS_BY_ID[originChainId]
  const destinationChain: Chain = CHAINS_BY_ID[destinationChainId]

  const { name: originChainName } = originChain
  const { name: destinationChainName } = destinationChain

  if (originChainName === 'Ethereum' && destinationChainName === 'Avalanche') {
    return (
      <WarningMessage
        header="USD Coin from Ethereum to Avalanche is not yet available."
        message="Synapse is currently opening liquidity routes for the new USDC to USDC cross-chain transfer protocol.
        Follow Twitter or Discord for updates when this route becomes available."
      />
    )
  }
  return <>{originChainName}</>
}

const WarningMessage = ({
  header,
  message,
}: {
  header: string
  message: string
}) => {
  return (
    <div className="bg-[#353038] text-white text-m p-3">
      <p>{header}</p>
      <p>{message}</p>
    </div>
  )
}
