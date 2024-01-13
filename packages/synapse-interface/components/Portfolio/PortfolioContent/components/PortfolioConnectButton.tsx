import { useMemo, useState } from 'react'
import { useDispatch } from 'react-redux'
import { switchNetwork } from '@wagmi/core'
import { setFromChainId } from '@/slices/bridge/reducer'

type PortfolioConnectButton = {
  portfolioChainId: number
  connectedChainId: number
}

export const PortfolioConnectButton = ({
  portfolioChainId,
  connectedChainId,
}: PortfolioConnectButton) => {
  const isCurrentlyConnectedNetwork: boolean = useMemo(() => {
    return portfolioChainId === connectedChainId
  }, [portfolioChainId, connectedChainId])

  return (
    <div data-test-id="portfolio-connect-button">
      {isCurrentlyConnectedNetwork ? (
        <ConnectedButton />
      ) : (
        <ConnectButton chainId={portfolioChainId} />
      )}
    </div>
  )
}

const buttonStyles = `
  flex gap-2 items-center text-sm
  px-3 py-1 rounded-full
  border border-transparent
  hover:border-zinc-500
`

const ConnectedButton = () => {
  return (
    <button
      data-test-id="connected-button"
      className={`${buttonStyles} hover:opacity-70 hover:cursor-not-allowed`}
    >
      <>
        <span className="w-2 h-2 bg-green-500 rounded-full" />
        Connected
      </>
    </button>
  )
}

const ConnectButton = ({ chainId }: { chainId: number }) => {
  const [isConnecting, setIsConnecting] = useState<boolean>(false)
  const dispatch = useDispatch()

  function scrollToTop(): void {
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }

  const handleConnectNetwork: () => Promise<void> = async () => {
    setIsConnecting(true)
    try {
      await switchNetwork({ chainId: chainId }).then((success) => {
        success && dispatch(setFromChainId(chainId))
        scrollToTop()
      })
    } catch (error) {
      error && setIsConnecting(false)
    }
  }

  return (
    <button
      data-test-id="connect-button"
      className={buttonStyles}
      onClick={handleConnectNetwork}
    >
      {isConnecting ? (
        <>
          <span className="w-2 h-2 border border-green-300 rounded-full" />
          Connecting…
        </>
      ) : (
        <>
          <span className="w-2 h-2 border border-indigo-300 rounded-full" />
          Connect
        </>
      )}
    </button>
  )
}
