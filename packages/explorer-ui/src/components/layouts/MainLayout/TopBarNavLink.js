import { NavLink, useLocation } from 'react-router-dom'

import { ExternalLinkIcon } from '@heroicons/react/outline'

import { BASE_PATH } from '@urls'

export default function TopBarNavLink({ labelText, to, className }) {
  const location = useLocation()

  const match =
    location.pathname.split('/')[1] === to.split('/')[1] && to !== '#'
  const isInternal = to[0] === '/' || to[0] === '#'
  const linkContent = (
    <div className={`py-2 px-2 ${className}`}>
      <span
        className={`
            transform-gpu transition-all duration-75
            ${
              !match &&
              `
              `
            }
          `}
      >
        {labelText}
        {/* {!isInternal && <ExternalLinkIcon className="h-4 w-4 inline ml-2 -mt-0.5" />} */}
      </span>
    </div>
  )

  const linkClassName = `
    group items-center px-2 my-2 font-light tracking-wide rounded-md
    bg-opacity-50
    transform-gpu transition-all duration-75
    dark:text-white
    dark:text-opacity-30
    dark:hover:text-gray-300
    dark:hover:bg-gray-800
  `

  if (isInternal) {
    return (
      <NavLink
        exact={BASE_PATH === to && to !== '#'}
        to={to}
        className={linkClassName}
        activeclassname={`
          !font-medium
          !text-opacity-100
        `}
        activestyle={{ textShadow: '0px 0px #00000000' }}
      >
        {linkContent}
      </NavLink>
    )
  } else {
    return (
      <a href={to} target="_blank" className={linkClassName}>
        {linkContent}
      </a>
    )
  }
}
