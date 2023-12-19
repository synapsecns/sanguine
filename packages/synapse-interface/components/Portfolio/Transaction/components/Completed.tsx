import { Address } from 'viem'
import { shortenAddress } from '@/utils/shortenAddress'
import { convertUnixTimestampToMonthAndDate } from '@/utils/time'
import { isTimestampToday } from '@/utils/time'
import { isValidAddress } from '@/utils/isValidAddress'

export const Completed = ({
  transactionCompletedTime,
  connectedAddress,
  destinationAddress,
  handleExplorerClick,
}: {
  transactionCompletedTime: number
  connectedAddress?: Address | string
  destinationAddress: string
  handleExplorerClick: () => void
}) => {
  const formattedTime: string =
    transactionCompletedTime &&
    convertUnixTimestampToMonthAndDate(transactionCompletedTime)

  const isToday: boolean = isTimestampToday(transactionCompletedTime)

  const isDestinationSender: boolean =
    String(connectedAddress) === String(destinationAddress)

  const isDestinationValid: boolean = isValidAddress(destinationAddress)

  return (
    <div
      data-test-id="completed"
      className="flex flex-col text-right text-[#C2C2D6] gap-1 text-sm whitespace-nowrap"
      onClick={handleExplorerClick}
    >
      {isDestinationValid && !isDestinationSender && (
        <div>to {shortenAddress(destinationAddress)} </div>
      )}
      {isToday ? (
        <div className="text-[#3BDD77] hover:underline cursor-pointer">
          Today
        </div>
      ) : (
        <div className="cursor-pointer hover:underline">{formattedTime}</div>
      )}
    </div>
  )
}
