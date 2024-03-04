const AugmentWithUnits = ({
  content,
  label,
}: {
  content: string
  label: string
}) => {
  return (
    <div className="text-right">
      {content}{' '}
      <span className="font-thin text-sm text-[#B5B2B9]">{label}</span>
    </div>
  )
}
export default AugmentWithUnits
