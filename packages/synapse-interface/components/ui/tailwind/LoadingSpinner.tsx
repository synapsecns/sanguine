export default function LoadingSpinner({
  className,
  shift = false,
}: {
  className?: string
  shift?: boolean
}) {
  return (
    <div
      className={`inline-flex items-center justify-center pr-3 ${className}`}
    >
      <div className={`dot-flashing ${shift ? 'left-[12px]' : ''}`}></div>
    </div>
  )
}
