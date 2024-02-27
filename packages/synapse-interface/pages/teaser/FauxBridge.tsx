import PulseDot from '@/components/icons/PulseDot'
import { CHAINS_ARR } from '@/constants/chains'
import * as BRIDGEABLE from '@constants/tokens/bridgeable'

const cardStyle =
  'text-black dark:text-white bg-zinc-100 dark:bg-zinc-900 p-3 rounded-md border border-zinc-200 dark:border-zinc-800 shadow-xl grid gap-4 max-w-sm'
const sectionStyle =
  'bg-zinc-50 dark:bg-zinc-800 rounded-md px-2.5 py-3 grid gap-3 grid-cols-2 border border-zinc-300 dark:border-transparent'
const selectStyle =
  'rounded bg-inherit dark:bg-zinc-700 border-zinc-300 dark:border-zinc-700 w-fit cursor-pointer hover:border-zinc-400 dark:hover:border-zinc-500'
const buttonStyle =
  'rounded px-4 py-1 bg-zinc-100 dark:bg-zinc-700 border border-zinc-200 dark:border-transparent hover:border-zinc-400 hover:dark:border-zinc-500 h-fit mr-1 cursor-pointer'
const buttonSelectStyle =
  'flex gap-1.5 items-center rounded px-3 py-1.5 bg-inherit dark:bg-zinc-700 border border-zinc-200 dark:border-transparent hover:border-zinc-400 hover:dark:border-zinc-500'
const inputWrapperStyle =
  'flex bg-white dark:bg-inherit border border-zinc-200 dark:border-zinc-700 rounded-md gap-0 p-1.5 col-span-2 gap-1.5 items-center'
const inputStyle =
  'bg-inherit border-none w-full p-1.5 text-xxl font-normal dark:font-light tracking-wide rounded'

const FauxBridge = () => {
  return (
    <div className={cardStyle}>
      <section className={sectionStyle}>
        <ChainList type="Chain" data="volume" />
        <div className="flex gap-2.5 items-center justify-self-end text-sm text-zinc-700 dark:text-zinc-300 mr-1 cursor-default">
          <PulseDot className="fill-green-500 stroke-green-500" /> Connected
        </div>
        <div className={inputWrapperStyle}>
          <ChainList type="Token" data="volume" />
          <input type="text" value="1000" className={inputStyle} />
          <button disabled className={buttonStyle}>
            Max
          </button>
        </div>
      </section>
      <section className={sectionStyle}>
        <ChainList type="Chain" data="count" />
        <div className={inputWrapperStyle}>
          <ChainList type="Token" data="count" />
          <input
            disabled
            type="text"
            value="1,000"
            className={inputStyle}
          ></input>
        </div>
      </section>
      <a
        href="#"
        className="border border-fuchsia-500 py-2.5 pl-2 rounded text-lg tracking-wider text-center hover:bg-purple-50 hover:dark:bg-fuchsia-950"
        onMouseEnter={(e) => {
          const target = e.target as HTMLAnchorElement
          target.querySelector('animate')?.beginElement()
        }}
      >
        Bridge
        <svg
          width="12"
          height="13"
          viewBox="0 -8 16 16"
          overflow="visible"
          strokeWidth="4"
          fill="none"
          preserveAspectRatio="xMaxYMid"
          className="inline ml-2 mb-1 stroke-zinc-800 dark:stroke-zinc-200"
          xmlns="http://www.w3.org/2000/svg"
        >
          <animate
            attributeName="width"
            values="12; 18; 12"
            dur=".5s"
            calcMode="spline"
            keySplines="0 0 0 1; .5 0 0 1"
          />
          <path d="m16,0 -16,0 m8,-8 8,8 -8,8" />
        </svg>
      </a>
    </div>
  )
}

export default FauxBridge

const ChainList = ({
  type,
  data,
}: {
  type: 'Chain' | 'Token'
  data: 'volume' | 'count'
}) => {
  let button: string
  let header: string
  let value: number
  let reduce: Function
  let format: Function
  switch (data) {
    case 'volume':
      button = `Volume by ${type}`
      header = '$ vol.'
      value = 1000000000 + Math.random() * 100000000
      reduce = () => (value *= 0.75)
      format = () => '$' + (value / 1000000).toFixed(1) + 'M'
      break
    case 'count':
      button = `Txns by ${type}`
      header = 'Txns'
      value = 10000 + Math.random() * 1000
      reduce = () => (value *= 0.9)
      format = () => {
        let str = value.toFixed()
        for (let i = 3; i < str.length; i += 4)
          str = `${str.slice(0, str.length - i)},${str.slice(-i)}`
        return str
      }
      break
  }

  let arr
  let key: string
  let img: string
  let name: string

  switch (type) {
    case 'Chain':
      arr = CHAINS_ARR
      key = 'id'
      img = 'chainImg'
      name = 'name'
      break
    case 'Token':
      arr = Object.values(BRIDGEABLE)
      key = 'symbol'
      img = 'icon'
      name = 'symbol'
      break
  }

  return (
    <div className="group relative">
      <button className={`ml-0.5 w-fit whitespace-nowrap ${buttonSelectStyle}`}>
        {type} <span className="text-xxs">▼</span>
      </button>
      <div className="pt-2 hidden group-hover:block absolute animate-slide-down origin-top w-max z-10 group-active:hidden">
        <table className="relative bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-700 whitespace-nowrap rounded border-separate border-spacing-0 max-h-[372px] overflow-y-scroll block shadow">
          <thead className="text-sm cursor-default">
            <tr>
              <th className="font-normal px-4 pt-3 pb-2 text-left sticky top-0 bg-white dark:bg-zinc-900/90 backdrop-blur-sm shadow border-b border-zinc-800 text-zinc-500">
                {type}
              </th>
              <th className="font-normal px-4 pt-3 pb-2 text-right sticky top-0 bg-white dark:bg-zinc-900/90 backdrop-blur-sm shadow border-b border-zinc-800 text-zinc-500">
                {header}
              </th>
            </tr>
          </thead>
          <tbody>
            {arr.map((item) => {
              reduce()
              return (
                <tr
                  key={item[key]}
                  className="hover:bg-zinc-50 hover:dark:bg-zinc-800 cursor-pointer"
                >
                  <td className="px-4 py-2.5 text-left flex items-center gap-2">
                    <img width="16" height="16" src={item[img].src} />
                    {item[name]}
                  </td>
                  <td className="px-4 py-2 text-right">{format()}</td>
                </tr>
              )
            })}
          </tbody>
          <tfoot />
        </table>
      </div>
    </div>
  )
}
