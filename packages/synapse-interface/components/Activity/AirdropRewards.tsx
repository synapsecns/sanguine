import { useState, useEffect } from 'react'
import Image from 'next/image'
import { Address, useAccount } from 'wagmi'
import { arbitrum } from 'viem/chains'
import { getErc20TokenTransfers } from '@/utils/actions/getErc20TokenTransfers'
import TransactionArrow from '../icons/TransactionArrow'
import arbitrumImg from '@assets/chains/arbitrum.svg'
import { formatBigIntToString } from '@/utils/bigint/format'
import { trimTrailingZeroesAfterDecimal } from '@/utils/trimTrailingZeroesAfterDecimal'
import { shortenAddress } from '@/utils/shortenAddress'

/** ARB Token */
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

  const cumulativeRewards = calculateTotalTransferValue(data)

  return {
    logs: logs ?? [],
    transactions: data,
    cumulativeRewards,
  }
}

const calculateTotalTransferValue = (data: any[]): bigint => {
  let total: bigint = 0n
  for (const item of data) {
    if (item.transferValue) {
      total += item.transferValue
    }
  }
  return total
}

const parseTokenValue = (rawValue: bigint, tokenDecimals: number) => {
  return trimTrailingZeroesAfterDecimal(
    formatBigIntToString(rawValue, tokenDecimals, 3)
  )
}

export const AirdropRewards = () => {
  const [rewards, setRewards] = useState<string>(undefined)
  const [transactions, setTransactions] = useState<any[]>([])
  const { address: connectedAddress } = useAccount()

  const fetchStipAirdropRewards = async (address: Address) => {
    const { transactions, cumulativeRewards } = await getArbStipRewards(address)

    const parsedCumulativeRewards = parseTokenValue(
      cumulativeRewards,
      ARB.decimals
    )

    setTransactions(transactions)
    setRewards(parsedCumulativeRewards)
  }

  useEffect(() => {
    if (connectedAddress) {
      fetchStipAirdropRewards(connectedAddress)
    } else {
      setRewards(undefined)
    }
  }, [connectedAddress])

  console.log('rewards:', rewards)
  console.log('transactions:', transactions)

  return (
    <div
      id="airdrop-rewards"
      className="flex border rounded-lg text-secondary border-surface bg-background"
    >
      <div className="text-green-500">Rebate</div>
      <TransactionArrow />
      <div>
        <NetworkDisplay name={ARB.name} icon={ARB.icon} />
        <TokenAmountDisplay
          symbol={ARB.symbol}
          icon={ARB.icon}
          amount={`+ ${rewards}`}
        />
      </div>
      <RewardsDialog transactions={transactions} />
    </div>
  )
}

const RewardsDialog = ({ transactions }: { transactions: any[] }) => {
  const [open, setOpen] = useState<boolean>(true)

  console.log('transactions:', transactions)

  return (
    <dialog open={open} className="absolute bg-background">
      {transactions.map((transaction) => (
        <AirdropTransaction
          transactionHash={transaction.transactionHash}
          value={transaction.transferValue}
          blockNumber={transaction.blockNumber}
        />
      ))}
    </dialog>
  )
}

const AirdropTransaction = ({
  transactionHash,
  value,
  blockNumber,
}: {
  transactionHash: string
  value: string
  blockNumber: string
}) => {
  return (
    <div className="flex justify-between text-white">
      <div>{shortenAddress(transactionHash)}</div>

      <div className="flex">
        <div className="text-green-500">+ {value} ARB</div>
        <div>{blockNumber}</div>
      </div>
    </div>
  )
}

const NetworkDisplay = ({ name, icon }: { name: string; icon: string }) => {
  return (
    <div id="network-display" className="flex items-center space-x-1.5">
      <Image src={icon} alt={`${name} icon`} className="w-4 h-4 rounded-full" />
      <div className="text-md">{name}</div>
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
    <div
      id="token-amount-display"
      className="flex items-center space-x-1.5 leading-none"
    >
      <Image
        src={icon}
        alt={`${symbol} icon`}
        className="w-5 h-5 rounded-full"
      />
      <div className="text-white text-md">{amount}</div>
      <div className="text-sm">{symbol}</div>
    </div>
  )
}
