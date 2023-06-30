import {
  usePortfolioBalancesAndAllowances,
  NetworkTokenBalancesAndAllowances,
  TokenWithBalanceAndAllowance,
} from '@/utils/hooks/usePortfolioBalances'
import { PortfolioTabManager } from './PortfolioTabManager'
import { useAccount, useNetwork, Address } from 'wagmi'
import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'
import { useState, useMemo } from 'react'
import { ConnectWalletButton } from './ConnectWalletButton'
import { CHAINS_BY_ID } from '@/constants/chains'
import HomeSvg from '../icons/HomeIcon'
import { Chain } from '@/utils/types'
import Image from 'next/image'
import { switchNetwork } from '@wagmi/core'

export enum PortfolioTabs {
  HOME = 'home',
  PORTFOLIO = 'portfolio',
}

function filterPortfolioBalancesWithBalances(
  balancesAndAllowances: NetworkTokenBalancesAndAllowances
): NetworkTokenBalancesAndAllowances {
  const filteredBalances: NetworkTokenBalancesAndAllowances = {}

  Object.entries(balancesAndAllowances).forEach(([key, tokenWithBalances]) => {
    const filteredTokenWithBalances = tokenWithBalances.filter(
      (token: TokenWithBalanceAndAllowance) => token.balance > Zero
    )

    if (filteredTokenWithBalances.length > 0) {
      filteredBalances[key] = filteredTokenWithBalances
    }
  })

  return filteredBalances
}

export const Portfolio = () => {
  const [tab, setTab] = useState<PortfolioTabs>(PortfolioTabs.HOME)

  const portfolioData: NetworkTokenBalancesAndAllowances =
    usePortfolioBalancesAndAllowances()

  const filteredPortfolioDataForBalances: NetworkTokenBalancesAndAllowances =
    filterPortfolioBalancesWithBalances(portfolioData)

  const { address } = useAccount()
  const { chain } = useNetwork()

  return (
    <div className="flex flex-col w-1/2">
      <PortfolioTabManager activeTab={tab} setTab={setTab} />
      <div className="border-t border-solid border-[#28282F] mt-4">
        {tab === PortfolioTabs.HOME && <HomeContent />}
        {tab === PortfolioTabs.PORTFOLIO && (
          <PortfolioContent
            connectedAddress={address}
            networkPortfolioWithBalances={filteredPortfolioDataForBalances}
          />
        )}
      </div>
    </div>
  )
}

type PortfolioContentProps = {
  connectedAddress: Address | string
  networkPortfolioWithBalances: NetworkTokenBalancesAndAllowances
}

const PortfolioContent = ({
  connectedAddress,
  networkPortfolioWithBalances,
}: PortfolioContentProps) => {
  return (
    <div className="">
      <PortfolioAssetHeader />
      {connectedAddress ? (
        Object.keys(networkPortfolioWithBalances).map((chainId) => {
          const tokens = networkPortfolioWithBalances[chainId]
          return (
            <SingleNetworkPortfolio chainId={Number(chainId)} tokens={tokens} />
          )
        })
      ) : (
        <UnconnectedPortfolioContent />
      )}
    </div>
  )
}

const UnconnectedPortfolioContent = () => {
  return (
    <>
      <p
        className={`
      text-[#CCCAD3BF] mt-6 mb-4 pb-6
      border-b border-solid border-[#28282F]
      `}
      >
        Your bridgable assets appear here when your wallet is connected.
      </p>
      <ConnectWalletButton />
    </>
  )
}

type SingleNetworkPortfolioProps = {
  chainId: number
  tokens: TokenWithBalanceAndAllowance[]
}

const SingleNetworkPortfolio = ({
  chainId,
  tokens,
}: SingleNetworkPortfolioProps) => {
  const currentChain: Chain = CHAINS_BY_ID[chainId]

  return (
    <div className="flex flex-col">
      <PortfolioNetwork
        displayName={currentChain.name}
        chainIcon={currentChain.chainImg}
        chainId={chainId}
      />
    </div>
  )
}

type PortfolioNetworkProps = {
  displayName: string
  chainIcon: string
  chainId: number
}

const PortfolioNetwork = ({
  displayName,
  chainIcon,
  chainId,
}: PortfolioNetworkProps) => {
  const { chain } = useNetwork()
  const isCurrentlyConnectedNetwork: boolean = useMemo(() => {
    return chainId === chain.id
  }, [chain.id])

  return (
    <div className="flex flex-row">
      <Image
        className="rounded-md w-7 h-7"
        alt={`${displayName} img`}
        src={chainIcon}
      />
      <div>{displayName}</div>
      {isCurrentlyConnectedNetwork ? (
        <ConnectedButton />
      ) : (
        <ConnectButton chainId={chainId} />
      )}
    </div>
  )
}

const ConnectedButton = () => {
  const buttonClassName = `
    h-8 flex items-center
    text-base text-white px-4 py-2 rounded-3xl
    text-center transform-gpu transition-all duration-75
    border-2 border-[#D747FF] radial-gradient-bg
    hover:cursor-default
  `

  return <button className={buttonClassName}>Connected</button>
}

const ConnectButton = ({ chainId }: { chainId: number }) => {
  const handleConnectNetwork = async () => {
    await switchNetwork({ chainId: chainId })
  }

  const buttonClassName = `
    h-8 flex items-center
    text-base text-white px-4 py-2 rounded-3xl
    text-center transform-gpu transition-all duration-75
    border-2 border-[#101018]
    hover:cursor-pointer
  `

  return (
    <button className={buttonClassName} onClick={handleConnectNetwork}>
      Connect
    </button>
  )
}

const PortfolioAssetHeader = () => {
  return (
    <div className="flex text-[#CCCAD3BF] my-2">
      <div className="w-1/2 text-left">Token</div>
      <div className="w-1/2 text-left">Amount</div>
    </div>
  )
}

const HomeContent = () => {
  return (
    <div className="my-4 font-thin text-white">
      <p className="mb-3">
        Synapse is the most widely used, extensible, and secure cross-chain
        communications network.
      </p>
      <p className="mb-5">
        Preview your route in the Bridge panel, and connect your wallet when
        you're ready to authorize your transaction.
      </p>
      <ConnectWalletButton />
    </div>
  )
}
