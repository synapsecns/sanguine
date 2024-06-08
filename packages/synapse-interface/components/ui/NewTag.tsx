import { joinClassNames } from '@/utils/joinClassNames'

export const NewTag = () => {
  const className = joinClassNames({
    space: 'px-2 py-[2px] rounded-md',
    border: 'border border-fuchsia-500',
    background: 'bg-gradient-to-r from-fuchsia-950 to-purple-900',
    font: 'text-sm ',
  })
  return <div className={className}>New!</div>
}
