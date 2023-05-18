export function InfoBlock({ title, logo, content, className = 'mt-0' }) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'div'.
    <div className={`flex flex-col text-center ${className}`}>
      // @ts-expect-error TS(2304): Cannot find name 'dd'.
      <dd className="self-center text-2xl font-bold text-left text-slate-300">
        // @ts-expect-error TS(18004): No value exists in scope for the shorthand propert... Remove this comment to see the full error message
        {content}
      </dd>
      // @ts-expect-error TS(2304): Cannot find name 'dt'.
      <dt className="text-gray-500">
        // @ts-expect-error TS(2304): Cannot find name 'span'.
        <span className="inline mr-2 align-middle">{logo}</span>
        // @ts-expect-error TS(2304): Cannot find name 'span'.
        <span className="text-sm">{title}</span>
      </dt>
    </div>
  )
}
