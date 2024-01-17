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
        focus:outline-none
        focus:ring-0
        focus:border-none
        border-none
        flex-grow
        h-full min-w-[70%]
        p-3
        rounded bg-zinc-200 dark:bg-zinc-800 custom-shadow
        font-normal text-sm
        placeholder-black placeholder-opacity-40
      `}
      placeholder={placeholder}
      onChange={(e) => onSearch(e.target.value)}
      value={searchStr}
    />
  )
}
