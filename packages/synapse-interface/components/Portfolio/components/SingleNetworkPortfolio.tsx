import React, { useState, useEffect } from 'react'
import { Address } from 'viem'
import { useDispatch } from 'react-redux'
import _ from 'lodash'
import { CHAINS_BY_ID } from '@/constants/chains'
import {
  TokenAndBalance,
  sortTokensByBalanceDescending,
} from '@/utils/actions/fetchPortfolioBalances'
import { Chain } from '@/utils/types'
import { PortfolioAccordion } from './PortfolioAccordion'
import { PortfolioConnectButton } from './PortfolioConnectButton'
import { EmptyPortfolioContent } from './EmptyPortfolioContent'
import { FetchState } from '@/slices/portfolio/actions'
import { PortfolioTokenAsset } from './PortfolioTokenAsset'

import { WarningMessage } from '../../Warning'
import { TWITTER_URL, DISCORD_URL } from '@/constants/urls'
import { setFromToken, setToToken } from '@/slices/bridge/reducer'
import { PortfolioNetwork } from './PortfolioNetwork'
import { ChainTokens } from './ChainTokens'

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

  const isLoading = fetchState === FetchState.LOADING

  const chain: Chain = CHAINS_BY_ID[portfolioChainId]
  const isUnsupportedChain: boolean = !chain

  const sortedTokens = sortTokensByBalanceDescending(portfolioTokens)
  const hasNoTokenBalance: boolean =
    _.isNull(portfolioTokens) || _.isEmpty(portfolioTokens)

  useEffect(() => {
    if (isUnsupportedChain) {
      dispatch(setFromToken(null))
      dispatch(setToToken(null))
    }
  }, [isUnsupportedChain])

  return (
    <div
      id="single-network-portfolio"
      className="flex flex-col mb-4 border rounded-lg border-white/10"
    >
      <PortfolioAccordion
        connectedChainId={connectedChainId}
        portfolioChainId={portfolioChainId}
        selectedFromChainId={selectedFromChainId}
        initializeExpanded={initializeExpanded}
        hasNoTokenBalance={hasNoTokenBalance}
        header={
          <PortfolioNetwork
            displayName={chain?.name}
            chainIcon={chain?.chainImg}
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
          <ChainTokens balanceTokens={sortedTokens} />
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
            connectedChain={chain}
          />
        )}
        {sortedTokens &&
          sortedTokens.length > 0 &&
          sortedTokens.map(({ token, balance }: TokenAndBalance) => (
            <PortfolioTokenAsset
              key={token.symbol}
              token={token}
              balance={balance}
              portfolioChainId={portfolioChainId}
              connectedChainId={connectedChainId}
            />
          ))}
      </PortfolioAccordion>
    </div>
  )
}
