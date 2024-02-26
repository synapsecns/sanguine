import { Fragment, useRef, useEffect, useState } from 'react'
import { CHAINS_BY_ID } from '@/constants/chains'
import PulseDot from '../icons/PulseDot'

const testUrl =
  'https://explorer.omnirpc.io/graphql?query=%7B%0A%20%20bridgeTransactions(useMv%3Atrue%2C%20pending%3A%20false%2C%20startTime%3A1688815939%2C%20page%3A%201)%20%7B%0A%20%20%20%20fromInfo%20%7B%0A%20%20%20%20%20%20chainID%0A%20%20%20%20%20%20destinationChainID%0A%20%20%20%20%20%20address%0A%20%20%20%20%20%20txnHash%0A%20%20%20%20%20%20value%0A%20%20%20%20%20%20formattedValue%0A%20%20%20%20%20%20USDValue%0A%20%20%20%20%20%20tokenAddress%0A%20%20%20%20%20%20tokenSymbol%0A%20%20%20%20%20%20blockNumber%0A%20%20%20%20%20%20time%0A%20%20%20%20%20%20formattedTime%0A%20%20%20%20%20%20formattedEventType%0A%20%20%20%20%20%20eventType%0A%20%20%20%20%7D%0A%20%20%20%20toInfo%20%7B%0A%20%20%20%20%20%20chainID%0A%20%20%20%20%20%20destinationChainID%0A%20%20%20%20%20%20address%0A%20%20%20%20%20%20txnHash%0A%20%20%20%20%20%20value%0A%20%20%20%20%20%20formattedValue%0A%20%20%20%20%20%20USDValue%0A%20%20%20%20%20%20tokenAddress%0A%20%20%20%20%20%20tokenSymbol%0A%20%20%20%20%20%20blockNumber%0A%20%20%20%20%20%20time%0A%20%20%20%20%20%20formattedTime%0A%20%20%20%20%20%20formattedEventType%0A%20%20%20%20%20%20eventType%0A%20%20%20%20%7D%0A%20%20%20%20kappa%0A%20%20%7D%0A%7D%0A'

export default function Ticker() {
  const tickerRef = useRef(null)
  const [txData, setTxData] = useState([])

  const isLoading = !txData.length

  let start
  let requestId
  let pauseStart
  let accelStart
  let decelStart
  const ACCEL_TIME_MS = 400
  const DECEL_TIME_MS = 400
  const PX_PER_SECOND = -30 / 1000

  const step = (timestamp) => {
    if (start === undefined) start = timestamp

    if (pauseStart) {
      start += timestamp - pauseStart
      pauseStart = undefined
    }

    const dl = tickerRef.current
    if (dl.querySelector('dt')) {
      const { left, width } = dl.firstElementChild.getBoundingClientRect()

      if (left < -width) {
        start -= width / PX_PER_SECOND
        dl.appendChild(dl.firstElementChild) // <dt>
        dl.appendChild(dl.firstElementChild) // <dd>
      }

      dl.style.left = `${PX_PER_SECOND * (timestamp - start)}px`
    }

    requestId = window.requestAnimationFrame(step)
  }

  const decelerate = (timestamp) => {
    if (decelStart === undefined) {
      tickerRef.current.style.transform = `translateX(${0}rem)`
      decelStart = timestamp
    }

    const elapsed = timestamp - decelStart

    const dist = -0.375 * Math.log10(9 * (elapsed / DECEL_TIME_MS) + 1)

    tickerRef.current.style.transform = `translateX(${dist}rem)`

    elapsed < DECEL_TIME_MS
      ? window.requestAnimationFrame(decelerate)
      : (decelStart = undefined)
  }

  const accelerate = (timestamp) => {
    if (accelStart === undefined) {
      tickerRef.current.style.transform = `translateX(${-0.375}rem)`
      accelStart = timestamp
    }
    const elapsed = timestamp - accelStart

    const dist = -0.375 + 0.375 * Math.log10(9 * (elapsed / ACCEL_TIME_MS) + 1)

    tickerRef.current.style.transform = `translateX(${Math.min(dist, 0)}rem)`

    elapsed < ACCEL_TIME_MS
      ? window.requestAnimationFrame(accelerate)
      : (accelStart = undefined)
  }

  const startTicker = () => {
    requestId = window.requestAnimationFrame(step)
    window.requestAnimationFrame(accelerate)
  }
  const stopTicker = () => {
    pauseStart = performance.now()
    window.cancelAnimationFrame(requestId)
    window.requestAnimationFrame(decelerate)
  }

  useEffect(() => {
    tickerRef.current.addEventListener('mouseenter', stopTicker)
    tickerRef.current.addEventListener('mouseleave', startTicker)
    requestId = window.requestAnimationFrame(step)

    fetchExplorerTxs(setTxData)

    return () => {
      tickerRef.current.removeEventListener('mouseenter', stopTicker)
      tickerRef.current.removeEventListener('mouseleave', startTicker)
      window.cancelAnimationFrame(requestId)
    }
  }, [])

  /* Easter egg: define custom <marquee> element */

  return (
    <article
      className={`flex w-full z-10 text-sm bg-zinc-50 dark:bg-zinc-950 absolute border-b border-zinc-300 dark:border-zinc-900 overflow-x-clip justify-between cursor-default ${
        isLoading ? 'opacity-70' : 'opacity-100'
      }`}
    >
      <div className="bg-zinc-50 dark:bg-zinc-950 px-4 py-1.5 border-r border-zinc-300 dark:border-zinc-800 flex items-center gap-2 z-10 bg-zinc-50">
        <PulseDot
          className={
            isLoading
              ? 'fill-zinc-500 stroke-none'
              : 'fill-green-500 stroke-green-500'
          }
        />
        {isLoading ? (
          <>Loading…</>
        ) : (
          <>
            {/* <span className="md:after:content-['_–_All_transactions']"> */}
            Live
            {/* </span> */}
            {/* <span className="text-xxs">▼</span> */}
          </>
        )}
      </div>
      <dl
        ref={tickerRef}
        className={`relative grid grid-flow-col grid-rows-[1fr_0] w-0 grow cursor-pointer whitespace-nowrap ${
          isLoading ? 'opacity-0' : 'opacity-100'
        } transition-opacity`}
      >
        {txData
          .filter((tx) => tx.toInfo.time > tx.fromInfo.time)
          .map((tx) => (
            <Fragment key={tx.fromInfo.txnHash}>
              <dt className="row-start-1">
                <a
                  href="#"
                  className="text-zinc-500 px-4 pt-2 hover:text-inherit hover:underline block"
                >
                  {`${formatAmount(tx.fromInfo.formattedValue)} ${
                    tx.fromInfo.tokenSymbol
                  } to ${CHAINS_BY_ID[tx.toInfo.chainID]?.name}`}
                </a>
              </dt>
              <dd className="row-start-2 animate-slide-down origin-top p-2 hidden [:hover_+_&]:block hover:block">
                <a
                  href="#"
                  className="absolute px-3 py-2 bg-zinc-50 dark:bg-zinc-950 border border-zinc-200 dark:border-zinc-900 rounded items-center grid gap-x-4 gap-y-1 shadow-sm"
                >
                  <ul className="inline">
                    <li>
                      {formatAmount(tx.fromInfo.formattedValue)}{' '}
                      {tx.fromInfo.tokenSymbol}
                    </li>
                    <li>{CHAINS_BY_ID[tx.fromInfo.chainID]?.name}</li>
                  </ul>
                  <RightCaret height="12" />
                  <ul className="inline">
                    <li>
                      {formatAmount(tx.toInfo.formattedValue)}{' '}
                      {tx.toInfo.tokenSymbol}
                    </li>
                    <li>{CHAINS_BY_ID[tx.toInfo.chainID]?.name}</li>
                  </ul>
                  <header className="text-zinc-500 row-start-2 col-span-3">
                    {formatTimestamp(tx)}
                  </header>
                </a>
              </dd>
            </Fragment>
          ))}
      </dl>
      <a
        href="#"
        className="bg-inherit px-4 py-1.5 border border-transparent border-l-zinc-300 dark:border-l-zinc-800 inline-block items-center z-10 md:before:content-['Explorer_'] md:before:mr-2.5 flex items-center hover:border-zinc-400 hover:dark:border-zinc-600 hover:rounded"
      >
        <RightCaret height="10" />
      </a>
    </article>
  )
}

