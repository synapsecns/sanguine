export default ({ children, className }) => (
  <table className={`min-w-full divide-y divide-gray-200 ${className} `}>
    {children}
  </table>
)
