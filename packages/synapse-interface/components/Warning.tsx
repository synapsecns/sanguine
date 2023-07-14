import { Chain, Token } from '@/utils/types'
import { CHAINS_BY_ID } from '@/constants/chains'
import { DISCORD_URL, TWITTER_URL } from '@/constants/urls'
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

  const isTokenUSDCAndUSDCe =
    (originTokenSymbol === 'USDC' && destinationTokenSymbol === 'USDCe') ||
    (originTokenSymbol === 'USDCe' && destinationTokenSymbol === 'USDC')
  const isChainAvalancheOrArbitrum =
    originChainName == 'Avalanche' ||
    originChainName == 'Arbitrum' ||
    destinationChainName == 'Avalanche' ||
    destinationChainName === 'Arbitrum'

  const isTokenUSDC =
    originTokenSymbol === 'USDC' && destinationTokenSymbol === 'USDC'
  const isChainOtherThanAvalancheArbitrumEthereum =
    (destinationChainName != 'Avalanche' &&
      destinationChainName != 'Arbitrum' &&
      destinationChainName != 'Ethereum') ||
    (originChainName != 'Avalanche' &&
      originChainName != 'Arbitrum' &&
      originChainName != 'Ethereum')

  const isTokenUSDCAndChainEthereumArbitrumAvalanche =
    isTokenUSDC &&
    ((originChainName === 'Ethereum' && destinationChainName === 'Arbitrum') ||
      (originChainName === 'Arbitrum' &&
        destinationChainName === 'Avalanche') ||
      (originChainName === 'Avalanche' && destinationChainName === 'Arbitrum'))

  const isChainHarmony =
    originChainName === 'Harmony' || destinationChainName === 'Harmony'
  const isChainFantom =
    originChainName === 'Fantom' || destinationChainName === 'Fantom'

  if (isTokenUSDCAndUSDCe && isChainAvalancheOrArbitrum) {
    return (
      <WarningMessage
        header="USDC and USDC.e are incompatible."
        message={
          <>
            <p className="mb-2">
              USDC can currently only be bridged with itself, via the new Circle
              cross-chain transfer protocol.
            </p>
            <p>
              Follow{' '}
              <a target="_blank" className="underline" href={TWITTER_URL}>
                Twitter
              </a>{' '}
              or{' '}
              <a target="_blank" className="underline" href={DISCORD_URL}>
                Discord
              </a>{' '}
              for updates as more CCTP routes become available.
            </p>
          </>
        }
      />
    )
  } else if (
    isTokenUSDC &&
    isChainAvalancheOrArbitrum &&
    isChainOtherThanAvalancheArbitrumEthereum
  ) {
    return (
      <WarningMessage
        header="Native USDC on Avalanche and Arbitrum may not be bridgeable to most chains, try another stablecoin instead."
        message={
          <>
            <p className="mb-2">
              Synapse is currently opening liquidity routes for Circle's Native
              USDC to other chains.
            </p>
            <p>
              Follow{' '}
              <a target="_blank" className="underline" href={TWITTER_URL}>
                Twitter
              </a>{' '}
              or{' '}
              <a target="_blank" className="underline" href={DISCORD_URL}>
                Discord
              </a>{' '}
              for updates when this route becomes available.
            </p>
          </>
        }
      />
    )
  } else if (isTokenUSDCAndChainEthereumArbitrumAvalanche) {
    return (
      <WarningMessage
        header="USDC now uses the Circle cross-chain transfer protocol."
        message={
          <>
            <p className="mb-2">
              CCTP transfers may take up to 20 minutes to complete.
            </p>
            <p>
              Follow{' '}
              <a target="_blank" className="underline" href={TWITTER_URL}>
                Twitter
              </a>{' '}
              or{' '}
              <a target="_blank" className="underline" href={DISCORD_URL}>
                Discord
              </a>{' '}
              for updates as more CCTP routes become available.
            </p>
          </>
        }
      />
    )
  } else if (isChainHarmony) {
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
  } else if (isChainFantom) {
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
    <div className="flex flex-col bg-[#353038] text-white text-sm p-3 rounded-md mt-4">
      <div className="mb-2 font-bold">{header}</div>
      <div>{message}</div>
    </div>
  )
}
