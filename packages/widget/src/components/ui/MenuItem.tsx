export const MenuItem = ({
  text,
  link,
  onClick,
}: {
  text: string
  link: string
  onClick?: () => any
}) => {
  return (
    <li
      className={`
      rounded cursor-pointer
      border border-solid border-transparent
      hover:border-[--synapse-focus]
      active:opacity-40
    `}
    >
      {onClick ? (
        <div
          onClick={onClick}
          className={`
            block pl-2 pr-3 py-2 whitespace-nowrap text-[--synapse-text-primary] no-underline after:content-['_↗'] after:text-xs after:text-[--synapse-secondary]
          `}
        >
          {text}
        </div>
      ) : (
        <a
          href={link ?? ''}
          onClick={onClick}
          target="_blank"
          rel="noreferrer"
          className={`
            block pl-2 pr-3 py-2 whitespace-nowrap text-[--synapse-text-primary] no-underline after:content-['_↗'] after:text-xs after:text-[--synapse-secondary]
          `}
        >
          {text}
        </a>
      )}
    </li>
  )
}
