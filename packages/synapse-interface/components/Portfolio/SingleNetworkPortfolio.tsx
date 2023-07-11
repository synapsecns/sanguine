import React, { useCallback, useMemo, useRef } from 'react'
import Image from 'next/image'
import { BigNumber } from 'ethers'
import { useDispatch } from 'react-redux'
import { useAccount } from 'wagmi'
import { switchNetwork } from '@wagmi/core'
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
import { EmptyPortfolioContent } from './PortfolioContent'
import { ROUTER_ADDRESS } from '@/utils/hooks/usePortfolioBalances'
import { FetchState } from '@/utils/hooks/usePortfolioBalances'
import { toast } from 'react-hot-toast'

type SingleNetworkPortfolioProps = {
  portfolioChainId: number
  connectedChainId: number
  selectedFromChainId: number
  portfolioTokens: TokenWithBalanceAndAllowance[]
  initializeExpanded: boolean
  fetchPortfolioBalancesCallback: () => Promise<void>
  fetchState: FetchState
  portfolioRef: React.RefObject<HTMLDivElement>
}

export const SingleNetworkPortfolio = ({
  portfolioChainId,
  connectedChainId,
  selectedFromChainId,
  portfolioTokens,
  initializeExpanded = false,
  fetchPortfolioBalancesCallback,
  fetchState,
  portfolioRef,
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
            ({ token, balance, allowance }: TokenWithBalanceAndAllowance) => (
              <PortfolioTokenAsset
                token={token}
                balance={balance}
                allowance={allowance}
                portfolioChainId={portfolioChainId}
                connectedChainId={connectedChainId}
                fetchPortfolioBalancesCallback={fetchPortfolioBalancesCallback}
                isApproved={true}
                portfolioRef={portfolioRef}
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
                portfolioRef={portfolioRef}
              />
            )
          )}
      </PortfolioAccordion>
    </div>
  )
}

type PortfolioTokenAssetProps = {
  token: Token
  balance: BigNumber
  allowance?: BigNumber
  portfolioChainId: number
  connectedChainId: number
  isApproved: boolean
  fetchPortfolioBalancesCallback: () => Promise<void>
  portfolioRef: React.RefObject<HTMLDivElement>
}

