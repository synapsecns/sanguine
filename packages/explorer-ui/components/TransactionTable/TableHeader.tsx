export const TableHeader = ({ headers }) => {
  return (
    <thead className="">
      <tr>
        {headers.map((header, index) => (
          <th
            key={index}
            scope="col"
            className="px-2 py-2 font-bold text-left text-white text-md "
          >
            {header}
          </th>
        ))}
      </tr>
    </thead>
  )
}
