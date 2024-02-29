export const TransactionStatus = ({
  string,
  className,
}: {
  string: string
  className?: string
}) => {
  return (
    <span id="transaction-status" className={className}>
      {string}
    </span>
  )
}
