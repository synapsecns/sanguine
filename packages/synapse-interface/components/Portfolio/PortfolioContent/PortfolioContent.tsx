import React, { useEffect, useState } from 'react'
import { Address, useAccount } from 'wagmi'
import Link from 'next/link'
import { NetworkTokenBalancesAndAllowances } from '@/utils/actions/fetchPortfolioBalances'
import {
  SingleNetworkPortfolio,
  PortfolioHeader,
} from './SingleNetworkPortfolio'
import { FetchState } from '@/slices/portfolio/actions'
import { ConnectWalletButton } from './components/ConnectWalletButton'
import { CHAINS_BY_ID } from '@/constants/chains'
import { Chain } from '@/utils/types'
import { DISCORD_URL, TWITTER_URL } from '@/constants/urls'
import { shortenAddress } from '@/utils/shortenAddress'

type PortfolioContentProps = {
  connectedAddress: Address | string
  connectedChainId: number
  selectedFromChainId: number
  networkPortfolioWithBalances: NetworkTokenBalancesAndAllowances
  fetchState: FetchState
  visibility: boolean
}

export const PortfolioContent = ({
  connectedAddress,
  connectedChainId,
  selectedFromChainId,
  networkPortfolioWithBalances,
  fetchState,
  visibility,
}: PortfolioContentProps) => {
  const { currentNetworkPortfolio, remainingNetworksPortfolios } =
    getCurrentNetworkPortfolio(
      selectedFromChainId,
      networkPortfolioWithBalances
    )

  const portfolioExists: boolean =
    Object.keys(networkPortfolioWithBalances).length > 0
  const currentChain: Chain = CHAINS_BY_ID[selectedFromChainId]
  const isUnsupportedChain: boolean = currentChain ? false : true

  const isInitialFetchLoading: boolean =
    !portfolioExists && fetchState === FetchState.LOADING

  return (
    <div
      data-test-id="portfolio-content"
      className={`${visibility ? 'block' : 'hidden'}`}
    >
      {!connectedAddress && <HomeContent />}
      {connectedAddress && isInitialFetchLoading && <LoadingPortfolioContent />}
      {currentNetworkPortfolio &&
        connectedAddress &&
        connectedChainId &&
        selectedFromChainId &&
        !isInitialFetchLoading && (
          <SingleNetworkPortfolio
            portfolioChainId={selectedFromChainId}
            connectedChainId={connectedChainId}
            selectedFromChainId={selectedFromChainId}
            portfolioTokens={currentNetworkPortfolio[selectedFromChainId]}
            initializeExpanded={true}
            fetchState={fetchState}
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
                portfolioChainId={Number(chainId)}
                connectedChainId={connectedChainId}
                selectedFromChainId={selectedFromChainId}
                portfolioTokens={tokens}
                initializeExpanded={false}
                fetchState={fetchState}
              />
            )
          }
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
      <p data-test-id="loading-portfolio-content" className="text-[#CCCAD3BF]">
        Loading assets...
      </p>
    </>
  )
}

const UnconnectedPortfolioContent = () => {
  return (
    <>
      <p
        data-test-id="unconnected-portfolio-content"
        className={`
        text-[#C2C2D6] mt-6 mb-4 pb-6 pl-2
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
  const { address } = useAccount()
  const shortened = shortenAddress(address, 3)
  return (
    <div data-test-id="empty-portfolio-content" className="p-1">
      <p className="text-[#C2C2D6] mb-4">
        No supported assets found {address && `for ${shortened}`}.
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
