import EventCard from './EventCard'
import DestinationTx from './DestinationTx'
import { BridgeWatcherTx } from '@types'
import Link from 'next/link'
import { ANALYTICS_KAPPA } from '@urls'
import { memo } from 'react'

const BridgeEvent = memo((fromEvent: BridgeWatcherTx) => {
  // Saving Event Link for when indexing occurs faster
  const EventLink = (
    <Link href={ANALYTICS_KAPPA + fromEvent.kappa} target="_blank">
      <u className="w-auto pt-1 text-sm text-gray-500 underline-offset-2 hover:text-gray-400">
        View on Explorer
      </u>{' '}
    </Link>
  )
  return (
    <div className="mb-3">
      <div className="flex items-center text-gray-500">
        <div className="flex-1 ">
          {fromEvent && <EventCard {...fromEvent} />}
        </div>
        <div className="flex-1 ">
          {fromEvent && <DestinationTx {...fromEvent} />}
        </div>
      </div>
      {/* {EventLink} */}
    </div>
  )
})
export default BridgeEvent
