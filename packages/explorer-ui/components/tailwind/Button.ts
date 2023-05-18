// TEST
const baseClassname = `
  text-white

  px-2 py-2 rounded-md
  transition-all duration-75
  focus:!outline-none active:!outline-none ring-none
  group

  dark:disabled:text-gray-500
  `

const fancyBgClassname = `
  bg-gradient-to-r from-purple-600 to-blue-600
  hover:from-purple-700 hover:to-blue-700
  active:from-purple-800 active:to-blue-800
  disabled:from-gray-300 disabled:to-gray-200

  dark:disabled:from-gray-700 dark:disabled:to-gray-600
  `

const bgClassname = `
  bg-indigo-600 hover:bg-indigo-800 active:bg-indigo-900
  disabled:bg-gray-300
  dark:disabled:bg-gray-700
  `

const outlineClassname = `
  bg-transparent active:bg-gray-50 disabled:bg-gray-300
  border border-gray-200 hover:border-gray-500 active:border-blue-500
  text-gray-600 hover:text-gray-800

  active:bg-gray-800
  dark:border-gray-700 dark:hover:border-purple-500
  dark:text-gray-500 dark:hover:text-gray-400
  `

export default function Button({
  className,
  children,
  fancy,
  outline,
  innerRef,
  ...props
}) {
  let btnStyleClassname
  if (fancy) {
    btnStyleClassname = fancyBgClassname
  } else if (outline) {
    btnStyleClassname = outlineClassname
  } else {
    btnStyleClassname = bgClassname
  }

  return (
    // @ts-expect-error TS(2304): Cannot find name 'button'.
    <button
      // @ts-expect-error TS(2349): This expression is not callable.
      className={`${baseClassname} ${btnStyleClassname} ${className}`}
      // @ts-expect-error TS(2304): Cannot find name 'ref'.
      ref={innerRef}
      // @ts-expect-error TS(2304): Cannot find name 'props'.
      {...props}
    // @ts-expect-error TS(2365): Operator '<' cannot be applied to types 'boolean' ... Remove this comment to see the full error message
    >
      // @ts-expect-error TS(18004): No value exists in scope for the shorthand propert... Remove this comment to see the full error message
      {children}
    </button>
  )
}
