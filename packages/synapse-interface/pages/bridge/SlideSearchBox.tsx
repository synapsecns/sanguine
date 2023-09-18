import { useEffect, useRef } from 'react'

export default function SlideSearchBox({
  searchStr,
  onSearch,
  placeholder,
  focusOnMount = true,
}: {
  searchStr: string
  onSearch: (str: string) => void
  placeholder: string
  focusOnMount?: boolean
}) {
  const inputRef = useRef<any>(null)

  useEffect(() => {
    if (focusOnMount) {
      inputRef.current?.focus()
    }
  }, [focusOnMount])

  return (
    <input
      data-test-id="slide-search-box"
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
        rounded bg-[#252226] custom-shadow
        font-normal  text-sm
      placeholder-white placeholder-opacity-40
      `}
      placeholder={placeholder}
      onChange={(e) => onSearch(e.target.value)}
      value={searchStr}
    />
  )
}
