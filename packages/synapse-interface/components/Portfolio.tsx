import Card from '@/components/ui/tailwind/Card'
import { useAccount, useProvider, useNetwork } from 'wagmi'
import { Address, erc20ABI, fetchBalance, fetchSigner } from '@wagmi/core'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'
import { useEffect, useMemo, useState } from 'react'
import { formatBNToString } from '@/utils/bignumber/format'
import Image from 'next/image'
import { CHAINS_ARR } from '@/constants/chains'
import { PageHeader } from '@/components/PageHeader'
import { ConnectButton } from '@rainbow-me/rainbowkit'
import { AddressZero, Zero } from '@ethersproject/constants'
import { approveToken } from '@/utils/approveToken'
import { getAccount } from '@wagmi/core'
import { Contract } from 'ethers'
import { getContract } from '@wagmi/core'

const ROUTER_ADDRESS = '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a'

const buttonStyle = {
  background:
    'linear-gradient(310.65deg, rgba(255, 0, 255, 0.2) -17.9%, rgba(172, 143, 255, 0.2) 86.48%)',
  borderRadius: '16px',
}

const fetchBalanceAndAllownaces = async (token, activeChain, address) => {
  let obj = {}

  try {
    if (token.addresses[activeChain.id] !== AddressZero) {
      const fetchedBalance = await fetchBalance({
        address: address,
        chainId: activeChain.id,
        token: token.addresses[activeChain.id] as Address,
      })

      obj['token'] = token
      obj['fetchedBalance'] = fetchedBalance

      try {
        const a = await getTokenAllowance(
          address,
          activeChain.id,
          token.addresses[activeChain.id],
          ROUTER_ADDRESS
        )

        obj['allowance'] = a
      } catch (error) {
        console.log(`[getTokenAllowance] error`, error)
        obj['allowance'] = Zero
      }

      return obj
    }
  } catch (error) {
    console.log(`error, `, error)
    return {
      token: token,
      fetchedBalance: null,
      allowance: null,
    }
  }
}

export const Portfolio = () => {
  const { address } = useAccount()
  const { chain } = useNetwork()
  const [isLoading, setIsLoading] = useState(false)

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
        return await fetchBalanceAndAllownaces(token, activeChain, address)
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
    setIsLoading(true)
    const fetchData = async () => {
      const b = await fetchAllBalances()
      setAllBalances(b)
    }

    fetchData()
    setIsLoading(false)
  }, [chain])

  return (
    <div className="flex flex-col w-1/3">
      <PageHeader title="Portfolio" subtitle={shortenAddress(address)} />
      <Card
        divider={false}
        className={`
          max-w-lg px-1 pb-0 mt-3 overflow-hidden
          rounded-xl
          bg-transparent
        text-white
        `}
      >
        <div className="flex items-center justify-between mb-4 text-lg">
          <div className="flex items-center space-x-2">
            <Image
              alt={`${activeChain.name} img`}
              className="rounded-md w-7 h-7"
              src={activeChain.chainImg}
            />
            <div>{activeChain.name}</div>
          </div>
          <div className="space-x-2 text-sm">
            <span className="text-green-500">●</span> Live
          </div>
        </div>
        {!isLoading &&
          filteredBalances.map((balance, i) => {
            return (
              <Balance balance={balance} chainId={activeChain.id} key={i} />
            )
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

const Balance = ({ balance, chainId }) => {
  const allowance = balance.allowance
  const account = getAccount()
  const address = account?.address

  let showBalance = formatBNToString(
    balance.fetchedBalance.value,
    balance.token.decimals[chainId],
    3
  )

  if (showBalance === '0.0') {
    showBalance = '\u2212'
  }

  const handleApprove = () => {
    approveToken(
      address,
      chainId,
      balance.token.addresses[chainId],
      balance.fetchedBalance.value
    )
      .then((res) => console.log(`res`, res))
      .catch((err) => console.log(`err`, err))
  }

  return (
    balance.token &&
    balance.fetchedBalance && (
      <div className="flex items-center justify-between h-10">
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
        <div className="flex items-center space-x-2">
          <div>{showBalance}</div>
          {allowance && allowance.gte(balance.fetchedBalance.value) ? (
            <button
              className={`
                h-8 border-[#AC8FFF] flex items-center border
                text-sm px-3 py-2 opacity-75 rounded-lg
                text-center transform-gpu transition-all duration-75
              `}
              disabled={true}
              style={buttonStyle}
            >
              Approved
            </button>
          ) : (
            <button
              className={`
                h-8 border-[#AC8FFF] flex items-center border
                text-sm px-3 py-2 hover:opacity-75 rounded-lg
                text-center transform-gpu transition-all duration-75
              `}
              style={buttonStyle}
              onClick={handleApprove}
            >
              Approve
            </button>
          )}
        </div>
      </div>
    )
  )
}

const getTokenAllowance = async (
  address: string,
  chainId: number,
  tokenAddress: string,
  routerAddress: string
) => {
  let wallet = await fetchSigner({
    chainId: 43114,
  })

  // if (!wallet) {
  //   wallet = await fetchSigner({
  //     chainId,
  //   })
  // }

  console.log(`wallet, wallet`, wallet)

  const erc20 = new Contract(tokenAddress, erc20ABI, wallet)

  const allowance = await erc20.allowance(address, routerAddress)
  return allowance
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
                // return (
                //   <button
                //     className={buttonClassName}
                //     style={buttonStyle}
                //     onClick={openChainModal}
                //   >
                //     <div className="flex items-center gap-2 text-white">
                //       Connected
                //     </div>
                //   </button>
                // )
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

// function where you pass it a token and it returns an allowance
// { needsAllowance: true, amount: BigNumber, }
