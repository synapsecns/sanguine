import React, { useState, useEffect } from 'react'
import { useDispatch } from 'react-redux'
import Image from 'next/image'
import { CHAINS_BY_ID } from '@/constants/chains'
import {
  ROUTER_ADDRESS,
  TokenWithBalanceAndAllowance,
  TokenWithBalanceAndAllowances,
  separateTokensByAllowance,
  sortTokensByBalanceDescending,
} from '@/utils/actions/fetchPortfolioBalances'
import { Chain } from '@/utils/types'
import PortfolioAccordion from './PortfolioAccordion'
import { PortfolioConnectButton } from './PortfolioConnectButton'
import { EmptyPortfolioContent } from './PortfolioContent'
import { FetchState } from '@/slices/portfolio/actions'
import { PortfolioTokenAsset } from './PortfolioTokenAsset'
import { QuestionMarkCircleIcon } from '@heroicons/react/outline'
import { WarningMessage } from '../Warning'
import { TWITTER_URL, DISCORD_URL } from '@/constants/urls'
import {
  setSupportedFromTokens,
  setSupportedToTokens,
  setToChainId,
  initialState,
} from '@/slices/bridge/reducer'

type SingleNetworkPortfolioProps = {
  portfolioChainId: number
  connectedChainId: number
  selectedFromChainId: number
  portfolioTokens: TokenWithBalanceAndAllowances[]
  initializeExpanded: boolean
  fetchState: FetchState
}

export const SingleNetworkPortfolio = ({
  portfolioChainId,
  connectedChainId,
  selectedFromChainId,
  portfolioTokens,
  initializeExpanded = false,
  fetchState,
}: SingleNetworkPortfolioProps) => {
  const dispatch = useDispatch()

  const currentChain: Chain = CHAINS_BY_ID[portfolioChainId]
  const isUnsupportedChain: boolean = currentChain ? false : true

  const [tokensWithAllowance, tokensWithoutAllowance] =
    separateTokensByAllowance(portfolioTokens)

  const sortedTokensWithAllowance: TokenWithBalanceAndAllowances[] =
    sortTokensByBalanceDescending(tokensWithAllowance)
  const sortedTokensWithoutAllowance: TokenWithBalanceAndAllowances[] =
    sortTokensByBalanceDescending(tokensWithoutAllowance)
  const sortedTokensForVisualizer: TokenWithBalanceAndAllowances[] =
    sortTokensByBalanceDescending(portfolioTokens)

  const hasNoTokenBalance: boolean =
    !portfolioTokens || portfolioTokens.length === 0

  const isLoading: boolean = fetchState === FetchState.LOADING

  useEffect(() => {
    if (isUnsupportedChain) {
      dispatch(setSupportedFromTokens([]))
      dispatch(setSupportedToTokens(initialState.supportedToTokens))
    }
  }, [isUnsupportedChain])

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
            isUnsupportedChain={isUnsupportedChain}
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
        {isUnsupportedChain && (
          <WarningMessage
            twClassName="!p-2 !mt-0"
            message={
              <p className="leading-6">
                This chain is not yet supported. New chain or token support can
                be discussed on{' '}
                <a target="_blank" className="underline" href={TWITTER_URL}>
                  Twitter
                </a>{' '}
                or{' '}
                <a target="_blank" className="underline" href={DISCORD_URL}>
                  Discord
                </a>
                .
              </p>
            }
          />
        )}
        <PortfolioHeader />
        {!isLoading && hasNoTokenBalance && <EmptyPortfolioContent />}
        {sortedTokensWithAllowance &&
          sortedTokensWithAllowance.length > 0 &&
          sortedTokensWithAllowance.map(
            ({ token, balance, allowances }: TokenWithBalanceAndAllowances) => (
              <PortfolioTokenAsset
                token={token}
                balance={balance}
                allowances={allowances}
                portfolioChainId={portfolioChainId}
                connectedChainId={connectedChainId}
                isApproved={true}
              />
            )
          )}
        {sortedTokensWithoutAllowance &&
          sortedTokensWithoutAllowance.length > 0 &&
          sortedTokensWithoutAllowance.map(
            ({ token, balance }: TokenWithBalanceAndAllowances) => (
              <PortfolioTokenAsset
                token={token}
                balance={balance}
                portfolioChainId={portfolioChainId}
                connectedChainId={connectedChainId}
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
  isUnsupportedChain: boolean
}

const PortfolioNetwork = ({
  displayName,
  chainIcon,
  isUnsupportedChain,
}: PortfolioNetworkProps) => {
  return (
    <div
      data-test-id="portfolio-network"
      className="flex flex-row justify-between flex-1 py-4 pl-2"
    >
      <div className="flex flex-row items-center">
        {isUnsupportedChain ? (
          <QuestionMarkCircleIcon className="w-6 h-6 mr-3 text-white rounded-md" />
        ) : (
          <Image
            className="w-6 h-6 mr-3 rounded-md"
            alt={`${displayName} img`}
            src={chainIcon}
          />
        )}
        <div className="text-lg font-medium text-white">
          {isUnsupportedChain ? 'Unsupported Network' : displayName}
        </div>
      </div>
    </div>
  )
}

const PortfolioTokenVisualizer = ({
  portfolioTokens,
}: {
  portfolioTokens: TokenWithBalanceAndAllowances[]
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
            {portfolioTokens.map((token: TokenWithBalanceAndAllowances) => {
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
