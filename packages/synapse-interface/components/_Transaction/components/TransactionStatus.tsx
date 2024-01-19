const TransactionStatus = ({
  string,
  className,
}: {
  string: string
  className?: string
}) => {
  return (
    <div id="transaction-status" className={className}>
      {string}
    </div>
  )
}
