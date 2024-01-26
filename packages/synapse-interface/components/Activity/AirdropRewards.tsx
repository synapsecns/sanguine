import { useState, useEffect } from 'react'
import Image from 'next/image'
import { Address, useAccount } from 'wagmi'
import { arbitrum } from 'viem/chains'
import { getErc20TokenTransfers } from '@/utils/actions/getErc20TokenTransfers'
import TransactionArrow from '../icons/TransactionArrow'
import arbitrumImg from '@assets/chains/arbitrum.svg'

/** Temp constant as we do not currently store this */
const ARB = {
  name: 'Arbitrum',
  symbol: 'ARB',
  decimals: 18,
  tokenAddress: '0x912CE59144191C1204E64559FE8253a0e49E6548' as Address,
  icon: arbitrumImg,
  network: arbitrum,
}

/** ARB STIP Rewarder */
const Rewarder = {
  address: '0x48fa1ebda1af925898c826c566f5bb015e125ead' as Address,
  startBlock: 174234366n, // Start of STIP Rewards on Arbitrum
}

const getArbStipRewards = async (connectedAddress: Address) => {
  const { logs, data } = await getErc20TokenTransfers(
    ARB.tokenAddress,
    Rewarder.address,
    connectedAddress,
    ARB.network,
    Rewarder.startBlock
  )
  return [logs, data]
}

export const AirdropRewards = () => {
  const [rewards, setRewards] = useState<any>(undefined)
  const { address: connectedAddress } = useAccount()

  useEffect(() => {
    if (connectedAddress) {
      ;(async () => {
        const [logs, data] = await getArbStipRewards(connectedAddress)

        console.log('Raw Rewards Transfer logs: ', logs)

        setRewards(data)
      })()
    } else {
      setRewards(undefined)
    }
  }, [connectedAddress])

  console.log('rewards:', rewards)

  return (
    <div
      id="airdrop-rewards"
      className="flex border rounded-lg text-secondary border-surface bg-background"
    >
      <div className="text-green-500">Rebate</div>
      <TransactionArrow />
      <div>
        <NetworkDisplay name={ARB.name} icon={ARB.icon} />
        <div>Token Img, Airdropped Total</div>
      </div>
    </div>
  )
}

const NetworkDisplay = ({ name, icon }: { name: string; icon: string }) => {
  return (
    <div id="network-display" className="flex items-center space-x-1.5">
      <Image src={icon} alt={`${name} icon`} className="w-4 h-4 rounded-full" />
      <div>{name}</div>
    </div>
  )
}

const TokenAmountDisplay = ({
  symbol,
  icon,
  amount,
}: {
  symbol: string
  icon: string
  amount: string
}) => {
  return (
    <div id="token-amount-display" className="flex items-center space-x-1.5">
      <Image
        src={icon}
        alt={`${symbol} icon`}
        className="w-4 h-4 rounded-full"
      />
      <div>{amount}</div>
      <div>{symbol}</div>
    </div>
  )
}
