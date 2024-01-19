import React, { useMemo } from 'react'
import { Address } from 'wagmi'
import Link from 'next/link'
import {
  NetworkTokenBalances,
  TokenAndBalance,
} from '@/utils/actions/fetchPortfolioBalances'
import { SingleNetworkPortfolio } from './SingleNetworkPortfolio'
import { FetchState } from '@/slices/portfolio/actions'
import { ConnectWalletButton } from './ConnectWalletButton'
import { Chain } from '@/utils/types'
import { DISCORD_URL, TWITTER_URL } from '@/constants/urls'
import { shortenAddress } from '@/utils/shortenAddress'

type PortfolioContentProps = {
  connectedAddress: Address | string
  connectedChainId: number
  selectedFromChainId: number
  networkPortfolioWithBalances: NetworkTokenBalances
  fetchState: FetchState
  visibility: boolean
  searchInputActive: boolean
  searchStatus: FetchState
  searchInput: string
}

export const PortfolioContent = ({
  connectedAddress,
  connectedChainId,
  selectedFromChainId,
  networkPortfolioWithBalances,
  fetchState,
  visibility,
  searchInputActive,
  searchStatus,
  searchInput,
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

  const showCurrentNetworkPortfolio: boolean = useMemo(() => {
    if (searchInputActive && currentNetworkPortfolio) {
      return Boolean(currentNetworkPortfolio[selectedFromChainId])
    } else {
      return Boolean(currentNetworkPortfolio)
    }
  }, [
    searchInputActive,
    currentNetworkPortfolio,
    networkPortfolioWithBalances,
    selectedFromChainId,
  ])

  const hasFilteredSearchResults: boolean = useMemo(() => {
    if (networkPortfolioWithBalances) {
      return Object.values(networkPortfolioWithBalances).length > 0
    } else {
      return false
    }
  }, [networkPortfolioWithBalances])

  return (
    <div
      data-test-id="portfolio-content"
      className={`${visibility ? 'block' : 'hidden'}`}
    >
      {!connectedAddress && !searchInputActive && <HomeContent />}
      {searchInputActive &&
        !hasFilteredSearchResults &&
        searchStatus !== FetchState.LOADING && (
          <NoSearchResultsContent searchStr={searchInput} />
        )}
      {connectedAddress && isInitialFetchLoading && <LoadingPortfolioContent />}
      {showCurrentNetworkPortfolio &&
        connectedAddress &&
        selectedFromChainId &&
        !isInitialFetchLoading && (
          <SingleNetworkPortfolio
            connectedAddress={connectedAddress as Address}
            portfolioChainId={selectedFromChainId as number}
            connectedChainId={connectedChainId as number}
            selectedFromChainId={selectedFromChainId as number}
            portfolioTokens={currentNetworkPortfolio[selectedFromChainId]}
            initializeExpanded={false}
            fetchState={fetchState as FetchState}
          />
        )}
      {connectedAddress &&
        !isInitialFetchLoading &&
        Object.keys(remainingNetworksPortfolios).map(
          (chainId: string, index: number) => {
            const tokens = remainingNetworksPortfolios[chainId]
            return (
              <SingleNetworkPortfolio
                key={chainId}
                connectedAddress={connectedAddress as Address}
                portfolioChainId={Number(chainId) as number}
                connectedChainId={connectedChainId as number}
                selectedFromChainId={selectedFromChainId as number}
                portfolioTokens={tokens as TokenAndBalance[]}
                initializeExpanded={false}
                fetchState={fetchState as FetchState}
              />
            )
          }
        )}
    </div>
  )
}

function getCurrentNetworkPortfolio(
  currentChainId: number,
  networks: NetworkTokenBalances
): {
  currentNetworkPortfolio: NetworkTokenBalances
  remainingNetworksPortfolios: NetworkTokenBalances
} {
  const currentNetworkPortfolio: NetworkTokenBalances = {
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
      <p data-test-id="loading-portfolio-content" className="text-[#CCCAD3BF]">
        Loading assets...
      </p>
    </>
  )
}

export const EmptyPortfolioContent = ({
  connectedAddress,
  connectedChain,
}: {
  connectedAddress: Address
  connectedChain: Chain
}) => {
  const shortened: string = shortenAddress(connectedAddress)
  return (
    <div data-test-id="empty-portfolio-content" className="p-4">
      <p className="text-[#C2C2D6] mb-4">
        No bridgeable assets found {connectedAddress && `for ${shortened}`} on{' '}
        {connectedChain?.name}.
      </p>
      <p className="text-[#C2C2D6] mb-4">
        Don't see a chain or token you want to bridge?
      </p>
      <a className="text-[#C2C2D6]">
        Let us know on
        <Link
          className="text-[#99E6FF] underline px-1"
          href={TWITTER_URL}
          target="_blank"
        >
          Twitter
        </Link>
        or
        <Link
          className="text-[#99E6FF] underline pl-1"
          href={DISCORD_URL}
          target="_blank"
        >
          Discord
        </Link>
        .
      </a>
    </div>
  )
}

export const HomeContent = () => {
  return (
    <div data-test-id="portfolio-home-content" className="text-white">
      <p className="mb-3">
        Synapse is the most widely used, extensible, and secure cross-chain
        communications network.
      </p>
      <p className="mb-5">
        Get route quotes in the Bridge panel, and connect your wallet when you
        are ready to submit a transaction.
      </p>
      <ConnectWalletButton />
    </div>
  )
}

export const NoSearchResultsContent = ({
  searchStr,
}: {
  searchStr: string
}) => {
  return (
    <div
      data-test-id="portfolio-no-search-results-content"
      className="text-white"
    >
      <p className="mb-3 break-words">No results found for '{searchStr}'.</p>
    </div>
  )
}
