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
  const { symbol: originTokenSymbol } = originToken
  const { symbol: destinationTokenSymbol } = destinationToken

  if (
    (originTokenSymbol === 'USDC' &&
      destinationTokenSymbol === 'USDCe' &&
      originChainName !== 'Ethereum') ||
    (originTokenSymbol === 'USDCe' &&
      destinationTokenSymbol === 'USDC' &&
      destinationChainName !== 'Ethereum')
  ) {
    return (
      <WarningMessage
        header="USDC and USDC.e are incompatible."
        message={
          <>
            <p className="mb-2">
              USD Coin can currently only be bridged with itself, via the new
              cross-chain transfer protocol.
            </p>
            <p>
              Follow Twitter or Discord for updates as more CCTP routes become
              available.
            </p>
          </>
        }
      />
    )
  } else if (
    originChainName === 'Ethereum' &&
    destinationChainName === 'Avalanche' &&
    originTokenSymbol === 'USDC' &&
    destinationTokenSymbol === 'USDC'
  ) {
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
    destinationChainName === 'Arbitrum' &&
    originTokenSymbol === 'USDC' &&
    destinationTokenSymbol === 'USDC'
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
  } else if (
    originChainName === 'Harmony' ||
    destinationChainName === 'Harmony'
  ) {
    return (
      <WarningMessage
        header="Warning! The Harmony bridge has been exploited."
        message={
          <>
            <p>
              Do not bridge via Harmony unless you understand the risks
              involved.
            </p>
          </>
        }
      />
    )
  } else if (
    originChainName === 'Fantom' ||
    destinationChainName === 'Fantom'
  ) {
    return (
      <WarningMessage
        header="Warning! The Fantom bridge has been exploited."
        message={
          <>
            <p>
              Do not bridge via Fantom unless you understand the risks involved.
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
