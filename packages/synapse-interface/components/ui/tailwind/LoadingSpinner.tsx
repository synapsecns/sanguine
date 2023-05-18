export default function LoadingSpinner({ className }: { className?: string }) {
  return (
    <div
      className={`inline-flex items-center justify-center pr-3 ${className}`}
    >
      <div className="dot-flashing left-[12px]"></div>
    </div>
  )
}
