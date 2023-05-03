import { twMerge } from 'tailwind-merge'

const baseClassname = `
  bg-bgBase
  pt-3 px-6 pb-6 rounded-md
  `

const titleBaseClassname = `
  font-medium text-lg mb-2 text-white
`

export default function Card({
  title,
  className,
  children,
  titleClassName,
  divider = true,
  image = null,
  ...props
}: {
  title?: any
  className?: string
  children: any
  titleClassName?: string
  divider: boolean
  image?: string
}) {
  const mergedClassName = twMerge(`${baseClassname} ${className}`)
  const mergedTitleClassname = twMerge(
    `${titleBaseClassname} ${titleClassName}`
  )

  let titleContent = null
  if (title) {
    titleContent = (
      <>
        <div className={mergedTitleClassname}>{title}</div>
        {divider ? <hr className="hidden" /> : ''}
      </>
    )
  }

  return (
    <div className={mergedClassName} {...props}>
      {image && <div className="flex justify-start mb-2">{image}</div>}
      {titleContent}
      {children}
    </div>
  )
}
