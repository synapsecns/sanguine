export default function StackIcon({ className }) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'svg'.
    <svg
      // @ts-expect-error TS(2349): This expression is not callable.
      className={`h-6 w-6 ${className}`}
      // @ts-expect-error TS(2304): Cannot find name 'fill'.
      fill="none"
      // @ts-expect-error TS(2304): Cannot find name 'viewBox'.
      viewBox="0 0 24 24"
      // @ts-expect-error TS(2304): Cannot find name 'stroke'.
      stroke="currentColor"
      // @ts-expect-error TS(2304): Cannot find name 'aria'.
      aria-hidden="true"
    >
      // @ts-expect-error TS(2304): Cannot find name 'path'.
      <path
        // @ts-expect-error TS(2304): Cannot find name 'strokeLinecap'.
        strokeLinecap='round'
        // @ts-expect-error TS(2304): Cannot find name 'strokeLinejoin'.
        strokeLinejoin='round'
        // @ts-expect-error TS(2304): Cannot find name 'strokeWidth'.
        strokeWidth={2}
        // @ts-expect-error TS(2304): Cannot find name 'd'.
        d='M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z'
      />
    </svg>
  )
}