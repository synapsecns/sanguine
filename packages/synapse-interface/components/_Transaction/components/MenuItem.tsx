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
      id="menu-item"
      className={`
        rounded-md cursor-pointer min-w-[150px]
        border border-solid border-transparent
        hover:border-[--synapse-focus] active:opacity-40
      `}
    >
      {onClick ? (
        <OptionButton
          onClick={onClick}
          text={text}
          className={``}
        />

      ) : (
        <a
          href={link ?? ''}
          onClick={onClick}
          target="_blank"
          rel="noreferrer"
          className={`
          `}
        >
          <OptionButton
              onClick={onClick}
              text={text}

            />
        </a>
      )}
    </li>
  )
}

export const OptionButton = ({
  icon,
  text,
  onClick,
}: {
  icon: any
  text: string
  onClick: () => void
}) => {
  return (
    <div
      data-test-id="option-button"
      onClick={onClick}
      className="flex hover:cursor-pointer hover:bg-slate-400/20 rounded-sm p-1 text-white/80 hover:text-white"
    >
      <div className="my-auto mr-1">{icon}</div>
      <div className="text-sm">{text}</div>
    </div>
  )
}
