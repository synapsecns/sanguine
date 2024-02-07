import React from 'react'
import { TransactionStatus } from '../Transaction'
import ProcessingIcon from '@/components/icons/ProcessingIcon'

export const EstimatedDuration = ({
  timeRemaining,
  transactionStatus,
}: {
  timeRemaining: number
  transactionStatus: TransactionStatus
}) => {
  return (
    <div
      data-test-id="estimated-duration"
      className="text-[#C2C2D6] text-sm flex flex-col"
    >
      {timeRemaining >= 0 ? (
        <React.Fragment>
          <div>
            {timeRemaining} - {timeRemaining + 1} min
          </div>
          {transactionStatus !== TransactionStatus.PENDING_WALLET_ACTION && (
            <ProcessingIcon className="fill-[#343036] mt-0.5" />
          )}
        </React.Fragment>
      ) : (
        <React.Fragment>
          <div>Waiting... </div>
          <ProcessingIcon className="fill-[#343036] mt-0.5" />
        </React.Fragment>
      )}
    </div>
  )
}
