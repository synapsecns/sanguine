import {
  usePortfolioBalancesAndAllowances,
  NetworkTokenBalancesAndAllowances,
  TokenWithBalanceAndAllowance,
} from '@/utils/hooks/usePortfolioBalances'
import { PortfolioTabManager } from './PortfolioTabManager'
import { useAccount, useNetwork, Address } from 'wagmi'
import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'
import { useState } from 'react'
import { ConnectWalletButton } from './ConnectWalletButton'
import { CHAINS_BY_ID } from '@/constants/chains'
import HomeSvg from '../icons/HomeIcon'
import { Chain } from '@/utils/types'

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

  console.log(
    'filteredPortfolioDataForBalances: ',
    filteredPortfolioDataForBalances
  )

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
            portfolioData={portfolioData}
          />
        )}
      </div>
    </div>
  )
}

type PortfolioContentProps = {
  connectedAddress: Address | string
  portfolioData: NetworkTokenBalancesAndAllowances
}

const PortfolioContent = ({
  connectedAddress,
  portfolioData,
}: PortfolioContentProps) => {
  console.log('portfolioData inside content: ', portfolioData)
  return (
    <div className="">
      <PortfolioAssetHeader />
      {connectedAddress ? (
        <></>
      ) : (
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
      )}
    </div>
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

  return <></>
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
