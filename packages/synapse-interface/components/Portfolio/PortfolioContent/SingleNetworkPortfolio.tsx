import React, { useState, useEffect } from 'react'
import { Address } from 'viem'
import { useDispatch } from 'react-redux'
import Image from 'next/image'
import { CHAINS_BY_ID } from '@/constants/chains'
import {
  TokenAndBalance,
  // separateTokensByAllowance,
  sortTokensByBalanceDescending,
} from '@/utils/actions/fetchPortfolioBalances'
import { Chain } from '@/utils/types'
import { PortfolioAccordion } from './components/PortfolioAccordion'
import { PortfolioConnectButton } from './components/PortfolioConnectButton'
import { EmptyPortfolioContent } from './PortfolioContent'
import { FetchState } from '@/slices/portfolio/actions'
import { PortfolioTokenAsset } from './components/PortfolioTokenAsset'
import { QuestionMarkCircleIcon } from '@heroicons/react/outline'
import { WarningMessage } from '../../Warning'
import { TWITTER_URL, DISCORD_URL } from '@/constants/urls'
import { setFromToken, setToToken } from '@/slices/bridge/reducer'

type SingleNetworkPortfolioProps = {
  connectedAddress: Address
  portfolioChainId: number
  connectedChainId: number
  selectedFromChainId: number
  portfolioTokens: TokenAndBalance[]
  initializeExpanded: boolean
  fetchState: FetchState
}

