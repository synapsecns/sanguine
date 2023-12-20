import { useMemo } from 'react'
import { DoubleDownArrow } from '@/components/icons/DoubleDownArrow'

export const Receipt = ({ quote, send, receive }) => {
  const estTime = useMemo(() => {
    return quote?.estimatedTime / 60
  }, [quote])

  return (
    quote ? (
      <details className="text-sm group text-right">
        <summary
          className="hover:bg-[--synapse-border] pl-2 pr-1 py-1 gap-1 rounded active:opacity-40 cursor-pointer list-none inline-flex items-center"
        >
          {estTime} min via Synapse <DoubleDownArrow />
        </summary>
        <dl className="receipt mt-1 mb-0 p-2 text-sm rounded border border-solid border-[--synapse-border] grid grid-cols-2">
          <dt className="text-left">Router</dt>
          <dd className="text-right">{quote?.bridgeModuleName}</dd>
          <dt className="text-left">Origin</dt>
          <dd className="text-right">Ethereum</dd>
          <dt className="text-left">Destination</dt>
          <dd className="text-right">Arbitrum</dd>
          <dt className="text-left">Send</dt>
          <dd className="text-right">{send}</dd>
          <dt className="text-left">Receive</dt>
          <dd className="text-right">{receive}</dd>
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
  )
}
