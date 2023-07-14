export default function LoadingSpinner({
  className,
  shift = false,
}: {
  className?: string
  shift?: boolean
}) {
  return (
    <div className={`flex relative left-[12px] ${className}`}>
      <div className={`dot-flashing ${shift ? 'left-[12px]' : ''}`}></div>
    </div>
  )
}
