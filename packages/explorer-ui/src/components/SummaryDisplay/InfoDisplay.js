import { InfoBlock } from '@components/misc/InfoBlock'

export function InfoDisplay({ arr }) {
  return (
    <div className="justify-center my-5 space-y-5">
      {arr.map(({ title, content, logo }, i) => (
        <div key={i}>
          <InfoBlock title={title} content={content} logo={logo} />
        </div>
      ))}
    </div>
  )
}
