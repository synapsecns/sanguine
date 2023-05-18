import { QuestionMarkCircleIcon } from '@heroicons/react/outline';
import { getChainUrl, getExplorerTxUrl } from '@urls';
import { CHAIN_EXPLORER_URLS, CHAIN_INFO_MAP } from '@constants/networks';
import Image from 'next/image';

interface ChainInfoProps {
  chainId?: any;
  imgClassName?: string;
  linkClassName?: string;
  textClassName?: string;
  txHash?: string;
  useExplorerLink?: boolean;
  noLink?: boolean;
}

export function ChainInfo({
                            chainId,
                            imgClassName = 'w-4 h-4 rounded-full',
                            linkClassName = 'float-right text-white transition ease-out hover:text-[#8FEBFF] px-1.5  rounded-md ease-in-out bg-[#191919]',
                            textClassName = 'pl-1 whitespace-nowrap text-sm text-white',
                            txHash,
                            useExplorerLink = false,
                            noLink = false,
                          }: ChainInfoProps) {
  const { chainName, chainImg } = CHAIN_INFO_MAP[chainId] ?? {};
  let link = '';

  if (txHash) {
    link = getExplorerTxUrl({ hash: txHash, chainId });
  }

  if (useExplorerLink) {
    link = getChainUrl({ chainId });
  }

  if (chainName) {
    return (
      <div className="w-full relative">
        <div className="flex justify-between ">
          <div className="flex flex-row w-[90%] items-center">
            <Image
              className={`inline mr-[.5rem] rounded-full ${imgClassName}`}
              src={chainImg}
              alt={chainName}
            />
            <p className={textClassName}>{chainName}</p>
          </div>
          {noLink ? null : (
            <div className="flex items-center">
              <a
                type="link"
                target="_blank"
                href={link}
                className={linkClassName}
              >
                ↗
              </a>
            </div>
          )}
        </div>
      </div>
    );
  } else {
    return (
      <div className="flex items-center">
        <QuestionMarkCircleIcon
          className={`inline mr-2 rounded-lg ${imgClassName}`}
          strokeWidth={1}
        />
        <span>--</span>
      </div>
    );
  }
}
