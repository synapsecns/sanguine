import { formatTimestampToDate } from '@utils/time'
import EventCard from './EventCard'
import DestinationTx from './DestinationTx'
import { BridgeWatcherTx } from '@types'
import Link from 'next/link'
import Button from '@tw/Button'
import { ANALYTICS_KAPPA } from '@urls'

const BridgeEvent = (fromEvent: BridgeWatcherTx) => {
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
      <Link href={ANALYTICS_KAPPA + fromEvent.kappa} target="_blank">
        <u className="w-auto  pt-1 text-sm  underline-offset-2 text-gray-500 hover:text-gray-400">
          View on Explorer
        </u>{' '}
      </Link>
    </div>
  )
}

export default BridgeEvent
