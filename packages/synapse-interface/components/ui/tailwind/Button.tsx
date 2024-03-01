import { twMerge } from 'tailwind-merge'

const baseClassName = `
  group
  cursor-pointer
  text-white
  rounded-md
  outline-none focus:outline-none active:outline-none ring-none
  transition-all duration-100 transform-gpu
  `
const fancyClassName = `
  w-full rounded-md px-4 py-3
  text-white text-opacity-100 transition-all
  hover:opacity-80
  active:opacity-70
  disabled:opacity-70 disabled:text-white/70
  disabled:!from-bgBase/10 disabled:!to-bgBase/10
  bg-gradient-to-r from-[#CF52FE] to-[#AC8FFF]
  `

export default function Button({
  className,
  disabled,
  children,
  fancy=false,
  ...props
}: {
  children: any
  disabled?: boolean
  className?: string
  fancy?: boolean
  style?: any
  onClick?: (e: any | undefined) => void
}) {
  const mergedClassName = twMerge(`
    ${baseClassName} ${className} ${fancy ? fancyClassName : ''}}
  `)

  return (
    <button className={mergedClassName} {...props} disabled={disabled}>
      {children}
    </button>
  )
}
