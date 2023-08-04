import { Token } from '@/utils/types'
import { DISCORD_URL, TWITTER_URL } from '@/constants/urls'
import {
  ARBITRUM,
  AVALANCHE,
  ETH,
  FANTOM,
  HARMONY,
} from '@/constants/chains/master'

interface WarningProps {
  originChainId: number
  destinationChainId: number
  originToken: Token | undefined
  destinationToken: Token
}

export const Warning = ({
  originChainId,
  destinationChainId,
  originToken,
  destinationToken,
}: WarningProps) => {
  if (!originToken || !destinationToken) {
    return null
  }

  const originTokenSymbol = originToken && originToken.symbol
  const destinationTokenSymbol = destinationToken && destinationToken.symbol

  const isTokenUSDCAndUSDCe =
    (originTokenSymbol === 'USDC' && destinationTokenSymbol === 'USDCe') ||
    (originTokenSymbol === 'USDCe' && destinationTokenSymbol === 'USDC')

  const isChainAvalancheOrArbitrum =
    originChainId === AVALANCHE.id ||
    originChainId === ARBITRUM.id ||
    destinationChainId === AVALANCHE.id ||
    destinationChainId === ARBITRUM.id

  const isTokenUSDC =
    originTokenSymbol === 'USDC' && destinationTokenSymbol === 'USDC'

  const isChainOtherThanAvalancheArbitrumEthereum =
    (destinationChainId != AVALANCHE.id &&
      destinationChainId != ARBITRUM.id &&
      destinationChainId != ETH.id) ||
    (originChainId != AVALANCHE.id &&
      originChainId != ARBITRUM.id &&
      originChainId != ETH.id)

  const isTokenUSDCAndChainEthereumArbitrumAvalanche =
    isTokenUSDC &&
    ((originChainId === ETH.id && destinationChainId === ARBITRUM.id) ||
      (originChainId === ARBITRUM.id && destinationChainId === AVALANCHE.id) ||
      (originChainId === AVALANCHE.id && destinationChainId === ARBITRUM.id))

  const isChainHarmony =
    originChainId === HARMONY.id || destinationChainId === HARMONY.id
  const isChainFantom =
    originChainId === FANTOM.id || destinationChainId === FANTOM.id

  if (isTokenUSDCAndUSDCe && isChainAvalancheOrArbitrum) {
    return (
      <WarningMessage
        header="USDC and USDC.e are incompatible on some routes."
        message={
          <>
            <p className="mb-2">
              If you don't see an output, try bridging to another stablecoin.
              For example, on some routes, USDC can currently only be bridged
              with itself, via the new Circle cross-chain transfer protocol.
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

export const WarningMessage = ({
  header,
  message,
  twClassName,
}: {
  header?: string
  message?: React.ReactNode
  twClassName?: string
}) => {
  return (
    <div
      className={`flex flex-col bg-[#353038] text-white text-sm p-3 rounded-md mt-4 ${twClassName}`}
    >
      {header && <div className="mb-2 font-bold">{header}</div>}
      {message && <div>{message}</div>}
    </div>
  )
}
