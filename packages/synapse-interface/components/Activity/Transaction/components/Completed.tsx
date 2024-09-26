import { Address } from 'viem'
import { useTranslations } from 'next-intl'

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

  const t = useTranslations('Completed')

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
        <div>
          {t('to')} {shortenAddress(destinationAddress)}{' '}
        </div>
      )}
      {isToday ? (
        <div>{t('Today')}</div>
      ) : (
        <div>{formattedTime ? formattedTime : t('Completed')}</div>
      )}
    </div>
  )
}
