import React, { useState } from 'react'
import Image from 'next/image'
import { CHAINS_BY_ID } from '@/constants/chains'
import { TokenWithBalanceAndAllowance } from '@/utils/hooks/usePortfolioBalances'
import { Chain } from '@/utils/types'
import PortfolioAccordion from './PortfolioAccordion'
import { PortfolioConnectButton } from './PortfolioConnectButton'
import { EmptyPortfolioContent } from './PortfolioContent'
import { FetchState } from '@/slices/portfolio/reducer'
import { PortfolioTokenAsset } from './PortfolioTokenAsset'

type SingleNetworkPortfolioProps = {
  portfolioChainId: number
  connectedChainId: number
  selectedFromChainId: number
  portfolioTokens: TokenWithBalanceAndAllowance[]
  initializeExpanded: boolean
  fetchPortfolioBalancesCallback: () => Promise<void>
  fetchState: FetchState
}

export const SingleNetworkPortfolio = ({
  portfolioChainId,
  connectedChainId,
  selectedFromChainId,
  portfolioTokens,
  initializeExpanded = false,
  fetchPortfolioBalancesCallback,
  fetchState,
}: SingleNetworkPortfolioProps) => {
  const currentChain: Chain = CHAINS_BY_ID[portfolioChainId]

  const [tokensWithAllowance, tokensWithoutAllowance] =
    separateTokensByAllowance(portfolioTokens)

  const sortedTokensWithAllowance: TokenWithBalanceAndAllowance[] =
    sortByBalanceDescending(tokensWithAllowance)
  const sortedTokensWithoutAllowance: TokenWithBalanceAndAllowance[] =
    sortByBalanceDescending(tokensWithoutAllowance)
  const sortedTokensForVisualizer: TokenWithBalanceAndAllowance[] =
    sortByBalanceDescending(portfolioTokens)

  const hasNoTokenBalance: boolean =
    !portfolioTokens || portfolioTokens.length === 0

  const isLoading: boolean = fetchState === FetchState.LOADING

  return (
    <div data-test-id="single-network-portfolio" className="flex flex-col">
      <PortfolioAccordion
        connectedChainId={connectedChainId}
        portfolioChainId={portfolioChainId}
        selectedFromChainId={selectedFromChainId}
        initializeExpanded={initializeExpanded}
        header={
          <PortfolioNetwork
            displayName={currentChain?.name}
            chainIcon={currentChain?.chainImg}
          />
        }
        expandedProps={
          <PortfolioConnectButton
            connectedChainId={connectedChainId}
            portfolioChainId={portfolioChainId}
          />
        }
        collapsedProps={
          <PortfolioTokenVisualizer
            portfolioTokens={sortedTokensForVisualizer}
          />
        }
      >
        <PortfolioHeader />
        {!isLoading && hasNoTokenBalance && <EmptyPortfolioContent />}
        {sortedTokensWithAllowance &&
          sortedTokensWithAllowance.length > 0 &&
          sortedTokensWithAllowance.map(
            ({ token, balance, allowance }: TokenWithBalanceAndAllowance) => (
              <PortfolioTokenAsset
                token={token}
                balance={balance}
                allowance={allowance}
                portfolioChainId={portfolioChainId}
                connectedChainId={connectedChainId}
                fetchPortfolioBalancesCallback={fetchPortfolioBalancesCallback}
                isApproved={true}
              />
            )
          )}
        {sortedTokensWithoutAllowance &&
          sortedTokensWithoutAllowance.length > 0 &&
          sortedTokensWithoutAllowance.map(
            ({ token, balance }: TokenWithBalanceAndAllowance) => (
              <PortfolioTokenAsset
                token={token}
                balance={balance}
                portfolioChainId={portfolioChainId}
                connectedChainId={connectedChainId}
                fetchPortfolioBalancesCallback={fetchPortfolioBalancesCallback}
                isApproved={false}
              />
            )
          )}
      </PortfolioAccordion>
    </div>
  )
}

