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
import { getCoinTextColorCombined } from '@styles/tokens'
import { BridgeWatcherTx } from '@types'
import { BRIDGE_REQUIRED_CONFIRMATIONS } from '@constants/bridge'
import {
  EmptySubTransactionItem,
  CheckingConfPlaceholder,
} from '../../../components/TransactionItems'
import _ from 'lodash'
export default function BlockCountdown({
  fromEvent,
  toEvent,
}: {
  fromEvent: BridgeWatcherTx
  toEvent?: BridgeWatcherTx
}) {
  const [time, setTime] = useState(Date.now())
  const chain = CHAINS_BY_ID[fromEvent.toChainId]
  const [confirmationDelta, setConfirmationDelta] = useState(0)

  useEffect(() => {
    const interval = setInterval(() => {
      if (toEvent?.kappa) {
        setTime(Date.now())
      }
    }, 10000)
    return () => {
      clearInterval(interval)
    }
  }, [])

  useEffect(() => {
    fetchBlockNumber({
      chainId: fromEvent.toChainId,
    }).then((newestBlockNumber) => {
      const delta = (newestBlockNumber ?? 0) - fromEvent.blockNumber
      setConfirmationDelta(delta > 0 ? delta : 0)
    })
  }, [time])

  const fromNetworkColorClassName = getNetworkTextColor(chain.color)

  return (
    <>
      <div className="flex-1">
        <div className={`flex items-center p-2 align-middle`}>
          {fromEvent && !toEvent && confirmationDelta != 0 && (
            <>
              <ChevronRightIcon
                className={`
                  w-5 h-5
                  place-self-center
                  ${fromNetworkColorClassName}
                  text-opacity-50
                `}
              />
              {
                <BlockCountdownCircle
                  clampedDiff={confirmationDelta}
                  fromChainConfirmations={
                    BRIDGE_REQUIRED_CONFIRMATIONS[fromEvent.toChainId]
                  }
                  fromNetworkColorClassName={fromNetworkColorClassName}
                />
              }
              <CheckingConfPlaceholder chainId={chain.id} />
              <ChevronRightIcon
                className={`
                  w-5 h-5 animate-pulse
                  place-self-center
                  text-gray-500
                `}
              />
              <EmptySubTransactionItem chainId={fromEvent.toChainId} />
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

function BlockCountdownCircle({
  clampedDiff,
  fromChainConfirmations,
  fromNetworkColorClassName,
}) {
  const dataMatrix = [
    [fromChainConfirmations - clampedDiff, 0, 0, 0],
    [clampedDiff, 0, 0, 0],
  ]
  console.log('clampedDiff', clampedDiff)
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
                              ${
                                i == 0
                                  ? `fill-current ${fromNetworkColorClassName}`
                                  : undefined
                              }
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
