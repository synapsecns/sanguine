export default function TabItem({ onClick, children, isActive }) {
  let statusClassname
  if (isActive) {
    statusClassname = `
      bg-gray-100          text-gray-700
      dark:bg-gray-900 dark:text-gray-400`
  } else {
    statusClassname = `
      text-gray-500          hover:bg-gray-200          hover:text-gray-700
      dark:text-gray-500 dark:hover:bg-gray-900 dark:hover:text-gray-400
      `
  }
  return (
    <div
      onClick={onClick}
      className={`
        p-2 w-full cursor-pointer
        text-center font-medium text-sm
        rounded-md place-self-center
        transition-all duration-75
        ${statusClassname}
      `}
    >
      {children}
    </div>
  )
}
