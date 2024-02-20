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
    'flex gap-4 items-center justify-between pl-2 pr-3 py-2 space-x-1.5 whitespace-nowrap text-[--synapse-text-primary] no-underline'

  return (
    <li
      id="menu-item"
      className={`
        rounded cursor-pointer list-none
        border border-solid border-transparent
        hover:border-[--synapse-focus] active:opacity-40
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
