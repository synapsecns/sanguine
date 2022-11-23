export default function Table({ children, className }) {
  return (
    <table className={`min-w-full divide-y divide-gray-200 ${className} `}>
      {children}
    </table>
  )
}
