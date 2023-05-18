import { TableRow } from './TableRow'

export function TableBody({ rows }) {
  return (
    <tbody className="transition duration-150 ease-in ">
      {rows.map((row) => (
        <TableRow items={row.items} key={row.key} />
      ))}
    </tbody>
  )
}
