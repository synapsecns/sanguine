import { Bridge } from '@synapsecns/widget'
import { StaticJsonRpcProvider } from '@ethersproject/providers'
import { BaseSyntheticEvent, useState } from 'react'

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

  const [customTheme, setCustomTheme] = useState({})

    function colorInputHandler(e: BaseSyntheticEvent) {
    function hexToRgb(hex: string) {
      // https://stackoverflow.com/questions/5623838/rgb-to-hex-and-hex-to-rgb
      var result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);
      return result ? { r: parseInt(result[1], 16) / 255, g: parseInt(result[2], 16) / 255, b: parseInt(result[3], 16) / 255, } : null;
    }
    function rgb2hsl({ r, g, b, a = 1 }: any) {
      // in: r,g,b in [0,1], out: h in [0,360) and s,l in [0,100] // https://stackoverflow.com/a/54071699
      let v = Math.max(r, g, b), c = v - Math.min(r, g, b), f = 1 - Math.abs(v + v - c - 1)
      let h = c && ((v === r) ? (g - b) / c : ((v === g) ? 2 + (b - r) / c : 4 + (r - g) / c))
      return { h: 60 * (h < 0 ? h + 6 : h), s: f ? 100 * c / f : 0, l: 100 * (v + v - c) / 2, a }
    }
    const hsla = rgb2hsl(hexToRgb(e.target.value))
    console.log(hsla)

    setCustomTheme(hsla.l < 50
      ?
        {
          '--h': hsla.h,
          '--s': `${hsla.s}%`,
          '--primary':    'hsl(var(--h), var(--s), 96%)',
          '--secondary':  'hsl(var(--h), var(--s), 86%)',
          '--small':      'hsl(var(--h), var(--s), 66%)',
          '--accent':     'hsl(var(--h), var(--s), 29%)',
          '--separator':  'hsl(var(--h), var(--s), 13%)',
          '--surface':    'hsl(var(--h), var(--s), 13%)',
          '--background': 'hsl(var(--h), var(--s), 7%)',
        }
      : 
        {
          '--h': hsla.h,
          '--s': `${hsla.s}%`,
          '--primary':    'hsl(var(--h), var(--s), 7%)',
          '--secondary':  'hsl(var(--h), var(--s), 41%)',
          '--small':      'hsl(var(--h), var(--s), 66%)',
          '--accent':     'hsl(var(--h), var(--s), 96%)',
          '--separator':  'hsl(var(--h), var(--s), 86%)',
          '--surface':    'hsl(var(--h), var(--s), 100%)',
          '--background': 'hsl(var(--h), var(--s), 96%)',
        }
    )
  }

  const customThemeDFK = {
    '--primary': 'rgb(62,31,5)',
    '--secondary': 'rgb(62,31,5)',
    '--small': 'rgb(224,228,203)',
    '--separator': 'rgb(216,172,130)',
    '--background': 'rgb(255,227,189)',
    '--surface': 'rgb(216,172,130)',
    '--accent': 'rgb(255,227,189)',
  }

  const customThemeWeird = {
    '--primary': 'red',
    '--secondary': 'green',
    '--small': 'yellow',
    '--separator': 'blue',
    '--background': 'orange',
    '--surface': 'purple',
    '--accent': 'gray',
  }

  return (
    <>
      <header>
        <h1>Synapse Widget</h1>
      </header>
      
      <main>
        <Bridge chainIds={chainIds} providers={providers} tokens={tokens} />
        <Bridge chainIds={chainIds} providers={providers} tokens={tokens} theme="night" />
        <Bridge chainIds={chainIds} providers={providers} tokens={tokens} customTheme={customThemeDFK}/>
        <Bridge chainIds={chainIds} providers={providers} tokens={tokens} customTheme={Object.keys(customTheme).length && customTheme }/>
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
      {/* </header> */}
      </main>

      <footer>
        <input id="color-picker" type="color" onInput={colorInputHandler} />
      </footer>
    </>
  )
}

export default App
