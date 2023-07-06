import React from 'react'
import { Address } from 'wagmi'
import {
  NetworkTokenBalancesAndAllowances,
  FetchState,
} from '@/utils/hooks/usePortfolioBalances'
import {
  SingleNetworkPortfolio,
  PortfolioAssetHeader,
} from './SingleNetworkPortfolio'
import { ConnectWalletButton } from './ConnectWalletButton'

type PortfolioContentProps = {
  connectedAddress: Address | string
  connectedChainId: number
  selectedFromChainId: number
  networkPortfolioWithBalances: NetworkTokenBalancesAndAllowances
  fetchState: FetchState
}

export const PortfolioContent = ({
  connectedAddress,
  connectedChainId,
  selectedFromChainId,
  networkPortfolioWithBalances,
  fetchState,
}: PortfolioContentProps) => {
  const { currentNetwork, remainingNetworks } = getCurrentNetworkPortfolio(
    selectedFromChainId,
    networkPortfolioWithBalances
  )

  return (
    <div data-test-id="portfolio-content">
      {currentNetwork && connectedChainId && (
        <SingleNetworkPortfolio
          portfolioChainId={selectedFromChainId}
          connectedChainId={connectedChainId}
          portfolioTokens={currentNetwork[selectedFromChainId]}
          initializeExpanded={true}
          fetchState={fetchState}
        />
      )}
      {connectedAddress ? (
        fetchState === FetchState.LOADING ? (
          <LoadingPortfolioContent />
        ) : (
          Object.keys(remainingNetworks).map(
            (chainId: string, index: number) => {
              const tokens = remainingNetworks[chainId]
              return (
                <SingleNetworkPortfolio
                  portfolioChainId={Number(chainId)}
                  connectedChainId={connectedChainId}
                  portfolioTokens={tokens}
                  initializeExpanded={false}
                  fetchState={fetchState}
                />
              )
            }
          )
        )
      ) : (
        <React.Fragment>
          <PortfolioAssetHeader />
          <UnconnectedPortfolioContent />
        </React.Fragment>
      )}
    </div>
  )
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

const LoadingPortfolioContent = () => {
  return (
    <>
      <p
        data-test-id="loading-portfolio-content"
        className={`
        text-[#CCCAD3BF] mt-6 mb-4 pb-6
          border-b border-solid border-[#28282F]
        `}
      >
        Loading portfolio balances...
      </p>
      <ConnectWalletButton />
    </>
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

export const HomeContent = () => {
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
