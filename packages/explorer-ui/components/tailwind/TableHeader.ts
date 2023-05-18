export default function TableHeader({ children }) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'thead'.
    <thead className="bg-gray-50">
      // @ts-expect-error TS(2304): Cannot find name 'tr'.
      <tr>{children}</tr>
    </thead>
  )
}
