export default function ChartPieIcon({ className }) {
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
        strokeLinecap="round"
        // @ts-expect-error TS(2304): Cannot find name 'strokeLinejoin'.
        strokeLinejoin="round"
        // @ts-expect-error TS(2304): Cannot find name 'strokeWidth'.
        strokeWidth={2}
        // @ts-expect-error TS(2304): Cannot find name 'd'.
        d="M11 3.055A9.001 9.001 0 1020.945 13H11V3.055z"
      />
      // @ts-expect-error TS(2304): Cannot find name 'path'.
      <path
        // @ts-expect-error TS(2304): Cannot find name 'strokeLinecap'.
        strokeLinecap="round"
        // @ts-expect-error TS(2304): Cannot find name 'strokeLinejoin'.
        strokeLinejoin="round"
        // @ts-expect-error TS(2304): Cannot find name 'strokeWidth'.
        strokeWidth={2}
        // @ts-expect-error TS(2304): Cannot find name 'd'.
        d="M20.488 9H15V3.512A9.025 9.025 0 0120.488 9z"
      />
    </svg>
  )
}

