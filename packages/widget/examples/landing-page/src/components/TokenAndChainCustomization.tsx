const tokenAndChainCustomizationCodeBlock = `
  import { Bridge, CustomRpcs, ETH, USDC, USDT } from ‘@synapsecns/widget’

  const MyApp = () => {
    const web3Provider = new ethers.BrowserProvider(window.ethereum)

    <Bridge
      web3Provider={web3Provider}
      targetTokens={[ETH, USDC, USDT]}
      targetChainIds={[42161, 43114]}
    />
  }
  `

export const TokenAndChainCustomization = () => {
  return (
    <>
      <p>
        While the Synapse Widget supports{' '}
        <a href="https://synapseprotocol.com" target="_blank" rel="noreferrer">
          hundreds of tokens and chains
        </a>
        , for a streamlined user experience, you can render a separate instance
        of the bridge for each user need.
      </p>
      <p>
        To further tailor the bridge widget to meet the specific demands of your
        project, additional optional <code>targetTokens</code> and{' '}
        <code>targetChainIds</code>
        parameters are provided. These allow for customizing which chain and
        tokens your consuming application will support bridging to. This is
        effectively a way to filter for specific tokens on destination chain
        your application's users bridge.
      </p>
      <p className="info">
        Note: Token naming convention is based on the tokens provided by
        <code>@synapsecns/widget</code>. For example, USDC on Metis is{' '}
        <code>METISUSDC</code> instead of simply <code>USDC</code>. The
        package's <code>src/constants/bridgeable.ts</code> file contains a
        detailed list of supported tokens and the chains they live on.
        Additionally, to see a detailed list of Synapse Protocol supported
        chains, please see <code>src/constants/chains.ts</code>.
      </p>
      <p>
        The source code can be found{' '}
        <a
          href="https://github.com/synapsecns/sanguine/tree/master/packages/widget"
          target="_blank"
        >
          here
        </a>
      </p>
      <pre>{tokenAndChainCustomizationCodeBlock}</pre>
      <p className="info">
        <strong>Note</strong>: Whitelisting one side of a transaction limits the
        other side to compatible tokens automatically.
      </p>
    </>
  )
}
