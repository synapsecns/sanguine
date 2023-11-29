import './App.css'
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

  console.log(`providres`, providers)

  return (
    <div className="App">
      <header className="App-header">
        <Bridge chainIds={chainIds} providers={providers} />
      </header>
    </div>
  )
}

export default App
