import Grid from '@tw/Grid'
import { useEffect, useState } from 'react'
import { useAccount } from 'wagmi'

import { DEFAULT_FROM_CHAIN } from '@/constants/swap'
import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import StandardPageContainer from '@layouts/StandardPageContainer'
import { getErc20TokenAllowance } from '@/actions/getErc20TokenAllowance'
import { approveToken } from '@/utils/approveToken'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { useConnectModal } from '@rainbow-me/rainbowkit'

const LifiPage = () => {
  const { address: currentAddress, chain, isConnected } = useAccount()
  const [connectedChainId, setConnectedChainId] = useState(0)
  const [address, setAddress] = useState(undefined)
  const { openConnectModal } = useConnectModal()

  useEffect(() => {
    setConnectedChainId(chain?.id ?? DEFAULT_FROM_CHAIN)
  }, [chain])

  useEffect(() => {
    setAddress(currentAddress)
  }, [currentAddress])

            const usdcAddress = '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48'
            const usdtAddress = '0xdac17f958d2ee523a2206206994597c13d831ec7'
            const wethAddress = '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2'

            const [usdcAllowance, setUsdcAllowance] = useState<bigint>(0n)
            const [usdtAllowance, setUsdtAllowance] = useState<bigint>(0n)
            const [wethAllowance, setWethAllowance] = useState<bigint>(0n)

            useEffect(() => {
              const fetchAllowances = async () => {
                if (address) {
                  const usdcAllowance = await getErc20TokenAllowance({
                    address: "0xbc6f5a4ed57f16af3db54da801aba8d1dc4ed675",
                    chainId: connectedChainId,
                    tokenAddress: usdcAddress,
                    spender: "0x1231deb6f5749ef6ce6943a275a1d3e7486f4eae",
                  })
                  setUsdcAllowance(usdcAllowance)

                  const usdtAllowance = await getErc20TokenAllowance({
                    address: "0xbc6f5a4ed57f16af3db54da801aba8d1dc4ed675",
                    chainId: connectedChainId,
                    tokenAddress: usdtAddress,
                    spender: "0x1231deb6f5749ef6ce6943a275a1d3e7486f4eae",
                  })
                  setUsdtAllowance(usdtAllowance)

                  const wethAllowance = await getErc20TokenAllowance({
                    address: "0xbc6f5a4ed57f16af3db54da801aba8d1dc4ed675",
                    chainId: connectedChainId,
                    tokenAddress: wethAddress,
                    spender: "0x1231deb6f5749ef6ce6943a275a1d3e7486f4eae",
                  })
                  setWethAllowance(wethAllowance)
                }
              }

              fetchAllowances()
            }, [address, connectedChainId])



  return (
    <LandingPageWrapper>
      <StandardPageContainer
        connectedChainId={connectedChainId}
        address={address}
      >
        <div className="flex justify-between">
          <div>
            <div className="text-2xl text-white">
              Revoke Li.fi Approvals
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
                <div>USDC Allowance At Risk: {usdcAllowance.toString()}</div>
                <div>USDT Allowance At Risk: {usdtAllowance.toString()}</div>
                <div>WETH Allowance at Risk: {wethAllowance.toString()}</div>
              </div>
            </div>
            {!isConnected && (
              <div className="flex flex-col justify-center h-full p-10">
                <TransactionButton
                  style={{
                    background:
                      'linear-gradient(90deg, rgba(128, 0, 255, 0.2) 0%, rgba(255, 0, 191, 0.2) 100%)',
                    border: '1px solid #9B6DD7',
                    borderRadius: '4px',
                  }}
                  label="Connect wallet"
                  pendingLabel="Connecting"
                  onClick={() =>
                    new Promise((resolve, reject) => {
                      try {
                        openConnectModal()
                        resolve(true)
                      } catch (e) {
                        reject(e)
                      }
                    })
                  }
                />
              </div>
            )}
            {usdcAllowance > 0n && (
              <TransactionButton
                className="btn btn-primary"
                pendingLabel="Revoking..."
                label="Revoke USDC Approval"
                onClick={async () => {
                  await approveToken(
                    "0x1231deb6f5749ef6ce6943a275a1d3e7486f4eae",
                    connectedChainId,
                    usdcAddress,
                    0n,
                  )
                  setUsdcAllowance(0n)
                }}
              />
            )}

            {usdtAllowance > 0n && (
              <TransactionButton
                className="btn btn-primary"
                pendingLabel="Revoking..."
                label="Revoke USDT Approval"
                onClick={async () => {
                  await approveToken(
                    "0x1231deb6f5749ef6ce6943a275a1d3e7486f4eae",
                    connectedChainId,
                    usdtAddress,
                    0n,
                  )
                  setUsdtAllowance(0n)
                }}
              />
            )}

          {wethAllowance > 0n && (
              <TransactionButton
                className="btn btn-primary"
                pendingLabel="Revoking..."
                label="Revoke WETH Approval"
                onClick={async () => {
                  await approveToken(
                    "0x1231deb6f5749ef6ce6943a275a1d3e7486f4eae",
                    connectedChainId,
                    wethAddress,
                    0n,
                  )
                  setWethAllowance(0n)
                }}
              />
            )}
          </Grid>
        </div>
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

export default LifiPage
