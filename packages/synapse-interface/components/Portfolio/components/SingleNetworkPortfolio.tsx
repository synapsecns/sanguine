import React, { useEffect } from 'react'
import { Address } from 'viem'
import { useDispatch } from 'react-redux'
import _, { isArray } from 'lodash'
import { useTranslations } from 'next-intl'

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
import { PortfolioTokenVisualizer } from './PortfolioTokenVisualizer'
import { PortfolioNetwork } from './PortfolioNetwork'
import { NON_BRIDGEABLE_GAS_TOKENS } from '@/constants/tokens'
import { GasTokenAsset } from './GasTokenAsset'

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

  const t = useTranslations('Portfolio')

  const isLoading = fetchState === FetchState.LOADING

  const chain: Chain = CHAINS_BY_ID[portfolioChainId]
  const isUnsupportedChain: boolean = !chain

  const sortedTokens = sortTokensByBalanceDescending(portfolioTokens)

  const [nonBridgeableGasTokens, bridgeableTokens] =
    filterOutNonBridgeableGasTokens(sortedTokens, portfolioChainId)

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
          <PortfolioTokenVisualizer
            portfolioTokens={sortedTokens}
            portfolioChainId={portfolioChainId}
          />
        }
      >
        {isUnsupportedChain && (
          <WarningMessage
            twClassName="!p-2 !mt-0"
            message={
              <p className="leading-6">
                {t(
                  'This chain is not yet supported; New chain or token support can be discussed on'
                )}{' '}
                <a target="_blank" className="underline" href={TWITTER_URL}>
                  {t('Twitter')}
                </a>{' '}
                {t('or')}{' '}
                <a target="_blank" className="underline" href={DISCORD_URL}>
                  {t('Discord')}
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
        {nonBridgeableGasTokens &&
          nonBridgeableGasTokens.length > 0 &&
          nonBridgeableGasTokens.map(({ token, balance }: TokenAndBalance) => (
            <GasTokenAsset key={token.symbol} token={token} balance={balance} />
          ))}
        {bridgeableTokens &&
          bridgeableTokens.length > 0 &&
          bridgeableTokens.map(({ token, balance }: TokenAndBalance) => (
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

const filterOutNonBridgeableGasTokens = (
  tokens: TokenAndBalance[],
  chainId: number
): [TokenAndBalance[], TokenAndBalance[]] => {
  const nonBridgableGasTokens = NON_BRIDGEABLE_GAS_TOKENS[chainId]

  let filteredNonBridgeableGasTokens: TokenAndBalance[] = []
  let remainingTokens: TokenAndBalance[] = []

  if (!nonBridgableGasTokens && !isArray(tokens)) {
    return [filteredNonBridgeableGasTokens, remainingTokens]
  }

  if (!nonBridgableGasTokens && isArray(tokens)) {
    remainingTokens = [...tokens]
  } else {
    const gasTokenAddresses = nonBridgableGasTokens?.flatMap(
      (token) => token.address
    )
    tokens?.forEach((token) => {
      if (gasTokenAddresses?.includes(token.tokenAddress)) {
        filteredNonBridgeableGasTokens.push(token)
      } else {
        remainingTokens.push(token)
      }
    })
  }

  return [filteredNonBridgeableGasTokens, remainingTokens]
}
