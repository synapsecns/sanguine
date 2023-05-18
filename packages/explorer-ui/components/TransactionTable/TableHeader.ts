export function TableHeader({headers}) {
  return(
    // @ts-expect-error TS(2304): Cannot find name 'thead'.
    <thead className="">
      // @ts-expect-error TS(2304): Cannot find name 'tr'.
      <tr>
        {headers.map((header) => (
          // @ts-expect-error TS(2304): Cannot find name 'th'.
          <th
            // @ts-expect-error TS(2304): Cannot find name 'scope'.
            scope="col"
            // @ts-expect-error TS(2304): Cannot find name 'className'.
            className="px-2 py-2 text-left text-md font-bold text-white "
          >
            // @ts-expect-error TS(2552): Cannot find name 'header'. Did you mean 'headers'?
            {header}
          </th>
        ))}
      </tr>
    </thead>
  )
}
