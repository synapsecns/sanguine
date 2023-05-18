import { ContainerCard } from '@components/ContainerCard';
import { ExclamationIcon } from '@heroicons/react/outline';

interface ErrorProps {
  text: string;
  param: string | string[];
  subtitle?: string;
}

export const Error = ({ text, param, subtitle }: ErrorProps) => {
  return (
    <ContainerCard
      className="px-10 mt-10"
      icon={<ExclamationIcon className="w-5 h-5 text-red-500" />}
      subtitle={subtitle}
    >
      <div className="mt-3 text-white">
        <div className="mt-5 mb-2 font-extralight">{text}</div>
        <div className="text-base text-gray-400 break-words font font-extralight">
          {param}
        </div>
        <div className="mt-2 font-extralight">
          Please click{' '}
          <a href={'/'} className="text-gray-400 hover:underline">
            here
          </a>{' '}
          to go back to the main page.
        </div>
      </div>
    </ContainerCard>
  );
};
