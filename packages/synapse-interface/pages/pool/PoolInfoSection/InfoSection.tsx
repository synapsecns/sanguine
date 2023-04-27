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
  let dividerClassName
  if (showDivider) {
    dividerClassName = 'divide-y divide-solid divide-[#3B363D]'
  } else {
    dividerClassName = ''
  }
  return (
    <>
      {title && <h3 className="mt-4 mb-1 text-sm">{title}</h3>}
      <ul
        className={`${
          showOutline ? 'border border-gray-700' : ''
        } text-default rounded-md ${dividerClassName} ${className}`}
      >
        {children}
      </ul>
    </>
  )
}
export default InfoSection
