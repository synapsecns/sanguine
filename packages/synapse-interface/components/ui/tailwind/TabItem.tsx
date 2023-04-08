const TabItem = ({
  onClick,
  children,
  isActive,
}: {
  onClick: any
  children: any
  isActive: any
}) => {
  let statusClassname
  if (isActive) {
    statusClassname = `
      bg-[#111111]
      text-white
      `
  } else {
    statusClassname = `
      bg-bgLight
      hover:bg-opacity-70
      text-secondaryTextColor text-opacity-50
      `
  }
  return (
    <div
      onClick={onClick}
      className={`
        flex justify-center items-center
        p-2 w-full cursor-pointer
        font-medium
        rounded-xl
        transition-all duration-75
        h-14
        ${statusClassname}
      `}
    >
      {children}
    </div>
  )
}
export default TabItem
