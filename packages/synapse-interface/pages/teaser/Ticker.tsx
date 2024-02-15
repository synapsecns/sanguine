import styles from './marquee.module.css'
import { generateTx } from './strings'

export default function Ticker() {
  return (
    <div className="absolute w-screen z-10 mb-12 bg-zinc-50 dark:bg-zinc-950 text-sm border-y border-zinc-200 dark:border-zinc-900 overflow-hidden">
      <button className="absolute bg-inherit px-4 py-1.5 border-r border-zinc-300 dark:border-zinc-800 flex items-center gap-2 z-10">
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
      <ul className={`flex whitespace-nowrap list-disc marker:text-zinc-500 ${styles.marquee}`}>
        {[...Array(20)].map((x, i) => {
          return <li className="pl-2 ml-7"><a href="#" className="text-zinc-500 hover:text-inherit hover:underline py-1.5 inline-block">{generateTx()}</a></li>
        })}
      </ul>
      <a href="#" className="absolute top-0 right-0 bg-inherit px-4 py-1.5 border-l border-zinc-300 dark:border-zinc-800 flex items-center gap-2 z-10">
        Explorer {'->'}
      </a>
    </div>
  )
}
