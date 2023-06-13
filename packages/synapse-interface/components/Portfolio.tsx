import Card from '@/components/ui/tailwind/Card'
import { useAccount, useNetwork } from 'wagmi'
import { fetchBalance } from '@wagmi/core'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'
import { useEffect, useMemo, useState } from 'react'
import { formatBNToString } from '@/utils/bignumber/format'
import Image from 'next/image'
import { CHAINS_ARR } from '@/constants/chains'
import { PageHeader } from '@/components/PageHeader'
import { ConnectButton } from '@rainbow-me/rainbowkit'
import { AddressZero } from '@ethersproject/constants'

export const Portfolio = () => {
  const { address } = useAccount()
  const { chain } = useNetwork()

  const [allBalances, setAllBalances] = useState([])
  const [showZeroBalances, setShowZeroBalances] = useState(false)

  const activeChain = CHAINS_ARR.find((ch) => ch.id === chain.id)

  const bridgeableTokens = BRIDGABLE_TOKENS[chain.id]

  const filteredBalances = useMemo(() => {
    return showZeroBalances
      ? allBalances
      : allBalances.filter(
          (token) => Number(token.fetchedBalance.formatted) !== 0
        )
  }, [allBalances, showZeroBalances])

  // TODO: NEEDS TO HANDLE ETH too + gas tokens, right now only doing erc20s
  const fetchAllBalances = async () => {
    const amounts = []

    if (bridgeableTokens) {
      const promises = bridgeableTokens.map(async (token) => {
        let obj = {}

        if (token.addresses[chain.id] !== AddressZero) {
          const fetchedBalance = await fetchBalance({
            address: address,
            chainId: chain.id,
            token: token.addresses[chain.id] as `0x${string}`,
          })

          obj['token'] = token
          obj['fetchedBalance'] = fetchedBalance

          return obj
        }
      })

      const results = await Promise.all(promises)

      results.forEach((result) => {
        if (result) amounts.push(result)
      })
    }

    return amounts.sort(
      (a, b) =>
        Number(b.fetchedBalance.formatted) - Number(a.fetchedBalance.formatted)
    )
  }

  useEffect(() => {
    const fetchData = async () => {
      const b = await fetchAllBalances()
      setAllBalances(b)
    }

    fetchData()
  }, [chain])

  return (
    <div className="flex flex-col w-1/3">
      <div className="flex items-center justify-between">
        <PageHeader title="Portfolio" subtitle={shortenAddress(address)} />
        <ConnectWallet />
      </div>
      <Card
        divider={false}
        className={`
          max-w-lg px-1 pb-0 mt-3 overflow-hidden
          rounded-xl
          bg-transparent 
        text-white
        `}
      >
        <div className="flex items-center mb-4 space-x-2 text-2xl">
          <Image
            alt={`${activeChain.name} img`}
            className="w-8 h-8 rounded-md"
            src={activeChain.chainImg}
          />
          <div>{activeChain.name}</div>
        </div>
        {filteredBalances.map((balance, i) => {
          return <Balance balance={balance} chain={chain} key={i} />
        })}
        <button
          className="mt-2 text-sm hover:opacity-80"
          onClick={() => setShowZeroBalances(!showZeroBalances)}
        >
          {showZeroBalances ? 'Hide' : 'Show'} zero balances
        </button>
      </Card>
    </div>
  )
}

const Balance = ({ balance, chain }) => {
  let showBalance = formatBNToString(
    balance.fetchedBalance.value,
    balance.token.decimals[chain.id],
    3
  )

  if (showBalance === '0.0') {
    showBalance = '\u2212'
  }

  return (
    balance.token &&
    balance.fetchedBalance && (
      <div className="flex items-center justify-between">
        <div className="flex items-center space-x-1">
          <div className="relative flex p-1 rounded-full">
            <Image
              alt={`${balance.token.symbol} img`}
              className="w-6 h-6 rounded-md"
              src={balance.token.icon}
            />
          </div>
          <div>{balance.token.symbol}</div>
        </div>
        <div>{showBalance}</div>
      </div>
    )
  )
}

