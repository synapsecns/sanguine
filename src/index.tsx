import './index.css'
;('use client')
import { WidgetProps } from 'types'
import { Widget } from './components/Widget'
import { Web3Provider } from 'providers/Web3Provider'
import { Provider } from 'react-redux'
import { store } from '@/state/store'

import * as BRIDGEABLE from '@/constants/bridgeable'

// TODO: After separating required fields need two types: BridgeProps (consumer) and WidgetProps (internal)
// TODO: Can we clarify browserProvider vs. web3Provider vs. web3Context

export const Bridge = ({
  chainIds,
  web3Provider,
  networkProviders,
  theme,
  customTheme,
  tokens,
  toChainId,
}: WidgetProps) => {
  return (
    <Web3Provider config={web3Provider}>
      <Provider store={store}>
        <Widget
          chainIds={chainIds}
          networkProviders={networkProviders}
          theme={theme}
          customTheme={customTheme}
          tokens={tokens}
          toChainId={toChainId}
        />
      </Provider>
    </Web3Provider>
  )
}

export const ETH = BRIDGEABLE.ETH
export const USDC = BRIDGEABLE.USDC
export const USDT = BRIDGEABLE.USDT
export const DAI = BRIDGEABLE.DAI
