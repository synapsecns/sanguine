import Link from 'next/link'
import { BASE_PATH } from '@urls'

export function TopBarNavLink({ labelText, to, className }:{ labelText: string, to: string, className?: string}) {

  const isInternal = to[0] === '/' || to[0] === '#'
  const linkContent = (
    <div className={`py-2 px-2 ${className}`}>
      <span
        className="transform-gpu transition-all duration-75"
      >
        {labelText}
      </span>
    </div>
  )

  const linkClassName = `
    group items-center px-2 my-2 font-normal tracking-wide
    transform-gpu transition-all duration-75
    text-white text-opacity-30
    hover:text-opacity-100
  `

  if (isInternal) {
    return (
      <Link
        href={to}
        className={linkClassName}
        // activeClassName="!text-opacity-100"
      >
        {linkContent}
      </Link>
    )
  } else {
    return (
      <a href={to} target="_blank" className={linkClassName}>
        {linkContent}
      </a>
    )
  }
}
