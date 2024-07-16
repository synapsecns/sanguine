import { useEffect, useState } from 'react'
import { useAccount, useSwitchChain } from 'wagmi'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import StandardPageContainer from '@layouts/StandardPageContainer'
import { getErc20TokenAllowance } from '@/actions/getErc20TokenAllowance'
import { approveToken } from '@/utils/approveToken'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import Image from 'next/image'
import { CHAINS_BY_ID } from '@/constants/chains'

const CHAIN_IDS = [1, 56, 42161, 10]
const LIFI_SPENDER = '0x1231deb6f5749ef6ce6943a275a1d3e7486f4eae'

const MAX_AMOUNT = 100000000000000000000000000000n

const TOKENS = {
  1: {
    // Ethereum Mainnet
    USDC: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
    USDT: '0xdac17f958d2ee523a2206206994597c13d831ec7',
    WETH: '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2',
  },
  42161: {
    // Arbitrum
    USDC: '0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8',
    USDT: '0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9',
    WETH: '0x82aF49447D8a07e3bd95BD0d56f35241523fBab1',
  },
  10: {
    // Optimism
    USDC: '0x7F5c764cBc14f9669B88837ca1490cCa17c31607',
    USDT: '0x94b008aA00579c1307B0EF2c499aD98a8ce58e58',
    WETH: '0x4200000000000000000000000000000000000006',
  },
  56: {
    // BSC
    USDC: '0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d',
    USDT: '0x55d398326f99059ff775485246999027b3197955',
    WETH: '0x4DB5a66E937A9F4473fA95b1cAF1d1E1D62E29EA',
  },
}

interface TokenAllowances {
  [chainId: number]: {
    [token: string]: bigint
  }
}

const LifiPage = () => {
  const { address, isConnected, chain } = useAccount()
  const { chains, switchChain: switchNetwork } = useSwitchChain()
  const { openConnectModal } = useConnectModal()
  const [hoveredToken, setHoveredToken] = useState(null)

  const [allowances, setAllowances] = useState<TokenAllowances>({})

  useEffect(() => {
    const fetchAllowances = async () => {
      if (address) {
        const newAllowances: TokenAllowances = {}

        for (const chainId of CHAIN_IDS) {
          newAllowances[chainId] = {}

          for (const [tokenName, tokenAddress] of Object.entries(
            TOKENS[chainId]
          )) {
            const allowance = await getErc20TokenAllowance({
              address,
              chainId,
              tokenAddress: tokenAddress as `0x${string}`,
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

  const handleRevoke = async (chainId: number, tokenName: string) => {
    if (chain?.id !== chainId) {
      await switchNetwork({ chainId: chainId })
    }
    const tokenAddress = TOKENS[chainId][tokenName]
    await approveToken(LIFI_SPENDER, chainId, tokenAddress, 0n)
    setAllowances((prev) => ({
      ...prev,
      [chainId]: {
        ...prev[chainId],
        [tokenName]: 0n,
      },
    }))
  }

  return (
    <LandingPageWrapper>
      <StandardPageContainer connectedChainId={chain?.id} address={address}>
        <div className="flex justify-between">
          <div>
            <div className="text-2xl text-white">
              Revoke Li.fi Approvals (Multi-chain)
            </div>
          </div>
        </div>
        <div className="py-6">
          <div className="pb-3 place-self-center">
            <div>
              <h3 className="text-lg">
                Li.fi / Jumper is investigating an ongoing exploit, and users
                should revoke approvals{' '}
                <a
                  className="underline"
                  target="_blank"
                  href="https://x.com/lifiprotocol/status/1813207291778215955"
                >
                  - Li.fi Tweet
                </a>
              </h3>
              <br />
              <h3 className="text-lg">
                Check to see if you have any approvals at risk below:
              </h3>
              <br />
              {isConnected ? (
                CHAIN_IDS.map((chainId) => (
                  <div key={chainId} className="mb-5 ">
                    <span className="flex items-center gap-1 mb-2">
                      <Image
                        loading="lazy"
                        src={CHAINS_BY_ID[chainId].chainImg}
                        alt={`${chain.name} img`}
                        width="20"
                        height="20"
                        className="w-5 h-5 max-w-fit"
                      />
                      <div className="text-xl">
                        {CHAINS_BY_ID[chainId].name}
                      </div>
                    </span>
                    {Object.entries(allowances[chainId] || {}).map(
                      ([tokenName, allowance]) => {
                        return (
                          <div
                            key={tokenName}
                            className="flex items-center space-x-8 h-[40px]"
                          >
                            <div className="flex justify-between w-1/2">
                              <div>{tokenName} </div>
                              <div
                                onMouseEnter={() =>
                                  setHoveredToken(
                                    allowance > MAX_AMOUNT ? tokenName : null
                                  )
                                }
                                onMouseLeave={() => setHoveredToken(null)}
                                className="hover:cursor-pointer"
                              >
                                {allowance > MAX_AMOUNT && <div>UNLIMITED</div>}
                                {hoveredToken &&
                                  hoveredToken === tokenName &&
                                  allowance > 0n && (
                                    <div className="absolute p-2 mt-4 bg-gray-500">
                                      {allowances[chainId][
                                        hoveredToken
                                      ].toString()}
                                    </div>
                                  )}
                              </div>
                            </div>
                            {allowance > 0n ? (
                              <TransactionButton
                                className="btn btn-primary rounded-md max-w-[250px] h-[40px]"
                                pendingLabel="Revoking..."
                                label={`Revoke ${tokenName} Approval`}
                                onClick={() => handleRevoke(chainId, tokenName)}
                              />
                            ) : (
                              <div className="w-[250px] flex justify-center">
                                -
                              </div>
                            )}
                          </div>
                        )
                      }
                    )}
                  </div>
                ))
              ) : (
                <div className="flex flex-col justify-center h-full p-10">
                  <TransactionButton
                    style={{
                      background:
                        'linear-gradient(90deg, rgba(128, 0, 255, 0.2) 0%, rgba(255, 0, 191, 0.2) 100%)',
                      border: '1px solid #9B6DD7',
                      borderRadius: '4px',
                    }}
                    label="Connect wallet to check for approvals"
                    pendingLabel="Connecting"
                    onClick={() =>
                      new Promise((resolve) => {
                        openConnectModal()
                        resolve(true)
                      })
                    }
                  />
                </div>
              )}
            </div>
          </div>
        </div>
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

export default LifiPage
