import {
  ChevronRightIcon,
  ChevronDoubleRightIcon,
} from '@heroicons/react/outline'
import { Arc } from '@visx/shape'
import { Chord } from '@visx/chord'
import { CHAINS_BY_ID } from '@/constants/chains'
import { useEffect, useState } from 'react'
import { getNetworkTextColor } from '@styles/chains'
import { fetchBlockNumber } from '@wagmi/core'
import { BridgeWatcherTx } from '@types'
import { BRIDGE_REQUIRED_CONFIRMATIONS } from '@constants/bridge'
import {
  EmptySubTransactionItem,
  CheckingConfPlaceholder,
} from '@components/TransactionItems'
import { memo } from 'react'
import _ from 'lodash'
const BlockCountdown = memo(
  ({
    fromEvent,
    toEvent,
    setCompletedConf,
  }: {
    fromEvent: BridgeWatcherTx
    toEvent?: BridgeWatcherTx
    setCompletedConf: (bool: boolean) => void
  }) => {
    const chain = fromEvent?.toChainId
      ? CHAINS_BY_ID[fromEvent.toChainId]
      : null
    const [confirmationDelta, setConfirmationDelta] = useState(-1)
    const [time, setTime] = useState(Date.now())

    useEffect(() => {
      const interval = setInterval(() => {
        setTime(Date.now())
      }, 5000)

      return () => {
        console.log('cleared')
        clearInterval(interval)
      }
    }, [])

    useEffect(() => {
      if (confirmationDelta === 0 || toEvent) {
        return
      }
      fetchBlockNumber({
        chainId: fromEvent?.chainId,
      }).then((newestBlockNumber) => {
        const delta =
          BRIDGE_REQUIRED_CONFIRMATIONS[fromEvent?.chainId] +
          (fromEvent.blockNumber - (newestBlockNumber ?? 0))
        setConfirmationDelta(delta > 0 ? delta : 0)
        if (delta <= 0) {
          setCompletedConf(true)
        }
      })
    }, [time])

    console.log('fromEvent?.toChainId: ', fromEvent?.toChainId)
    console.log('toEvent: ', toEvent)
    console.log('confirmationDelta:', confirmationDelta)
    return (
      <>
        <div className="flex-1">
          <div className={`flex items-center p-2 align-middle`}>
            {fromEvent?.toChainId && !toEvent && confirmationDelta != 0 && (
              <>
                <ChevronRightIcon
                  className={`
                  w-5 h-5
                  place-self-center
                  ${getNetworkTextColor(chain?.color)}
                  text-opacity-50
                `}
                />

                <BlockCountdownCircle
                  clampedDiff={confirmationDelta}
                  fromChainConfirmations={
                    BRIDGE_REQUIRED_CONFIRMATIONS[fromEvent?.chainId]
                  }
                  coloring={getNetworkTextColor(chain?.color)}
                />

                <CheckingConfPlaceholder chain={chain} />
                <ChevronRightIcon
                  className={`
                  w-5 h-5 animate-pulse
                  place-self-center
                  text-gray-500
                `}
                />
                <EmptySubTransactionItem chainId={fromEvent?.toChainId} />
              </>
            )}
            {confirmationDelta == 0 && (
              <div className="items-center flex-shrink-0 align-middle">
                <ChevronDoubleRightIcon
                  className={`
                w-5 h-5
                place-self-center
                text-gray-500'
                text-opacity-50
              `}
                />
              </div>
            )}
          </div>
        </div>
      </>
    )
  }
)

const BlockCountdownCircle = ({
  clampedDiff,
  fromChainConfirmations,
  coloring,
}) => {
  const dataMatrix = [
    [fromChainConfirmations - clampedDiff, 0, 0, 0],
    [clampedDiff, 0, 0, 0],
  ]
  return (
    <svg
      viewBox="0 0 200 200"
      xmlns="http://www.w3.org/2000/svg"
      className={`
            stroke-current stroke-[8px] text-gray-600
            bg-transparent
            fill-none
            w-16 h-16  -mb-3
          `}
    >
      <text
        textAnchor="middle"
        x="100"
        y="96"
        className="tracking-wider text-gray-400 fill-current stroke-0"
        style={{ fontSize: 44 }}
      >
        {clampedDiff}
      </text>
      <g transform="translate(100, 80)">
        <Chord matrix={dataMatrix} padAngle={0}>
          {({ chords }) => (
            <g>
              {chords.groups
                .filter((group) => group.value != 0)
                .map((group, i) => (
                  <Arc
                    key={`key-${i}`}
                    data={group}
                    innerRadius={72}
                    outerRadius={74}
                    className={`
                              ${i == 0 ? `fill-current ${coloring}` : null}
                              transform-gpu transition-all
                            `}
                  />
                ))}
            </g>
          )}
        </Chord>
      </g>
    </svg>
  )
}

export default BlockCountdown
