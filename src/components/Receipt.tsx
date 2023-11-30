import { useMemo, useState } from 'react'
import { DoubleUpArrow } from '@/components/DoubleUpArrow'
import { DoubleDownArrow } from '@/components/DoubleDownArrow'

export const Receipt = ({ quote, send, receive }) => {
  const [isExpanded, setIsExpanded] = useState<boolean>(false)
  const estTime = useMemo(() => {
    return quote.estimatedTime / 60
  }, [quote])

  const handleToggle = () => {
    setIsExpanded(!isExpanded)
  }

  return (
    <div>
      <div className="flex items-center justify-end">
        <div className="text-sm">
          {estTime} min via <span className="text-widget-accent">Synapse</span>
        </div>
        <div onClick={handleToggle}>
          {isExpanded ? <DoubleUpArrow /> : <DoubleDownArrow />}
        </div>
      </div>
      {isExpanded && (
        <div className="p-2 mt-2 text-sm border border-widget-separator">
          <div className="flex items-center justify-between">
            <div>Router</div>
            <div className="text-widget-accent">{quote.bridgeModuleName}</div>
          </div>
          <div className="flex items-center justify-between">
            <div>Origin</div>
            <div>Ethereum</div>
          </div>
          <div className="flex items-center justify-between">
            <div>Destination</div>
            <div>Arbitrum</div>
          </div>
          <div className="flex items-center justify-between">
            <div>Send</div>
            <div>{send}</div>
          </div>
          <div className="flex items-center justify-between">
            <div>Receive</div>
            <div>{receive}</div>
          </div>
        </div>
      )}
    </div>
  )
}
