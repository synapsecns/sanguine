import { TableRow } from "./TableRow";

export function TableBody({ rows }) {

  return (
    // @ts-expect-error TS(2552): Cannot find name 'tbody'. Did you mean 'Body'?
    <tbody className="transition duration-150 ease-in ">

      {rows.map((row) => (

          // @ts-expect-error TS(2749): 'TableRow' refers to a value, but is being used as... Remove this comment to see the full error message
          <TableRow items={row.items} key={row.key} />
      ))

      }

    </tbody>
  )

}
