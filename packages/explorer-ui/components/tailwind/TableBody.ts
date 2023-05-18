export default function TableBody({ children }) {
  // @ts-expect-error TS(2552): Cannot find name 'tbody'. Did you mean 'Body'?
  return <tbody className="divide-y divide-gray-200">{children}</tbody>
}
