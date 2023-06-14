import StateManagedBridge from './state-managed-bridge'

<<<<<<< HEAD
export default StateManagedBridge
=======
const Home = () => {
  const { address: currentAddress } = useAccount()
  const { chain } = useNetwork()
  const [connectedChainId, setConnectedChainId] = useState(0)
  const [address, setAddress] = useState(undefined)

  useEffect(() => {
    setConnectedChainId(chain?.id ?? DEFAULT_FROM_CHAIN)
  }, [chain])
  useEffect(() => {
    setAddress(currentAddress)
  }, [currentAddress])

  return (
    <>
      <Head>
        <title>Synapse</title>
        <meta
          name="description"
          content="Synapse is the most widely used, extensible, secure cross-chain communications network."
        />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      {connectedChainId !== 0 ? (
        <BridgePage address={address} fromChainId={connectedChainId} />
      ) : null}
    </>
  )
}

export default Home
>>>>>>> c0cd699b (Adds wallet analytics provider)
