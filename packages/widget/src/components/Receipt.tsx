import { useMemo } from 'react'

import { DoubleDownArrow } from '@/components/icons/DoubleDownArrow'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'

export const Receipt = ({ quote, loading, send, receive }) => {
  const { originChainId, destinationChainId } = useBridgeState()

  const estTime = useMemo(() => {
    if (!quote.estimatedTime) {
      return null
    }

    if (quote?.estimatedTime > 60) {
      return Math.ceil(quote?.estimatedTime / 60) + ' minutes'
    } else {
      return quote?.estimatedTime + ' seconds'
    }
  }, [quote])

  const isValidQuote = Boolean(quote.outputAmount)

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
        <dt className="text-left">Receive</dt>
        <dd className="m-0 text-right justify-self-end">
          {loading ? '...' : isValidQuote ? receive : '-'}
        </dd>
      </dl>
    </details>
  )
}
