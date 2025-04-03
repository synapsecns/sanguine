import React, { useState, useEffect } from 'react';
import ExternalLinkIcon from './icons/ExternalLinkIcon';
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider';

const STORAGE_KEY = 'cortex-promo-dismissed';

export const CortexPromoPopup: React.FC = () => {
  // Initial state based on client-side check to avoid hydration mismatch
  const [isDismissed, setIsDismissed] = useState(true);
  const [isInitialized, setIsInitialized] = useState(false);
  
  useEffect(() => {
    // Check if the popup has been dismissed before
    const isDismissedInStorage = localStorage.getItem(STORAGE_KEY) === 'true';
    setIsDismissed(isDismissedInStorage);
    setIsInitialized(true);
  }, []);
  
  const handleDismiss = () => {
    localStorage.setItem(STORAGE_KEY, 'true');
    setIsDismissed(true);
    
    // Track dismiss event
    segmentAnalyticsEvent('cortex_promo_dismissed', {
      action: 'dismiss',
      location: 'popup',
    });
  };

  // Only show when initialized and not dismissed
  if (!isInitialized || isDismissed) return null;

  return (
    <div className="fixed bottom-5 right-5 z-50 max-w-[280px] rounded-md bg-bgBase border border-[#3B363D] overflow-hidden animate-fadeIn">
      <div className="p-3">
        <button 
          className="absolute top-3 right-3 text-secondaryTextColor hover:text-primaryTextColor transition-colors duration-200"
          onClick={handleDismiss}
          aria-label="Dismiss"
        >
          <svg width="10" height="10" viewBox="0 0 12 12" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M1 1L11 11M1 11L11 1" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
          </svg>
        </button>
        
        <div className="flex items-center gap-2 mb-2">
          <div className="shrink-0">
            <svg width="24" height="24" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
              <g clip-path="url(#clip0_1_70)">
                <path d="M24 48C37.2548 48 48 37.2548 48 24C48 10.7452 37.2548 0 24 0C10.7452 0 0 10.7452 0 24C0 37.2548 10.7452 48 24 48Z" fill="#13152D"/>
                <path d="M35 11H13C11.8954 11 11 11.8954 11 13V35C11 36.1046 11.8954 37 13 37H35C36.1046 37 37 36.1046 37 35V13C37 11.8954 36.1046 11 35 11Z" fill="#121254" stroke="#FF6633" stroke-width="2"/>
                <path d="M15 15H33V33H15V15Z" fill="#121254" stroke="#00E500" stroke-width="2"/>
                <path d="M19 19H29V29H19V19Z" fill="#121254" stroke="#1A8CFF" stroke-width="2"/>
              </g>
              <defs>
                <clipPath id="clip0_1_70">
                  <rect width="48" height="48" fill="white"/>
                </clipPath>
              </defs>
            </svg>
          </div>
          <span className="text-primaryTextColor font-medium text-sm">Try bridging in Cortex</span>
        </div>
        <p className="text-secondaryTextColor text-xs mb-3">A new onchain experience powered by AI.</p>
        
        <a 
          href="https://cortexprotocol.com/agent" 
          target="_blank"
          rel="noopener noreferrer"
          className="inline-flex items-center justify-center w-full px-5 py-2 text-xs font-medium text-primaryTextColor bg-[#564f58] border border-transparent hover:border-[#AC8FFF] rounded-sm transition-all duration-200 cursor-pointer"
          onClick={() => {
            segmentAnalyticsEvent('cortex_promo_clicked', {
              action: 'click',
              location: 'popup',
              destination: 'cortex_agent'
            });
          }}
        >
          <span className="flex items-center">
            Launch Cortex Agent
            <ExternalLinkIcon className="ml-1.5 h-3 w-3 relative -mt-px" />
          </span>
        </a>
      </div>
    </div>
  );
};

export default CortexPromoPopup;