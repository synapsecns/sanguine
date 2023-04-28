export function PageHeader({
  title,
  subtitle,
  className,
}: {
  title: string
  subtitle: string
  className?: string
}) {
  return (
    <div className={className}>
      <div className="text-2xl font-medium text-white">{title}</div>
      <div className="text-base text-white text-opacity-50">{subtitle}</div>
    </div>
  )
}
