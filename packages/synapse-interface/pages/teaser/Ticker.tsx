import { generateTx } from '../../utils/fakeDataGen/teaserMarquee'
import PulseDot from './icons/PulseDot'
import { useRef, useEffect } from 'react'

const txs = new Array()
for (let i = 0; i < 6; i++) txs.push(generateTx())
console.log(txs)

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

  const originDateFormatted = `${originHour % 12}:${originMinute}`

  const destinationDateFormatted = `${
    destinationHour % 12
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

/* Ticker – Easter egg: define custom <marquee> element */

export default function Ticker() {
  const tickerRef = useRef(null)

  const sleep = (time) => new Promise((resolve) => setTimeout(resolve, time))

  let start, previousTimeStamp
  let done = false
  let offset = 0

  function step(timeStamp) {
    if (start === undefined) {
      start = timeStamp
    }
    const elapsed = timeStamp - start

    if (previousTimeStamp !== timeStamp) {
      const { left, width } = tickerRef.current.getBoundingClientRect()
      // Math.min() is used here to make sure the element stops at exactly 200px
      const count = -(elapsed * 0.03) + offset // 1px/frame @ 30fps
      console.log(Math.round(tickerRef.current.firstChild.offsetWidth))
      if (count + width < window.outerWidth) {
        offset += tickerRef.current.firstChild.offsetWidth
        console.log(tickerRef.current.firstChild.firstChild.innerText)
        var x = tickerRef.current.appendChild(tickerRef.current.firstChild) // <dt>
        var x = tickerRef.current.appendChild(tickerRef.current.firstChild) // <dd>
        // done = false
      }
      tickerRef.current.style.left = `${count}px`
    }

    if (true || elapsed < 2000) {
      // Stop the animation after 2 seconds
      // previousTimeStamp = timeStamp
      if (!done) {
        window.requestAnimationFrame(step)
      }
    }
  }

  useEffect(() => {
    window.requestAnimationFrame(step)
  })

  return (
    <article
      className="absolute w-screen z-10 text-sm overflow-x-clip overflow-y-visible"
      style={{ counterReset: 'txn' }}
    >
      <dl
        ref={tickerRef}
        className="grid grid-flow-col whitespace-nowrap list-disc marker:text-zinc-500 absolute"
      >
        {txs.map((tx, i) => {
          return (
            <>
              <dt className="relative group row-start-1 bg-zinc-50 dark:bg-zinc-950 border-y border-zinc-200 dark:border-zinc-900">
                <a
                  href="#"
                  className="text-zinc-500 px-4 hover:text-inherit hover:underline py-1.5 inline-block"
                >
                  {`${tx.origin.formattedAmount} ${tx.origin.payload} to ${tx.destination.chain}`}
                </a>
              </dt>
              <dd className="row-start-2 animate-slide-down origin-top relative p-2 hidden [:hover_+_&]:block hover:block">
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
      <button className="top-0 left-0 absolute bg-inherit px-4 py-1.5 border-r border-zinc-300 dark:border-zinc-800 flex items-center gap-2 z-10 bg-zinc-50 dark:bg-zinc-950">
        <PulseDot />
        <span className="md:after:content-['_–_All_transactions']">Live</span>
        <span className="text-xxs">▼</span>
      </button>
      <a
        href="#"
        className="absolute top-0 right-0 bg-inherit px-4 py-1.5 border-l border-zinc-300 dark:border-zinc-800 inline-block items-center gap-2 z-10 bg-zinc-50 dark:bg-zinc-950 md:before:content-['Explorer_']"
      >
        {'->'}
      </a>
    </article>
  )
}
