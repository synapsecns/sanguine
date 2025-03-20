import { PopulatedTransaction } from 'ethers'

type TxWithStringValue = Omit<PopulatedTransaction, 'value'> & {
  value: string
}

export const stringifyTxValue = ({
  tx,
  preserveTx,
}: {
  tx: PopulatedTransaction | undefined
  preserveTx: boolean
}): TxWithStringValue | null => {
  if (!tx || !preserveTx) {
    return null
  }
  return {
    ...tx,
    value: tx.value.toString(),
  }
}
