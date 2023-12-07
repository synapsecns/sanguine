import './index.css'
;('use client')
import { WidgetProps } from 'types'
import { Widget } from './components/Widget'
import { Web3Provider } from 'providers/Web3Provider'
import { Provider } from 'react-redux'
import { store } from '@/state/store'

import { createPublicClient, createWalletClient, custom, http } from 'viem'
import { mainnet, arbitrum } from 'viem/chains'

import * as BRIDGEABLE from '@/constants/bridgeable'

export const viemPublicClient = createPublicClient({
  chain: arbitrum,
  transport: http(),
})

export const viemWalletClient = createWalletClient({
  chain: arbitrum,
  transport: custom(window.ethereum),
})

export const Bridge = ({
  chainIds,
  web3Provider,
  networkProviders,
  theme,
  customTheme,
  tokens,
}: WidgetProps) => {
  return (
    <Web3Provider>
      <Provider store={store}>
        <Widget
          chainIds={chainIds}
          web3Provider={web3Provider}
          networkProviders={networkProviders}
          theme={theme}
          customTheme={customTheme}
          tokens={tokens}
        />
      </Provider>
    </Web3Provider>
  )
}

export const ETH = BRIDGEABLE.ETH
export const USDC = BRIDGEABLE.USDC
export const USDT = BRIDGEABLE.USDT
export const DAI = BRIDGEABLE.DAI
