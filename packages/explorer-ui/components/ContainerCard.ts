import Card from '@components/tailwind/Card'

export function ContainerCard({
  title,
  subtitle,
  icon,
  children,
  className,
  subtitleClassName,
  titleClassName,
  ...props
}) {
  return (
    // @ts-expect-error TS(2749): 'Card' refers to a value, but is being used as a t... Remove this comment to see the full error message
    <Card
      // @ts-expect-error TS(2349): This expression is not callable.
      className={`
        text-gray-500
        border border-indigo-500 bg-gray-900
        hover:border-purple-500 ${className}
      `}
      // @ts-expect-error TS(2304): Cannot find name 'props'.
      {...props}
    >
      // @ts-expect-error TS(2749): 'ContainerTitle' refers to a value, but is being u... Remove this comment to see the full error message
      <ContainerTitle
        // @ts-expect-error TS(2304): Cannot find name 'title'.
        title={title}
        // @ts-expect-error TS(2304): Cannot find name 'subtitle'.
        subtitle={subtitle}
        // @ts-expect-error TS(2304): Cannot find name 'subtitleClassName'.
        subtitleClassName={subtitleClassName}
        // @ts-expect-error TS(2304): Cannot find name 'icon'.
        icon={icon}
        // @ts-expect-error TS(2304): Cannot find name 'titleClassName'.
        titleClassName={titleClassName}
      />
      // @ts-expect-error TS(18004): No value exists in scope for the shorthand propert... Remove this comment to see the full error message
      {children}
    </Card>
  )
}

function ContainerTitle({
  icon,
  title,
  subtitle,
  subtitleClassName,
  titleClassName = 'text-transparent bg-clip-text bg-gradient-to-r from-purple-600 to-blue-600',
}) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'div'.
    <div className="flex items-center">
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="mr-2 align-middle">{icon}</div>
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className={`flex-grow font-medium ${titleClassName}`}>{title}</div>
      // @ts-expect-error TS(2304): Cannot find name 'span'.
      <span className={`text-sm text-gray-400 ${subtitleClassName}`}>
        // @ts-expect-error TS(18004): No value exists in scope for the shorthand propert... Remove this comment to see the full error message
        {subtitle}
      </span>
    </div>
  )
}
