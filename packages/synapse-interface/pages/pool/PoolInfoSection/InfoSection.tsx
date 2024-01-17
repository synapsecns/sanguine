const InfoSection = ({
  title,
  children,
  className = '',
  showDivider = false,
  showOutline = true,
}: {
  title?: string
  children: any
  className?: string
  showDivider?: boolean
  showOutline: boolean
}) => {
  return (
    <>
      {title && <h3 className="mt-4 mb-1 text-sm">{title}</h3>}
      <ul
        className={`
          rounded-md
          ${showOutline && 'border border-gray-700'}
          ${showDivider && 'divide-y divide-solid divide-[#3B363D]'}
          ${className}
        `}
      >
        {children}
      </ul>
    </>
  )
}
export default InfoSection
