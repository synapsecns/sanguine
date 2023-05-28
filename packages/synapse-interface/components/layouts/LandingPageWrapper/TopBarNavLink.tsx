import Link from 'next/link'
import { useRouter } from 'next/router'
import { BASE_PATH } from '@/constants/urls'

export function TopBarNavLink({
  labelText,
  to,
  className,
  match,
}: {
  labelText: string
  to: string
  className?: string
  match?: string
}) {
  const router = useRouter()

  const isInternal = to[0] === '/' || to[0] === '#'
  const linkContent = (
    <div className={`py-2 px-2 ${className}`}>
      <span className="transition-all duration-75 transform-gpu">
        {labelText}
      </span>
    </div>
  )

  const linkClassName = `
    group items-center px-2 my-2 font-normal tracking-wide
    transform-gpu transition-all duration-75
    text-white ${
      match && router.asPath.includes(match)
        ? 'text-opacity-100'
        : 'text-opacity-30'
    }
    hover:text-opacity-100
  `

  if (isInternal) {
    return (
      <Link
        href={to}
        className={linkClassName}
        // activeClassName="!text-opacity-100"
        data-test-id="nav-link"
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
