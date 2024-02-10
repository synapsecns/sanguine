import MiniMaxButton from './buttons/MiniMaxButton'

const InteractiveInputRow = ({
  title,
  isConnected,
  balanceStr,
  onClickBalance,
  value,
  placeholder,
  onChange,
  disabled,
  icon,
}: {
  title: string
  isConnected: boolean
  balanceStr: string
  onClickBalance: (e) => void
  value: string
  placeholder: string
  onChange: (e) => void
  disabled: boolean
  icon: string
}) => {
  return (
    <div className="flex flex-col rounded-sm bg-bgBase/10 ring-1 ring-white/10">
      <div className="border-none rounded-md">
        <div className="flex space-x-2">
          <div className="flex items-center flex-grow w-full h-16 pl-3 ">
            <div className="sm:mt-[-1px]">
              <div
                className={`
                group rounded-sm
                bg-bgBase/10 ring-1 ring-white/10
              `}
              >
                <div
                  className={`flex justify-center md:justify-start items-center rounded-md py-1.5 px-2 w-36`}
                >
                  <div className="self-center flex-shrink-0 hidden mr-2 sm:block">
                    <div
                      className={`
                      relative flex p-1 rounded-full
                    `}
                    >
                      <img className="w-8 h-8 " src={icon} />
                    </div>
                  </div>
                  <div className="text-left">
                    <h4 className="text-lg text-white">
                      <span>{title}</span>
                    </h4>
                  </div>
                </div>
              </div>
            </div>
            <div
              className={`
                flex flex-grow items-center
                mx-3 w-full h-16 -mt-3
                border-none
                relative overflow-hidden
              `}
            >
              <input
                autoComplete="off"
                className={`
                    ${isConnected ? '-mt-2' : '-mt-0'}
                    focus:outline-none
                    focus:ring-0
                    focus:border-none
                    border-none bg-transparent
                    p-0
                    w-[300px] sm:min-w-[170px] sm:w-full scrollbar-none
                  placeholder:text-[#88818C] text-white
                    text-opacity-80 text-lg md:text-2xl lg:text-2xl font-medium
                    overflow-hidden
                `}
                value={value}
                placeholder={placeholder}
                onChange={(e) => {
                  onChange(e)
                }}
                name="inputRow"
              />
              <div>
                {isConnected && (
                  <div className="hidden md:block">
                    <MiniMaxButton
                      disabled={disabled}
                      onClickBalance={disabled ? undefined : onClickBalance}
                    />
                  </div>
                )}
              </div>
              {isConnected && (
                <label
                  htmlFor="inputRow"
                  className="absolute bottom-0 text-sm text-secondaryTextColor hover:text-opacity-70 hover:cursor-pointer"
                  onClick={onClickBalance}
                >
                  {balanceStr}
                  <span className=" text-secondaryTextColor"> available</span>
                </label>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
export default InteractiveInputRow
