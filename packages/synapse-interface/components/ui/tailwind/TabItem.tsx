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
      bg-bgBase
      text-white
      ${className}
      `
  } else {
    statusClassname = `
      bg-[#111111]
      hover:bg-opacity-70
      text-secondaryTextColor text-opacity-50
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
        font-base
        h-14
      `}
    >
      {children}
    </div>
  )
}
export default TabItem