const RightCaret = ({ height }) => {
  const width = height / 2
  return (
    <svg
      width={width}
      height={height}
      viewBox={`0 0 ${width} ${height}`}
      fill="none"
      strokeWidth={height / 6}
      strokeLinejoin="round"
      strokeLinecap="round"
      overflow="visible"
      className="stroke-zinc-500"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path d={`M0,0 ${width},${width} 0,${height}`} />
    </svg>
  )
}

function fetchExplorerTxs(setState) {
  fetch(testUrl)
    .then((response) => {
      if (!response.ok)
        throw new Error(`HTTP error! Status: ${response.status}`)

      return response.blob()
    })
    .then(async (response) =>
      setState(JSON.parse(await response.text()).data.bridgeTransactions)
    )
}

const formatAmount = (amount) => {
  const MAX_DECIMALS = 4

  let [, left, right] =
    amount.toFixed(MAX_DECIMALS).match(/(\d+)\.?(\d*)/) ?? new Array(3).fill('')

  if (left === '0') {
    left = ''
  } else {
    for (let i = 3; i < left.length; i += 4)
      left = `${left.slice(0, left.length - i)},${left.slice(-i)}`
  }

  return left.length < MAX_DECIMALS
    ? `${left}.${right.slice(0, MAX_DECIMALS - left.length)}`
    : left
}

const formatTimestamp = (tx) => {
  const { fromInfo, toInfo } = tx

  /* From time */
  const fromDate = new Date(fromInfo.time)
  const fromHour = fromDate.getHours()
  const fromMinute =
    (fromDate.getMinutes() < 10 ? '0' : '') + fromDate.getMinutes()

  const fromDateFormatted = `${
    fromHour === 12 ? 12 : fromHour % 12
  }:${fromMinute}`

  /* To time */
  const toDate = new Date(toInfo.time)
  const toHour = toDate.getHours()
  const toMinute = (toDate.getMinutes() < 10 ? '0' : '') + toDate.getMinutes()

  const toDateFormatted = `${toHour === 12 ? 12 : toHour % 12}:${toMinute}${
    toHour < 12 ? 'am' : 'pm'
  }`

  /* Time range */
  const timeRange =
    fromHour === toHour && fromMinute === toMinute
      ? toDateFormatted
      : `${fromDateFormatted}–${toDateFormatted}`

  /* Elapsed */
  const seconds = toInfo.time - fromInfo.time
  const minutes = Math.round(seconds / 60)
  const roundedSeconds = (Math.round(seconds / 15) * 15) % 60

  const elapsedTime =
    minutes === 0
      ? (roundedSeconds === 0 ? seconds : roundedSeconds) + 's'
      : minutes + 'm' + (roundedSeconds ? ` ${roundedSeconds}` : '')

  return `${timeRange} (${elapsedTime})`
}
