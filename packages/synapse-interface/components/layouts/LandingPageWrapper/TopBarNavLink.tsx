import { NextRouter, useRouter } from 'next/router'
import Link from 'next/link'

export function TopBarNavLink({
  labelText,
  to,
  match,
}: {
  labelText: string
  to: string
  match?: string | { startsWith: string }
}) {
  const router = useRouter()

  const isInternal = to[0] === '/' || to[0] === '#'

  const LinkComponent = isInternal ? Link : 'a'

  const isRouteMatched = checkIsRouteMatched(router, match)

  return (
    <LinkComponent
      href={to}
      className={`
        px-2 tracking-wide transform-gpu transition-all duration-75 text-white hover:text-opacity-100
        ${isRouteMatched ? 'text-opacity-100' : 'text-opacity-30'}
      `}
      target={!isInternal ? '_blank' : undefined}
      rel={!isInternal ? 'noopener noreferrer' : undefined}
      data-test-id="nav-link"
    >
      <LinkContent labelText={labelText} />
    </LinkComponent>
  )
}

const LinkContent = ({ labelText }: { labelText: string }) => {
  return (
    <div className={`py-2 px-2`}>
      <span className="transition-all duration-75 transform-gpu whitespace-nowrap">
        {labelText}
      </span>
    </div>
  )
}

export const checkIsRouteMatched = (
  router: NextRouter,
  match?: string | { startsWith: string }
) => {
  if (!match) return false

  if (router.asPath === '/') {
    return match === '/'
  } else if (typeof match === 'string') {
    return (
      router.asPath === match ||
      router.asPath === '/' ||
      router.asPath.startsWith(match + '?')
    )
  } else if (match.startsWith && typeof match.startsWith === 'string') {
    return router.asPath.startsWith(match.startsWith)
  }

  return false
}
