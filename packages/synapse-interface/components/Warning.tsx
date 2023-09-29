import { DOGE, FANTOM, HARMONY } from '@/constants/chains/master'
import { useBridgeState } from '@/slices/bridge/hooks'

export const Warning = () => {
  const { fromChainId, toChainId } = useBridgeState()

  const isChainHarmony = [fromChainId, toChainId].includes(HARMONY.id)
  const isChainFantom = [fromChainId, toChainId].includes(FANTOM.id)
  const isChainDoge = [fromChainId, toChainId].includes(DOGE.id)

  if (isChainHarmony) {
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
