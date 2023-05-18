export function InfoBlock({ title, logo, content, className = 'mt-0' }) {
  return (
    <div className={`flex flex-col text-center ${className}`}>
      <dd className="self-center text-2xl font-bold text-left text-slate-300">
        {content}
      </dd>
      <dt className="text-gray-500">
        <span className="inline mr-2 align-middle">{logo}</span>
        <span className="text-sm">{title}</span>
      </dt>
    </div>
  )
}