type PortfolioNetworkProps = {
  displayName: string
  chainIcon: string
}

const PortfolioNetwork = ({
  displayName,
  chainIcon,
}: PortfolioNetworkProps) => {
  return (
    <div
      data-test-id="portfolio-network"
      className="flex flex-row justify-between flex-1 py-4 pl-2"
    >
      <div className="flex flex-row items-center">
        <Image
          className="w-6 h-6 mr-3 rounded-md"
          alt={`${displayName} img`}
          src={chainIcon}
        />
        <div className="text-lg font-medium text-white">{displayName}</div>
      </div>
    </div>
  )
}

const PortfolioTokenVisualizer = ({
  portfolioTokens,
}: {
  portfolioTokens: TokenWithBalanceAndAllowance[]
}) => {
  const [isHovered, setIsHovered] = useState(false)
  const hasOneToken = portfolioTokens && portfolioTokens.length > 0
  const hasTwoTokens = portfolioTokens && portfolioTokens.length > 1
  const numOverTwoTokens =
    portfolioTokens && portfolioTokens.length - 2 > 0
      ? portfolioTokens.length - 2
      : 0

  return (
    <div
      data-test-id="portfolio-token-visualizer"
      className="flex flex-row items-center hover-trigger"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      {hasOneToken && (
        <Image
          loading="lazy"
          className="w-6 h-6 rounded-md"
          alt={`${portfolioTokens[0].token.symbol} img`}
          src={portfolioTokens[0].token.icon}
        />
      )}
      {hasTwoTokens && (
        <Image
          loading="lazy"
          className="w-6 h-6 ml-1 rounded-md"
          alt={`${portfolioTokens[1].token.symbol} img`}
          src={portfolioTokens[1].token.icon}
        />
      )}
      {numOverTwoTokens > 0 && (
        <div className="ml-1 text-white">+ {numOverTwoTokens}</div>
      )}
      <div className="relative inline-block">
        {isHovered && (
          <div
            className={`
            absolute z-50 hover-content p-2 text-white
            border border-solid border-[#252537]
            bg-[#101018] rounded-md`}
          >
            {portfolioTokens.map((token: TokenWithBalanceAndAllowance) => {
              const tokenSymbol = token.token.symbol
              const balance = token.parsedBalance
              return (
                <div className="whitespace-nowrap">
                  {balance} {tokenSymbol}
                </div>
              )
            })}
          </div>
        )}
      </div>
    </div>
  )
}

export const PortfolioHeader = () => {
  return (
    <div
      data-test-id="portfolio-asset-header"
      className="flex text-[#CCCAD3BF] my-2 pl-2"
    >
      <div className="flex flex-row justify-between w-2/3 text-left">
        <div className="pl-4">Token</div>
        <div className="pr-2">Amount</div>
      </div>
      <div className="w-1/3 text-left" />
    </div>
  )
}

function separateTokensByAllowance(
  tokens: TokenWithBalanceAndAllowance[]
): [TokenWithBalanceAndAllowance[], TokenWithBalanceAndAllowance[]] {
  const tokensWithAllowance: TokenWithBalanceAndAllowance[] = []
  const tokensWithoutAllowance: TokenWithBalanceAndAllowance[] = []

  tokens &&
    tokens.forEach((token) => {
      // allowance is null for native gas tokens
      if (token.allowance === null) {
        tokensWithAllowance.push(token)
      } else if (token.allowance > 0n) {
        tokensWithAllowance.push(token)
      } else {
        tokensWithoutAllowance.push(token)
      }
    })

  return [tokensWithAllowance, tokensWithoutAllowance]
}

function sortByBalanceDescending(
  tokens: TokenWithBalanceAndAllowance[]
): TokenWithBalanceAndAllowance[] {
  return (
    tokens &&
    tokens.sort(
      (a: TokenWithBalanceAndAllowance, b: TokenWithBalanceAndAllowance) =>
        b.parsedBalance > a.parsedBalance ? 1 : -1
    )
  )
}
