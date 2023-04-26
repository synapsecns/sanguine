const InfoListItem = ({
  labelText,
  content,
  className = '',
}: {
  labelText: string
  content: any
  className?: string
}) => {
  return (
    <li
      className={`pl-3 pr-4 py-2 text-sm w-full flex border-gray-200 ${className}`}
    >
      <div className="text-white">{labelText} </div>
      <div className="self-center ml-auto text-white">{content}</div>
    </li>
  )
}

export default InfoListItem
