export function HorizontalDivider() {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'div'.
    <div className="divide-y-[1px] divide-white/10">
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div></div>
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div></div>
    </div>
  )
}
