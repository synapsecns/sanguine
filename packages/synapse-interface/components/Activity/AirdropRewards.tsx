import Image from 'next/image'
import Link from 'next/link'
import _ from 'lodash'
import { useAppSelector } from '@/store/hooks'
import { useState, useEffect, useRef } from 'react'
import { Address, useAccount } from 'wagmi'
import { arbitrum } from 'viem/chains'
import { trimTrailingZeroesAfterDecimal } from '@/utils/trimTrailingZeroesAfterDecimal'
import { getErc20TokenTransfers } from '@/utils/actions/getErc20TokenTransfers'
import { formatBigIntToString } from '@/utils/bigint/format'
import { shortenAddress } from '@/utils/shortenAddress'
import { ARBITRUM } from '@/constants/chains/master'
import { CloseButton } from '../StateManagedBridge/components/CloseButton'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'
import TransactionArrow from '../icons/TransactionArrow'
import arbitrumImg from '@assets/chains/arbitrum.svg'

/** ARB Token */
const ARB = {
  name: 'Arbitrum',
  symbol: 'ARB',
  decimals: 18,
  tokenAddress: '0x912CE59144191C1204E64559FE8253a0e49E6548' as Address,
  icon: arbitrumImg,
  network: arbitrum,
  explorerUrl: ARBITRUM.explorerUrl,
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
  const [rewards, setRewards] = useState<string>('0')
  const [transactions, setTransactions] = useState<any[]>([])
  const { address: connectedAddress } = useAccount()
  const { arbPrice } = useAppSelector((state) => state.priceData)

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

  // console.log('rewards:', rewards)
  // console.log('transactions:', transactions)

  const [open, setOpen] = useState<boolean>(false)

  const handleOpen = () => setOpen(true)
  const handleClose = () => setOpen(false)

  return (
    <div
      id="airdrop-rewards"
      className="flex items-center mb-2 border rounded-lg cursor-pointer text-secondary border-surface bg-background"
      onClick={handleOpen}
    >
      <RewardsTitle icon={ARB.icon} />
      <TransactionArrow className="stroke-surface fill-transparent" />
      <div className="flex justify-between flex-1">
        <RewardAmountDisplay
          symbol={ARB.symbol}
          icon={ARB.icon}
          tokenAmount={''}
          dollarAmount={''}
          // amount={`+ ${rewards}`}
        />

        <div>Now - Mar 31</div>
      </div>
      <RewardsDialog
        open={open}
        setOpen={setOpen}
        onClose={handleClose}
        transactions={transactions}
      />
    </div>
  )
}

const RewardsDialog = ({
  open,
  setOpen,
  onClose,
  transactions,
}: {
  open: boolean
  setOpen: (value: React.SetStateAction<boolean>) => void
  onClose
  transactions: any[]
}) => {
  const dialogRef = useRef(null)

  useCloseOnOutsideClick(dialogRef, onClose)

  return (
    <dialog
      id="rewards-dialog"
      ref={dialogRef}
      open={open}
      className="fixed top-[40%] z-10 p-4 text-white border rounded-lg bg-background w-96 border-separator cursor-default"
    >
      <div className="space-y-4">
        <div className="flex justify-between mb-2">
          <div className="text-2xl">ARB Rewards</div>
          <CloseButton onClick={onClose} />
        </div>

        <p>
          Through Mar 31, ARB rewards are automatically applied to select routes
          to and from Arbitrum.
        </p>

        <p>
          Click{' '}
          <Link
            href="https://synapse.mirror.xyz/NpzSkXDUlistuxNQaMwP6HQ9k2gVJsI-G1Y7-gaLxfQ"
            target="_blank"
            className="underline text-blueText"
          >
            here
          </Link>{' '}
          for full route and rebate inforamtion.
        </p>

        <div className="flex text-left">
          <div className="w-1/2">
            <div className="text-greenText">Total Arb</div>
            <div className="text-greenText">+0</div>
          </div>
          <div className="w-1/2">
            <div>Days remaining</div>
            <div>-</div>
          </div>
        </div>

        <AirdropTxHeader />
        {_.isEmpty(transactions) ? (
          <div className="text-center text-secondary">
            Individual rebates will appear here
          </div>
        ) : (
          transactions.map((transaction) => (
            <AirdropTransaction
              transactionHash={transaction.transactionHash}
              value={parseTokenValue(transaction.transferValue, ARB.decimals)} // TODO: Make dynamic so we do not hardcode decimals
              blockNumber={transaction.blockNumber.toString()}
              explorerUrl={ARB.explorerUrl}
            />
          ))
        )}
      </div>
    </dialog>
  )
}

const AirdropTxHeader = () => {
  return (
    <div className="grid grid-cols-3 text-white border-none">
      <div className="text-greenText">ARB</div>
      <div>Value</div>
      <div>Tx Hash</div>
    </div>
  )
}

const AirdropTransaction = ({
  transactionHash,
  value,
  blockNumber,
  explorerUrl,
}: {
  transactionHash: string
  value: string
  blockNumber: string
  explorerUrl: string
}) => {
  return (
    <div className="grid grid-cols-3 text-white">
      <div className="text-greenText">+ {value} ARB</div>
      <div>${value}</div>
      <Link
        href={getBlockExplorerTransactionLink({ explorerUrl, transactionHash })}
        referrerPolicy="no-referrer"
        target="_blank"
      >
        {shortenAddress(transactionHash, 5)}
      </Link>
    </div>
  )
}

// TODO: Check if pattern works with other explorers, can move to utils
export const getBlockExplorerTransactionLink = ({
  explorerUrl,
  transactionHash,
}: {
  explorerUrl: string
  transactionHash: string
}) => {
  return `${explorerUrl}/tx/${transactionHash}`
}

const RewardsTitle = ({ icon }) => {
  return (
    <div id="rewards-title" className="flex items-center space-x-1.5">
      <Image
        src={icon}
        alt="reward chain icon"
        className="w-4 h-4 rounded-full"
      />
      <div className="text-md">Rewards</div>
    </div>
  )
}

const RewardAmountDisplay = ({
  symbol,
  icon,
  tokenAmount,
  dollarAmount,
}: {
  symbol: string
  icon: string
  tokenAmount: string
  dollarAmount: string
}) => {
  return (
    <div
      id="token-amount-display"
      className="flex items-center space-x-1.5 leading-none"
    >
      {/* <div className="text-sm">{symbol}</div> */}
      <div className="text-white text-md">+${dollarAmount}</div>
      <div>
        ({tokenAmount} {symbol})
      </div>
      {/* <Image
        src={icon}
        alt={`${symbol} icon`}
        className="w-5 h-5 rounded-full"
      /> */}
    </div>
  )
}
