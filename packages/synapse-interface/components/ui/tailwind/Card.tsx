import { twMerge } from 'tailwind-merge'

const baseClassname = "p-4 rounded-md"
const titleBaseClassname = "font-medium text-lg text-zinc-900 dark:text-zinc-300"

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
    <div {...props} className={mergedClassName}>
      {image && <div className="flex justify-start mb-2">{image}</div>}
      {titleContent}
      {children}
    </div>
  )
}
