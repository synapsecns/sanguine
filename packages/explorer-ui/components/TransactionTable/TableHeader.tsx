export function TableHeader({ headers }) {
  return (
    <thead className="">
      <tr>
        {headers.map((header) => (
          <th
            scope="col"
            className="px-2 py-2 text-left text-md font-bold text-white "
          >
            {header}
          </th>
        ))}
      </tr>
    </thead>
  )
}
