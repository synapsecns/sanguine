export default function LoadingDots({
  className,
  shift = false,
}: {
  className?: string
  shift?: boolean
}) {
  return (
    <div data-test-id="loading-dots" className={`flex relative ${className}`}>
      <div className={`dot-flashing ${shift ? 'left-[12px]' : ''}`}></div>
    </div>
  )
}
