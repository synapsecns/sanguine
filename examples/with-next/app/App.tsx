import { Bridge, USDC, USDT, DAI, ETH, USDCe } from '@abtestingalpha/widget'
import { useEthereumWallet } from './hooks/useEthereumWallet'

const tokens = [USDC, USDT, DAI, ETH, USDCe]

function App() {
  const { web3Provider } = useEthereumWallet()

  if (!web3Provider) {
    return null
  }

  return (
    <main className="flex items-center justify-center h-screen">
      <div className="w-[33%]">
        <Bridge
          web3Provider={web3Provider}
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
