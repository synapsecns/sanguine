import { Bridge, USDC, USDT, DAI } from '@synapsecns/widget'
import { StaticJsonRpcProvider } from '@ethersproject/providers'
import { useEthereumWallet } from './hooks/useEthereumWallet'
import { BaseSyntheticEvent, useState } from 'react'

const tokens = [USDC, USDT, DAI]

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

  // const providers = [ethersProvider, aribtrumProvider]
  // const chainIds = [1, 42161]

  const providers = [aribtrumProvider, ethersProvider, polygonProvider]
  const chainIds = [42161, 1, 137]

  const [customTheme, setCustomTheme] = useState({})

  function createCustomTheme() {
    const colorPicker = document.getElementById(
      'color-picker'
    ) as HTMLInputElement | null

    const accentColorPicker = document.getElementById(
      'accent-color-picker'
    ) as HTMLInputElement | null

    setCustomTheme({
      bgColor: colorPicker?.value,
      accentColor: accentColorPicker?.value,
      /* Overrides */
      // '--synapse-text-primary': hslString(h, s, l * 0.96 + 96, a),
      // '--synapse-text-secondary': hslString(h, s, l * 0.86 + 86, a),
      // '--synapse-bg-select': hslString(h, s, l * 0.25 + 25, a),
      // '--synapse-bg-surface': hslString(h, s, l * 0.12 + 12, a),
      // '--synapse-bg-root': hslString(h, s, l * 0.07 + 7, a),
      // '--synapse-border': hslString(h, s, l * 0.12 + 12, a),
      // '--synapse-border-hover': hslString(h, s, l * 0.66 + 66, a),
      // '--synapse-accent': accentColorPicker?.value ?? '#ffffff',
    })
  }

  const customThemeDFK = {
    '--synapse-text-primary': 'rgb(62,31,5)',
    '--synapse-text-secondary': 'rgb(62,31,5)',
    '--synapse-bg-select': 'rgb(255,227,189)',
    '--synapse-bg-surface': 'rgb(216,172,130)',
    '--synapse-bg-background': 'rgb(255,227,189)',
    '--synapse-border': 'rgb(216,172,130)',
    '--synapse-border-hover': 'rgb(224,228,203)',
    '--synapse-accent': 'rgb(62,31,5)',
  }

  const customThemeWeird = {
    '--synapse-text-primary': 'red',
    '--synapse-text-secondary': 'green',
    '--synapse-bg-select': 'gray',
    '--synapse-bg-surface': 'purple',
    '--synapse-bg-background': 'orange',
    '--synapse-border': 'blue',
    '--synapse-border-hover': 'yellow',
    '--synapse-accent': 'red',
  }

  const { web3Provider, connectedAddress, connectedNetwork } =
    useEthereumWallet()

  return (
    <>
      <header style={{ placeContent: 'space-between' }}>
        <img width="160" src="/synapse-logo.svg" alt="Synapse logo" />
        <a href="https://synapseprotocol.com" target="_blank" rel="noreferrer">
          EVM Bridge
        </a>
      </header>

      <main style={{ flexDirection: 'column' }}>
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
          customTheme={customTheme}
        />
        <div>
          <input id="color-picker" type="color" onInput={createCustomTheme} />
          {/* &nbsp;
          <label htmlFor='color-picker'>Background</label>
          &nbsp;
          &nbsp;
          <input
            id="accent-color-picker"
            type="color"
            onInput={createCustomTheme}
          />
          &nbsp;
          <label htmlFor='accent-color-picker'>Accent</label> */}
        </div>
        {/* <Bridge
          chainIds={chainIds}
          web3Provider={web3Provider}
          networkProviders={providers}
          tokens={tokens}
          theme="night"
        /> */}
        <h2>Customize</h2>
        <h3>Token list</h3>
        TODO: Describe how to customize the source or destination token lists
        <h3>Appearance</h3>
        <h4>Dark mode</h4>
        To override the default light theme, set the bgColor property to dark.
        <pre>customTheme = &#123; bgColor: 'dark' &#125;</pre>
        <h4>Auto-palette</h4>
        Generate a palette based on your brand colors by setting bgColor to any
        hex, rgb, or hsl color string. Hex values must contain 6 characters.
        <pre>
          customTheme = &#123;
          <br /> bgColor: '#000A14'
          <br /> bgColor: 'rgb(0 10 20)'
          <br /> bgColor: 'hsl(210deg 100% 4%)'
          <br />
          &#125;
        </pre>
        {/* <h4>Accent Color</h4>
        Add an accent color to text links and button hover states by setting accentColor to any hex, rgb, or hsl color string.
        <pre>
        customTheme = &#123;
          <br />  accentColor: '#d557ff'
          <br />  accentColor: 'rgb(213 87 255)'
          <br />  accentColor: 'hsl(285deg 100% 67%)'
          <br />&#125;
        </pre> */}
        <h4>Overrides</h4>
        The following CSS variables can be added to your CustomTheme to override
        the generated values. Any valid CSS color string can be used, including
        var() aliases.
        <pre>
          customTheme = &#123;
          <br /> --synapse-text-primary: 'white'
          <br /> --synapse-text-secondary: '#cccccc'
          <br />
          <br /> --synapse-bg-select: 'hsl(210deg 100% 50%)'
          <br /> --synapse-bg-surface: 'hsl(210deg 100% 12.5%)'
          <br /> --synapse-bg-root: 'inherit'
          <br />
          <br /> --synapse-border: 'hsl(210deg 100% 25%)'
          <br /> --synapse-border-hover: 'hsl(285deg 100% 33%)'
          <br />
          <br /> --synapse-accent: 'var(--my-brand-color)'
          <br />
          &#125;
        </pre>
        {/* <h4>Typography — WIP, not reflected in code</h4>
        <dl>
          <dt>--synapse-font-size</dt><dd>100%</dd>
          <dt>--synapse-font-family-display</dt><dd>system-ui</dd>
          <dt>--synapse-font-family-text</dt><dd>system-ui</dd>
          <dt>--synapse-font-weight-display</dt><dd>600 (semibold)</dd>
          <dt>--synapse-font-weight-text</dt><dd>500 (medium)</dd>
        </dl> */}
      </main>

      <footer></footer>
    </>
  )
}

export default App
