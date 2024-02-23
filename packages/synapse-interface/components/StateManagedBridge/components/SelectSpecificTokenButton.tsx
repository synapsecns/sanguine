import _ from 'lodash'
import { Token } from '@/utils/types'
import { useBridgeState } from '@/slices/bridge/hooks'
import { findChainIdsWithPausedToken } from '@/constants/tokens'
import LoadingDots from '@tw/LoadingDots'

import { getUnderlyingBridgeTokens } from '@/utils/getUnderlyingBridgeTokens'
import { ARBITRUM, AVALANCHE, ETH } from '@/constants/chains/master'

import { SelectTokenButton } from '@/components/bridgeSwap/SelectTokenButton'

export const SelectSpecificTokenButton = ({
  showAllChains,
  isOrigin,
  token,
  active,
  selectedToken,
  onClick,
  alternateBackground = false,
  isLoadingExchangeRate = false,
  exchangeRate,
  isBestExchangeRate = false,
  estimatedDurationInSeconds,
}: {
  showAllChains?: boolean
  isOrigin: boolean
  token: Token
  active: boolean
  selectedToken: Token
  onClick: () => void
  alternateBackground?: boolean
  isLoadingExchangeRate?: boolean
  exchangeRate?: string
  isBestExchangeRate?: boolean
  estimatedDurationInSeconds?: number
}) => {
  const { fromChainId, toChainId } = useBridgeState()

  return (
    <SelectTokenButton
      showAllChains={showAllChains}
      token={token}
      active={active}
      selectedToken={selectedToken}
      chainId={isOrigin ? fromChainId : toChainId}
      isOrigin={isOrigin}
      onClick={onClick}
      isEligible={isTokenEligible(token)}
      pausedChainIds={findChainIdsWithPausedToken(token.routeSymbol)}
      alternateBackground={alternateBackground}
    >
      {isLoadingExchangeRate ? (
        <LoadingDots className="mr-8 opacity-50" />
      ) : (
        <>
          {exchangeRate && (
            isBestExchangeRate
              ? <OptionTag type={BestOptionType.RATE} />
              :
                <OptionDetails
                  exchangeRate={exchangeRate}
                  estimatedDurationInSeconds={estimatedDurationInSeconds}
                />
          )}
        </>
      )}
    </SelectTokenButton>
  )
}

export enum BestOptionType {
  RATE = 'Best rate',
  SPEED = 'Fastest',
}

export const OptionTag = ({ type }: { type: BestOptionType }) => {
  return (
    <div
      data-test-id="option-tag"
      className="flex px-3 py-0.5 mr-3 text-sm whitespace-nowrap text-primary rounded-xl"
      style={{
        background:
          'linear-gradient(to right, rgba(128, 0, 255, 0.2), rgba(255, 0, 191, 0.2))',
      }}
    >{`${type}`}</div>
  )
}

export const OptionDetails = ({
  exchangeRate,
  estimatedDurationInSeconds,
}: {
  exchangeRate: string
  estimatedDurationInSeconds: number
}) => {
  let showTime
  let timeUnit

  if (estimatedDurationInSeconds > 60) {
    showTime = Math.floor(estimatedDurationInSeconds / 60)
    timeUnit = 'min'
  } else {
    showTime = estimatedDurationInSeconds
    timeUnit = 'seconds'
  }

  return (
    <div data-test-id="option-details" className="flex flex-col">
      <div className="flex items-center font-normal">
        <div className="flex text-sm text-secondary whitespace-nowrap">
          1&nbsp;:&nbsp;
        </div>
        <div className="mb-[1px] text-primary">{exchangeRate}</div>
      </div>
      <div className="text-xs text-right text-secondary">
        {showTime} {timeUnit}
      </div>
    </div>
  )
}




/*
Synapse:Bridge
  Tokens: nETH, nUSD, GMX
  From Any to ARB: all txs (don't limit this to "user has to receive ETH / USDC / ...")
  ARB to ETH txs: nETH, nUSD
  ARB to AVAX txs: GMX

Synapse:CCTP
  Tokens: USDC
  Any to ARB: all txs
  ARB to ETH: all txs

Synapse: RFQ
  Tokens: USDC
  Any to ARB: all txs
  ARB to ETH: all txs
*/

const isTokenEligible = (token: Token) => {
  const { fromChainId, toChainId, bridgeQuote } = useBridgeState()

  const underlyingBridgeTokens = getUnderlyingBridgeTokens(token, fromChainId)

  if (!underlyingBridgeTokens) {
    return false
  }

  const includesUSDC = underlyingBridgeTokens.includes('USDC')

  return (
    (includesUSDC && toChainId === ARBITRUM.id) ||
    (includesUSDC &&
      fromChainId === ARBITRUM.id &&
      toChainId === ETH.id) ||
    (includesUSDC && toChainId === ARBITRUM.id) ||
    (includesUSDC &&
      fromChainId === ARBITRUM.id &&
      toChainId === ETH.id) ||
    (_.some(['nETH', 'nUSD', 'GMX'], (value) =>
      _.includes(underlyingBridgeTokens, value)
    ) &&
      toChainId === ARBITRUM.id) ||
    (_.some(['nETH', 'nUSD'], (value) =>
      _.includes(underlyingBridgeTokens, value)
    ) &&
      fromChainId === ARBITRUM.id &&
      toChainId === ETH.id) ||
    (_.some(['GMX'], (value) => _.includes(underlyingBridgeTokens, value)) &&
      fromChainId === ARBITRUM.id &&
      toChainId === AVALANCHE.id)
  )
}

