import { DoubleDownArrow } from '@/components/icons/DoubleDownArrow'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { formatBigIntToString } from '@/utils/formatBigIntToString'

const formatBridgeFee = (
  nativeFee: bigint | undefined,
  nativePrecision: number | undefined
) => {
  if (!nativePrecision || typeof nativeFee !== 'bigint' || nativeFee <= 0n) {
    return null
  }

  const fullPrecision = formatBigIntToString(
    nativeFee,
    nativePrecision,
    nativePrecision
  )

  if (!fullPrecision) {
    return null
  }

  const [whole, fraction = ''] = fullPrecision.split('.')
  const trimmedFraction = fraction.replace(/0+$/, '')

  if (!trimmedFraction) {
    return whole
  }

  const firstNonZero = trimmedFraction.search(/[1-9]/)
  const decimalPlaces =
    whole === '0'
      ? Math.min(trimmedFraction.length, Math.max(4, firstNonZero + 2))
      : Math.min(trimmedFraction.length, 4)
  const visibleFraction = trimmedFraction.slice(0, decimalPlaces)

  return visibleFraction ? `${whole}.${visibleFraction}` : whole
}

const getEstimatedTimeLabel = (estimatedTime?: number | null) => {
  if (!estimatedTime) {
    return null
  }

  if (estimatedTime > 60) {
    return `${Math.ceil(estimatedTime / 60)} minutes`
  }

  return `${estimatedTime} seconds`
}

const getFormattedBridgeFee = ({
  nativeFee,
  originChainId,
  isValidQuote,
  loading,
}: {
  nativeFee: bigint | undefined
  originChainId?: number
  isValidQuote: boolean
  loading: boolean
}) => {
  if (!originChainId || !isValidQuote || loading) {
    return null
  }

  return formatBridgeFee(
    nativeFee,
    CHAINS_BY_ID[originChainId]?.nativeCurrency.decimals
  )
}

export const Receipt = ({ quote, loading, send, receive }) => {
  const { originChainId, destinationChainId } = useBridgeState()

  const estTime = getEstimatedTimeLabel(quote?.estimatedTime)
  const isValidQuote = Boolean(quote.outputAmount)
  const nativeFeeSymbol = originChainId
    ? CHAINS_BY_ID[originChainId]?.nativeCurrency.symbol
    : null
  const formattedBridgeFee = getFormattedBridgeFee({
    nativeFee: quote?.nativeFee,
    originChainId,
    isValidQuote,
    loading,
  })
  const shouldShowBridgeFee = Boolean(formattedBridgeFee)

  return (
    <details className="text-sm text-right group">
      <summary className="hover:bg-[--synapse-select-border] pl-2 pr-1 py-1 gap-1 rounded active:opacity-40 cursor-pointer list-none inline-flex items-center">
        {loading ? (
          <>
            fetching... <DoubleDownArrow />
          </>
        ) : isValidQuote ? (
          <>
            {' '}
            {estTime} via {quote?.bridgeModuleName} <DoubleDownArrow />{' '}
          </>
        ) : (
          <div className="text-sm text-right text-[--synapse-secondary]">
            Powered by&nbsp;
            <a
              href="https://synapseprotocol.com"
              target="_blank"
              className="text-[--synapse-text] no-underline hover:underline active:opacity-40"
            >
              Synapse
            </a>
          </div>
        )}
      </summary>
      <dl className="receipt mt-1 mb-0 p-2 text-sm rounded border border-solid border-[--synapse-select-border] grid grid-cols-[auto_auto] gap-1">
        <dt className="text-left">Router</dt>
        <dd className="m-0 text-right justify-self-end">
          {loading ? '...' : isValidQuote ? quote?.bridgeModuleName : '-'}
        </dd>
        <dt className="text-left">Origin</dt>
        <dd className="m-0 text-right justify-self-end">
          {originChainId ? CHAINS_BY_ID[originChainId]?.name : '-'}
        </dd>
        <dt className="text-left">Destination</dt>
        <dd className="m-0 text-right justify-self-end">
          {destinationChainId ? CHAINS_BY_ID[destinationChainId]?.name : '-'}
        </dd>
        <dt className="text-left">Send</dt>
        <dd className="m-0 text-right justify-self-end">{send ? send : '-'}</dd>
        {shouldShowBridgeFee && (
          <>
            <dt className="text-left">Bridge fee</dt>
            <dd className="m-0 text-right justify-self-end">
              {formattedBridgeFee}
              {nativeFeeSymbol ? ` ${nativeFeeSymbol}` : ''}
            </dd>
          </>
        )}
        <dt className="text-left">Receive</dt>
        <dd className="m-0 text-right justify-self-end">
          {loading ? '...' : isValidQuote ? receive : '-'}
        </dd>
      </dl>
    </details>
  )
}
