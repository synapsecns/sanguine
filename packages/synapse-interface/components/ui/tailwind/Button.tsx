import { twMerge } from 'tailwind-merge'

const baseClassname = `
  group
  cursor-pointer
  text-white
  rounded-md
  outline-none focus:outline-none active:outline-none ring-none
  transition-all duration-100 transform-gpu
  `

export default function Button({
  className,
  disabled,
  children,
  ...props
}: {
  children: any
  disabled?: boolean
  className?: string
  style?: any
  onClick?: (e: any | undefined) => void
}) {
  const mergedClassName = twMerge(`${baseClassname} ${className}`)

  return (
    <button className={mergedClassName} {...props} disabled={disabled}>
      {children}
    </button>
  )
}
