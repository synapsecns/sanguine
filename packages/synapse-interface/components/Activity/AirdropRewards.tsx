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
import { ArrowUpRightIcon } from '../icons/ArrowUpRightIcon'
import { QuestionMarkCircleIcon } from '@heroicons/react/outline'
import { HoverContent } from '../Portfolio/PortfolioContent/SingleNetworkPortfolio'

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

  /** Dialog state */
  const [open, setOpen] = useState<boolean>(false)
  const handleToggle = () => setOpen(!open)
  const handleClose = () => setOpen(false)

  return (
    <>
      <div
        id="airdrop-rewards"
        className="flex items-center mb-2 border  cursor-pointer text-white border-greenText bg-[#0A381B]"
        onClick={handleToggle}
      >
        <RewardsTitle icon={ARB.icon} />
        <TransactionArrow className="stroke-greenText fill-transparent" />
        <div className="flex justify-between flex-1 p-3">
          <RewardsAmountDisplay
            symbol={ARB.symbol}
            tokenAmount={rewards}
            dollarAmount={convertTokensToDollarValue(rewards, arbPrice)}
          />

          <div className="flex items-center space-x-2">
            <div>Now - Mar 29</div>
            <HoverContentIcon>
              <p>
                Through Mar 29, ARB rewards are automatically
                <br />
                applied to select routes to and from Arbitrum.
              </p>

              <p>Click for more info.</p>
            </HoverContentIcon>
          </div>
        </div>
      </div>
      <RewardsDialog
        open={open}
        setOpen={setOpen}
        onClose={handleClose}
        transactions={transactions}
        rewards={rewards}
        tokenPrice={arbPrice}
      />
    </>
  )
}

const DialogWrapper = ({ open, children }) => {
  useEffect(() => {
    if (open) {
      // Disable scroll on the body when dialog is open
      document.body.style.overflow = 'hidden'
    }

    // Clean up; Re-enable scroll when dialog unmounts
    return () => {
      document.body.style.overflow = 'auto'
    }
  }, [open])

  return (
    <div
      className={`${
        open &&
        'fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-80'
      }`}
    >
      {children}
    </div>
  )
}

const RewardsDialog = ({
  open,
  setOpen,
  onClose,
  transactions,
  rewards,
  tokenPrice,
}: {
  open: boolean
  setOpen: (value: React.SetStateAction<boolean>) => void
  onClose
  transactions: any[]
  rewards
  tokenPrice
}) => {
  const dialogRef = useRef(null)

  useCloseOnOutsideClick(dialogRef, onClose)

  const maxHeight = window.innerHeight > 768 ? 300 : 200

  return (
    <DialogWrapper open={open}>
      <dialog
        id="rewards-dialog"
        ref={dialogRef}
        open={open}
        className="absolute z-50 max-w-md p-4 m-auto text-white border rounded-lg cursor-default bg-background border-separator"
      >
        <div className="space-y-4">
          <div className="flex justify-between mb-2">
            <div className="text-2xl">ARB Rewards</div>
            <CloseButton onClick={onClose} />
          </div>

          <p>
            Through Mar 29, ARB rewards are automatically applied to select
            routes to and from Arbitrum.
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
            for full route and rebate information.
          </p>

          <div className="flex flex-wrap-reverse">
            <div className="mr-4 min-w-1/2">
              <div className="text-lg text-greenText">Total Arb</div>
              <div className="flex space-x-1">
                <div className="text-2xl text-greenText">+{rewards}</div>
                <div className="text-2xl text-secondary">
                  (${convertTokensToDollarValue(rewards, tokenPrice)})
                </div>
              </div>
            </div>

            <div>
              <div className="text-lg">Days remaining</div>
              <div className="text-2xl">{calculateDaysUntilStipEnds()}</div>
            </div>
          </div>

          <AirdropTxHeader />
        </div>
        {_.isEmpty(transactions) ? (
          <div className="h-64 pt-6 text-center text-secondary">
            Individual rebates will appear here.
          </div>
        ) : (
          <div className={`overflow-y-scroll`} style={{ maxHeight: maxHeight }}>
            {transactions.map((transaction) => (
              <AirdropTransaction
                transactionHash={transaction.transactionHash}
                tokenValue={parseTokenValue(
                  transaction.transferValue,
                  ARB.decimals
                )}
                tokenPrice={tokenPrice}
                explorerUrl={ARB.explorerUrl}
              />
            ))}
          </div>
        )}
      </dialog>
    </DialogWrapper>
  )
}

