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
      <div className="text-2xl font-medium">{title}</div>
      <div className="opacity-50">{subtitle}</div>
    </div>
  )
}
