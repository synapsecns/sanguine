const TabItem = ({
  onClick,
  children,
  isActive,
  className,
}: {
  onClick: any
  children: any
  isActive: boolean
  className?: string
}) => {
  let statusClassname
  if (isActive) {
    statusClassname = `
      border-transparent
      text-white
      ${className}
      `
  } else {
    statusClassname = `
      bg-slate-900/50
      hover:bg-bgBase/10
      text-secondaryTextColor/50
      border-bgBase/10
      ${className}
      `
  }
  return (
    <div
      onClick={onClick}
      className={`
        border-b border-l border-r
        ${statusClassname}
        flex justify-center items-center
        w-full cursor-pointer
        font-base hover:text-white
        border-collapse
        h-14
      `}
    >
      {children}
    </div>
  )
}
export default TabItem