const AirdropTransaction = ({
  transactionHash,
  tokenValue,
  tokenPrice,
  explorerUrl,
}: {
  transactionHash: string
  tokenValue: string
  tokenPrice: string | number
  explorerUrl: string
}) => {
  const [isHovered, setIsHovered] = useState<boolean>(false)
  return (
    <Link
      id="airdrop-transaction"
      href={getBlockExplorerTransactionLink({ explorerUrl, transactionHash })}
      referrerPolicy="no-referrer"
      target="_blank"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
      className="grid grid-cols-3 py-1.5 text-white hover:bg-tint cursor-pointer"
    >
      <div className="text-greenText">+ {tokenValue}</div>
      <div>${convertTokensToDollarValue(tokenValue, tokenPrice)}</div>
      <div className="flex items-center ml-auto mr-4 text-right">
        {shortenAddress(transactionHash, 5)}
        {isHovered && (
          <ArrowUpRightIcon className="w-4 h-4 stroke-[3px] ml-1" />
        )}
      </div>
    </Link>
  )
}

const AirdropTxHeader = () => {
  return (
    <div className="grid grid-cols-3 text-white border-none">
      <div className="text-greenText">ARB</div>
      <div>Value</div>
      <div className="mr-4 text-right">Tx Hash</div>
    </div>
  )
}

const RewardsTitle = ({ icon }) => {
  return (
    <div id="rewards-title" className="flex items-center space-x-1.5 p-3">
      <Image
        src={icon}
        alt="reward chain icon"
        className="w-4 h-4 rounded-full"
      />
      <div className="text-md">Rewards</div>
    </div>
  )
}

const RewardsAmountDisplay = ({
  symbol,
  tokenAmount,
  dollarAmount,
}: {
  symbol: string
  tokenAmount: string
  dollarAmount: string
}) => {
  return (
    <div
      id="rewards-amount-display"
      className="flex flex-wrap items-center space-x-1.5 leading-1"
    >
      <div className="text-white text-md">+${dollarAmount}</div>
      <div className="text-sm text-secondary">
        ({tokenAmount} {symbol})
      </div>
    </div>
  )
}

export const HoverContentIcon = ({ children }) => {
  const [isHovered, setIsHovered] = useState<boolean>(false)
  return (
    <div
      id="hover-content-icon"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      <QuestionMarkCircleIcon
        className="w-5 h-5 fill-white stroke-[#0A381B]"
        onMouseEnter={() => setIsHovered(true)}
        onMouseLeave={() => setIsHovered(false)}
      />
      <HoverContent isHovered={isHovered}>{children}</HoverContent>
    </div>
  )
}

/** Helper Functions */
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

const convertTokensToDollarValue = (
  tokenAmount: number | string,
  tokenPrice: number | string
) => {
  return (Number(tokenAmount) * Number(tokenPrice)).toFixed(2)
}

const calculateDaysUntilStipEnds = () => {
  const currentDate = new Date()
  const targetDate = new Date('2024-03-29')

  const timeDifference = Number(targetDate) - Number(currentDate)
  const daysDifference = Math.floor(timeDifference / (1000 * 60 * 60 * 24))

  return daysDifference
}

export const getBlockExplorerTransactionLink = ({
  explorerUrl,
  transactionHash,
}: {
  explorerUrl: string
  transactionHash: string
}) => {
  return `${explorerUrl}/tx/${transactionHash}`
}
