import Grid from '@tw/Grid'
import { useEffect, useState } from 'react'
import { useAccount, useAccountEffect, useSwitchChain } from 'wagmi'


import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import StandardPageContainer from '@layouts/StandardPageContainer'
import { getErc20TokenAllowance } from '@/actions/getErc20TokenAllowance'
import { approveToken } from '@/utils/approveToken'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { useConnectModal } from '@rainbow-me/rainbowkit'

const CHAIN_IDS = [1, 42161, 10]
const LIFI_SPENDER = "0x1231deb6f5749ef6ce6943a275a1d3e7486f4eae"

const TOKENS = {
  USDC: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
  USDT: '0xdac17f958d2ee523a2206206994597c13d831ec7',
  WETH: '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2'
} as const;

interface TokenAllowances {
  [chainId: number]: {
    [token: string]: bigint
  }
}

const LifiPage = () => {
  const { address, isConnected, chain } = useAccount()
  const { chains, switchChain: switchNetwork } = useSwitchChain()
  const { openConnectModal } = useConnectModal()

  const [allowances, setAllowances] = useState<TokenAllowances>({})

  useEffect(() => {
    const fetchAllowances = async () => {
      if (address) {
        const newAllowances: TokenAllowances = {}

        for (const chainId of CHAIN_IDS) {
          newAllowances[chainId] = {}

          for (const [tokenName, tokenAddress] of Object.entries(TOKENS)) {
            const allowance = await getErc20TokenAllowance({
              address,
              chainId,
              tokenAddress,
              spender: LIFI_SPENDER,
            })
            newAllowances[chainId][tokenName] = allowance
          }
        }

        setAllowances(newAllowances)
      }
    }

    if (isConnected) {
      fetchAllowances()
    }
  }, [address, isConnected])

  const handleRevoke = async (chainId: number, tokenName: string, tokenAddress: string) => {
    if (chain?.id !== chainId) {
      await switchNetwork({chainId: chainId})
    }
    await approveToken(LIFI_SPENDER, chainId, tokenAddress, 0n)
    setAllowances(prev => ({
      ...prev,
      [chainId]: {
        ...prev[chainId],
        [tokenName]: 0n
      }
    }))
  }

  return (
    <LandingPageWrapper>
      <StandardPageContainer
        connectedChainId={chain?.id}
        address={address}
      >
        <div className="flex justify-between">
          <div>
            <div className="text-2xl text-white">
              Revoke Li.fi Approvals (Multi-chain)
            </div>
          </div>
        </div>
        <div className="py-6">
          <Grid
            cols={{ xs: 1 }}
            gap={6}
            className="justify-center px-2 py-16 sm:px-6 md:px-8"
          >
            <div className="pb-3 place-self-center">
              <div>
                <h3>Li.fi / Jumper is investigating an ongoing exploit, and users should revoke approvals <a className="underline" target="_blank" href="https://x.com/lifiprotocol/status/1813196697641570635">- Li.fi Tweet</a></h3>
                <br />
                <h3>Check to see if you have any approvals at risk below:</h3>
                <br />
                {isConnected ? (
                  CHAIN_IDS.map(chainId => (
                    <div key={chainId}>
                      <h4>Chain ID: {chainId}</h4>
                      {Object.entries(allowances[chainId] || {}).map(([tokenName, allowance]) => (
                        <div key={tokenName}>
                          {tokenName} Allowance: {allowance.toString()}
                          {allowance > 0n && (
                            <TransactionButton
                              className="btn btn-primary ml-2"
                              pendingLabel="Revoking..."
                              label={`Revoke ${tokenName} Approval`}
                              onClick={() => handleRevoke(chainId, tokenName, TOKENS[tokenName])}
                            />
                          )}
                        </div>
                      ))}
                    </div>
                  ))
                ) : (
                  <div className="flex flex-col justify-center h-full p-10">
                    <TransactionButton
                      style={{
                        background: 'linear-gradient(90deg, rgba(128, 0, 255, 0.2) 0%, rgba(255, 0, 191, 0.2) 100%)',
                        border: '1px solid #9B6DD7',
                        borderRadius: '4px',
                      }}
                      label="Connect wallet to check for approvals"
                      pendingLabel="Connecting"
                      onClick={() => new Promise((resolve) => {
                        openConnectModal()
                        resolve(true)
                      })}
                    />
                  </div>
                )}
              </div>
            </div>
          </Grid>
        </div>
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

export default LifiPage
