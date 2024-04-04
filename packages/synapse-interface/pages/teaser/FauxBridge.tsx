import PulseDot from '@/components/icons/PulseDot'
import { CHAINS_ARR } from '@/constants/chains'
import * as BRIDGEABLE from '@constants/tokens/bridgeable'
import { TOKENS_SORTED_BY_SWAPABLETYPE } from '@/constants/tokens'
import * as WALLET_ICONS from '@components/WalletIcons'

const cardStyle =
  'text-black dark:text-white bg-zinc-100 dark:bg-zinc-900/95 p-3 rounded-md border border-zinc-200 dark:border-zinc-800 shadow-xl grid gap-4 max-w-sm'
const sectionStyle =
  'relative bg-zinc-50 dark:bg-zinc-800 rounded-md px-2.5 py-3 grid gap-3 grid-cols-2 border border-zinc-300 dark:border-transparent'
const buttonStyle =
  'rounded px-4 py-1 bg-zinc-100 dark:bg-zinc-700 border border-zinc-200 dark:border-transparent hover:border-zinc-400 hover:dark:border-zinc-500 h-fit mr-1 cursor-pointer focus:border-zinc-400 focus:dark:borer-zinc-500'
const buttonSelectStyle =
  'flex gap-1.5 items-center rounded px-3 py-1.5 bg-inherit dark:bg-zinc-700 border border-zinc-200 dark:border-transparent hover:border-zinc-400 hover:dark:border-zinc-500 active:opacity-70 focus:ring-1 focus:ring-zinc-500 focus:border-transparent'
const inputWrapperStyle =
  'relative flex bg-white dark:bg-inherit border border-zinc-200 dark:border-zinc-700 rounded-md gap-0 p-1.5 col-span-2 gap-1.5 items-center'
const inputStyle =
  'bg-inherit border-none w-full p-1.5 text-xxl font-normal dark:font-light tracking-wide rounded'

export default () => {
  return (
    <div className={cardStyle}>
      <section className={sectionStyle}>
        <Select type="Chain" data="volume" />
        <SupportedWallets />
        <div className={inputWrapperStyle}>
          <Select type="Token" data="volume" />
          <input type="text" placeholder="1000" className={inputStyle} />
          <HistoricMax />
        </div>
      </section>
      <section className={sectionStyle}>
        <Select type="Chain" data="count" />
        <div className={inputWrapperStyle}>
          <Select type="Token" data="count" />
          <input disabled readOnly type="text" className={inputStyle}></input>
        </div>
      </section>
      <BridgeButton />
    </div>
  )
}

const Select = ({
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
      reduce = () => (value *= 0.85)
      format = () => {
        if (value >= 1000000) return '$' + (value / 1000000).toFixed(1) + 'M'
        let str = value.toFixed(0)
        if (value >= 1000) {
          for (let i = 3; i < str.length; i += 4)
            str = `${str.slice(0, str.length - i)},${str.slice(-i)}`
          return '$' + str
        }
        return '$' + value.toFixed(2)
      }
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
      key = 'name'
      img = 'icon'
      name = 'symbol'
      break
  }

  return (
    <div>
      <button
        className={`peer ml-0.5 w-fit whitespace-nowrap ${buttonSelectStyle}`}
      >
        {type} <span className="text-xxs">▼</span>
      </button>
      <div className="absolute pt-1.5 hidden hover:block peer-hover:block peer-focus:block animate-slide-down origin-top-left z-10 whitespace-nowrap">
        <table className="rounded border-separate border-spacing-0 shadow bg-white dark:bg-zinc-900/95 backdrop-blur-sm border border-zinc-200 dark:border-zinc-700 block max-h-[372px] overflow-y-scroll">
          <thead className="text-sm cursor-default">
            <tr>
              <th className="font-normal px-4 pt-3 pb-2 sticky top-0 bg-white/90 dark:bg-zinc-900/95 backdrop-blur-sm border-b border-zinc-300 dark:border-zinc-800 text-zinc-500 text-left ">
                {type}
              </th>
              <th className="font-normal px-4 pt-3 pb-2 sticky top-0 bg-white/90 dark:bg-zinc-900/95 backdrop-blur-sm border-b border-zinc-300 dark:border-zinc-800 text-zinc-500 text-right ">
                {header}
              </th>
            </tr>
          </thead>
          <tbody>
            {arr.map((item, i) => {
              reduce()
              return (
                <tr
                  key={item[key] === 'Wrapped ETH' ? item.symbol : item[key]}
                  className="hover:bg-zinc-100 hover:dark:bg-zinc-800 cursor-default"
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

const SupportedWallets = () => (
  <div className="group relative self-center justify-self-end">
    <div className="absolute mb-2 bottom-full right-0 px-3 py-2 bg-white dark:bg-zinc-900/50 backdrop-blur-sm border border-zinc-200 dark:border-zinc-800 rounded items-center hidden group-hover:grid gap-2 shadow whitespace-nowrap animate-slide-up origin-bottom-right">
      <header className="text-sm text-zinc-500">Supported wallets</header>
      <ul className="flex gap-4">
        {Object.values(WALLET_ICONS).map((icon, i) => (
          <li key={i}>{icon({ width: 24, height: 24 })}</li>
        ))}
      </ul>
    </div>
    <button className="peer flex gap-2.5 items-center justify-self-end text-sm text-zinc-700 dark:text-zinc-300 mr-1 cursor-pointer">
      <PulseDot className="fill-green-500 stroke-green-500" /> Connected
    </button>
  </div>
)

const HistoricMax = () => (
  <div className="group relative">
    <div className="absolute mb-2 bottom-full right-0 hidden group-hover:block animate-slide-up origin-bottom-right">
      <a
        href="https://ftmscan.com/tx/0x18199d88fe9fc8baa0f0f02d216c0e7998e1e59aaef6e0ea7a7a35d8dd6bc90b"
        target="_blank"
        className="px-3 py-2 bg-white dark:bg-zinc-900/95 backdrop-blur-sm border border-zinc-200 dark:border-zinc-800 rounded items-center grid gap-x-4 gap-y-1 shadow whitespace-nowrap text-sm"
      >
        <ul>
          <li>40,668 ETH</li>
          <li>Fantom</li>
        </ul>
        <RightAngle height="12" />
        <ul>
          <li>40,668 ETH</li>
          <li>Ethereum</li>
        </ul>
        <header className="text-zinc-500 col-span-3">Jan 29, 2022 – #1</header>
      </a>
    </div>
    <button className={`peer ${buttonStyle}`}>Max</button>
  </div>
)

const RightAngle = ({ height }) => {
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

const BridgeButton = () => (
  <div className="group relative">
    <div className="absolute mb-1 w-full bottom-full hidden group-hover:block animate-slide-up origin-bottom">
      <div className="text-sm text-zinc-500 dark:text-zinc-400 w-fit m-auto px-3 py-1 bg-white dark:bg-zinc-900/95 border border-zinc-200 dark:border-zinc-800 rounded shadow">
        Visit Bridge
      </div>
    </div>
    <a
      href="#"
      className="border border-fuchsia-500 py-2.5 pl-2 block rounded text-lg tracking-wider text-center hover:bg-purple-50 hover:dark:bg-fuchsia-950"
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