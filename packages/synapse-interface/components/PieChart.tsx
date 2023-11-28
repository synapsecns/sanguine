export const PieChart = ({
  activeAmount,
  totalAmount,
}: {
  activeAmount: number
  totalAmount: number
}) => {
  const calculatedActiveDegrees: number = (activeAmount / totalAmount) * 360
  const calculatedInactiveDegrees: number = 360 - calculatedActiveDegrees

  return (
    <div
      data-test-id="pie-chart"
      className="w-4 h-4 rounded-[50%] m-auto"
      style={{
        backgroundImage: `
          conic-gradient(#99E6FF 0deg, #99E6FF ${calculatedActiveDegrees}deg, #343036 ${calculatedActiveDegrees}deg, #343036 ${calculatedInactiveDegrees}deg)`,
      }}
    ></div>
  )
}
