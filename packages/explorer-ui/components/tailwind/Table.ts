export default function Table({ children, className }) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'table'.
    <table className={`min-w-full divide-y divide-gray-200 ${className} `}>
      // @ts-expect-error TS(18004): No value exists in scope for the shorthand propert... Remove this comment to see the full error message
      {children}
    </table>
  )
}
