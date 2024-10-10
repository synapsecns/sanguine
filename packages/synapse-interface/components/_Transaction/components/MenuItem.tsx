import Image from 'next/image'

export const MenuItem = ({
  text,
  link,
  iconUrl,
  onClick,
}: {
  text: string
  link: string
  iconUrl?: string
  onClick?: () => any
}) => {
  const className =
    'flex items-center justify-between gap-4 py-2 pl-2 pr-3 space-x-4 no-underline whitespace-nowrap'

  return (
    <li
      id="menu-item"
      className={`
        rounded cursor-pointer list-none
        border border-solid border-transparent
         active:opacity-40
      `}
    >
      {onClick ? (
        <div onClick={onClick} className={className}>
          {text}
          <div className="mb-0.5">↗</div>
        </div>
      ) : (
        <a
          href={link ?? ''}
          onClick={onClick}
          target="_blank"
          rel="noreferrer"
          className={className}
        >
          <div className="flex space-x-1.5 items-center">
            {iconUrl && (
              <Image
                src={iconUrl}
                height={14}
                width={14}
                alt={`${text} icon`}
                className="mt-px"
              />
            )}
            <div>{text}</div>
          </div>
          <div className="mb-0.5">↗</div>
        </a>
      )}
    </li>
  )
}
