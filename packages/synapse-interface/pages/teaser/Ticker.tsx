import { generateTx } from '../../utils/fakeDataGen/teaserMarquee'
import PulseDot from './icons/PulseDot'
import { useRef, useEffect, useState } from 'react'

const txs = new Array()
for (let i = 0; i < 6; i++) txs.push(generateTx())

const formatTimestamp = (tx) => {
  const { origin, destination } = tx

  const originDate = new Date(origin.timestamp)
  const originHour = originDate.getHours()
  const originMinute =
    (originDate.getMinutes() < 10 ? '0' : '') + originDate.getMinutes()

  const destinationDate = new Date(destination.timestamp)
  const destinationHour = destinationDate.getHours()
  const destinationMinute =
    (destinationDate.getMinutes() < 10 ? '0' : '') +
    destinationDate.getMinutes()

  const seconds = Math.round((destination.timestamp - origin.timestamp) / 1000)
  const minutes = Math.round(seconds / 60)
  const secondsModulo = (Math.round(seconds / 15) * 15) % 60

  const originDateFormatted = `${
    originHour === 12 ? 12 : originHour % 12
  }:${originMinute}`

  const destinationDateFormatted = `${
    destinationHour === 12 ? 12 : destinationHour % 12
  }:${destinationMinute}${destinationHour < 12 ? 'am' : 'pm'}`

  const durationFormatted =
    minutes === 0
      ? (secondsModulo === 0 ? seconds : secondsModulo) + 's'
      : minutes + 'm' + (secondsModulo ? ` ${secondsModulo}` : '')

  const timeRange =
    originHour === destinationHour && originMinute === destinationMinute
      ? destinationDateFormatted
      : `${originDateFormatted}–${destinationDateFormatted}`

  return `${timeRange} (${durationFormatted})`
}

export default function Ticker() {
  const tickerRef = useRef(null)
  let left = 0

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
    const { left, width } = dl.firstElementChild.getBoundingClientRect()

    if (left < -width) {
      start -= width / PX_PER_SECOND
      dl.appendChild(dl.firstElementChild) // <dt>
      dl.appendChild(dl.firstElementChild) // <dd>
    }

    dl.style.left = `${PX_PER_SECOND * (timestamp - start)}px`
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

    return () => {
      tickerRef.current.removeEventListener('mouseenter', stopTicker)
      tickerRef.current.removeEventListener('mouseleave', startTicker)
      window.cancelAnimationFrame(requestId)
    }
  }, [])

  /* Ticker – Easter egg: define custom <marquee> element */

  return (
    <article className="flex w-full z-10 text-sm bg-zinc-50 dark:bg-zinc-950 absolute border-b border-zinc-300 dark:border-zinc-900 overflow-x-clip">
      <button className="bg-zinc-50 dark:bg-zinc-950 px-4 py-1.5 border-r border-zinc-300 dark:border-zinc-800 flex items-center gap-2 z-10 bg-zinc-50">
        <PulseDot />
        <span className="md:after:content-['_–_All_transactions']">Live</span>
        <span className="text-xxs">▼</span>
      </button>
      <dl
        ref={tickerRef}
        className="relative grid grid-flow-col grid-rows-[1fr_0] w-0 grow cursor-pointer whitespace-nowrap"
      >
        {txs.map((tx, i) => {
          return (
            <>
              <dt className="row-start-1">
                <a
                  href="#"
                  className="text-zinc-500 px-4 hover:text-inherit hover:underline py-1.5 block"
                >
                  {`${tx.origin.formattedAmount} ${tx.origin.payload} to ${tx.destination.chain}`}
                </a>
              </dt>
              <dd className="row-start-2 animate-slide-down origin-top p-2 hidden [:hover_+_&]:block hover:block">
                <a
                  href="#"
                  className="absolute px-3 py-2 bg-zinc-50 dark:bg-zinc-950 border border-zinc-200 dark:border-zinc-900 rounded items-center grid gap-x-4 gap-y-1 shadow-sm"
                >
                  <ul className="inline">
                    <li>
                      {tx.origin.formattedAmount} {tx.origin.payload}
                    </li>
                    <li>{tx.origin.chain}</li>
                  </ul>
                  <svg
                    width="6"
                    height="12"
                    viewBox="0 0 6 12"
                    fill="none"
                    stroke-width="2"
                    stroke-linejoin="round"
                    stroke-linecap="round"
                    overflow="visible"
                    className="stroke-zinc-500"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path d="M0,0 6,6 0,12" />
                  </svg>
                  <ul className="inline">
                    <li>
                      {tx.destination.formattedAmount} {tx.destination.payload}
                    </li>
                    <li>{tx.destination.chain}</li>
                  </ul>
                  <header className="text-zinc-500 row-start-2 col-span-3">
                    {formatTimestamp(tx)}
                  </header>
                </a>
              </dd>
            </>
          )
        })}
      </dl>
      <a
        href="#"
        className="bg-inherit px-4 py-1.5 border-l border-zinc-300 dark:border-zinc-800 inline-block items-center gap-2 z-10 md:before:content-['Explorer_']"
      >
        {'->'}
      </a>
    </article>
  )
}
