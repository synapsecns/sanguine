import { Address } from 'viem'
import { shortenAddress } from '@/utils/shortenAddress'
import { convertUnixTimestampToMonthAndDate } from '@/utils/time'
import { isTimestampToday } from '@/utils/time'

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

  const destinationIsSender: boolean =
    String(connectedAddress) === String(destinationAddress)

  return (
    <div
      data-test-id="completed"
      className="flex flex-col text-right text-[#C2C2D6] gap-1 text-sm whitespace-nowrap"
      onClick={handleExplorerClick}
    >
      {!destinationIsSender && (
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
