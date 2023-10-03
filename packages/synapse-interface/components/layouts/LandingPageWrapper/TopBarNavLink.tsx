import { useRouter } from 'next/router'
import Link from 'next/link'

export function TopBarNavLink({
  labelText,
  to,
  className,
  match,
}: {
  labelText: string
  to: string
  className?: string
  match?: string | RegExp | { startsWith: string; endsWith: string }
}) {
  const router = useRouter()

  const isInternal = to[0] === '/' || to[0] === '#'
  const linkContent = (
    <div className={`py-2 px-2 ${className}`}>
      <span className="transition-all duration-75 transform-gpu whitespace-nowrap">
        {labelText}
      </span>
    </div>
  )

  const linkClassName = `
    group items-center px-2 my-2 font-normal tracking-wide
    transform-gpu transition-all duration-75
    text-white ${
      match &&
      (typeof match === 'string'
        ? router.asPath.includes(match)
        : match instanceof RegExp
        ? match.test(router.asPath)
        : router.asPath.startsWith(match.startsWith) &&
          router.asPath.endsWith(match.endsWith))
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
