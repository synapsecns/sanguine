import { Bridge, USDC, USDT, DAI, ETH } from '@synapsecns/widget'
import { StaticJsonRpcProvider } from '@ethersproject/providers'
import { useEthereumWallet } from './hooks/useEthereumWallet'
import { BaseSyntheticEvent, useState } from 'react'
import Header from './Header'
import Footer from './Footer'

const tokens = [USDC, USDT, DAI, ETH]

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

  // const providers = [ethersProvider, aribtrumProvider]
  // const chainIds = [1, 42161]

  const providers = [
    aribtrumProvider,
    ethersProvider,
    polygonProvider,
    optimismProvider,
  ]
  const chainIds = [42161, 1, 137, 10]

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
      <Header />

      <main>
        <header>
          <h1>Install the Synapse Bridge</h1>
          <pre>npm synapse-widget</pre>
          <p>
            Easily onboard new users by adding a custom instance of the Synapse
            Bridge to your React project
          </p>
          <div id="example-container">
            <div id="bridge-container">
              <Bridge
                chainIds={chainIds}
                web3Provider={web3Provider}
                networkProviders={providers}
                tokens={tokens}
                customTheme={customTheme}
                toChainId={137}
              />
            </div>
            <input id="color-picker" type="color" onInput={createCustomTheme} />
          </div>
        </header>
        {/* <Bridge
          chainIds={chainIds}
          web3Provider={web3Provider}
          networkProviders={providers}
          tokens={tokens}
          theme="night"
        /> */}
        {/* <hr /> */}
        <article>
          <h2>⬇️&nbsp; Install</h2>
          <p>Install the Synapse Widget in your Next.js or React project</p>
          <pre>npm synapse-widget</pre>

          <h2>⚙️&nbsp; Setup</h2>
          <h3>Supported tokens</h3>
          <p>
            While the Synapse Widget supports{' '}
            <a
              href="https://synapseprotocol.com"
              target="_blank"
              rel="noreferrer"
            >
              hundreds of tokens and chains
            </a>
            , for a streamlined user experience, you can render a separate
            instance of the bridge for each user need.
          </p>
          <p>
            For example: separate <code>BridgeIn</code> and{' '}
            <code>BridgeOut</code> functions allow you to define the tokens you
            support sending and receiving.
          </p>
          <p className="info">
            <strong>Note</strong>: Whitelisting one side of a transaction limits
            the other side to compatible tokens automatically.
          </p>
          <pre>
            {`// Bridge in
tokenList = {
  source: [], destination: [ token, token, token ]
}

// Bridge out
tokenList = {
  source: [ token, token, token ], destination: [],
}`}
          </pre>

          <h2>🎨&nbsp; Customize</h2>
          <h3>Dark mode</h3>
          <p>
            To override the default light theme, set &nbsp;<code>bgColor</code>
            &nbsp; to &nbsp;<code>'dark'</code>&nbsp;.
          </p>
          <pre>customTheme = &#123; bgColor: 'dark' &#125;</pre>
          <h3>Auto-palette</h3>
          <p>
            Generate a palette based on your brand colors by setting bgColor to
            any hex, rgb, or hsl color string. Hex values must contain 6
            characters.
          </p>
          <pre>
            {`customTheme = {
  bgColor: '#000A14'
  bgColor: 'rgb(0 10 20)'
  bgColor: 'hsl(210deg 100% 4%)'
}`}
          </pre>
          {/* <h3>Accent Color</h3>
          Add an accent color to text links and button hover states by setting accentColor to any hex, rgb, or hsl color string.
          <pre>
          customTheme = &#123;
            <br />  accentColor: '#d557ff'
            <br />  accentColor: 'rgb(213 87 255)'
            <br />  accentColor: 'hsl(285deg 100% 67%)'
            <br />&#125;
          </pre> */}
          <h3>Overrides</h3>
          <p>
            The following CSS variables can be added to your CustomTheme to
            override the generated values. Any valid CSS color string can be
            used, including var() aliases.
          </p>
          <pre>
            {`customTheme = {
  --synapse-text-primary: 'white'
  --synapse-text-secondary: '#cccccc'

  --synapse-bg-select: 'hsl(210deg 100% 50%)'
  --synapse-bg-surface: 'hsl(210deg 100% 12.5%)'
  --synapse-bg-root: 'inherit'

  --synapse-border: 'hsl(210deg 100% 25%)'
  --synapse-border-hover: 'hsl(285deg 100% 33%)'
}`}
          </pre>
          {/* <h3>Typography — WIP, not reflected in code</h3>
          <dl>
            <dt>--synapse-font-size</dt><dd>100%</dd>
            <dt>--synapse-font-family-display</dt><dd>system-ui</dd>
            <dt>--synapse-font-family-text</dt><dd>system-ui</dd>
            <dt>--synapse-font-weight-display</dt><dd>600 (semibold)</dd>
            <dt>--synapse-font-weight-text</dt><dd>500 (medium)</dd>
          </dl> */}
          <h2>🙋&nbsp; Support</h2>
          <p>
            For help and feedback, reach out to our Support team in the{' '}
            <a href="#" target="_blank" rel="noreferrer">
              Synapse Discord channel.
            </a>
          </p>
        </article>
      </main>

      <Footer />
    </>
  )
}

export default App
