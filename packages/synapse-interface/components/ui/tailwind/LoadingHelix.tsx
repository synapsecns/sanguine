import _ from 'lodash'

export function LoadingHelix({
  className,
  shift = false,
}: {
  className?: string
  shift?: boolean
})  {
  const dotCount = 26
  const dots = Array.from({ length: dotCount }, (_, i) => i + 1)

  return (
    <div className="translate-x-1/2">
     <div className="loader scale-[0.6] left-[90px] w-full ">
       {dots.map(dot => (
         <div key={dot} className="dot"></div>
       ))}
     </div>
    </div>
  )
};


