import React, { useMemo, useCallback, useState } from 'react'
import { useAppDispatch } from '@/store/hooks'
import { useAccount } from 'wagmi'
import {
  setFromChainId,
  setFromToken,
  updateFromValue,
} from '@/slices/bridge/reducer'
import { Token } from '@/utils/types'
import { formatBigIntToString } from '@/utils/bigint/format'
import { CHAINS_BY_ID } from '@/constants/chains'
import { inputRef } from '../../../StateManagedBridge/InputContainer'
import { approveToken } from '@/utils/approveToken'
import { Address, switchNetwork } from '@wagmi/core'
import Image from 'next/image'
import { toast } from 'react-hot-toast'
import {
  ROUTER_ADDRESS,
  CCTP_ROUTER_ADDRESS,
  Allowances,
} from '@/utils/actions/fetchPortfolioBalances'
import { useBridgeState } from '@/slices/bridge/hooks'
import { fetchAndStoreSingleTokenAllowance } from '@/slices/portfolio/hooks'
import { AVALANCHE, ETH, ARBITRUM } from '@/constants/chains/master'
import { USDC } from '@/constants/tokens/bridgeable'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'

const handleFocusOnInput = () => {
  inputRef.current.focus()
}

function checkCCTPChainConditions(
  fromChainId: number,
  toChainId: number
): boolean {
  const CctpPairs = new Set([
    `${ETH.id}-${ARBITRUM.id}`,
    `${ARBITRUM.id}-${ETH.id}`,
    `${ETH.id}-${AVALANCHE.id}`,
    `${AVALANCHE.id}-${ETH.id}`,
    `${ARBITRUM.id}-${AVALANCHE.id}`,
    `${AVALANCHE.id}-${ARBITRUM.id}`,
  ])

  return CctpPairs.has(`${fromChainId}-${toChainId}`)
}

function checkIfUsingCCTP({
  fromChainId,
  fromToken,
  toChainId,
  toToken,
}: {
  fromChainId: number
  fromToken: Token
  toChainId: number
  toToken: Token
}): boolean {
  const isTokensUSDC: boolean = fromToken === USDC && toToken === USDC
  const isSupportedCCTPChains: boolean = checkCCTPChainConditions(
    fromChainId,
    toChainId
  )

  return isTokensUSDC && isSupportedCCTPChains
}

type PortfolioTokenAssetProps = {
  token: Token
  balance: bigint
  allowances?: Allowances
  portfolioChainId: number
  connectedChainId: number
  isApproved: boolean
}

