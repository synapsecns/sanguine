export const WarningMessage = ({
  header,
  message,
  twClassName,
}: {
  header?: string
  message?: React.ReactNode
  twClassName?: string
}) => {
  return (
    <div
      className={`flex flex-col bg-[#353038] text-white text-sm p-3 rounded-md ${twClassName}`}
    >
      {header && <div className="mb-2 font-bold">{header}</div>}
      {message && <div>{message}</div>}
    </div>
  )
}
