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
        message={
          <>
            <p className="mb-2">
              Synapse is currently opening liquidity routes for the new USDC to
              USDC cross-chain transfer protocol.
            </p>
            <p>
              Follow Twitter or Discord for updates when this route becomes
              available.
            </p>
          </>
        }
      />
    )
  } else if (
    originChainName === 'Ethereum' &&
    destinationChainName === 'Arbitrum'
  ) {
    return (
      <WarningMessage
        header="USD Coin now uses the free cross-chain transfer protocol."
        message={
          <>
            <p className="mb-2">
              CCTP transfers may take up to 20 minutes to complete.
            </p>
            <p>
              Follow Twitter or Discord for updates as more CCTP routes become
              available.
            </p>
          </>
        }
      />
    )
  }
}

const WarningMessage = ({
  header,
  message,
}: {
  header: string
  message: React.ReactNode
}) => {
  return (
    <div className="flex flex-col bg-[#353038] text-white text-m p-3 rounded-md mt-4">
      <div className="mb-2 font-bold">{header}</div>
      <div>{message}</div>
    </div>
  )
}
