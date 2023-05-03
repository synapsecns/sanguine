const AugmentWithUnits = ({
  content,
  label,
}: {
  content: string
  label: string
}) => {
  return (
    <div className="text-right">
      {content} <span className="text-gray-500">{label}</span>
    </div>
  )
}
export default AugmentWithUnits