const PortfolioTokenAsset = ({
  token,
  balance,
  allowance,
  portfolioChainId,
  connectedChainId,
  isApproved,
  fetchPortfolioBalancesCallback,
  portfolioRef,
}: PortfolioTokenAssetProps) => {
  function scrollToRef() {
    if (portfolioRef.current) {
      portfolioRef.current.scrollIntoView({ behavior: 'smooth' })
    }
  }

  const dispatch = useDispatch()
  const { address } = useAccount()
  const { icon, symbol, decimals, addresses } = token

  function hasOnlyZeros(input: string): boolean {
    return /^0+(\.0+)?$/.test(input)
  }

  const parsedBalance: string = useMemo(() => {
    const formattedBalance = formatBNToString(
      balance,
      decimals[portfolioChainId],
      3
    )
    return balance.gt(0) && hasOnlyZeros(formattedBalance)
      ? '< 0.001'
      : formattedBalance
  }, [balance, portfolioChainId])

  const parsedAllowance: string =
    allowance && formatBNToString(allowance, decimals[portfolioChainId], 3)

  const currentChainName: string = CHAINS_BY_ID[portfolioChainId].name

  const tokenAddress: string = addresses[portfolioChainId]

  const isCurrentlyConnected: boolean = portfolioChainId === connectedChainId

  const hasAllowanceButLessThanBalance: boolean =
    allowance && balance.gt(allowance)

  const isDisabled: boolean = false

  const handleTotalBalanceInputCallback = useCallback(() => {
    return //remove this when callback is ready to implement
    if (!isDisabled) {
      dispatch(setFromToken(token))
      dispatch(setFromChainId(portfolioChainId))
      dispatch(updateFromValue(balance))
    }
  }, [isDisabled, token, balance])

  const handleSelectFromTokenCallback = useCallback(() => {
    dispatch(setFromChainId(portfolioChainId))
    dispatch(setFromToken(token))
    scrollToRef()
  }, [token, isDisabled, portfolioChainId])

  const handleApproveCallback = useCallback(async () => {
    if (isCurrentlyConnected) {
      dispatch(setFromChainId(portfolioChainId))
      dispatch(setFromToken(token))
      await approveToken(ROUTER_ADDRESS, connectedChainId, tokenAddress).then(
        (success) => {
          success && fetchPortfolioBalancesCallback()
        }
      )
    } else {
      try {
        await switchNetwork({ chainId: portfolioChainId })
        await approveToken(ROUTER_ADDRESS, portfolioChainId, tokenAddress).then(
          (success) => {
            success && fetchPortfolioBalancesCallback()
          }
        )
      } catch (error) {
        toast.error(
          `Failed to approve ${token.symbol} token on ${currentChainName} network`,
          {
            id: 'approve-in-progress-popup',
            duration: 5000,
          }
        )
      }
    }
  }, [
    token,
    address,
    tokenAddress,
    connectedChainId,
    portfolioChainId,
    isCurrentlyConnected,
    isDisabled,
  ])

  return (
    <div
      data-test-id="portfolio-token-asset"
      className="flex flex-row flex-wrap items-center py-2 text-white"
    >
      <div className="flex flex-row justify-between w-2/3">
        <div
          onClick={handleSelectFromTokenCallback}
          className={`
            flex flex-row px-2 py-2
            hover:cursor-pointer
            hover:bg-[#272731]
          `}
        >
          <Image
            loading="lazy"
            alt={`${symbol} img`}
            className="w-6 h-6 mr-2 rounded-md"
            src={icon}
          />
          <div>{symbol}</div>
        </div>
        <div
          onClick={handleTotalBalanceInputCallback}
          className="py-2 cursor-default"
        >
          {parsedBalance}
        </div>
      </div>
      <div className="flex flex-row items-center w-1/3 text-left">
        <PortfolioAssetActionButton
          token={token}
          connectedChainId={connectedChainId}
          portfolioChainId={portfolioChainId}
          sendCallback={handleSelectFromTokenCallback}
          approveCallback={handleApproveCallback}
          isApproved={isApproved}
          isDisabled={isDisabled}
        />
      </div>
      {hasAllowanceButLessThanBalance && (
        <a
          onClick={handleApproveCallback}
          className={`
            text-[#A3A3C2] text-xs pt-1 pl-2
            hover:text-[#75E6F0]
            hover:underline
            hover:cursor-pointer
            active:opacity-[67%]
          `}
        >
          {parsedAllowance} approved ({parsedBalance} available)
        </a>
      )}
    </div>
  )
}

type PortfolioAssetActionButtonProps = {
  token: Token
  connectedChainId: number
  portfolioChainId: number
  sendCallback: () => void
  approveCallback: () => Promise<void>
  isApproved: boolean
  isDisabled: boolean
}

const PortfolioAssetActionButton = ({
  token,
  connectedChainId,
  portfolioChainId,
  sendCallback,
  approveCallback,
  isApproved,
  isDisabled,
}: PortfolioAssetActionButtonProps) => {
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
            border border-[#D747FF]
            hover:bg-[#272731]
            active:opacity-[67%]
          `}
          onClick={sendCallback}
        >
          Send
        </button>
      ) : (
        <button
          data-test-id="portfolio-asset-action-button"
          className={`
            ${buttonClassName}
            border border-[#3D3D5C]
            hover:border-[#A3A3C2]
            hover:bg-[#272731]
            active:border-[#A3A3C2]
            active:opacity-[67%]
          `}
          onClick={approveCallback}
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
    </div>
  )
}

export const PortfolioAssetHeader = () => {
  return (
    <div
      data-test-id="portfolio-asset-header"
      className="flex text-[#CCCAD3BF] my-2 pl-2"
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
