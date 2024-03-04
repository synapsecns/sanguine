import LoadingDots from '@tw/LoadingDots'

export default function ButtonLoadingDots({
  className=''
}: {
  className?: string
}) {
  return <LoadingDots className={`opacity-50 ${className}`} />
}
