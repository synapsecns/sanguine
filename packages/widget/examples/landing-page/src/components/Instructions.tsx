import { Install, Developer, Support } from './Icons'
import { PackageInstall } from './PackageInstall'

export function Instructions() {
  return (
    <article>
      <h2>
        <Install />
        Install
      </h2>
      <p>Install the Synapse Widget in your Next.js or React project</p>
      <PackageInstall />
      <h2>
        <Developer />
        Setup
      </h2>
      <h3>Supported tokens</h3>
      <p>
        While the Synapse Widget supports{' '}
        <a href="https://synapseprotocol.com" target="_blank" rel="noreferrer">
          hundreds of tokens and chains
        </a>
        , for a streamlined user experience, you can render a separate instance
        of the bridge for each user need.
      </p>
      <p>
        For example: separate <code>BridgeIn</code> and <code>BridgeOut</code>{' '}
        functions allow you to define the tokens you support sending and
        receiving.
      </p>
      <p className="info">
        <strong>Note</strong>: Whitelisting one side of a transaction limits the
        other side to compatible tokens automatically.
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
        To override the default light theme, set <code>bgColor</code> to{' '}
        <code>'dark'</code>.
      </p>
      <pre>customTheme = &#123; bgColor: 'dark' &#125;</pre>
      <h4>Auto-palette</h4>
      <p>
        Generate a palette based on your brand colors by setting bgColor to any
        hex, rgb, or hsl color string. Hex values must contain 6 characters.
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
        The following CSS variables can be added to your CustomTheme to override
        the generated values. Any valid CSS color string can be used, including
        var() aliases.
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
        Select and button elements can be specifically overriddden to introduce
        brand colors or custom styles.
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
      <h2>
        <Support />
        Support
      </h2>
      <p>
        For help and feedback, reach out to our Support team in the{' '}
        <a href="#" target="_blank" rel="noreferrer">
          Synapse Discord channel.
        </a>
      </p>
    </article>
  )
}
