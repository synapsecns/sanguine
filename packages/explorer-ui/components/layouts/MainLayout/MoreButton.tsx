import Button from '@components/tailwind/Button';
import { DotsHorizontalIcon } from '@heroicons/react/outline';
import React from 'react';


interface MoreButtonProps {
  open: boolean;
  onClick?: () => void;
  className?: string;
}

const MoreButton: React.FC<MoreButtonProps> = ({
                                                 open,
                                                 onClick,
                                                 className,
                                                 ...props
                                               }) => {
  return (
    <Button
      onClick={onClick}
      className={`
        w-full cursor-pointer rounded-lg px-4 py-4 group border-none dark:hover:bg-[#111111] ${className} text-sm
        ${open && ' bg-[#111111]'}
      `}
      outline={true}
      {...props}
    >
      <div className="space-x-2">
        <div className="inline-block rounded-md">
          <DotsHorizontalIcon
            className={`
              ${open && 'opacity-100'}
              inline-block w-4 h-4 text-white dark:text-white group-hover:opacity-100
            `}
          />
        </div>
      </div>
    </Button>
  );
};

export default MoreButton;
