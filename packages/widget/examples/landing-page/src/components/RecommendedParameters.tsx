const recommendedParametersCodeBlock = `
  import { Bridge, CustomRpcs } from ‘@synapsecns/widget’

  const customRpcs: CustomRpcs =  {
    1: 'https://ethereum.my-custom-rpc.com',
    10: 'https://optimism.my-custom-rpc.com',
    42161: 'https://arbitrum.my-custom-rpc.com',
  }

  const MyApp = () => {
    const web3Provider = new ethers.BrowserProvider(window.ethereum)

    return (
      <Bridge
        web3Provider={web3Provider}
        customRpcs={customRpcs}
      />
    )
  }
`

export const RecommendedParameters = () => {
  return (
    <>
      <p>
        The Bridge widget is a React component designed for straightforward integration into any React-based project. Apart from a <code>web3Provider</code>, it requires no initial parameters or web3 setup to begin operation. The widget bridges across all networks where the <a href="https://synapseprotocol.com" target="_blank">Synapse Protocol</a> is active.
      </p>
      <p>
        While the widget is provides basic primary and fallback JSON-RPC endpoints, we encourage developers to specify their own for enhanced performance. This is done through a <code>customRpcs</code> parameter object with chain ID keys and associated RPC endpoint values.
      </p>
      <pre>{recommendedParametersCodeBlock}</pre>
    </>
  )
}