export const PortfolioTokenAsset = ({
  token,
  balance,
  allowances,
  portfolioChainId,
  connectedChainId,
  isApproved,
}: PortfolioTokenAssetProps) => {
  const dispatch = useAppDispatch()
  const { fromChainId, fromToken, toChainId, toToken } = useBridgeState()
  const { address } = useAccount()
  const { icon, symbol, decimals, addresses } = token as Token

  const parsedBalance: string = useMemo(() => {
    const formattedBalance = formatBigIntToString(
      balance,
      decimals[portfolioChainId],
      3
    )
    return balance > 0n && hasOnlyZeroes(formattedBalance)
      ? '< 0.001'
      : formattedBalance
  }, [balance, portfolioChainId])

  const isCCTP: boolean = checkIfUsingCCTP({
    fromChainId,
    fromToken,
    toChainId,
    toToken,
  })

  const tokenRouterAddress: string = isCCTP
    ? CCTP_ROUTER_ADDRESS
    : ROUTER_ADDRESS

  const bridgeAllowance: bigint = allowances?.[tokenRouterAddress]

  const parsedAllowance: string =
    bridgeAllowance &&
    formatBigIntToString(bridgeAllowance, decimals[portfolioChainId], 3)

  const currentChainName: string = CHAINS_BY_ID[portfolioChainId].name

  const tokenAddress: string = addresses[portfolioChainId]

  const isCurrentlyConnected: boolean = portfolioChainId === connectedChainId

  const isTokenSelected: boolean = useMemo(() => {
    return fromToken === token && fromChainId === portfolioChainId
  }, [fromChainId, fromToken, token, portfolioChainId])

  const hasAllowanceButLessThanBalance: boolean =
    bridgeAllowance && balance > bridgeAllowance

  const isDisabled: boolean = false

  const handleTotalBalanceInputCallback = useCallback(async () => {
    await dispatch(setFromChainId(portfolioChainId as number))
    await dispatch(setFromToken(token as Token))
    await dispatch(
      await updateFromValue(
        formatBigIntToString(
          balance,
          token.decimals[portfolioChainId]
        ) as string
      )
    )
    handleFocusOnInput()
  }, [isDisabled, token, balance, portfolioChainId])

  const handleSelectFromTokenCallback = useCallback(() => {
    dispatch(setFromChainId(portfolioChainId as number))
    dispatch(setFromToken(token as Token))
    handleFocusOnInput()
  }, [token, isDisabled, portfolioChainId])

  const handleApproveCallback = useCallback(async () => {
    if (isCurrentlyConnected) {
      dispatch(setFromChainId(portfolioChainId as number))
      dispatch(setFromToken(token as Token))
      try {
        await approveToken(
          tokenRouterAddress,
          connectedChainId,
          tokenAddress
        ).then((success) => {
          dispatch(
            fetchAndStoreSingleTokenAllowance({
              routerAddress: tokenRouterAddress as Address,
              tokenAddress: tokenAddress as Address,
              address: address as Address,
              chainId: portfolioChainId as number,
            })
          )
        })
      } catch (error) {
        toast.error(
          `Failed to approve ${token.symbol} token on ${currentChainName} network`,
          {
            id: 'approve-in-progress-popup',
            duration: 5000,
          }
        )
      }
    } else {
      try {
        await switchNetwork({ chainId: portfolioChainId })
        await approveToken(
          tokenRouterAddress,
          portfolioChainId,
          tokenAddress
        ).then((success) => {
          success &&
            dispatch(
              fetchAndStoreSingleTokenAllowance({
                routerAddress: tokenRouterAddress as Address,
                tokenAddress: tokenAddress as Address,
                address: address as Address,
                chainId: portfolioChainId as number,
              })
            )
        })
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
    tokenRouterAddress,
  ])

  return (
    <div
      data-test-id="portfolio-token-asset"
      className={`
        flex flex-row items-center p-2 text-white
        ${
          isTokenSelected
            ? 'bg-tint border-y border-surface'
            : 'border-y border-transparent'
        }
      `}
    >
      <div className="flex flex-row justify-between w-2/3">
        <div
          onClick={handleSelectFromTokenCallback}
          className={`
            flex flex-row px-2 py-2 mb-auto
            cursor-pointer hover:bg-tint active:opacity-[67%]
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
        <div className="flex flex-col">
          <div
            onClick={handleTotalBalanceInputCallback}
            className={`
              p-2 ml-auto cursor-pointer
              hover:bg-tint active:opacity-[67%]
            `}
          >
            {parsedBalance}
          </div>
          {hasAllowanceButLessThanBalance && (
            <HoverClickableText
              defaultText={`${parsedAllowance} ${
                isCCTP ? 'approved (CCTP)' : 'approved'
              }`}
              hoverText="Increase Limit"
              callback={handleApproveCallback}
            />
          )}
        </div>
      </div>
      <div className="flex flex-row items-center w-1/3 py-2 mb-auto text-left">
        <PortfolioAssetActionButton
          selectCallback={handleSelectFromTokenCallback}
          approveCallback={handleApproveCallback}
          isApproved={isApproved}
          isDisabled={isDisabled}
        />
      </div>
    </div>
  )
}

export const HoverClickableText = ({
  defaultText,
  hoverText,
  callback,
}: {
  defaultText: string
  hoverText: string
  callback: () => void
}) => {
  const [isHovered, setIsHovered] = useState<boolean>(false)
  return (
    <div
      data-test-id="hover-clickable-text"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
      onClick={callback}
      className={`
        group px-2
        text-[#A3A3C2]
        hover:text-[#75E6F0]
        hover:underline
        hover:cursor-pointer
        active:opacity-[67%]
      `}
    >
      <div className="text-[14px]">{isHovered ? hoverText : defaultText}</div>
    </div>
  )
}

type PortfolioAssetActionButtonProps = {
  selectCallback: () => void
  approveCallback: () => Promise<void>
  isApproved: boolean
  isDisabled: boolean
}

const PortfolioAssetActionButton = ({
  selectCallback,
  approveCallback,
  isApproved,
  isDisabled,
}: PortfolioAssetActionButtonProps) => {
  const buttonClassName = `
    flex ml-auto justify-center
    py-1 px-6 ml-2 rounded-sm
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
            border border-synapsePurple mr-1
            hover:bg-tint active:opacity-[67%]
          `}
          onClick={selectCallback}
        >
          Select
        </button>
      ) : (
        <button
          data-test-id="portfolio-asset-action-button"
          className={`
            ${buttonClassName}
            border border-separator
            hover:bg-tint active:opacity-[67%]
          `}
          onClick={approveCallback}
        >
          Approve
        </button>
      )}
    </React.Fragment>
  )
}
