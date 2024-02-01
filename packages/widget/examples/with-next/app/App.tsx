import { Bridge, USDC, USDT, DAI, ETH, USDCe } from '@synapsecns/widget'
import { useEthereumWallet } from './hooks/useEthereumWallet'

const targetTokens = [USDC, USDT, DAI, ETH, USDCe]

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
          targetTokens={targetTokens}
          targetChainIds={[137]}
          customTheme={{ bgColor: 'light' }}
          container={true}
        />
      </div>
    </main>
  )
}

export default App
