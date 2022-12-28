import { TableRow }  from "./TableRow";

export function TableBody({rows}) {
  return (
    <tbody>
      {rows.map((row) => (
          <TableRow items={row.items} key={row.key}/>
        ))
      }
    </tbody>
  )

}
