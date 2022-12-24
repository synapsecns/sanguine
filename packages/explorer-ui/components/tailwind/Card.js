import {twMerge} from 'tailwind-merge'

export default function Card({
  title,
  className,
  children,
  titleClassName,
  divider = true,
  ...props
}) {
  let mergedClassName = twMerge(`
    bg-gray-800 shadow-lg pt-3 px-6 pb-6 rounded-lg ${className ?? ''}
  `)
  let titleContent = ''
  if (title) {
    titleContent = (
      <>
        <div
          className={'font-medium text-lg mb-2 text-gray-400 ' + titleClassName}
        >
          {title}
        </div>
        {divider ? <hr className="hidden" /> : ''}
      </>
    )
  }

  return (
    <div className={mergedClassName} {...props}>
      {titleContent}
      {children}
    </div>
  )
}
