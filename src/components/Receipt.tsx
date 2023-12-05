import { useMemo, useState } from 'react'
import { DoubleUpArrow } from '@/components/DoubleUpArrow'
import { DoubleDownArrow } from '@/components/DoubleDownArrow'

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
      { quote ? (
        <div onClick={handleToggle} className={`cursor-s-resize hover:bg-[--separator] flex self-end pl-2 pr-1 py-1 gap-1 rounded active:opacity-40 ${isExpanded ? 'cursor-n-resize' : 'cursor-s-resize'}`}>
            {estTime} min via
            <a href="https://synapseprotocol.com" target="_blank" className="text-[--brand] cursor-alias text-[--strong]">
              Synapse
            </a>
            {isExpanded ? <DoubleUpArrow /> : <DoubleDownArrow />}
          </div>
        ) : (
          <div className={`flex self-end pl-2 pr-1 py-1 gap-1 text-[--secondary] cursor-default`}>
            Powered by
            <a href="https://synapseprotocol.com" target="_blank" className="text-[--brand] cursor-alias active:opacity-40 ">
              Synapse
            </a>
          </div>
        )
      }
      </div>
      {isExpanded && (
        <div className="p-2 text-sm border border-[--separator]">
          <div className="flex justify-between">
            <div>Router</div>
            <div className="text-[--primary]">{quote.bridgeModuleName}</div>
          </div>
          <div className="flex justify-between">
            <div>Origin</div>
            <div>Ethereum</div>
          </div>
          <div className="flex justify-between">
            <div>Destination</div>
            <div>Arbitrum</div>
          </div>
          <div className="flex justify-between">
            <div>Send</div>
            <div>{send}</div>
          </div>
          <div className="flex justify-between">
            <div>Receive</div>
            <div>{receive}</div>
          </div>
        </div>
      )}
    </>
  )
}
