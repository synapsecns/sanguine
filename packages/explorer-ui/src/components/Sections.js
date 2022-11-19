export const PrimaryContainer = ({ children }) => {
  return (
    <div
      className="min-h-screen overflow-hidden bg-gray-900 bg-no-repeat"
      style={{ backgroundImage: "url('/assets/wavylinesstretch.svg')", backgroundPosition: 'top, center' }}
    >
      <div className="p-8 m-auto max-w-7xl">{children}</div>
    </div>
  )
}

export const Card = ({ children }) => {
  return (
    <div className="p-4 bg-[#21283a] rounded-lg shadow-lg border-grey-100 shadow-md hover:shadow-xl">{children}</div>
  )
}

export const GradientCard = ({ children }) => {
  return (
    <div className="p-4 bg-[#21283a] rounded-lg shadow-lg border-grey-100 shadow-md hover:shadow-xl from-[#111827] to-[#21283a]  bg-gradient-to-t">
      {children}
    </div>
  )
}

export const SmallGradientCard = ({ children }) => {
  return (
    <div className="p-4 bg-[#21283a] rounded-lg shadow-lg border-grey-100 shadow-md hover:shadow-xl from-[#111827] to-[#21283a]  bg-gradient-to-t">
      {children}
    </div>
  )
}
