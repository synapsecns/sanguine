// import PulseDot from '@/components/icons/PulseDot'
import { useTranslations } from 'next-intl'

import { CHAINS_ARR } from '@/constants/chains'
import * as BRIDGEABLE from '@constants/tokens/bridgeable'
import { TOKENS_SORTED_BY_SWAPABLETYPE } from '@/constants/tokens'
// import * as WALLET_ICONS from '@components/WalletIcons'

const cardStyle =
  'text-black dark:text-white bg-zinc-100 dark:bg-bgBase p-3 rounded-[.75rem] border border-zinc-200 dark:border-zinc-800 shadow-xl grid gap-4 max-w-sm' // dark:bg-zinc-900/95
const sectionStyle =
  'relative bg-zinc-50 dark:bg-bgLight rounded-md px-2.5 py-3 grid gap-1 grid-cols-2 border border-zinc-300 dark:border-transparent' // dark:bg-zinc-800
const buttonStyle =
  'rounded px-4 py-1 bg-zinc-100 dark:bg-separator border border-zinc-200 dark:border-transparent hover:border-zinc-400 hover:dark:border-zinc-500 h-fit mr-1 cursor-pointer focus:border-zinc-400 focus:dark:borer-zinc-500' // dark:bg-zinc-700
const chainSelectStyle =
  'flex gap-2 items-center rounded px-2 py-1 bg-inherit border border-zinc-200 dark:border-transparent hover:border-zinc-400 hover:dark:border-zinc-500 active:opacity-70 focus:ring-1 focus:ring-zinc-500 focus:border-transparent'
const tokenSelectStyle =
  'flex gap-2 items-center rounded px-2 py-1.5 bg-inherit dark:bg-separator border border-zinc-200 dark:border-transparent hover:border-zinc-400 hover:dark:border-zinc-500 active:opacity-70 focus:ring-1 focus:ring-zinc-500 focus:border-transparent text-lg ' // dark:bg-zinc-700
const inputWrapperStyle =
  'relative flex bg-white dark:bg-inherit border border-zinc-200 dark:border-separator rounded-md gap-0 p-1.5 col-span-2 gap-1.5 items-center' // dark:border-zinc-700
const inputStyle =
  'bg-inherit border-none w-full p-1.5 text-xxl font-normal dark:font-medium tracking-wide rounded placeholder:text-secondary pointer-events-none'

export default () => {
  return (
    <div className={cardStyle}>
      <section className={sectionStyle}>
        <Select
          type="Chain"
          label="From"
          defaultName="Ethereum"
          data="volume"
        />
        <SupportedWallets />
        <div className={inputWrapperStyle}>
          <Select type="Token" defaultName="USDC" data="volume" />
          <input type="text" placeholder="0" className={inputStyle} />
          <HistoricMax />
        </div>
      </section>
      <section className={sectionStyle}>
        <Select type="Chain" label="To" defaultName="Arbitrum" data="count" />
        <div className={inputWrapperStyle}>
          <Select type="Token" defaultName="USDC" data="count" />
          <input disabled readOnly type="text" className={inputStyle}></input>
        </div>
      </section>
      <BridgeButton />
    </div>
  )
}

