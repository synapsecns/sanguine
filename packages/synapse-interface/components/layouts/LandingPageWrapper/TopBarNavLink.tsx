import { NextRouter, useRouter } from 'next/router'
import Link from 'next/link'
import classNames from 'classnames'

export function TopBarNavLink({
  labelText,
  to,
  match,
}: {
  labelText: string
  to: string
  match?: string | RegExp | { startsWith: string }
}) {
  const router = useRouter()

  const isInternal = to[0] === '/' || to[0] === '#'

  const LinkComponent = isInternal ? Link : 'a'

  const isRouteMatched = checkIsRouteMatched(router, match)

  return (
    <LinkComponent
      href={to}
      className={classNames(
        'px-2 tracking-wide transform-gpu transition-all duration-75 text-white',
        {
          'text-opacity-100': isRouteMatched,
          'text-opacity-30': !isRouteMatched,
          'hover:text-opacity-100': true,
        }
      )}
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
  match?: string | RegExp | { startsWith: string }
) => {
  if (!match) return false

  if (typeof match === 'string') {
    return (
      router.asPath === match ||
      router.asPath === '/' ||
      router.asPath.startsWith(match + '?')
    )
  } else if (match instanceof RegExp) {
    return match.test(router.asPath)
  } else if (match.startsWith && typeof match.startsWith === 'string') {
    return router.asPath.startsWith(match.startsWith)
  }

  return false
}
