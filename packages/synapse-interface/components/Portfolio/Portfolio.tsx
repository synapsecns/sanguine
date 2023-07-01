import {
  usePortfolioBalancesAndAllowances,
  NetworkTokenBalancesAndAllowances,
  TokenWithBalanceAndAllowance,
} from '@/utils/hooks/usePortfolioBalances'
import { PortfolioTabManager } from './PortfolioTabManager'
import {
  PortfolioAssetHeader,
  SingleNetworkPortfolio,
} from './SingleNetworkPortfolio'
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
    <div data-test-id="portfolio" className="flex flex-col w-1/2">
      <PortfolioTabManager activeTab={tab} setTab={setTab} />
      <div className="border-t border-solid border-[#28282F] mt-4">
        {tab === PortfolioTabs.HOME && <HomeContent />}
        {tab === PortfolioTabs.PORTFOLIO && (
          <PortfolioContent
            connectedAddress={address}
            connectedChainId={chain.id}
            networkPortfolioWithBalances={filteredPortfolioDataForBalances}
          />
        )}
      </div>
    </div>
  )
}

type PortfolioContentProps = {
  connectedAddress: Address | string
  connectedChainId: number
  networkPortfolioWithBalances: NetworkTokenBalancesAndAllowances
}

function getCurrentNetworkPortfolio(
  currentChainId: number,
  networks: NetworkTokenBalancesAndAllowances
): {
  currentNetwork: NetworkTokenBalancesAndAllowances
  remainingNetworks: NetworkTokenBalancesAndAllowances
} {
  const currentNetwork: NetworkTokenBalancesAndAllowances = {
    [currentChainId]: networks[currentChainId],
  }

  const remainingNetworks = { ...networks }
  delete remainingNetworks[currentChainId]

  return {
    currentNetwork,
    remainingNetworks,
  }
}

const PortfolioContent = ({
  connectedAddress,
  connectedChainId,
  networkPortfolioWithBalances,
}: PortfolioContentProps) => {
  const { currentNetwork, remainingNetworks } = getCurrentNetworkPortfolio(
    connectedChainId,
    networkPortfolioWithBalances
  )

  return (
    <div data-test-id="portfolio-content">
      {currentNetwork && (
        <SingleNetworkPortfolio
          portfolioChainId={connectedChainId}
          connectedChainId={connectedChainId}
          portfolioTokens={currentNetwork[connectedChainId]}
          initializeExpanded={true}
        />
      )}
      {connectedAddress ? (
        Object.keys(remainingNetworks).map((chainId: string, index: number) => {
          const tokens = remainingNetworks[chainId]
          const isExpanded = index === 0
          return (
            <SingleNetworkPortfolio
              portfolioChainId={Number(chainId)}
              connectedChainId={connectedChainId}
              portfolioTokens={tokens}
              initializeExpanded={isExpanded}
            />
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
        data-test-id="unconnected-portfolio-content"
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
    <div data-test-id="portfolio-home-content" className="my-4 text-white">
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
