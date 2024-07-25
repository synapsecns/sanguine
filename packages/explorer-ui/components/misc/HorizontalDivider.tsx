interface HorizontalDividerProps {
  className?: string
}

export function HorizontalDivider({ className }: HorizontalDividerProps) {
  return (
    <div className={`divide-y-[1px] divide-white/10 ${className}`}>
      <div></div>
      <div></div>
    </div>
  )
}
