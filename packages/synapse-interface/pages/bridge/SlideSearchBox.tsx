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
        focus:outline-none focus:ring-0 focus:border-none
        border-none bg-transparent
        w-full h-full p-3 text-sm

      `}
      placeholder={placeholder}
      onChange={(e) => onSearch(e.target.value)}
      value={searchStr}
    />
  )
}
