import styles from './ticker.module.css'
import { generateTx } from '../../utils/fakeDataGen/teaserMarquee'
import PulseDot from './PulseDot'

const txs = new Array()
for (let i = 0; i < 20; i++) txs.push(generateTx())

const formatTimestamp = (tx) => {
  const { origin, destination } = tx
  const originDate = new Date(origin.timestamp)
  const originHour = originDate.getHours()
  const originMinute = originDate.getMinutes()
  const seconds = Math.round((destination.timestamp - origin.timestamp) / 1000)
  const minutes = Math.round(seconds / 60)
  const secondsModulo = (Math.round(seconds / 15) * 15) % 60

  const originDateFormatted = `${originHour}:${originMinute}${
    originHour < 12 ? 'am' : 'pm'
  }`
  const durationFormatted =
    minutes + 'm' + (secondsModulo ? ` ${secondsModulo}` : '')

  return `${originDateFormatted} (${durationFormatted})`
}

/* Ticker – Easter egg: define custom <marquee> element */

export default function Ticker() {
  return (
    <article
      className="absolute w-screen z-10 text-sm overflow-x-clip overflow-y-visible"
      style={{ counterReset: 'txn' }}
    >
      <dl
        className={`grid whitespace-nowrap list-disc marker:text-zinc-500 ${styles.ticker}`}
      >
        {txs.map((tx, i) => {
          return (
            <>
              <dt
                className="relative group row-start-1 bg-zinc-50 dark:bg-zinc-950 border-y border-zinc-200 dark:border-zinc-900"
                style={{ gridColumnStart: i + 1 }}
              >
                <a
                  href="#"
                  className="text-zinc-500 px-4 hover:text-inherit hover:underline py-1.5 inline-block"
                >
                  {`${tx.origin.formattedAmount} ${tx.origin.payload} to ${tx.destination.chain}`}
                </a>
              </dt>
              <dd
                className="row-start-2 animate-slide-down origin-top relative p-2"
                style={{ gridColumnStart: i + 1 }}
              >
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
