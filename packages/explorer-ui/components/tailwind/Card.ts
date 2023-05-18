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
    // @ts-expect-error TS(2322): Type 'boolean' is not assignable to type 'string'.
    titleContent = (
      <>
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div
          // @ts-expect-error TS(2365): Operator '>' cannot be applied to types '{ 'font-m... Remove this comment to see the full error message
          className={'font-medium text-lg mb-2 text-gray-400 ' + titleClassName}
        >
          {title}
        </div>
        // @ts-expect-error TS(2304): Cannot find name 'hr'.
        {divider ? <hr className="hidden" /> : ''}
      </>
    )
  }

  return (
    // @ts-expect-error TS(2304): Cannot find name 'div'.
    <div className={mergedClassName} {...props}>
      {titleContent}
      {children}
    </div>
  )
}
