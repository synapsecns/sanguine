import './index.css'
import { WidgetProps } from 'types'
import { Widget } from './components/Widget'
import { Web3Provider } from 'providers/Web3Provider'
import { Provider } from 'react-redux'
import { store } from '@/state/store'

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
