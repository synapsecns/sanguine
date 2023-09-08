import { useMemo } from 'react'

import { DISCORD_URL, TWITTER_URL } from '@/constants/urls'
import { AVALANCHE, DOGE, FANTOM, HARMONY } from '@/constants/chains/master'
import { useBridgeState } from '@/slices/bridge/hooks'

import { CCTP_ROUTER_ADDRESS } from '@/utils/actions/fetchPortfolioBalances'

export const Warning = () => {
  const { fromChainId, toChainId, fromToken, bridgeQuote } = useBridgeState()

  const isCCTP = useMemo(() => {
    return (
      fromChainId !== AVALANCHE.id &&
      fromToken.swapableType === 'USD' &&
      bridgeQuote.routerAddress.toLowerCase() ===
        CCTP_ROUTER_ADDRESS.toLowerCase()
    )
  }, [bridgeQuote, fromToken])

  const isChainHarmony = [fromChainId, toChainId].includes(HARMONY.id)
  const isChainFantom = [fromChainId, toChainId].includes(FANTOM.id)
  const isChainDoge = [fromChainId, toChainId].includes(DOGE.id)

  if (isCCTP) {
    return (
      <WarningMessage
        header="This route uses the Circle cross-chain transfer protocol."
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