export const PortfolioPreview = () => {
  const [showContent, setShowContent] = useState('home')

  const renderCardContent = () => {
    if (showContent === 'home') {
      return (
        <div className="">
          <div className="flex items-center mb-4 space-x-2 text-lg">
            Synapse is the most widely used, extensible, and secure cross-chain
            communications network.
          </div>
          <div className="mb-5">
            Preview your route in the Bridge panel, and connect your wallet when
            you're ready to authorize your transaction.
          </div>
          <ConnectWallet />
        </div>
      )
    } else if (showContent === 'portfolio') {
      return (
        <div>
          <div className="flex items-center mb-5 space-x-2 text-lg">
            Your bridgable assets appear here when your wallet is connected.
          </div>
          <ConnectWallet />
        </div>
      )
    }
  }

  return (
    <div className="flex flex-col w-1/3">
      <div className="flex items-center space-x-2">
        <div
          onClick={() => setShowContent('home')}
          className="hover:opacity-80 hover:cursor-pointer"
        >
          <div className="flex items-center h-[32px]">
            <HomeSvg />
          </div>
          <div
            className={`mt-2 border-b-3 ${
              showContent === 'home'
                ? 'border-indigo-500'
                : 'border-transparent'
            }`}
          />
        </div>
        <div
          onClick={() => setShowContent('portfolio')}
          className="hover:opacity-80 hover:cursor-pointer"
        >
          <div className="flex items-center h-[32px]">
            <PageHeader title="Portfolio" subtitle="" />
          </div>
          <div
            className={`mt-2 border-b-3 ${
              showContent === 'portfolio'
                ? 'border-indigo-500'
                : 'border-transparent'
            }`}
          />
        </div>
      </div>
      <Card
        divider={false}
        className={`
          max-w-lg px-1 pb-0 mt-3 overflow-hidden
          rounded-xl
          bg-transparent 
        text-white
        `}
      >
        {renderCardContent()}
      </Card>
    </div>
  )
}

function ConnectWallet() {
  const buttonClassName = `
    h-10 border-[#AC8FFF] flex items-center border
    text-base px-4 py-3 hover:opacity-75 rounded-lg
    text-center transform-gpu transition-all duration-75
  `
  const buttonStyle = {
    background:
      'linear-gradient(310.65deg, rgba(255, 0, 255, 0.2) -17.9%, rgba(172, 143, 255, 0.2) 86.48%)',
    borderRadius: '16px',
  }

  return (
    <div>
      <ConnectButton.Custom>
        {({ account, chain, openConnectModal, mounted, openChainModal }) => {
          return (
            <>
              {(() => {
                if (!mounted || !account || !chain) {
                  return (
                    <button
                      className={buttonClassName}
                      style={buttonStyle}
                      onClick={openConnectModal}
                    >
                      Connect Wallet
                    </button>
                  )
                }
                return (
                  <button
                    className={buttonClassName}
                    style={buttonStyle}
                    onClick={openChainModal}
                  >
                    <div className="flex items-center gap-2 text-white">
                      Connected
                    </div>
                  </button>
                )
              })()}
            </>
          )
        }}
      </ConnectButton.Custom>
    </div>
  )
}

export function shortenAddress(address: string, chars = 6) {
  const start = address.slice(0, chars + 2)
  const end = address.slice(-chars)
  return `${start}...${end}`
}

const HomeSvg = () => {
  return (
    <svg
      width="24"
      height="25"
      viewBox="0 0 24 25"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        fill-rule="evenodd"
        clip-rule="evenodd"
        d="M20 10H4V21.5C4 22.0523 4.44772 22.5 5 22.5H9V17.5C9 16.9477 9.44772 16.5 10 16.5H14C14.5523 16.5 15 16.9477 15 17.5V22.5H19C19.5523 22.5 20 22.0523 20 21.5V10Z"
        fill="#D9D9D9"
      />
      <path
        d="M11.3598 3.03349C11.7307 2.72445 12.2693 2.72445 12.6402 3.03349L21.8781 10.7318C22.5967 11.3305 22.1732 12.5 21.238 12.5H2.76205C1.82675 12.5 1.40335 11.3305 2.12187 10.7318L11.3598 3.03349Z"
        fill="#D9D9D9"
      />
    </svg>
  )
}
