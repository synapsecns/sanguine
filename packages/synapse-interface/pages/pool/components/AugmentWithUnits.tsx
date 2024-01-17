const AugmentWithUnits = ({
  content,
  label,
}: {
  content: string
  label: string
}) => {
  return (
    <div className="text-right">
      {content} {label}
    </div>
  )
}
export default AugmentWithUnits
