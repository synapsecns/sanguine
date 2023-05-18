export default function TableHeaderCell({ children }) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'th'.
    <th
      // @ts-expect-error TS(2304): Cannot find name 'scope'.
      scope="col"
      // @ts-expect-error TS(2304): Cannot find name 'className'.
      className="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase"
    >
      {children}
    </th>
  )
}