export const SingleNetworkPortfolio = ({
  connectedAddress,
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

  // const [tokensWithAllowance, tokensWithoutAllowance] =
  //   separateTokensByAllowance(portfolioTokens)

  // const sortedTokensWithAllowance: TokenWithBalanceAndAllowances[] =
  //   sortTokensByBalanceDescending(tokensWithAllowance)
  // const sortedTokensWithoutAllowance: TokenWithBalanceAndAllowances[] =
  //   sortTokensByBalanceDescending(tokensWithoutAllowance)
  const sortedTokensForVisualizer: TokenAndBalance[] =
    sortTokensByBalanceDescending(portfolioTokens)

  const hasNoTokenBalance: boolean =
    !portfolioTokens || portfolioTokens.length === 0

  const isLoading: boolean = fetchState === FetchState.LOADING

  useEffect(() => {
    if (isUnsupportedChain) {
      dispatch(setFromToken(null))
      dispatch(setToToken(null))
    }
  }, [isUnsupportedChain])

  return (
    <div
      data-test-id="single-network-portfolio"
      className="flex flex-col mb-4 border rounded-md border-surface"
    >
      <PortfolioAccordion
        connectedChainId={connectedChainId}
        portfolioChainId={portfolioChainId}
        selectedFromChainId={selectedFromChainId}
        initializeExpanded={initializeExpanded}
        hasNoTokenBalance={hasNoTokenBalance}
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
        {!isLoading && hasNoTokenBalance && (
          <EmptyPortfolioContent
            connectedAddress={connectedAddress}
            connectedChain={currentChain}
          />
        )}
        {sortedTokensForVisualizer &&
          sortedTokensForVisualizer.length > 0 &&
          sortedTokensForVisualizer.map(
            ({ token, balance }: TokenAndBalance) => (
              <PortfolioTokenAsset
                key={token.symbol}
                token={token}
                balance={balance}
                portfolioChainId={portfolioChainId}
                connectedChainId={connectedChainId}
                isApproved={true}
              />
            )
          )}
        {/* {sortedTokensWithoutAllowance &&
          sortedTokensWithoutAllowance.length > 0 &&
          sortedTokensWithoutAllowance.map(
            ({ token, balance }: TokenWithBalanceAndAllowances) => (
              <PortfolioTokenAsset
                key={token.symbol}
                token={token}
                balance={balance}
                portfolioChainId={portfolioChainId}
                connectedChainId={connectedChainId}
                isApproved={false}
              />
            )
          )} */}
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
      className="flex flex-row justify-between flex-1 py-4 cursor-pointer"
    >
      <div className="flex flex-row items-center px-4">
        {isUnsupportedChain ? (
          <QuestionMarkCircleIcon className="w-6 h-6 mr-2 text-white rounded-md" />
        ) : (
          <Image
            className="w-6 h-6 mr-2 rounded-md"
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

export const PortfolioTokenVisualizer = ({
  portfolioTokens,
}: {
  portfolioTokens: TokenAndBalance[]
}) => {
  const [isT1Hovered, setIsT1Hovered] = useState<boolean>(false)
  const [isT2Hovered, setIsT2Hovered] = useState<boolean>(false)
  const [isT3Hovered, setIsT3Hovered] = useState<boolean>(false)

  const hasNoTokens: boolean =
    !portfolioTokens || (portfolioTokens && portfolioTokens.length === 0)
  const hasOneToken: boolean = portfolioTokens && portfolioTokens.length > 0
  const hasTwoTokens: boolean = portfolioTokens && portfolioTokens.length > 1
  const numOverTwoTokens: number =
    portfolioTokens && portfolioTokens.length - 2 > 0
      ? portfolioTokens.length - 2
      : 0
  const hasOnlyOneToken: boolean =
    portfolioTokens && portfolioTokens.length === 1
  const hasOnlyTwoTokens: boolean =
    portfolioTokens && portfolioTokens.length === 2

  if (hasNoTokens) {
    return (
      <div
        data-test-id="portfolio-token-visualizer"
        className="flex flex-row items-center mr-4 cursor-pointer hover-trigger text-secondary"
      >
        -
      </div>
    )
  }
  return (
    <div
      data-test-id="portfolio-token-visualizer"
      className="flex flex-row items-center space-x-2 cursor-pointer hover-trigger"
    >
      {hasOneToken && (
        <div>
          <Image
            loading="lazy"
            className="w-6 h-6 rounded-md"
            alt={`${portfolioTokens[0].token.symbol} img`}
            src={portfolioTokens[0].token.icon}
            onMouseEnter={() => setIsT1Hovered(true)}
            onMouseLeave={() => setIsT1Hovered(false)}
          />
          <div className="relative">
            <HoverContent isHovered={isT1Hovered}>
              <div className="whitespace-nowrap">
                {portfolioTokens[0]?.parsedBalance}{' '}
                {portfolioTokens[0]?.token.symbol}
              </div>
            </HoverContent>
          </div>
        </div>
      )}
      {hasOnlyOneToken && (
        <div className="text-white whitespace-nowrap">
          {portfolioTokens[0].parsedBalance} {portfolioTokens[0].token.symbol}
        </div>
      )}
      {hasTwoTokens && (
        <div>
          <Image
            loading="lazy"
            className="w-6 h-6 rounded-md"
            alt={`${portfolioTokens[1].token.symbol} img`}
            src={portfolioTokens[1].token.icon}
            onMouseEnter={() => setIsT2Hovered(true)}
            onMouseLeave={() => setIsT2Hovered(false)}
          />
          <div className="relative">
            <HoverContent isHovered={isT2Hovered}>
              <div className="whitespace-nowrap">
                {portfolioTokens[1]?.parsedBalance}{' '}
                {portfolioTokens[1]?.token.symbol}
              </div>
            </HoverContent>
          </div>
        </div>
      )}
      {numOverTwoTokens > 0 && (
        <div
          className="text-white"
          onMouseEnter={() => setIsT3Hovered(true)}
          onMouseLeave={() => setIsT3Hovered(false)}
        >
          + {numOverTwoTokens}
        </div>
      )}
      <div className="relative inline-block">
        <HoverContent isHovered={isT3Hovered}>
          {portfolioTokens?.map((token: TokenAndBalance, key: number) => {
            if (key > 1) {
              const tokenSymbol = token.token.symbol
              const balance = token.parsedBalance
              return (
                <div className="whitespace-nowrap" key={key}>
                  {balance} {tokenSymbol}
                </div>
              )
            }
          })}
        </HoverContent>
      </div>
    </div>
  )
}

export const HoverContent = ({
  isHovered,
  children,
}: {
  isHovered: boolean
  children: React.ReactNode
}) => {
  if (isHovered) {
    return (
      <div
        className={`
          absolute z-50 hover-content p-2 text-white
          border border-solid border-[#252537]
          bg-[#101018] rounded-md text-left
        `}
      >
        {children}
      </div>
    )
  }
}
