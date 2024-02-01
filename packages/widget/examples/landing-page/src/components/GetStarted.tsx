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
        The demo landing page app, for example, defines this provider from the
        <code>ethers</code> library. However, the component supports any similar
        provider:
      </p>
      <pre>{getStartedCodeBlock}</pre>
      <p>
        Your site should now display a fully operational bridge widget
        integrating the routes and tokens supported by the Synapse protocol. By
        utilizing Synapse's multiple routers, you will be able to find the best
        quotes to support your bridging use case.
      </p>
    </>
  )
}
