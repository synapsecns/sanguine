import React, { useCallback } from 'react'
import Image from 'next/image'
import { BigNumber } from 'ethers'
import { useDispatch } from 'react-redux'
import { useAccount } from 'wagmi'
import { Zero } from '@ethersproject/constants'
import {
  setFromToken,
  setFromChainId,
  updateFromValue,
} from '@/slices/bridgeSlice'
import { CHAINS_BY_ID } from '@/constants/chains'
import { TokenWithBalanceAndAllowance } from '@/utils/hooks/usePortfolioBalances'
import { usePortfolioBalancesAndAllowances } from '@/utils/hooks/usePortfolioBalances'
import { approveToken } from '@/utils/approveToken'
import { formatBNToString } from '@/utils/bignumber/format'
import { Chain, Token } from '@/utils/types'
import PortfolioAccordion from './PortfolioAccordion'
import { PortfolioConnectButton } from './PortfolioConnectButton'
import { ROUTER_ADDRESS } from '@/utils/hooks/usePortfolioBalances'
import { FetchState } from '@/utils/hooks/usePortfolioBalances'
import { toast } from 'react-hot-toast'

type SingleNetworkPortfolioProps = {
  portfolioChainId: number
  connectedChainId: number
  selectedFromChainId: number
  portfolioTokens: TokenWithBalanceAndAllowance[]
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
                selectedFromChainId={selectedFromChainId}
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
                selectedFromChainId={selectedFromChainId}
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
  selectedFromChainId: number
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
  const isDisabled: boolean = false
  const filteredOpacity: string = 'opacity-50 cursor-default'

  const handleTotalBalanceInputCallback = useCallback(() => {
    // if (!isDisabled) {
    //   dispatch(setFromToken(token))
    //   dispatch(setFromChainId(portfolioChainId))
    //   dispatch(updateFromValue(balance))
    // }
  }, [isDisabled, token, balance])

  return (
    <div
      data-test-id="portfolio-token-asset"
      className="flex flex-row items-center py-2 text-white "
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
          className="cursor-default"
        >
          {parsedBalance}
        </div>
      </div>
      <div className="flex flex-row items-center w-1/3 text-left">
        <PortfolioAssetActionButton
          connectedChainId={connectedChainId}
          portfolioChainId={portfolioChainId}
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
  portfolioChainId: number
  token: Token
  isApproved: boolean
  isDisabled: boolean
}

const PortfolioAssetActionButton = ({
  connectedChainId,
  portfolioChainId,
  token,
  isApproved,
  isDisabled,
}: PortfolioAssetActionButtonProps) => {
  const dispatch = useDispatch()
  const { address } = useAccount()
  const { fetchPortfolioBalances } = usePortfolioBalancesAndAllowances()
  const currentChainName: string = CHAINS_BY_ID[portfolioChainId].name
  const isCurrentlyConnected: boolean = portfolioChainId === connectedChainId
  const tokenAddress: string = token.addresses[connectedChainId]

  const handleBridgeCallback = useCallback(() => {
    if (!isDisabled) {
      dispatch(setFromChainId(portfolioChainId))
      dispatch(setFromToken(token))
    }
  }, [token, isDisabled, portfolioChainId])

  const handleApproveCallback = useCallback(async () => {
    if (!isDisabled && isCurrentlyConnected) {
      dispatch(setFromToken(token))
      return await approveToken(
        ROUTER_ADDRESS,
        connectedChainId,
        tokenAddress
      ).then((success) => {
        success && fetchPortfolioBalances()
      })
    } else {
      toast.error(
        `Connect to ${currentChainName} network to approve ${token.symbol} token`,
        {
          id: 'approve-in-progress-popup',
          duration: Infinity,
        }
      )
    }
  }, [
    connectedChainId,
    tokenAddress,
    address,
    isDisabled,
    portfolioChainId,
    token,
    isCurrentlyConnected,
  ])

  const buttonClassName = `
    flex ml-auto justify-center
    py-1 px-6 ml-2 rounded-3xl
    transform-gpu transition-all duration-75
    ${isDisabled ? 'hover:cursor-default' : 'hover:cursor-pointer'}
  `

  const activeButtonClass = `active:opacity-[67%]`

  return (
    <React.Fragment>
      {isApproved ? (
        <button
          data-test-id="portfolio-asset-action-button"
          className={`
            ${buttonClassName}
            ${activeButtonClass}
            border border-[#D747FF]
            hover:bg-[#272731]
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
            ${activeButtonClass}
            border border-[#3D3D5C]
          hover:border-[#A3A3C2]
          hover:bg-[#272731]
          active:border-[#A3A3C2]
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
      className="flex flex-row justify-between flex-1 py-4 pl-2"
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
    <div
      data-test-id="portfolio-token-visualizer"
      className="flex flex-row items-center"
    >
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
