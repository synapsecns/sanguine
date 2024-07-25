import { TableRow } from './TableRow'

export const TableBody = ({ rows }) => {
  return (
    <tbody className="transition duration-150 ease-in">
      {rows.map((row, index) => (
        <TableRow
          key={index}
          items={row.items}
          className={`hover:bg-[#0F172A] hover:cursor-pointer ${
            index % 2 === 0 ? 'bg-[#101018]' : 'bg-[#100C13]'
          } ${index === 0 ? 'border-t-0' : ''} ${
            index === rows.length - 1 ? 'border-b-0' : ''
          }`}
        />
      ))}
    </tbody>
  )
}
