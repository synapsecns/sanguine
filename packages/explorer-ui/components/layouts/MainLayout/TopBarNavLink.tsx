interface TopBarNavLinkProps {
  labelText: string
  to?: string
  className?: string
}
const TopBarNavLink: React.FC<TopBarNavLinkProps> = ({
  labelText,
  to,
  className,
}) => {
  const linkContent = (
    <div className={`py-2 px-2 ${className}`}>
      <span
        className={`
        text-gray-300
            transform-gpu transition-all duration-75
          `}
      >
        {labelText}
      </span>
    </div>
  )

  const linkClassName = `
    group items-center px-2 my-2 font-light tracking-wide rounded-md
    bg-opacity-50
    transform-gpu transition-all duration-75
    dark:text-white
    dark:text-opacity-30
    dark:hover:text-gray-300
    dark:hover:bg-gray-800
  `
  return (
    <a href={to} target="_blank" className={linkClassName}>
      {linkContent}
    </a>
  )
  // }
}

export default TopBarNavLink
