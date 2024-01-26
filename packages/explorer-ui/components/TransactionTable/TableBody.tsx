import { TableRow } from './TableRow'

export function TableBody({ rows }) {
  return (
    <tbody className="transition duration-150 ease-in ">
      {rows.map((row, index) => (
        <TableRow key={index} items={row.items} />
      ))}
    </tbody>
  )
}
