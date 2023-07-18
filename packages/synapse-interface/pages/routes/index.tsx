import _ from 'lodash'
import { CHAINS_BY_ID } from '@/constants/chains'
import {
  BridgeRoutes,
  EXISTING_BRIDGE_ROUTES,
} from '@/constants/existing-bridge-routes'
import { useEffect, useState } from 'react'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import * as ALL_COINS from '@/constants/tokens/master'
import StandardPageContainer from '@/components/layouts/StandardPageContainer'

const CHAIN_IDS = Object.keys(CHAINS_BY_ID)

const findCoinByTicker = (ticker) => {
  for (let coin of Object.values(ALL_COINS)) {
    if (
      coin.symbol.toLowerCase().trim() === ticker.toLowerCase() ||
      ticker === coin?.swapableType ||
      ALL_COINS[ticker] === coin
    ) {
      return coin
    }
  }

  return null
}

function findTickersForRoute(
  fromChainId: number,
  toChainId: number,
  routes: BridgeRoutes
): string[] {
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
}

interface RouteMatrixProps {
  routes: BridgeRoutes
}

const ImageAndChain = ({ chainId }) => {
  const { chainImg, name } = CHAINS_BY_ID[chainId]
  return (
    <div className="flex items-center space-x-2" key={name}>
      <img src={chainImg.src} className="w-4 h-4" />
      <div>{name}</div>
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
      <div className="mb-10">
        <input
          type="text"
          placeholder="Filter by symbol"
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
                const tickers = findTickersForRoute(
                  Number(fromId),
                  Number(toId),
                  routes
                ).filter(
                  (ticker) =>
                    ticker !== 'KLAYTN_USDT' &&
                    ticker !== 'KLAYTN_USDC' &&
                    ticker !== 'KLAYTN_DAI' &&
                    ticker !== 'DOGECHAIN_BUSD' &&
                    ticker !== 'WBNB' &&
                    ticker !== 'WMATIC'
                )

                if (Number(toId) === 10 && Number(fromId) === 250) {
                  console.log(`fromId`, fromId)
                  console.log(`toId`, toId)
                  console.log(tickers)
                }

                const filteredTickers = tickers.filter((ticker) =>
                  ticker.toLowerCase().includes(filter.toLowerCase())
                )

                return (
                  <td key={`to-${toId}`} className="h-[50px] w-[75px]">
                    <div className="flex flex-wrap">
                      {filteredTickers.map((ticker) => {
                        const coin = findCoinByTicker(ticker)

                        return (
                          <img
                            key={coin?.name}
                            src={coin?.icon.src}
                            className="w-4 h-4"
                            alt={coin?.name}
                            title={coin?.name}
                          />
                        )
                      })}
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
      <StandardPageContainer address={''} connectedChainId={0}>
        <RouteMatrix routes={EXISTING_BRIDGE_ROUTES} />
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

export default RoutesPage
