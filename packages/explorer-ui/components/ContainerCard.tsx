import Card from '@components/tailwind/Card';

interface ContainerCardProps {
  title?: any;
  subtitle?: any;
  icon: any;
  children: any;
  className?: string;
  subtitleClassName?: string;
  titleClassName?: string;
}

export function ContainerCard({
                                title,
                                subtitle,
                                icon,
                                children,
                                className,
                                subtitleClassName,
                                titleClassName,
                                ...props
                              }: ContainerCardProps) {
  return (
    <Card
      title={title}
      className={`
        text-gray-500
        border border-indigo-500 bg-gray-900
        hover:border-purple-500 ${className}
      `}
      titleClassName={titleClassName}
      {...props}
    >
      <ContainerTitle
        title={title}
        subtitle={subtitle}
        subtitleClassName={subtitleClassName}
        icon={icon}
        titleClassName={titleClassName}
      />
      {children}
    </Card>
  );
}

interface ContainerTitleProps {
  icon: any;
  title: any;
  subtitle: any;
  subtitleClassName: string;
  titleClassName?: string;
}

function ContainerTitle({
                          icon,
                          title,
                          subtitle,
                          subtitleClassName,
                          titleClassName = 'text-transparent bg-clip-text bg-gradient-to-r from-purple-600 to-blue-600',
                        }: ContainerTitleProps) {
  return (
    <div className="flex items-center">
      <div className="mr-2 align-middle">{icon}</div>
      <div className={`flex-grow font-medium ${titleClassName}`}>{title}</div>
      <span className={`text-sm text-gray-400 ${subtitleClassName}`}>
        {subtitle}
      </span>
    </div>
  );
}
