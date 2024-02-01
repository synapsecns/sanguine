const recommendedParametersCodeBlock = `
  import { Bridge, CustomRpcs } from ‘@synapsecns/widget’

  const customRpcs: CustomRpcs =  {
    1: 'https://ethereum.my-custom-rpc.com',
    10: 'https://optimism.my-custom-rpc.com',
    42161: 'https://arbitrum.my-custom-rpc.com',
  }

  const MyApp = () => {
    const web3Provider = new ethers.BrowserProvider(window.ethereum)

    <Bridge
      web3Provider={web3Provider}
      customRpcs={customRpcs}
    />
  }
`

export const RecommendedParameters = () => {
  return (
    <>
      <p>
        The bridge widget is a React component designed for straightforward
        integration into any React-based project. Engineered for immediate
        functionality, and apart from a <code>web3Provider</code>, it requires
        no initial parameters or web3 setup to begin operation. The widget
        facilitates bridging across all networks where the Synapse Protocol is
        active.
      </p>
      <p>
        While the widget is primed for immediate use without configuration as it
        provides some basic primary and fallback JSON-RPC endpoints, we
        encourage developers to specify their own for enhancd performance. This
        can be done by including a <code>customRpcs</code> parameter in the
        format of an object with chain ids as keys and their associated RPC
        endpoints as values.
      </p>
      <pre>{recommendedParametersCodeBlock}</pre>
    </>
  )
}
