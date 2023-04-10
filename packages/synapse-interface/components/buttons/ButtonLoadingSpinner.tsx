import LoadingSpinner from '@tw/LoadingSpinner'

export default function ButtonLoadingSpinner({
  className,
}: {
  className?: string
}) {
  return <LoadingSpinner className={`opacity-50 ${className}`} />
}
