export default function LoadingSpinner({ className }) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'svg'.
    <svg
      // @ts-expect-error TS(2349): This expression is not callable.
      className={`inline-flex animate-spin h-5 w-5 text-indigo-600 ${className}`}
      // @ts-expect-error TS(2304): Cannot find name 'fill'.
      fill="none"
      // @ts-expect-error TS(2304): Cannot find name 'viewBox'.
      viewBox="0 0 24 24"
    >
      // @ts-expect-error TS(2304): Cannot find name 'circle'.
      <circle
        // @ts-expect-error TS(2304): Cannot find name 'className'.
        className="opacity-25"
        // @ts-expect-error TS(2304): Cannot find name 'cx'.
        cx="12"
        // @ts-expect-error TS(2304): Cannot find name 'cy'.
        cy="12"
        // @ts-expect-error TS(2304): Cannot find name 'r'.
        r="10"
        // @ts-expect-error TS(2304): Cannot find name 'stroke'.
        stroke="currentColor"
        // @ts-expect-error TS(2304): Cannot find name 'strokeWidth'.
        strokeWidth="4"
      />
      // @ts-expect-error TS(2304): Cannot find name 'path'.
      <path
        // @ts-expect-error TS(2304): Cannot find name 'className'.
        className="opacity-75"
        // @ts-expect-error TS(2304): Cannot find name 'fill'.
        fill="currentColor"
        // @ts-expect-error TS(2304): Cannot find name 'd'.
        d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
      />
    </svg>
  )
}
