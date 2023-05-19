import { twMerge } from 'tailwind-merge';

interface CardProps {
  title?: string;
  className?: string;
  children: React.ReactNode;
  titleClassName?: string;
  divider?: boolean;

  onClick?: any,
}

export default function Card({
                               title,
                               className,
                               children,
                               titleClassName,
                               divider = true,
                               ...props
                             }: CardProps): JSX.Element {
  const mergedClassName = twMerge(`
    bg-gray-800 shadow-lg pt-3 px-6 pb-6 rounded-lg ${className ?? ''}
  `);

  let titleContent: JSX.Element | null = null;
  if (title) {
    titleContent = (
      <>
        <div className={`font-medium text-lg mb-2 text-gray-400 ${titleClassName}`}>
          {title}
        </div>
        {divider ? <hr className="hidden" /> : null}
      </>
    );
  }

  return (
    <div className={mergedClassName} {...props}>
      {titleContent}
      {children}
    </div>
  );
}
