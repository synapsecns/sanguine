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

  const linkClassName = `
    px-4 py-2 ${
      match &&
      (typeof match === 'string'
        ? router.asPath.includes(match)
        : match instanceof RegExp
        ? match.test(router.asPath)
        : router.asPath.startsWith(match.startsWith) &&
          router.asPath.endsWith(match.endsWith))
        ? 'opacity-100 cursor-default'
        : 'opacity-40'
    }
    hover:opacity-100
  `

  return isInternal
  ? (
    <Link
      href={to}
      className={linkClassName}
      data-test-id="nav-link"
    >
      {labelText}
    </Link>
  ) : (
    <a href={to} target="_blank" className={linkClassName}>
      {labelText}
    </a>
  )
}
