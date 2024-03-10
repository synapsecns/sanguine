const tokenAndChainCustomizationCodeBlock = `
  import { Bridge, CustomRpcs, ETH, USDC, USDT } from ‘@synapsecns/widget’

  const MyApp = () => {
    const web3Provider = new ethers.BrowserProvider(window.ethereum)

    return (
      <Bridge
        web3Provider={web3Provider}
        targetTokens={[ETH, USDC, USDT]}
        targetChainIds={[42161, 43114]}
      />
    )
  }
  `

export const TokenAndChainCustomization = () => {
  return (
    <>
      <p>
        While the Synapse Widget supports <a href="https://synapseprotocol.com" target="_blank" rel="noreferrer">hundreds of tokens and chains</a>, you can instruct the Bridge to prioritize routes to your project for a streamlined experience.
      </p>
      <p>
        Optional <code>targetTokens</code> and{' '}
        <code>targetChainIds</code>
        parameters describe the chain and tokens your consuming application supports. This is effectively a way to filter for specific tokens when onboarding or offboarding users.
      </p>
      <div className="info">
        <p>
          Note: Token naming convention is based on the tokens provided by <code>@synapsecns/widget</code>. For example, USDC on Metis is <code>METISUSDC</code> instead of simply <code>USDC</code>.
        </p>
        <p>
          See <code>src/constants/bridgeable.ts</code> for a detailed list of supported tokens and their chains, and <code>src/constants/chains.ts</code> for supported chains.
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
      </div>

      <pre>{tokenAndChainCustomizationCodeBlock}</pre>
      <p className="info">
        <strong>Note</strong>: Setting one side of a transaction automatically filters the other side to show compatible tokens only.
      </p>
    </>
  )
}
