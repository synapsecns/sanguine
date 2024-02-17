
export const HoverContent = ({
  isHovered,
  children,
  className
}: {
  isHovered: boolean
  children: React.ReactNode
  className?: string
}) => {
  if (isHovered) {
    return (
     <div className="absolute !overflow-visible ">
        <div
            className={`
            absolute z-[100] hover-content py-2 px-3 text-white
            border border-white/20 bg-bgBase/10 backdrop-blur-xl
            rounded-md text-left -translate-x-[69%]
            ${className}
            `}
        >
        <div className="pr-6">
            {children}
        </div>

      </div>
     </div>
    )
  }
}

