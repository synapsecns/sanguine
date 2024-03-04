export const OptionButton = ({
  icon,
  text,
  onClick,
}: {
  icon?: any
  text: string
  onClick: () => void
}) => {
  return (
    <div
      data-test-id="option-button"
      onClick={onClick}
      className="flex hover:cursor-pointer hover:bg-slate-400/20 rounded-sm p-1 text-white/80 hover:text-white"
    >
      <div className="my-auto mr-1">{icon}</div>
      <div className="text-sm">{text}</div>
    </div>
  )
}
