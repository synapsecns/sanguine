export default function InfoIcon() {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'svg'.
    <svg
      // @ts-expect-error TS(2304): Cannot find name 'className'.
      className="w-4 text-default hover:text-gray-500"
      // @ts-expect-error TS(2304): Cannot find name 'viewBox'.
      viewBox="0 0 20 20"
      // @ts-expect-error TS(2304): Cannot find name 'fill'.
      fill="currentColor"
    >
      // @ts-expect-error TS(2304): Cannot find name 'path'.
      <path
        // @ts-expect-error TS(2304): Cannot find name 'fillRule'.
        fillRule="evenodd"
        // @ts-expect-error TS(2304): Cannot find name 'd'.
        d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
        // @ts-expect-error TS(2304): Cannot find name 'clipRule'.
        clipRule="evenodd"
      />
    </svg>
  )
}