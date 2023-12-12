import { useMemo, useState } from 'react'
import { DoubleUpArrow } from '@/components/icons/DoubleUpArrow'
import { DoubleDownArrow } from '@/components/icons/DoubleDownArrow'

export const Receipt = ({ quote, send, receive }) => {
  const [isExpanded, setIsExpanded] = useState<boolean>(false)
  const estTime = useMemo(() => {
    return quote?.estimatedTime / 60
  }, [quote])

  const handleToggle = () => {
    setIsExpanded(!isExpanded)
  }

  return (
    <>
      <div className="flex justify-end text-sm">
        {quote ? (
          <div
            onClick={handleToggle}
            className="hover:bg-[--synapse-border] flex self-end pl-2 pr-1 py-1 gap-1 rounded active:opacity-40 cursor-pointer"
          >
            {estTime} min via Synapse
            {isExpanded ? <DoubleUpArrow /> : <DoubleDownArrow />}
          </div>
        ) : (
          <div
            className="flex self-end pl-2 pr-1 py-1 gap-1 text-[--synapse-text-secondary] cursor-default"
          >
            Powered by
            <a
              href="https://synapseprotocol.com"
              target="_blank"
              className="underline hover:text-[--synapse-text-primary] active:opacity-40"
            >
              Synapse
            </a>
          </div>
        )}
      </div>
      {isExpanded && (
        <dl className="receipt p-2 text-sm rounded border border-[--synapse-border] grid grid-cols-2">
          <dt>Router</dt>
          <dd className="text-right">{quote?.bridgeModuleName}</dd>
          <dt>Origin</dt>
          <dd className="text-right">Ethereum</dd>
          <dt>Destination</dt>
          <dd className="text-right">Arbitrum</dd>
          <dt>Send</dt>
          <dd className="text-right">{send}</dd>
          <dt>Receive</dt>
          <dd className="text-right">{receive}</dd>
        </dl>
      )}
    </>
  )
}
