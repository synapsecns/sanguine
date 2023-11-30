import { Bridge } from 'synapse-widget'
import { StaticJsonRpcProvider } from '@ethersproject/providers'

const tokens = [
  {
    tokenAddress: '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
    symbol: 'USDC',
    chainId: 1,
    decimals: 6,
  },
  {
    tokenAddress: '0xaf88d065e77c8cc2239327c5edb3a432268e5831',
    symbol: 'USDC',
    chainId: 42161,
    decimals: 6,
  },
  {
    tokenAddress: '0x6b175474e89094c44da98b954eedeac495271d0f',
    symbol: 'DAI',
    chainId: 1,
    decimals: 18,
  },
  {
    tokenAddress: '0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1',
    symbol: 'DAI',
    chainId: 42161,
    decimals: 18,
  },
  {
    tokenAddress: '0xdac17f958d2ee523a2206206994597c13d831ec7',
    symbol: 'USDT',
    chainId: 1,
    decimals: 6,
  },
  {
    tokenAddress: '0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9',
    symbol: 'USDT',
    chainId: 42161,
    decimals: 6,
  },
]

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
        <Bridge chainIds={chainIds} providers={providers} tokens={tokens} />
        {/* <Bridge
          chainIds={chainIds}
          providers={providers}
          theme="night"
          tokens={tokens}
        /> */}
        {/* <Bridge
          chainIds={chainIds}
          providers={providers}
          customTheme={customTheme}
          tokens={tokens}
        /> */}
      </header>
    </div>
  )
}

export default App
