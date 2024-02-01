import { Bridge, USDC, USDT, DAI, ETH, CustomRpcs } from '@synapsecns/widget'
import { useEthereumWallet } from './hooks/useEthereumWallet'

const targetTokens = [USDC, USDT, DAI, ETH]

const customRpcs: CustomRpcs = {
  1: 'https://eth.llamarpc.com',
  42161: 'https://arbitrum.llamarpc.com',
}

function App() {
  const { web3Provider } = useEthereumWallet()

  return (
    <>
      <main>
        <div style={{ display: 'grid', placeItems: 'center', height: '100vh' }}>
          <Bridge
            web3Provider={web3Provider}
            customRpcs={customRpcs}
            targetTokens={targetTokens}
            customTheme={{ bgColor: 'light' }}
            container={true}
            targetChainIds={[42161]}
          />
        </div>
      </main>
    </>
  )
}

export default App
