import { TableRow } from './TableRow'

export function TableBody({ rows }) {
  return (
    <tbody className="transition duration-150 ease-in">
      {rows.map((row, index) => (
        <TableRow
          key={index}
          items={row.items}
          className={` hover:border-y border-slate-800 hover:bg-slate-900 hover:cursor-pointer ${
            index === 0 ? 'border-t-0' : ''} ${index === rows.length - 1 ? 'border-b-0' : ''}`
          }
        />
      ))}
    </tbody>
  )
}
