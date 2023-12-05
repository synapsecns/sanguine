import { Bridge } from '@synapsecns/widget'
import { StaticJsonRpcProvider } from '@ethersproject/providers'
import { useEthereumWallet } from './hooks/useEthereumWallet'
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

  function createCustomTheme(e: BaseSyntheticEvent) {
    function hexToRgb(hex: string) {
      // https://stackoverflow.com/questions/5623838/rgb-to-hex-and-hex-to-rgb
      var result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex)
      return result
        ? {
            r: parseInt(result[1], 16) / 255,
            g: parseInt(result[2], 16) / 255,
            b: parseInt(result[3], 16) / 255,
          }
        : null
    }
    function rgb2hsl({ r, g, b, a = 1 }: any) {
      // in: r,g,b in [0,1], out: h in [0,360) and s,l in [0,100] // https://stackoverflow.com/a/54071699
      let v = Math.max(r, g, b),
        c = v - Math.min(r, g, b),
        f = 1 - Math.abs(v + v - c - 1)
      let h =
        c &&
        (v === r ? (g - b) / c : v === g ? 2 + (b - r) / c : 4 + (r - g) / c)
      return {
        h: 60 * (h < 0 ? h + 6 : h),
        s: f ? (100 * c) / f : 0,
        l: (100 * (v + v - c)) / 2,
        a,
      }
    }

    const colorPicker = document.getElementById(
      'color-picker'
    ) as HTMLInputElement | null
    const hsla = rgb2hsl(hexToRgb(colorPicker?.value ?? '#000000'))

    const accentColorPicker = document.getElementById(
      'accent-color-picker'
    ) as HTMLInputElement | null

    setCustomTheme(
      hsla.l < 50
        ? {
            '--strong': `hsl(${hsla.h}deg ${hsla.s}% ${
              hsla.l * 1.0 + 100
            }% / 100%)`,
            '--primary': `hsl(${hsla.h}deg ${hsla.s}% ${
              hsla.l * 0.96 + 96
            }% / 100%)`,
            '--secondary': `hsl(${hsla.h}deg ${hsla.s}% ${
              hsla.l * 0.86 + 86
            }% / 100%)`,
            '--small': `hsl(${hsla.h}deg ${hsla.s}% ${
              hsla.l * 0.66 + 66
            }% / 100%)`,
            '--accent': `hsl(${hsla.h}deg ${hsla.s}% ${
              hsla.l * 0.25 + 25
            }% / 100%)`,
            '--separator': `hsl(${hsla.h}deg ${hsla.s}% ${
              hsla.l * 0.12 + 12
            }% / 100%)`,
            '--surface': `hsl(${hsla.h}deg ${hsla.s}% ${
              hsla.l * 0.12 + 12
            }% / 100%)`,
            '--background': `hsl(${hsla.h}deg ${hsla.s}% ${
              hsla.l * 0.07 + 7
            }% / 100%)`,
            '--brand': accentColorPicker?.value ?? '#000000',
          }
        : {
            '--strong': `hsl(${hsla.h}deg ${hsla.s}% ${
              Math.min(100, hsla.l * 0.0) * 0.0
            }% / 100%)`,
            '--primary': `hsl(${hsla.h}deg ${hsla.s}% ${
              Math.min(100, hsla.l * 1.07) * 0.07
            }% / 100%)`,
            '--secondary': `hsl(${hsla.h}deg ${hsla.s}% ${
              Math.min(100, hsla.l * 1.41) * 0.41
            }% / 100%)`,
            '--small': `hsl(${hsla.h}deg ${hsla.s}% ${
              Math.min(100, hsla.l * 1.66) * 0.66
            }% / 100%)`,
            '--accent': `hsl(${hsla.h}deg ${hsla.s}% ${
              Math.min(100, hsla.l * 1.96) * 0.96
            }% / 100%)`,
            '--separator': `hsl(${hsla.h}deg ${hsla.s}% ${
              Math.min(100, hsla.l * 1.86) * 0.86
            }% / 100%)`,
            '--surface': `hsl(${hsla.h}deg ${hsla.s}% ${
              Math.min(100, hsla.l * 2.0) * 1.0
            }% / 100%)`,
            '--background': `hsl(${hsla.h}deg ${hsla.s}% ${
              Math.min(100, hsla.l * 1.96) * 0.96
            }% / 100%)`,
            '--brand': accentColorPicker?.value ?? '#000000',
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

  const { web3Provider, connectedAddress, connectedNetwork } =
    useEthereumWallet()

  console.log('Consumer address:', connectedAddress)
  console.log('Consumer web3 provider: ', web3Provider)
  console.log('Consumer connectedNetwork: ', connectedNetwork)

  return (
    <>
      <header>
        <img width="160" src="/synapse-logo.svg" alt="Synapse logo" />
        <a href="https://synapseprotocol.com" target="_blank" rel="noreferrer">
          EVM Bridge
        </a>
      </header>

      <main style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
        <header>
          <h1>Synapse Widget</h1>
          <code>npm synapse-widget</code>
          <p>
            Easily onboard new users by adding a custom instance of the Synapse
            Bridge to your React project
          </p>
        </header>

        <Bridge
          chainIds={chainIds}
          web3Provider={web3Provider}
          networkProviders={providers}
          tokens={tokens}
          customTheme={Object.keys(customTheme)?.length && customTheme}
        <div>
          {/* <label htmlFor='color-picker'>Background</label> */}
          <input id="color-picker" type="color" onInput={createCustomTheme} />
          &nbsp;
          {/* <label htmlFor='accent-color-picker'>Accent</label> */}
          <input
            id="accent-color-picker"
            type="color"
            onInput={createCustomTheme}
          />
        </div>
        {/* <Bridge
          chainIds={chainIds}
          web3Provider={web3Provider}
          networkProviders={providers}
          tokens={tokens}
          theme="night"
        /> */}

        <h2>Customize</h2>
        <h3>Color Mode — WIP, not reflected in code</h3>
        <dl>
          <dt>theme</dt>
          <dd>
            <ul style={{ display: 'flex', gap: '1rem' }}>
              <li>auto</li>|<li>dark</li>|<li>light</li>
            </ul>
          </dd>
        </dl>
        <h3>Color Values — WIP, not reflected in code</h3>
        {/* <h4>Text</h4> */}
        <dl>
          <dt>--synapse-text-strong</dt>
          <dd>hsl(0deg 0% 0% / 100%)</dd>
          <dt>--synapse-text-primary</dt>
          <dd>hsl(0deg 0% 0% / 100%)</dd>
          <dt>--synapse-text-secondary</dt>
          <dd>hsl(0deg 0% 0% / 100%)</dd>
        </dl>
        {/* <h4>Objects</h4> */}
        <dl>
          <dt>--synapse-bg-button</dt>
          <dd>hsl(0deg 0% 0% / 100%)</dd>
          <dt>--synapse-bg-select</dt>
          <dd>hsl(0deg 0% 0% / 100%)</dd>
          <dt>--synapse-bg-card</dt>
          <dd>hsl(0deg 0% 0% / 100%)</dd>
          <dt>--synapse-bg-root</dt>
          <dd>hsl(0deg 0% 0% / 100%)</dd>
        </dl>
        {/* <h4>Color</h4> */}
        <dl>
          <dt>--synapse-border</dt>
          <dd>hsl(0deg 0% 0% / 100%)</dd>
          <dt>--synapse-accent</dt>
          <dd>hsl(0deg 0% 0% / 100%)</dd>
        </dl>
        <h3>Typography — WIP, not reflected in code</h3>
        <dl>
          <dt>--synapse-font-size</dt>
          <dd>100%</dd>
          <dt>--synapse-font-family-display</dt>
          <dd>system-ui</dd>
          <dt>--synapse-font-family-text</dt>
          <dd>system-ui</dd>
          <dt>--synapse-font-weight-display</dt>
          <dd>600 (semibold)</dd>
          <dt>--synapse-font-weight-text</dt>
          <dd>500 (medium)</dd>
        </dl>
        {/* <Bridge chainIds={chainIds} providers={providers} tokens={tokens} /> */}
        {/* <Bridge chainIds={chainIds} providers={providers} tokens={tokens} theme="night" />
        <Bridge chainIds={chainIds} providers={providers} tokens={tokens} customTheme={customThemeDFK}/> */}
      </main>

      <footer></footer>
    </>
  )
}

export default App
