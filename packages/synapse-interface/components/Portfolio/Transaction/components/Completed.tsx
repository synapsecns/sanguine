import { Address } from 'viem'
import { shortenAddress } from '@/utils/shortenAddress'
import { convertUnixTimestampToMonthAndDate } from '@/utils/time'
import { isTimestampToday } from '@/utils/time'
import { isValidAddress, getValidAddress } from '@/utils/isValidAddress'

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
    getValidAddress(connectedAddress) === getValidAddress(destinationAddress)

  const isDestinationValid: boolean = isValidAddress(destinationAddress)

  return (
    <div
      data-test-id="completed"
      onClick={handleExplorerClick}
      className={`
        flex flex-col text-right gap-1 text-sm whitespace-nowrap
        ${
          isToday
            ? 'text-[#3BDD77] hover:underline cursor-pointer'
            : 'text-[#C2C2D6] cursor-pointer hover:underline'
        }
        `}
    >
      {isDestinationValid && !isDestinationSender && (
        <div>to {shortenAddress(destinationAddress)} </div>
      )}
      {isToday ? (
        <div>Today</div>
      ) : (
        <div>{formattedTime ? formattedTime : 'Completed'}</div>
      )}
    </div>
  )
}
