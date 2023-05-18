export default function ForumIcon({ className }) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'svg'.
    <svg
      // @ts-expect-error TS(2349): This expression is not callable.
      className={`w-4 ${className}`}
      // @ts-expect-error TS(2304): Cannot find name 'fill'.
      fill='none'
      // @ts-expect-error TS(2304): Cannot find name 'viewBox'.
      viewBox='0 0 24 24'
      // @ts-expect-error TS(2304): Cannot find name 'stroke'.
      stroke='currentColor'
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
        d='M17 8h2a2 2 0 012 2v6a2 2 0 01-2 2h-2v4l-4-4H9a1.994 1.994 0 01-1.414-.586m0 0L11 14h4a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2v4l.586-.586z'
      />
    </svg>
  )
}