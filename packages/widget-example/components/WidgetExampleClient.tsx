import { Bridge } from '@synapsecns/widget'

import { useEthereumWallet } from '../hooks/useEthereumWallet'

const WidgetExampleClient = () => {
  const {
    connectedAddress,
    connectedNetwork,
    error,
    hasInjectedWallet,
    isInitializing,
    web3Provider,
  } = useEthereumWallet()

  if (!hasInjectedWallet) {
    return (
      <div className="widget-status-panel">
        <h2>Injected wallet required</h2>
        <p>
          Install or enable an injected Ethereum wallet such as MetaMask, then
          reload the page to render the widget.
        </p>
      </div>
    )
  }

  if (isInitializing) {
    return (
      <div className="widget-status-panel">
        <h2>Preparing browser wallet</h2>
        <p>Initializing the injected provider for the bridge widget.</p>
      </div>
    )
  }

  if (!web3Provider) {
    return (
      <div className="widget-status-panel">
        <h2>Wallet not ready</h2>
        <p>{error ?? 'The widget is waiting for a usable browser wallet.'}</p>
      </div>
    )
  }

  return (
    <div className="widget-stack">
      <div className="wallet-meta">
        <span>{connectedAddress || 'Read-only mode'}</span>
        <span>
          {connectedNetwork?.name ||
            (connectedNetwork?.chainId
              ? `Chain ${connectedNetwork.chainId}`
              : 'Wallet available')}
        </span>
      </div>
      {error ? <div className="wallet-note">{error}</div> : null}
      <div className="widget-frame">
        <Bridge
          web3Provider={web3Provider}
          container={true}
          customTheme={{ bgColor: 'light' }}
        />
      </div>
    </div>
  )
}

export default WidgetExampleClient
