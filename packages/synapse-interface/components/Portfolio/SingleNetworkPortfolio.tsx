import React, { useMemo, useCallback } from 'react'
import Image from 'next/image'
import { BigNumber } from 'ethers'
import { useDispatch } from 'react-redux'
import { useNetwork, useAccount } from 'wagmi'
import { switchNetwork } from '@wagmi/core'
import { Zero } from '@ethersproject/constants'
import {
  setFromToken,
  setFromChainId,
  updateFromValue,
} from '@/slices/bridgeSlice'
import { CHAINS_BY_ID } from '@/constants/chains'
import { TokenWithBalanceAndAllowance } from '@/utils/hooks/usePortfolioBalances'
import { approveToken } from '@/utils/approveToken'
import { formatBNToString } from '@/utils/bignumber/format'
import { Chain, Token } from '@/utils/types'
import PortfolioAccordion from './PortfolioAccordion'
import { ROUTER_ADDRESS } from '@/utils/hooks/usePortfolioBalances'
import { FetchState } from '@/utils/hooks/usePortfolioBalances'

type SingleNetworkPortfolioProps = {
  portfolioChainId: number
  connectedChainId: number
  portfolioTokens: TokenWithBalanceAndAllowance[]
  initializeExpanded: boolean
  fetchState: FetchState
}

export const SingleNetworkPortfolio = ({
  portfolioChainId,
  connectedChainId,
  portfolioTokens,
  initializeExpanded = false,
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
    <div
      data-test-id="single-network-portfolio"
      className="flex flex-col border-b border-solid border-[#28282F]"
    >
      <PortfolioAccordion
        connectedChainId={connectedChainId}
        portfolioChainId={portfolioChainId}
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
          <>
            <PortfolioTokenVisualizer
              portfolioTokens={sortedTokensForVisualizer}
            />
            <PortfolioConnectButton
              portfolioChainId={portfolioChainId}
              connectedChainId={connectedChainId}
            />
          </>
        }
      >
        <PortfolioAssetHeader />
        {!isLoading && hasNoTokenBalance && <EmptyPortfolioContent />}
        {sortedTokensWithAllowance &&
          sortedTokensWithAllowance.length > 0 &&
          sortedTokensWithAllowance.map(
            ({ token, balance }: TokenWithBalanceAndAllowance) => (
              <PortfolioTokenAsset
                token={token}
                balance={balance}
                portfolioChainId={portfolioChainId}
                connectedChainId={connectedChainId}
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
                isApproved={false}
              />
            )
          )}
      </PortfolioAccordion>
    </div>
  )
}

