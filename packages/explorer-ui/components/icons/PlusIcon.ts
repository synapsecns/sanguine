export default function PlusIcon({className}) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'svg'.
    <svg
      // @ts-expect-error TS(2349): This expression is not callable.
      className={`w-4 mr-1 ${className}`}
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
        d='M12 6v6m0 0v6m0-6h6m-6 0H6'
      />
    </svg>
  )
}