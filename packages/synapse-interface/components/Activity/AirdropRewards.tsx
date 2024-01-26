import { useState, useEffect } from 'react'
import { useAccount } from 'wagmi'
import { arbitrum } from 'viem/chains'
import { getErc20TokenTransfers } from '@/utils/actions/getErc20TokenTransfers'
import TransactionArrow from '../icons/TransactionArrow'

const arbTokenAddress = '0x912CE59144191C1204E64559FE8253a0e49E6548' // on Arbitrum
const rewarderAddress = '0x48fa1ebda1af925898c826c566f5bb015e125ead' // on Arbitrum
const network = arbitrum
const startBlock = 174234366n // Start of STIP Rewards on Arbitrum

export const AirdropRewards = () => {
  const [rewards, setRewards] = useState<any>(undefined)
  const { address: connectedAddress } = useAccount()

  useEffect(() => {
    if (connectedAddress) {
      ;(async () => {
        const { logs, data } = await getErc20TokenTransfers(
          arbTokenAddress,
          rewarderAddress,
          connectedAddress,
          network,
          startBlock
        )

        setRewards(data)
      })()
    }
  }, [connectedAddress])

  return (
    <div
      id="airdrop-rewards"
      className="flex border rounded-lg text-secondary border-surface bg-background"
    >
      <div className="text-green-500">Rebate</div>
      <TransactionArrow />
      <div>
        <div>Token Img, Token Name</div>
        <div>Token Img, Airdropped Total</div>
      </div>
    </div>
  )
}