const EmptyPortfolioContent = () => {
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

type PortfolioTokenAssetProps = {
  token: Token
  balance: BigNumber
  portfolioChainId: number
  connectedChainId: number
  isApproved: boolean
}

const PortfolioTokenAsset = ({
  token,
  balance,
  portfolioChainId,
  connectedChainId,
  isApproved,
}: PortfolioTokenAssetProps) => {
  const dispatch = useDispatch()
  const { icon, symbol, decimals, addresses } = token
  const parsedBalance: string = formatBNToString(
    balance,
    decimals[portfolioChainId],
    3
  )
  const isDisabled: boolean = portfolioChainId !== connectedChainId
  const filteredOpacity: string = 'opacity-50 cursor-default'

  const handleTotalBalanceInputCallback = useCallback(() => {
    if (!isDisabled) {
      dispatch(setFromToken(token))
      dispatch(setFromChainId(portfolioChainId))
      dispatch(updateFromValue(balance))
    }
  }, [isDisabled, token, balance])

  return (
    <div
      data-test-id="portfolio-token-asset"
      className={`
        flex flex-row items-center text-white py-1
        ${isDisabled ? filteredOpacity : 'opacity-100'}
        `}
    >
      <div className="flex flex-row justify-between w-2/3">
        <div className="flex flex-row">
          <Image
            alt={`${symbol} img`}
            className="w-6 h-6 mr-2 rounded-md"
            src={icon}
          />
          <div>{symbol}</div>
        </div>
        <div
          onClick={handleTotalBalanceInputCallback}
          className={`
          ${isDisabled ? 'cursor-default' : 'cursor-pointer'}
          `}
        >
          {parsedBalance}
        </div>
      </div>
      <div className="flex flex-row items-center w-1/3 text-left">
        <PortfolioAssetActionButton
          connectedChainId={connectedChainId}
          token={token}
          isApproved={isApproved}
          isDisabled={isDisabled}
        />
      </div>
    </div>
  )
}

type PortfolioAssetActionButtonProps = {
  connectedChainId: number
  token: Token
  isApproved: boolean
  isDisabled: boolean
}

const PortfolioAssetActionButton = ({
  connectedChainId,
  token,
  isApproved,
  isDisabled,
}: PortfolioAssetActionButtonProps) => {
  const { address } = useAccount()
  const dispatch = useDispatch()
  const tokenAddress = token.addresses[connectedChainId]

  const handleBridgeCallback = useCallback(() => {
    if (!isDisabled) {
      dispatch(setFromToken(token))
    }
  }, [token, isDisabled])

  const handleApproveCallback = useCallback(() => {
    if (!isDisabled) {
      return approveToken(ROUTER_ADDRESS, connectedChainId, tokenAddress)
    }
  }, [connectedChainId, tokenAddress, address, isDisabled])

  const buttonClassName = `
    flex ml-auto justify-center
    py-1 px-6 ml-2 rounded-3xl
    transform-gpu transition-all duration-75
    ${isDisabled ? 'hover:cursor-default' : 'hover:cursor-pointer'}
  `
  return (
    <React.Fragment>
      {isApproved ? (
        <button
          data-test-id="portfolio-asset-action-button"
          className={`
            ${buttonClassName}
            border-2 border-[#D747FF]
          `}
          onClick={handleBridgeCallback}
        >
          Send
        </button>
      ) : (
        <button
          data-test-id="portfolio-asset-action-button"
          className={`
            ${buttonClassName}
            border-2 border-[#28282F] border-opacity-50
          `}
          onClick={handleApproveCallback}
        >
          Approve
        </button>
      )}
    </React.Fragment>
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
      className="flex flex-row justify-between flex-1"
    >
      <div className="flex flex-row items-center">
        <Image
          className="mr-4 rounded-md w-7 h-7"
          alt={`${displayName} img`}
          src={chainIcon}
        />
        <div className="font-medium text-white text-18">{displayName}</div>
      </div>
    </div>
  )
}

const PortfolioTokenVisualizer = ({
  portfolioTokens,
}: {
  portfolioTokens: TokenWithBalanceAndAllowance[]
}) => {
  const hasOneToken = portfolioTokens && portfolioTokens.length > 0
  const hasTwoTokens = portfolioTokens && portfolioTokens.length > 1
  const numOverTwoTokens =
    portfolioTokens && portfolioTokens.length - 2 > 0
      ? portfolioTokens.length - 2
      : 0

  return (
    <div data-test-id="portfolio-token-visualizer" className="flex flex-row">
      {hasOneToken && (
        <Image
          className="w-6 h-6 rounded-md"
          alt={`${portfolioTokens[0].token.symbol} img`}
          src={portfolioTokens[0].token.icon}
        />
      )}
      {hasTwoTokens && (
        <Image
          className="w-6 h-6 ml-1 rounded-md"
          alt={`${portfolioTokens[1].token.symbol} img`}
          src={portfolioTokens[1].token.icon}
        />
      )}
      {numOverTwoTokens > 0 && (
        <div className="ml-1 text-white">+ {numOverTwoTokens}</div>
      )}
    </div>
  )
}

type PortfolioConnectButton = {
  portfolioChainId: number
  connectedChainId: number
}

const PortfolioConnectButton = ({
  portfolioChainId,
  connectedChainId,
}: PortfolioConnectButton) => {
  const isCurrentlyConnectedNetwork: boolean = useMemo(() => {
    return portfolioChainId === connectedChainId
  }, [portfolioChainId, connectedChainId])

  return (
    <div data-test-id="portfolio-connect-button" className="ml-2">
      {isCurrentlyConnectedNetwork ? (
        <ConnectedButton />
      ) : (
        <ConnectButton chainId={portfolioChainId} />
      )}
    </div>
  )
}

const ConnectedButton = () => {
  const buttonClassName = `
  h-9 flex items-center justify-end w-36
  text-base text-white px-4 py-2 rounded-3xl
  text-center transform-gpu transition-all duration-75
  hover:cursor-default
  `

  return (
    <button data-test-id="connected-button" className={buttonClassName}>
      <div className="flex flex-row text-sm">
        <div
          className={`
          my-auto ml-auto mr-2
          w-2 h-2
          bg-green-500 rounded-full`}
        />
        Connected
      </div>
    </button>
  )
}

const ConnectButton = ({ chainId }: { chainId: number }) => {
  const dispatch = useDispatch()

  const handleConnectNetwork = async () => {
    await switchNetwork({ chainId: chainId }).then((success) => {
      success && dispatch(setFromChainId(chainId))
    })
  }

  const buttonClassName = `
    h-9 flex items-right justify-end w-36
    text-base text-white px-4 py-2 rounded-3xl
    text-center transform-gpu transition-all duration-75
    `

  return (
    <button
      data-test-id="connect-button"
      className={buttonClassName}
      onClick={handleConnectNetwork}
    >
      <div className="flex flex-row text-sm">
        <div
          className={`
          my-auto ml-auto mr-2
          text-transparent w-2 h-2
          border border-indigo-300 border-solid rounded-full`}
        />
        Connect
      </div>
    </button>
  )
}

export const PortfolioAssetHeader = () => {
  return (
    <div
      data-test-id="portfolio-asset-header"
      className="flex text-[#CCCAD3BF] my-2"
    >
      <div className="flex flex-row justify-between w-2/3 text-left">
        <div>Token</div>
        <div>Amount</div>
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
      } else if (token.allowance.gt(Zero)) {
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
