import { useEffect, useRef } from 'react'

import { joinClassNames } from '@/utils/joinClassNames'

const isMobileDevice = () => {
  return window.innerWidth <= 768
}

const classNames = {
  text: 'text-sm font-normal',
  placeholder: 'placeholder-white/40',
  focus: 'focus:ring-0',
  border: 'border-none',
  flex: 'flex-grow',
  space: 'px-2 py-1.5',
  background: 'bg-[#252226]',
}

export function SlideSearchBox({
  searchStr,
  onSearch,
  placeholder,
}: {
  searchStr: string
  onSearch: (str: string) => void
  placeholder: string
}) {
  const inputRef = useRef(null)

  useEffect(() => {
    if (!isMobileDevice()) {
      inputRef.current?.focus()
    }
  }, [])

  return (
    <input
      ref={inputRef}
      className={joinClassNames(classNames)}
      placeholder={placeholder}
      onChange={(e) => onSearch(e.target.value)}
      value={searchStr}
    />
  )
}
