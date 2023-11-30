import { Bridge } from 'synapse-widget'
import { StaticJsonRpcProvider } from '@ethersproject/providers'

function App() {
  const ethersProvider = new StaticJsonRpcProvider(
    'https://eth.llamarpc.com',
    1
  )
  const aribtrumProvider = new StaticJsonRpcProvider(
    'https://arbitrum.llamarpc.com',
    42161
  )

  const providers = [ethersProvider, aribtrumProvider]
  const chainIds = [1, 42161]

  const customTheme = {
    primary: 'rgb(62,31,5)',
    secondary: 'rgb(146,150,167)',
    small: 'rgb(224,228,203)',
    separator: 'rgb(216,172,130)',
    background: 'rgb(255,227,189)',
    surface: 'rgb(216,172,130)',
    accent: 'rgb(35,152,186)',
  }

  return (
    <div className="App">
      <header className="App-header">
        <Bridge chainIds={chainIds} providers={providers} />
        {/* <Bridge chainIds={chainIds} providers={providers} theme="night" /> */}
        {/* <Bridge
          chainIds={chainIds}
          providers={providers}
          customTheme={customTheme}
        /> */}
      </header>
    </div>
  )
}

export default App
