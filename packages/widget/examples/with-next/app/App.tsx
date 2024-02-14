import { Bridge } from '@synapsecns/widget'
import { useEthereumWallet } from './hooks/useEthereumWallet'

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
          customTheme={{ bgColor: 'light' }}
        />
      </div>
    </main>
  )
}

export default App
