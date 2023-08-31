import _ from 'lodash'
import { CHAINS_BY_ID } from '@constants/chains'
import {
  BridgeRoutes,
  EXISTING_BRIDGE_ROUTES,
} from '@constants/existingBridgeRoutes'
import { useState } from 'react'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'

import * as ALL_COINS from '@constants/tokens/bridgeable'
import { Token } from '@/utils/types'
import { getSymbol } from '@/utils/routeMaker/generateRoutePossibilities'

const CHAIN_IDS = Object.keys(CHAINS_BY_ID)

function findTokensForRoute(
  fromChainId: number,
  toChainId: number,
  routes: BridgeRoutes
): Token[] {
  const tickers: string[] = []

  if (fromChainId === toChainId) {
    return []
  }

  for (const tickerChain in routes) {
    const ticker = tickerChain.split('-')[0]
    const chainId = Number(tickerChain.split('-')[1])

    const reachableChains = routes[tickerChain].map((chain) =>
      Number(chain.split('-')[1])
    )

    if (chainId === fromChainId && reachableChains.includes(toChainId)) {
      tickers.push(ticker)
    }
  }

  return tickers
    .filter((t) => t !== 'WETH')
    .map((t) => Object.values(ALL_COINS).find((coin) => coin.routeSymbol === t))
    .filter(Boolean)
}

interface RouteMatrixProps {
  routes: BridgeRoutes
}

const ImageAndChain = ({ chainId }) => {
  const { chainImg, shortName } = CHAINS_BY_ID[chainId]

  return (
    <div className="flex items-center space-x-2" key={shortName}>
      <img src={chainImg.src} className="w-4 h-4" />
      <div>{shortName}</div>
    </div>
  )
}

const RouteMatrix: React.FC<RouteMatrixProps> = ({ routes }) => {
  const [filter, setFilter] = useState<string>('')

  const handleFilterChange = _.debounce((event) => {
    setFilter(event.target.value.trim())
  }, 300)

  return (
    <div>
      <div className="mb-16">
        <input
          type="text"
          placeholder="Filter by token symbol"
          onChange={handleFilterChange}
        />
      </div>
      <table className="text-white">
        <thead className="">
          <tr className="">
            <th></th>
            {CHAIN_IDS.map((toId) => (
              <th key={`to-${toId}`} className="rotate-90">
                <ImageAndChain chainId={toId} />
              </th>
            ))}
          </tr>
        </thead>
        <tbody className="">
          <tr className="h-10"></tr>
          {CHAIN_IDS.map((fromId) => (
            <tr key={`from-${fromId}`}>
              <th className="pr-5">
                <ImageAndChain chainId={fromId} />
              </th>
              {CHAIN_IDS.map((toId) => {
                const tokens = _.sortBy(
                  findTokensForRoute(Number(fromId), Number(toId), routes),
                  'priorityRank'
                )

                const filteredTokens = tokens.filter((token) =>
                  token.symbol.toLowerCase().includes(filter.toLowerCase())
                )

                return (
                  <td key={`to-${toId}`} className="h-[50px] w-[75px]">
                    <div className="flex flex-wrap">
                      <AvailableTokens
                        tokens={filteredTokens}
                        fromId={Number(fromId)}
                        toId={Number(toId)}
                      />
                    </div>
                  </td>
                )
              })}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

const RoutesPage = () => {
  return (
    <LandingPageWrapper>
      <div className="p-4">
        <RouteMatrix routes={EXISTING_BRIDGE_ROUTES} />
      </div>
    </LandingPageWrapper>
  )
}

const AvailableTokens = ({
  tokens,
  fromId,
  toId,
}: {
  tokens: Token[]
  fromId: number
  toId: number
}) => {
  const [isHovered, setIsHovered] = useState(false)
  const hasOneToken = tokens.length > 0
  const hasMultipleTokens = tokens.length > 1
  const numOverTwoTokens = tokens.length - 2 > 0 ? tokens.length - 2 : 0

  return (
    <div
      data-test-id="portfolio-token-visualizer"
      className="flex flex-row items-center space-x-1 hover-trigger w-[75px]"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      {hasOneToken && (
        <img
          key={tokens[0]?.name}
          src={tokens[0]?.icon.src}
          className="w-4 h-4"
          alt={tokens[0]?.name}
          title={tokens[0]?.name}
        />
      )}
      {hasMultipleTokens && (
        <img
          key={tokens[1]?.name}
          src={tokens[1]?.icon.src}
          className="w-4 h-4"
          alt={tokens[1]?.name}
          title={tokens[1]?.name}
        />
      )}
      {numOverTwoTokens > 0 && (
        <div className="ml-1 text-xs text-white">+ {numOverTwoTokens}</div>
      )}
      <div className="relative inline-block">
        {isHovered && (
          <div
            className={`
              absolute z-50 hover-content p-2 text-white
              border border-solid border-[#252537]
              bg-[#101018] rounded-md w-[200px]
            `}
          >
            <div className="flex items-center space-x-1 text-sm">
              <div className="flex items-center space-x-1">
                <img
                  src={CHAINS_BY_ID[fromId]?.chainImg.src}
                  className="w-3 h-3"
                />
                <div>{CHAINS_BY_ID[fromId].shortName}</div>
              </div>
              <div>to</div>
              <div className="flex items-center space-x-1">
                <img
                  src={CHAINS_BY_ID[toId]?.chainImg.src}
                  className="w-3 h-3"
                />
                <div>{CHAINS_BY_ID[toId].shortName}</div>
              </div>
            </div>
            <div className="mb-2" />
            {tokens.map((token) => {
              return (
                <div className="mb-1">
                  <div className="flex items-center space-x-1">
                    <img className="w-4 h-4" src={token?.icon.src} />
                    <div className="text-sm">{token.symbol}</div>
                  </div>
                  <div className="flex items-center ml-2 text-sm opacity-80">
                    {'-> '}
                    {destinationTokens({
                      fromChainId: fromId,
                      fromToken: token,
                      toChainId: toId,
                    })}
                  </div>
                </div>
              )
            })}
          </div>
        )}
      </div>
    </div>
  )
}

const destinationTokens = ({
  fromChainId,
  fromToken,
  toChainId,
}: {
  fromChainId: number
  fromToken: Token
  toChainId: number
}) => {
  return EXISTING_BRIDGE_ROUTES[`${fromToken.routeSymbol}-${fromChainId}`]
    .filter((v) => v.endsWith(`-${toChainId}`))
    .map(getSymbol)
    .join(', ')
}

export default RoutesPage