const Select = ({
  type,
  defaultName,
  label,
  data,
}: {
  type: 'Chain' | 'Token'
  defaultName?: string
  label?: string
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

  const defaultItem = arr.find((a) => a[name] === defaultName)

  const t = useTranslations('FauxBridge')

  return (
    <div className="flex-none">
      <button
        className={`peer ml-0.5 w-fit whitespace-nowrap ${
          type === 'Chain' ? chainSelectStyle : tokenSelectStyle
        }`}
      >
        <img width="20" height="20" src={defaultItem[img].src} />
        <span>
          {label && (
            <div className="text-sm text-left text-secondary">{label}</div>
          )}
          {defaultName ?? type}
        </span>{' '}
        <span className="text-xs text-secondary">{t('▼')}</span>
      </button>
      <div className="absolute pt-1.5 hidden hover:block peer-hover:block peer-focus:block animate-slide-down origin-top-left z-10 whitespace-nowrap">
        <table className="rounded border-separate border-spacing-0 shadow bg-white dark:bg-bgBase backdrop-blur-sm border border-zinc-200 dark:border-zinc-700 block max-h-[372px] overflow-y-scroll">
          <thead className="text-sm cursor-default">
            <tr>
              <th className="sticky top-0 px-4 pt-3 pb-2 font-normal text-left border-b bg-white/90 dark:bg-bgBase backdrop-blur-sm border-zinc-300 dark:border-zinc-800 text-secondary">
                {type}
              </th>
              {/* <th className="sticky top-0 px-4 pt-3 pb-2 font-normal text-right border-b bg-white/90 dark:bg-bgBase backdrop-blur-sm border-zinc-300 dark:border-zinc-800 text-zinc-500">
                {header}
              </th> */}
            </tr>
          </thead>
          <tbody>
            {arr.map((item, i) => {
              reduce()
              return (
                <tr
                  key={item[key] === 'Wrapped ETH' ? item.symbol : item[key]}
                  className="cursor-default hover:bg-zinc-100 hover:dark:bg-bgLight"
                >
                  <td className="px-4 py-2.5 text-left flex items-center gap-2">
                    <img width="16" height="16" src={item[img].src} />
                    {item[name]}
                  </td>
                  {/* <td className="px-4 py-2 text-right">{format()}</td> */}
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

const SupportedWallets = () => {
  const t = useTranslations('FauxBridge')

  return (
    <div className="relative self-center group justify-self-end">
      <div className="absolute right-0 items-center hidden gap-2 px-3 py-2 mb-2 origin-bottom-right bg-white border rounded shadow bottom-full dark:bg-zinc-900/50 backdrop-blur-sm border-zinc-200 dark:border-zinc-800 group-hover:grid whitespace-nowrap animate-slide-up">
        <header className="text-sm text-zinc-500">
          {t('Supported wallets')}
        </header>
        <ul className="flex gap-4">
          {/* {Object.values(WALLET_ICONS).map((icon, i) => (
          <li key={i}>{icon({ width: 24, height: 24 })}</li>
        ))} */}
        </ul>
      </div>
      <button className="peer flex gap-2.5 items-center justify-self-end text-sm text-zinc-700 dark:text-zinc-300 mr-1 cursor-pointer">
        {/* <PulseDot className="fill-green-500 stroke-green-500" /> Connected */}
      </button>
    </div>
  )
}

const HistoricMax = () => {
  const t = useTranslations('FauxBridge')
  return (
    <div className="relative pointer-events-none group">
      <div className="absolute right-0 hidden mb-2 origin-bottom-right bottom-full group-hover:block animate-slide-up">
        <a
          href="https://ftmscan.com/tx/0x18199d88fe9fc8baa0f0f02d216c0e7998e1e59aaef6e0ea7a7a35d8dd6bc90b"
          target="_blank"
          className="grid items-center px-3 py-2 text-sm bg-white border rounded shadow dark:bg-zinc-900/95 backdrop-blur-sm border-zinc-200 dark:border-zinc-800 gap-x-4 gap-y-1 whitespace-nowrap"
        >
          <ul>
            <li>{t('40,668 ETH')}</li>
            <li>{t('Fantom')}</li>
          </ul>
          <RightAngle height="12" />
          <ul>
            <li>{t('40,668 ETH')}</li>
            <li>{t('Ethereum')}</li>
          </ul>
          <header className="col-span-3 text-zinc-500">
            {t('Jan 29, 2022 - #1')}
          </header>
        </a>
      </div>
      <button className={`peer ${buttonStyle}`}>{t('Max')}</button>
    </div>
  )
}

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

const BridgeButton = () => {
  const t = useTranslations('FauxBridge')

  return (
    <div className="relative group">
      <div className="absolute hidden w-full mb-1 origin-bottom bottom-full group-hover:block animate-slide-up">
        <div className="px-3 py-1 m-auto text-sm bg-white border rounded shadow text-zinc-500 dark:text-zinc-400 w-fit dark:bg-zinc-900/95 border-zinc-200 dark:border-zinc-800">
          {t('Enter Bridge')}
        </div>
      </div>
      <a
        href="/"
        className="border border-fuchsia-500 py-2.5 pl-2 block rounded tracking-wider text-center hover:bg-purple-50 hover:dark:bg-fuchsia-950"
        onMouseEnter={(e) => {
          const target = e.target as HTMLAnchorElement
          target.querySelector('animate')?.beginElement()
        }}
      >
        {t('Bridge')}
        <svg
          width="10"
          height="10"
          viewBox="0 -8 16 16"
          overflow="visible"
          strokeWidth="4"
          fill="none"
          preserveAspectRatio="xMaxYMid"
          className="inline ml-1.5 mb-0.5 stroke-zinc-800 dark:stroke-zinc-200"
          xmlns="http://www.w3.org/2000/svg"
        >
          <animate
            attributeName="width"
            values="10; 14; 10"
            dur=".4s"
            calcMode="spline"
            keySplines="0 0 0 1; .5 0 0 1"
          />
          <path d="m16,0 -16,0 m8,-8 8,8 -8,8" />
        </svg>
      </a>
    </div>
  )
}
