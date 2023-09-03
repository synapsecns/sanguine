export const DeadlineInput = ({
  deadlineMinutes,
  setDeadlineMinutes,
}: {
  deadlineMinutes: string
  setDeadlineMinutes: (value: string) => void
}) => {
  return (
    <div className="flex h-16 pb-4 space-x-2 text-left">
      <div
        className={`
          flex flex-grow items-center
          h-14 w-full
        bg-bgLight
          border border-transparent
          hover:border-gradient-br-magenta-melrose-bgLight hover:border-solid
          focus-within:border-gradient-br-magenta-melrose-bgLight focus-within:border-solid
          pl-1
          py-0.5 rounded-md
        `}
      >
        <input
          pattern="[0-9.]+"
          className={`
              ml-4 mr-4
              focus:outline-none
              bg-transparent
              w-[300px]
              sm:min-w-[300px]
              max-w-[calc(100%-92px)]
              sm:w-full
              text-lg
             placeholder-[#716e74]
             text-white text-opacity-80
            `}
          placeholder="Custom deadline..."
          onChange={(e) => setDeadlineMinutes(e.target.value)}
          value={deadlineMinutes}
        />
        <span className="hidden text-lg text-white md:block opacity-30">
          mins
        </span>
      </div>
    </div>
  )
}
