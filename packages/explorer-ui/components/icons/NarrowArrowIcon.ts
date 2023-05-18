export default function NarrowArrowIcon({ className }) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'svg'.
    <svg
      // @ts-expect-error TS(2349): This expression is not callable.
      className={`w-6 ${className}`}
      // @ts-expect-error TS(2304): Cannot find name 'fill'.
      fill="none"
      // @ts-expect-error TS(2304): Cannot find name 'viewBox'.
      viewBox="0 0 24 24"
      // @ts-expect-error TS(2304): Cannot find name 'stroke'.
      stroke="currentColor"
    >
      // @ts-expect-error TS(2304): Cannot find name 'path'.
      <path
        // @ts-expect-error TS(2304): Cannot find name 'strokeLinecap'.
        strokeLinecap="round"
        // @ts-expect-error TS(2304): Cannot find name 'strokeLinejoin'.
        strokeLinejoin="round"
        // @ts-expect-error TS(2304): Cannot find name 'strokeWidth'.
        strokeWidth={2}
        // @ts-expect-error TS(2304): Cannot find name 'd'.
        d="M17 8l4 4m0 0l-4 4m4-4H3"
      />
    </svg>
  )
}