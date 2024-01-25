import { useMemo } from 'react'
import { DoubleDownArrow } from '@/components/icons/DoubleDownArrow'

export const Receipt = ({ quote, send, receive }) => {
  const estTime = useMemo(() => {
    if (!quote.estimatedTime) return null

    if (quote?.estimatedTime > 60) {
      return Math.ceil(quote?.estimatedTime / 60) + ' minutes'
    } else {
      return quote?.estimatedTime + ' seconds'
    }
  }, [quote])

  return quote ? (
    <details className="text-sm text-right group">
      <summary className="hover:bg-[--synapse-border] pl-2 pr-1 py-1 gap-1 rounded active:opacity-40 cursor-pointer list-none inline-flex items-center">
        {estTime ? (
          <>
            {' '}
            {estTime} via {quote?.bridgeModuleName} <DoubleDownArrow />{' '}
          </>
        ) : null}
      </summary>
      <dl className="receipt mt-1 mb-0 p-2 text-sm rounded border border-solid border-[--synapse-border] grid grid-cols-[auto_auto]">
        <dt className="text-left">Router</dt>
        <dd className="m-0 text-right justify-self-end">
          {quote?.bridgeModuleName}
        </dd>
        <dt className="text-left">Origin</dt>
        <dd className="m-0 text-right justify-self-end">Ethereum</dd>
        <dt className="text-left">Destination</dt>
        <dd className="m-0 text-right justify-self-end">Arbitrum</dd>
        <dt className="text-left">Send</dt>
        <dd className="m-0 text-right justify-self-end">{send}</dd>
        <dt className="text-left">Receive</dt>
        <dd className="m-0 text-right justify-self-end">{receive}</dd>
      </dl>
    </details>
  ) : (
    <div className="text-sm text-right p-1 text-[--synapse-secondary]">
      Powered by&nbsp;
      <a
        href="https://synapseprotocol.com"
        target="_blank"
        className="text-[--synapse-text] no-underline hover:underline active:opacity-40"
      >
        Synapse
      </a>
    </div>
  )
}
