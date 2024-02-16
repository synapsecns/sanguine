import { useRouter } from 'next/router'
import Link from 'next/link'
import classNames from 'classnames'

export function TopBarNavLink({
  labelText,
  to,
  className,
  match,
}: {
  labelText: string
  to: string
  className?: string
  match?: string | RegExp | { startsWith: string }
}) {
  const router = useRouter()

  const isInternal = to[0] === '/' || to[0] === '#'

  const isRouteMatched =
    match &&
    (typeof match === 'string'
      ? router.asPath === match ||
        router.asPath === '/' ||
        router.asPath.startsWith(match + '?')
      : match instanceof RegExp
      ? match.test(router.asPath)
      : router.asPath.startsWith(match.startsWith))

  const linkClassName = classNames(
    'px-2 tracking-wide transform-gpu transition-all duration-75 text-white',
    {
      'text-opacity-100': isRouteMatched,
      'text-opacity-30': !isRouteMatched,
      'hover:text-opacity-100': true,
    },
    className
  )

  const LinkComponent = isInternal ? Link : 'a'

  return (
    <LinkComponent
      href={to}
      className={linkClassName}
      target={!isInternal ? '_blank' : undefined}
      rel={!isInternal ? 'noopener noreferrer' : undefined}
      data-test-id="nav-link"
    >
      <LinkContent className={className} labelText={labelText} />
    </LinkComponent>
  )
}

const LinkContent = ({
  className,
  labelText,
}: {
  className: string
  labelText: string
}) => {
  return (
    <div className={`py-2 px-2 ${className}`}>
      <span className="transition-all duration-75 transform-gpu whitespace-nowrap">
        {labelText}
      </span>
    </div>
  )
}
