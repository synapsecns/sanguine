import { useEffect, useRef } from 'react'

export default function SlideSearchBox({
  searchStr,
  onSearch,
  placeholder,
}: {
  searchStr: string
  onSearch: (str: string) => void
  placeholder: string
}) {
  const inputRef = useRef<any>(null)
  useEffect(() => inputRef.current?.focus(), [])

  return (
    <input
      ref={inputRef}
      className={`
        text-primaryTextColor
        focus:outline-none
        focus:ring-0
        focus:border-none
        border-none
        flex-grow
        h-full min-w-[70%]
        py-2 p-2
        rounded-md bg-slate-900/90 custom-shadow
        font-normal  text-sm
      placeholder-white placeholder-opacity-40
      `}
      placeholder={placeholder}
      onChange={(e) => onSearch(e.target.value)}
      value={searchStr}
    />
  )
}
