export default ({ className, children, ...props }) => (
  <td
    className={`px-6 py-4 whitespace-nowrap text-sm text-gray-500 ${className}`}
    {...props}
  >
    {children}
  </td>
)
