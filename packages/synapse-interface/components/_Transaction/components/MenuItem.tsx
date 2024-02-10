import { OptionButton } from "@/components/buttons/OptionButton"
export function MenuItem({
  text,
  link,
  onClick,
}: {
  text: string
  link: string
  onClick?: () => any
}) {
  return (
    <li
      id="menu-item"
      className={`rounded cursor-pointer min-w-[150px]`}
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
