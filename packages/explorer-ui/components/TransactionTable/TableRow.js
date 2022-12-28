export function TableRow({items, key}) {
  return (
    <tr key={key}>
      {items.map((item) =>
        <td className="whitespace-nowrap px-2 py-2 text-sm  text-white">
          {item}
        </td>
        )
      }
    </tr>
  )
}
