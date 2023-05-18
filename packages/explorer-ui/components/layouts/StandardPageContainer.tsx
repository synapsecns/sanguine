import React, { ReactNode } from 'react';

interface StandardPageContainerProps {
  title?: string;
  subtitle?: string;
  children: ReactNode;
  rightContent?: ReactNode;
}

export function StandardPageContainer({
                                        title,
                                        subtitle,
                                        children,
                                        rightContent,
                                      }: StandardPageContainerProps) {
  return (
    <main className="relative z-0 flex-1 h-full overflow-y-auto focus:outline-none">
      <div className="items-center px-4 py-8 mx-auto mt-4 2xl:w-5/6 sm:mt-6 sm:px-8 md:px-12 md:pb-14">
        <span
          className={`
            flex items-center
            text-5xl font-medium text-default
            font-bold
            text-white
          `}
        >
          {title}
        </span>
        {rightContent}
        <div className="mt-1 text-sm font-medium text-gray-500 dark:text-gray-600">
          {subtitle ?? ''}
        </div>
        {children}
      </div>
    </main>
  );
}
