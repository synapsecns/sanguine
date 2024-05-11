export function getTransactionStatusInfo(statusCode: number) {
  const status = statusCode as TransactionStatus

  const statusLabel = TransactionStatus[
    status
  ] as keyof typeof TransactionStatus
  const notes = TransactionStatusNotes[status] || []
  return { code: statusLabel, notes }
}

enum TransactionStatus {
  Ready = 0,
  AlreadyExecuted = 1,
  EntryAwaitingResponses = 2,
  EntryConflict = 3,
  ReceiverNotICApp = 4,
  ReceiverZeroRequiredResponses = 5,
  TxWrongDstChainId = 6,
  UndeterminedRevert = 7,
}

const TransactionStatusNotes: { [key in TransactionStatus]: string[] } = {
  [TransactionStatus.Ready]: [],
  [TransactionStatus.AlreadyExecuted]: ['`firstArg` is the transaction ID'],
  [TransactionStatus.EntryAwaitingResponses]: [
    '`firstArg` is the number of responses received',
    '`secondArg` is the number of responses required',
  ],
  [TransactionStatus.EntryConflict]: [
    '`firstArg` is the address of the module. This is either one of the modules that the app trusts, or the Guard module used by the app',
  ],
  [TransactionStatus.ReceiverNotICApp]: ['`firstArg` is the receiver address'],
  [TransactionStatus.ReceiverZeroRequiredResponses]: [
    'the app config requires zero responses for the transaction',
  ],
  [TransactionStatus.TxWrongDstChainId]: [
    '`firstArg` is the destination chain ID',
  ],
  [TransactionStatus.UndeterminedRevert]: [
    'the transaction will revert for another reason',
  ],
}
