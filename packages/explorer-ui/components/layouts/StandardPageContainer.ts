export function StandardPageContainer({
  title,
  subtitle,
  children,
  rightContent,
}) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'main'.
    <main className="relative z-0 flex-1 h-full overflow-y-auto focus:outline-none">
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="items-center px-4 py-8 mx-auto mt-4 2xl:w-5/6 sm:mt-6 sm:px-8 md:px-12 md:pb-14">
        // @ts-expect-error TS(2304): Cannot find name 'span'.
        <span
          // @ts-expect-error TS(2304): Cannot find name 'className'.
          className={`
            flex items-center
            text-5xl font-medium text-default
            font-bold
            text-white
          `}
        // @ts-expect-error TS(2365): Operator '<' cannot be applied to types 'boolean' ... Remove this comment to see the full error message
        >
          {title}
        </span>
        {rightContent}
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="mt-1 text-sm font-medium text-gray-500 dark:text-gray-600">
          {subtitle ?? ''}
        </div>
        // @ts-expect-error TS(2304): Cannot find name 'children'.
        {children}
      </div>
    </main>
  )
}
