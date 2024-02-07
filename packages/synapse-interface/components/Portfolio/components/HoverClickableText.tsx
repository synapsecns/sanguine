import { useState } from 'react'

export const HoverClickableText = ({
  defaultText,
  hoverText,
  callback,
}: {
  defaultText: string
  hoverText: string
  callback: () => void
}) => {
  const [isHovered, setIsHovered] = useState<boolean>(false)
  return (
    <div
      id="hover-clickable-text"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
      onClick={callback}
      className={`
        group px-2
        text-[#A3A3C2]
        hover:text-[#75E6F0]
        hover:underline
        hover:cursor-pointer
        active:opacity-70
      `}
    >
      <div className="text-sm">{isHovered ? hoverText : defaultText}</div>
    </div>
  )
}
