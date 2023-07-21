import React, { useEffect } from 'react'
import { Address } from 'wagmi'
import { NetworkTokenBalancesAndAllowances } from '@/utils/hooks/fetchPortfolioBalances'
import {
  SingleNetworkPortfolio,
  PortfolioHeader,
} from './SingleNetworkPortfolio'
import { FetchState } from '@/slices/portfolio/reducer'
import { ConnectWalletButton } from './ConnectWalletButton'

type PortfolioContentProps = {
  connectedAddress: Address | string
  connectedChainId: number
  selectedFromChainId: number
  networkPortfolioWithBalances: NetworkTokenBalancesAndAllowances
  fetchState: FetchState
  bridgeTxHashes: string[]
}

export const PortfolioContent = ({
  connectedAddress,
  connectedChainId,
  selectedFromChainId,
  networkPortfolioWithBalances,
  fetchState,
  bridgeTxHashes,
}: PortfolioContentProps) => {
  const { currentNetworkPortfolio, remainingNetworksPortfolios } =
    getCurrentNetworkPortfolio(
      selectedFromChainId,
      networkPortfolioWithBalances
    )
  const portfolioExists: boolean =
    Object.keys(networkPortfolioWithBalances).length > 0

  const isInitialFetchLoading: boolean =
    !portfolioExists && fetchState === FetchState.LOADING

  return (
    <div data-test-id="portfolio-content">
      {currentNetworkPortfolio && connectedChainId && selectedFromChainId && (
        <SingleNetworkPortfolio
          portfolioChainId={selectedFromChainId}
          connectedChainId={connectedChainId}
          selectedFromChainId={selectedFromChainId}
          portfolioTokens={currentNetworkPortfolio[selectedFromChainId]}
          initializeExpanded={true}
          fetchState={fetchState}
        />
      )}
      {connectedAddress ? (
        isInitialFetchLoading ? (
          <LoadingPortfolioContent />
        ) : (
          Object.keys(remainingNetworksPortfolios).map(
            (chainId: string, index: number) => {
              const tokens = remainingNetworksPortfolios[chainId]
              return (
                <SingleNetworkPortfolio
                  portfolioChainId={Number(chainId)}
                  connectedChainId={connectedChainId}
                  selectedFromChainId={selectedFromChainId}
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
          <PortfolioHeader />
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
  currentNetworkPortfolio: NetworkTokenBalancesAndAllowances
  remainingNetworksPortfolios: NetworkTokenBalancesAndAllowances
} {
  const currentNetworkPortfolio: NetworkTokenBalancesAndAllowances = {
    [currentChainId]: networks[currentChainId],
  }

  const remainingNetworksPortfolios = { ...networks }
  delete remainingNetworksPortfolios[currentChainId]

  return {
    currentNetworkPortfolio,
    remainingNetworksPortfolios,
  }
}

const LoadingPortfolioContent = () => {
  return (
    <>
      <p
        data-test-id="loading-portfolio-content"
        className={`
        text-[#CCCAD3BF] mt-6 mb-4 pb-6 pl-2
          border-b border-solid border-[#3D3D5C]
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
        text-[#CCCAD3BF] mt-6 mb-4 pb-6 pl-2
          border-b border-solid border-[#3D3D5C]
        `}
      >
        Your bridgable assets appear here when your wallet is connected.
      </p>
      <ConnectWalletButton />
    </>
  )
}

export const EmptyPortfolioContent = () => {
  return (
    <>
      <p
        data-test-id="empty-portfolio-content"
        className={`
        text-[#CCCAD3BF] py-4
        `}
      >
        No balances found.
      </p>
    </>
  )
}

export const HomeContent = () => {
  return (
    <div
      data-test-id="portfolio-home-content"
      className="border-t border-solid border-[#3D3D5C] text-white pl-2"
    >
      <p className="my-3">
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
