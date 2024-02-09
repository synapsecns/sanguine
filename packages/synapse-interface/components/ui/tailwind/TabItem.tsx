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

      text-white
      ${className}
      `
  } else {
    statusClassname = `
      bg-bgBase/20
      hover:bg-bgBase/40
      text-secondaryTextColor/50
      ${className}
      `
  }
  return (
    <div
      onClick={onClick}
      className={`
        ${statusClassname}
        flex justify-center items-center
        w-full cursor-pointer
        font-base hover:text-white
        h-14
      `}
    >
      {children}
    </div>
  )
}
export default TabItem
