import { Bridge, USDC, USDT, DAI, ETH, USDCe } from '@abtestingalpha/widget'
import { StaticJsonRpcProvider } from '@ethersproject/providers'
import { useEthereumWallet } from './hooks/useEthereumWallet'

const tokens = [USDC, USDT, DAI, ETH, USDCe]

function App() {
  const ethersProvider = new StaticJsonRpcProvider(
    'https://eth.llamarpc.com',
    1
  )
  const aribtrumProvider = new StaticJsonRpcProvider(
    'https://arbitrum.llamarpc.com',
    42161
  )
  const polygonProvider = new StaticJsonRpcProvider(
    'https://polygon.llamarpc.com',
    137
  )

  const optimismProvider = new StaticJsonRpcProvider(
    'https://mainnet.optimism.io',
    10
  )

  const providers = [
    aribtrumProvider,
    ethersProvider,
    polygonProvider,
    optimismProvider,
  ]
  const chainIds = [42161, 1, 137, 10]

  const { web3Provider } = useEthereumWallet()

  if (!web3Provider) {
    return null
  }

  return (
    <main className="flex items-center justify-center h-screen">
      <div className="w-[33%]">
        <Bridge
          chainIds={chainIds}
          web3Provider={web3Provider}
          networkProviders={providers}
          tokens={tokens}
          theme="light"
          container={true}
          toChainId={137}
        />
      </div>
    </main>
  )
}

export default App
