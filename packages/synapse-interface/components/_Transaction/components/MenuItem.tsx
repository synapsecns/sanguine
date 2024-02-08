export const MenuItem = ({
  text,
  link,
  onClick,
}: {
  text: string
  link: string
  onClick?: () => any
}) => {

  const className = 'flex gap-4 items-center justify-between pl-2 pr-3 py-2 whitespace-nowrap text-[--synapse-text-primary] no-underline'

  return (
    <li
      id="menu-item"
      className={`
        rounded cursor-pointer
        border border-solid border-transparent
        hover:border-[--synapse-focus] active:opacity-40
      `}
    >
      {onClick ? (
        <div
          onClick={onClick}
          className={className}
        >
          {text}<div className="mb-0.5">↗</div>
        </div>
      ) : (
        <a
          href={link ?? ''}
          onClick={onClick}
          target="_blank"
          rel="noreferrer"
          className={className}
        >
          {text}<div className="mb-0.5">↗</div>
        </a>
      )}
    </li>
  )
}
