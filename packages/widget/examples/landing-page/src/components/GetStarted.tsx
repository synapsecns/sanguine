const getStartedCodeBlock = `
  import { Bridge } from ‘@synapsecns/widget’

  const MyApp = () => {
    const web3Provider = new ethers.BrowserProvider(window.ethereum)

    <Bridge
      web3Provider={web3Provider}
    />
  }
`

export const GetStarted = () => {
  return (
    <>
      <p>
        To get started, import the <code>Widget</code> React component into your
        App. You will need a <code>web3Provider</code> parameter to pass to the
        widget.{' '}
      </p>
      <p>
        While this demo uses a provider from the <code>ethers</code> library, the component supports any similar provider:
      </p>
      <pre>{getStartedCodeBlock}</pre>
      <p>
        Your site should now display a fully operational bridge widget
        integrating the routes and tokens supported by the Synapse protocol. By
        utilizing Synapse's multiple routers, you'll receive the best
        quotes for your bridging use case.
      </p>
    </>
  )
}
