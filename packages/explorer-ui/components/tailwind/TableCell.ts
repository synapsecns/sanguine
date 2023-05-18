export default function TableCell({ className, children, ...props }) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'td'.
    <td
      // @ts-expect-error TS(2349): This expression is not callable.
      className={`px-6 py-4 whitespace-nowrap text-sm text-gray-500 ${className}`}
      // @ts-expect-error TS(2304): Cannot find name 'props'.
      {...props}
    // @ts-expect-error TS(2365): Operator '<' cannot be applied to types 'boolean' ... Remove this comment to see the full error message
    >
      // @ts-expect-error TS(18004): No value exists in scope for the shorthand propert... Remove this comment to see the full error message
      {children}
    </td>
  )
}
