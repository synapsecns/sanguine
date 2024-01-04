import './index.css'
;('use client')
import { WidgetProps, CustomFallbackRpcs } from 'types'
import { Widget } from './components/Widget'
import { Web3Provider } from 'providers/Web3Provider'
import { Provider } from 'react-redux'
import { store } from '@/state/store'

import * as BRIDGEABLE from '@/constants/bridgeable'

import TransactionsUpdater from '@/state/slices/transactions/updater'
import { SynapseProvider } from 'providers/SynapseProvider'
import { CHAINS, CHAINS_ARRAY } from './constants/chains'

// TODO: After separating required fields need two types: BridgeProps (consumer) and WidgetProps (internal)
// TODO: Can we clarify browserProvider vs. web3Provider vs. web3Context

export const Bridge = ({
  web3Provider,
  fallbackRpcs,
  theme,
  customTheme,
  container,
  tokens,
  toChainId,
}: WidgetProps) => {
  return (
    <Web3Provider config={web3Provider}>
      <SynapseProvider chains={CHAINS_ARRAY} fallbackRpcs={fallbackRpcs}>
        <Provider store={store}>
          <TransactionsUpdater />
          <Widget
            theme={theme}
            customTheme={customTheme}
            container={container}
            tokens={tokens}
            toChainId={toChainId}
          />
        </Provider>
      </SynapseProvider>
    </Web3Provider>
  )
}

export const ETH = BRIDGEABLE.ETH
export const USDC = BRIDGEABLE.USDC
export const USDT = BRIDGEABLE.USDT
export const DAI = BRIDGEABLE.DAI
export const USDCe = BRIDGEABLE.USDCe
