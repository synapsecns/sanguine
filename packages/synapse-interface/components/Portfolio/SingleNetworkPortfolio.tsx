import React, { useMemo } from 'react'
import { useNetwork } from 'wagmi'
import { switchNetwork } from '@wagmi/core'
import { Zero } from '@ethersproject/constants'
import { TokenWithBalanceAndAllowance } from '@/utils/hooks/usePortfolioBalances'
import { Chain } from '@/utils/types'
import { CHAINS_BY_ID } from '@/constants/chains'
import Image from 'next/image'
import { Token } from '@/utils/types'
import { BigNumber } from 'ethers'
import { formatBNToString } from '@/utils/bignumber/format'
import PortfolioAccordion from './PortfolioAccordion'

type SingleNetworkPortfolioProps = {
  portfolioChainId: number
  connectedChainId: number
  tokens: TokenWithBalanceAndAllowance[]
  initializeExpanded: boolean
}

export const SingleNetworkPortfolio = ({
  portfolioChainId,
  connectedChainId,
  tokens,
  initializeExpanded = false,
}: SingleNetworkPortfolioProps) => {
  const currentChain: Chain = CHAINS_BY_ID[portfolioChainId]

  const [tokensWithAllowance, tokensWithoutAllowance] =
    separateTokensByAllowance(tokens)

  const sortedTokensWithAllowance: TokenWithBalanceAndAllowance[] =
    sortByBalanceDescending(tokensWithAllowance)
  const sortedTokensWithoutAllowance: TokenWithBalanceAndAllowance[] =
    sortByBalanceDescending(tokensWithoutAllowance)

  const shouldShowDivider: boolean =
    sortedTokensWithAllowance.length > 0 &&
    sortedTokensWithoutAllowance.length > 0

  return (
    <div
      data-test-id="single-network-portfolio"
      className="flex flex-col border-b border-solid border-[#28282F] pt-4 pb-2 mt-2"
    >
      <PortfolioAccordion
        initializeExpanded={initializeExpanded}
        header={
          <PortfolioNetwork
            displayName={currentChain.name}
            chainIcon={currentChain.chainImg}
            chainId={portfolioChainId}
          />
        }
      >
        <PortfolioAssetHeader />
        {sortedTokensWithAllowance.length > 0 &&
          sortedTokensWithAllowance.map(
            ({ token, balance, allowance }: TokenWithBalanceAndAllowance) => (
              <PortfolioTokenAsset
                token={token}
                balance={balance}
                chainId={portfolioChainId}
                connectedChainId={connectedChainId}
                isApproved={true}
              />
            )
          )}
        {shouldShowDivider && (
          <div className="border-b border-solid border-[#28282F] my-2" />
        )}
        {sortedTokensWithoutAllowance.length > 0 &&
          sortedTokensWithoutAllowance.map(
            ({ token, balance, allowance }: TokenWithBalanceAndAllowance) => (
              <PortfolioTokenAsset
                token={token}
                balance={balance}
                chainId={portfolioChainId}
                connectedChainId={connectedChainId}
                isApproved={false}
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
  chainId: number
  connectedChainId: number
  isApproved: boolean
}

const PortfolioTokenAsset = ({
  token,
  balance,
  chainId,
  connectedChainId,
  isApproved,
}: PortfolioTokenAssetProps) => {
  const { icon, symbol, decimals, addresses } = token
  const parsedBalance = formatBNToString(balance, decimals[chainId], 3)
  const isDisabled = chainId !== connectedChainId
  const filteredOpacity = 'opacity-50 cursor-default'
  return (
    <div
      data-test-id="Portfolio Token Asset"
      className={`
        flex flex-row items-center my-2 text-white
        ${isDisabled && filteredOpacity}
        `}
    >
      <div
        className={`
          flex flex-row w-1/2 text-left
          ${!isApproved && 'opacity-50'}
          `}
      >
        <Image
          alt={`${symbol} img`}
          className="w-6 h-6 mr-2 rounded-md"
          src={icon}
        />
        <div>{symbol}</div>
      </div>
      <div className="flex flex-row items-center w-1/2 text-left">
        <div className={!isApproved && 'opacity-50'}>{parsedBalance}</div>
        <PortfolioAssetActionButton isApproved={isApproved} />
      </div>
    </div>
  )
}

const PortfolioAssetActionButton = ({
  isApproved,
}: {
  isApproved: boolean
}) => {
  const buttonClassName = `
    flex ml-auto justify-center
    w-36 py-1 rounded-3xl
    transform-gpu transition-all duration-75
    hover:cursor-pointer
  `
  return (
    <React.Fragment>
      {isApproved ? (
        <button
          className={`
            ${buttonClassName}
            border-2 border-[#D747FF]
          `}
        >
          Bridge
        </button>
      ) : (
        <button
          className={`
            ${buttonClassName}
            border-2 border-[#28282F] border-opacity-50
          `}
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
  chainId: number
}

const PortfolioNetwork = ({
  displayName,
  chainIcon,
  chainId,
}: PortfolioNetworkProps) => {
  const { chain } = useNetwork()
  const isCurrentlyConnectedNetwork: boolean = useMemo(() => {
    return chainId === chain.id
  }, [chain.id])

  return (
    <div className="flex flex-row justify-between flex-1 mb-4">
      <div className="flex flex-row items-center">
        <Image
          className="mr-4 rounded-md w-7 h-7"
          alt={`${displayName} img`}
          src={chainIcon}
        />
        <div className="font-medium text-white text-18">{displayName}</div>
      </div>
      {isCurrentlyConnectedNetwork ? (
        <ConnectedButton />
      ) : (
        <ConnectButton chainId={chainId} />
      )}
    </div>
  )
}

const ConnectedButton = () => {
  const buttonClassName = `
  h-8 flex items-center
  text-base text-white px-4 py-2 rounded-3xl
  text-center transform-gpu transition-all duration-75
  border-2 border-[#D747FF] radial-gradient-bg
  hover:cursor-default
  `

  return <button className={buttonClassName}>Connected</button>
}

const ConnectButton = ({ chainId }: { chainId: number }) => {
  const handleConnectNetwork = async () => {
    await switchNetwork({ chainId: chainId })
  }

  const buttonClassName = `
    h-8 flex items-center
    text-base text-white px-4 py-2 rounded-3xl
    text-center transform-gpu transition-all duration-75
    border-2 border-[#101018]
    hover:cursor-pointer
  `

  return (
    <button className={buttonClassName} onClick={handleConnectNetwork}>
      Connect
    </button>
  )
}

export const PortfolioAssetHeader = () => {
  return (
    <div className="flex text-[#CCCAD3BF] my-2">
      <div className="w-1/2 text-left">Token</div>
      <div className="w-1/2 text-left">Amount</div>
    </div>
  )
}

function separateTokensByAllowance(
  tokens: TokenWithBalanceAndAllowance[]
): [TokenWithBalanceAndAllowance[], TokenWithBalanceAndAllowance[]] {
  const tokensWithAllowance: TokenWithBalanceAndAllowance[] = []
  const tokensWithoutAllowance: TokenWithBalanceAndAllowance[] = []

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
  return tokens.sort(
    (a: TokenWithBalanceAndAllowance, b: TokenWithBalanceAndAllowance) =>
      b.balance.gt(a.balance) ? 1 : -1
  )
}
