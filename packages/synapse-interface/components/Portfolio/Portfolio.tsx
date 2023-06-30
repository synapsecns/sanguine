import {
  usePortfolioBalancesAndAllowances,
  NetworkTokenBalancesAndAllowances,
  TokenWithBalanceAndAllowance,
} from '@/utils/hooks/usePortfolioBalances'
import { PortfolioTabManager } from './PortfolioTabManager'
import {
  PortfolioAssetHeader,
  SingleNetworkPortfolio,
} from './SinglePortfolioNetwork'
import { useAccount, useNetwork, Address } from 'wagmi'
import { Zero } from '@ethersproject/constants'
import React, { useState } from 'react'
import { ConnectWalletButton } from './ConnectWalletButton'

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
      {connectedAddress ? (
        Object.keys(networkPortfolioWithBalances).map((chainId) => {
          const tokens = networkPortfolioWithBalances[chainId]
          return (
            <SingleNetworkPortfolio chainId={Number(chainId)} tokens={tokens} />
          )
        })
      ) : (
        <React.Fragment>
          <PortfolioAssetHeader />
          <UnconnectedPortfolioContent />
        </React.Fragment>
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
