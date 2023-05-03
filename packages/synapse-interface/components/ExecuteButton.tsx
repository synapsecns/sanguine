import { TransactionButton } from '@components/buttons/SubmitTxButton'
const ExecuteButton = ({ onClick, disabled, title, className, children }) => {
  return (
    <TransactionButton
      className={btnClassName}
      disabled={tokenInputSum.eq(0)}
      onClick={() => buttonAction()}
      onSuccess={() => {
        postButtonAction()
      }}
      label={btnLabel}
      pendingLabel={pendingLabel}
    />
  )
}
