import { OptionButton } from "@/components/buttons/OptionButton"
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
