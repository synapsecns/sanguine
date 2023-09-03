import { DISCORD_URL, TWITTER_URL } from '@/constants/urls'
import {
  ARBITRUM,
  AVALANCHE,
  DOGE,
  ETH,
  FANTOM,
  HARMONY,
} from '@/constants/chains/master'
import { useBridgeState } from '@/slices/bridge/hooks'

export const Warning = () => {
  const { fromChainId, toChainId, fromToken, toToken } = useBridgeState()

  const fromTokenSymbol = fromToken && fromToken?.symbol
  const toTokenSymbol = toToken && toToken?.symbol

  const isTokenUSDCAndUSDCe =
    (fromToken?.symbol === 'USDC' && toToken?.symbol === 'USDCe') ||
    (fromToken?.symbol === 'USDCe' && toToken?.symbol === 'USDC')

  const isChainAvalancheOrArbitrum =
    fromChainId === AVALANCHE.id ||
    fromChainId === ARBITRUM.id ||
    toChainId === AVALANCHE.id ||
    toChainId === ARBITRUM.id

  const isTokenUSDC = fromTokenSymbol === 'USDC' && toTokenSymbol === 'USDC'

  const isChainOtherThanAvalancheArbitrumEthereum =
    (toChainId != AVALANCHE.id &&
      toChainId != ARBITRUM.id &&
      toChainId != ETH.id) ||
    (fromChainId != AVALANCHE.id &&
      fromChainId != ARBITRUM.id &&
      fromChainId != ETH.id)

  const isTokenUSDCAndChainEthereumArbitrumAvalanche =
    isTokenUSDC &&
    ((fromChainId === ETH.id && toChainId === ARBITRUM.id) ||
      (fromChainId === ARBITRUM.id && toChainId === AVALANCHE.id) ||
      (fromChainId === AVALANCHE.id && toChainId === ARBITRUM.id))

  const isChainHarmony = [fromChainId, toChainId].includes(HARMONY.id)
  const isChainFantom = [fromChainId, toChainId].includes(FANTOM.id)
  const isChainDoge = [fromChainId, toChainId].includes(DOGE.id)

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
  } else if (isChainDoge) {
    return (
      <WarningMessage
        header="Alert: Transactions to Dogechain are temporarily paused."
        message={
          <>
            <p>
              You may still bridge funds from Dogechain to any supported
              destination chain.
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
