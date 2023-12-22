import { Bridge, USDC, USDT, DAI, ETH } from '@abtestingalpha/widget'
import { StaticJsonRpcProvider } from '@ethersproject/providers'
import { useEthereumWallet } from './hooks/useEthereumWallet'
import { BaseSyntheticEvent, SyntheticEvent, useState } from 'react'
import Header from './Header'
import Footer from './Footer'
import { Install, Developer, Support } from './icons'

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
  const [container, setContainer] = useState(true)

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
      // '--synapse-text': hslString(h, s, l * 0.96 + 96, a),
      // '--synapse-secondary': hslString(h, s, l * 0.86 + 86, a),
      // '--synapse-select-bg': hslString(h, s, l * 0.25 + 25, a),
      // '--synapse-surface': hslString(h, s, l * 0.12 + 12, a),
      // '--synapse-root': hslString(h, s, l * 0.07 + 7, a),
      // '--synapse-border': hslString(h, s, l * 0.12 + 12, a),
      // '--synapse-focus': hslString(h, s, l * 0.66 + 66, a),
      // '--synapse-accent': accentColorPicker?.value ?? '#ffffff',
    })
  }

  const customThemeDFK = {
    '--synapse-text': 'rgb(62,31,5)',
    '--synapse-secondary': 'rgb(62,31,5)',
    '--synapse-select-bg': 'rgb(255,227,189)',
    '--synapse-surface': 'rgb(216,172,130)',
    '--synapse-background': 'rgb(255,227,189)',
    '--synapse-border': 'rgb(216,172,130)',
    '--synapse-focus': 'rgb(224,228,203)',
    '--synapse-accent': 'rgb(62,31,5)',
  }

  const customThemeWeird = {
    '--synapse-text': 'red',
    '--synapse-secondary': 'green',
    '--synapse-select-bg': 'gray',
    '--synapse-surface': 'purple',
    '--synapse-background': 'orange',
    '--synapse-border': 'blue',
    '--synapse-focus': 'yellow',
    '--synapse-accent': 'red',
  }

  const { web3Provider, connectedAddress, connectedNetwork } =
    useEthereumWallet()

  const toggleContainer = (e: React.ChangeEvent<HTMLInputElement>) =>
    setContainer(e.target.checked)

  const bridgeContainerDisplayProperty = container ? 'grid' : 'block'

  return (
    <>
      <Header />

      <main>
        <header>
          <h1>Install the Synapse Bridge</h1>
          <pre>npm synapse-widget</pre>
          <p>
            Easily onboard new users by adding a custom instance of the Synapse
            Bridge to your React project.
          </p>
          <div id="example-container">
            <div
              id="bridge-container"
              style={{ display: bridgeContainerDisplayProperty, }}
            >
              <Bridge
                chainIds={chainIds}
                web3Provider={web3Provider}
                networkProviders={providers}
                tokens={tokens}
                customTheme={customTheme}
                container={container}
                toChainId={137}
              />
            </div>
            <div style={{ display: 'flex', justifyContent: 'space-between'}}>
              <div style={{ display: 'flex', gap: '1rem'}}>
                <input id="color-picker" type="color" onInput={createCustomTheme} />
                <div className="flex items-center gap-1">
                  <label>Container</label> <input type="checkbox" checked={container} onChange={toggleContainer}/>
                </div>
              </div>
              <span className="desktop-only">Drag to resize</span>
            </div>
          </div>
        </header>
        <article>
          <h2><Install />Install</h2>
          <p>Install the Synapse Widget in your Next.js or React project</p>
          <pre style={{ fontSize: '100%' }}>npm synapse-widget</pre>

          <h2><Developer />Setup</h2>
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

          <h3>Appearance</h3>
          <h4>Dark mode</h4>
          <p>
            To override the default light theme, set <code>bgColor</code> to <code>'dark'</code>.
          </p>
          <pre>customTheme = &#123; bgColor: 'dark' &#125;</pre>
          <h4>Auto-palette</h4>
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
          {/* <h4>Accent Color</h4>
          Add an accent color to text links and button hover states by setting accentColor to any hex, rgb, or hsl color string.
          <pre>
          customTheme = &#123;
            <br />  accentColor: '#d557ff'
            <br />  accentColor: 'rgb(213 87 255)'
            <br />  accentColor: 'hsl(285deg 100% 67%)'
            <br />&#125;
          </pre> */}
          <h4>Global Overrides</h4>
          <p>
            The following CSS variables can be added to your CustomTheme to
            override the generated values. Any valid CSS color string can be
            used, including var() aliases.
          </p>
          <pre>
            {`customTheme = {
  --synapse-text: 'white'
  --synapse-secondary: '#cccccc'
  --synapse-focus: 'hsl(285deg 100% 33%)'
  --synapse-border: 'hsl(210deg 100% 25%)'
  --synapse-object: 'hsl(210deg 100% 50%)'
  --synapse-surface: 'hsl(210deg 100% 12.5%)'
  --synapse-root: 'inherit'
}`}
          </pre>
          <h4>Object Overrides</h4>
          <p>
            Select and button elements can be specifically overriddden to introduce brand colors or custom styles.
          </p>
          <pre>
            {`customTheme = {
  --synapse-select-bg: 'var(--synapse-object)'
  --synapse-select-text: 'white'
  --synapse-select-border: 'var(--synapse-object)'

  --synapse-button-bg: 'var(--synapse-object)'
  --synapse-button-text: 'white'
  --synapse-button-border: 'var(--synapse-object)'
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
          <h2><Support />Support</h2>
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
