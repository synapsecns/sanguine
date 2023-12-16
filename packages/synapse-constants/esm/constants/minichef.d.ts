import { Address } from 'wagmi';
export interface ChainIdAddressMapping {
    [ChainId: number]: Address;
}
export declare const MINICHEF_ADDRESSES: ChainIdAddressMapping;
