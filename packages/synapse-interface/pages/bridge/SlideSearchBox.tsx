import { joinClassNames } from '@/utils/joinClassNames'
import { useEffect, useRef } from 'react'

const className = joinClassNames({
  text: 'text-sm font-normal',
  placeholder: 'placeholder-white/40',
  focus: 'focus:ring-0',
  border: 'border-none',
  flex: 'flex-grow',
  space: 'px-2 py-1.5',
  background: 'bg-[#252226]',
  // shadow: 'custom-shadow',
})

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
      className={className}
      placeholder={placeholder}
      onChange={(e) => onSearch(e.target.value)}
      value={searchStr}
    />
  )
}
