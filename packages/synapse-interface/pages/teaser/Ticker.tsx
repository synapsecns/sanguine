import styles from './ticker.module.css'
import { generateTx } from './strings'

const txs = new Array()
for (let i = 0; i < 20; i++) txs.push(generateTx())

export default function Ticker() {
  return (
    <article
      className="absolute w-screen z-10 mb-12 text-sm overflow-x-clip overflow-y-visible"
      style={{ counterReset: 'txn' }}
    >

      <dl className={`grid whitespace-nowrap list-disc marker:text-zinc-500 ${styles.ticker}`}>
        {txs.map((tx, i) => {
          return (
            <>
              <dt
                className={`relative group px-4 row-start-1 bg-zinc-50 dark:bg-zinc-950 border-y border-zinc-200 dark:border-zinc-900 ${styles.tx}`}
                style={{ gridColumnStart: i + 1}}
              >
                <a
                  href="#"
                  className="text-zinc-500 hover:text-inherit hover:underline py-1.5 inline-block"
                >
                  {`${tx.origin.formattedAmount} ${tx.origin.payload} to ${tx.destination.chain}`}
                </a>
              </dt>
              <dd
                className="row-start-2 animate-slide-down origin-top relative"
                style={{ gridColumnStart: i + 1}}
              >
                <div className="absolute top-2 left-2 px-3 py-2 bg-zinc-50 dark:bg-zinc-950 border border-zinc-200 dark:border-zinc-900 rounded items-center  flex gap-4">
                  <ul className="inline">
                    <li>{tx.origin.formattedAmount} {tx.origin.payload}</li>
                    <li>{tx.origin.chain}</li>
                    <li>{tx.origin.timestamp}</li>
                  </ul>
                  <svg width="6" height="12" viewBox="0 0 6 12" fill="none" stroke-width="2" stroke-linejoin="round" stroke-linecap="round" overflow="visible" className="stroke-zinc-500" xmlns="http://www.w3.org/2000/svg">
                    <path d="M0,0 6,6 0,12" />
                  </svg>
                  <ul className="inline">
                    <li>{tx.destination.formattedAmount} {tx.destination.payload}</li>
                    <li>{tx.destination.chain}</li>
                    <li>{tx.destination.timestamp}</li>
                  </ul>
                </div>
              </dd>
            </>
          )
        })}
      </dl>
      <button className="top-0 left-0 absolute bg-inherit px-4 py-1.5 border-r border-zinc-300 dark:border-zinc-800 flex items-center gap-2 z-10 bg-zinc-50 dark:bg-zinc-950">
        <svg width="8" height="8" viewBox="-4 -4 8 8" overflow="visible" className="fill-green-500 stroke-green-500" xmlns="http://www.w3.org/2000/svg">
          <circle r="4">
            <animate
              attributeName="stroke-width"
              values="0; 16"
              dur="1.5s"
              repeatCount="indefinite"
            />
            <animate
              attributeName="stroke-opacity"
              values=".5; 0"
              dur="1.5s"
              repeatCount="indefinite"
            />
          </circle>
        </svg>
        Live – All transactions <span className="text-xxs">▼</span>
      </button>
      <a href="#" className="absolute top-0 right-0 bg-inherit px-4 py-1.5 border-l border-zinc-300 dark:border-zinc-800 flex items-center gap-2 z-10 bg-zinc-50 dark:bg-zinc-950">
        Explorer {'->'}
      </a>
    </article>
  )
}
