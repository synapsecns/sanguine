import _ from 'lodash'
import { useEffect, useState, memo } from 'react'
import { fetchBlockNumber } from '@wagmi/core'
import {
  ChevronRightIcon,
  ChevronDoubleRightIcon,
} from '@heroicons/react/outline'
import { Arc } from '@visx/shape'
import { Chord } from '@visx/chord'
import { BridgeWatcherTx } from '@types'
import { getNetworkTextColor } from '@styles/chains'
import { CHAINS_BY_ID } from '@/constants/chains'
import { BRIDGE_REQUIRED_CONFIRMATIONS } from '@constants/bridge'
import {
  EmptySubTransactionItem,
  CheckingConfPlaceholder,
} from '@components/TransactionItems'

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
    const chain = fromEvent?.chainId ? CHAINS_BY_ID[fromEvent.chainId] : null
    const [confirmationDelta, setConfirmationDelta] = useState(-1)
    const [time, setTime] = useState(Date.now())

    useEffect(() => {
      const interval = setInterval(() => {
        setTime(Date.now())
      }, 5000)

      return () => {
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
        // if error with rpc getting block number, don't run the following code
        if (!newestBlockNumber) {
          return
        }
        // get number of blocks since from event blocknumber
        const blockDifference = newestBlockNumber - fromEvent.blockNumber

        // get number of blocks since event block number - required confirmations
        const blocksSinceConfirmed =
          blockDifference - BRIDGE_REQUIRED_CONFIRMATIONS[fromEvent?.chainId]

        // if blocks since confirmed is less than 0, thats how many blocks left to confirm
        setConfirmationDelta(
          blocksSinceConfirmed >= 0 ? 0 : blocksSinceConfirmed * -1
        )
        if (blocksSinceConfirmed >= 0) {
          setCompletedConf(true)
        }
      })
    }, [time])

    return (
      <>
        <div className="flex-1">
          <div className={`flex items-center align-middle`}>
            {fromEvent?.toChainId && confirmationDelta > 0 && (
              <>
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
        w-16 h-16   -mb-3
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
