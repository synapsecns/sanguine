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
  return (
    <div
      onClick={onClick}
      className={`
        flex place-content-center
        w-full p-4
        ${isActive
          ? "bg-zinc-100 dark:bg-zinc-800"
          : "bg-zinc-200 dark:bg-zinc-900 hover:bg-opacity-70 cursor-pointer"
        }
        ${className}
      `}
    >
      {children}
    </div>
  )
}
export default TabItem
