export const formatBNToString = (
    bn: bigint,
    nativePrecision: number,
    decimalPlaces?: number
  ) => {
    // Convert to string and add padding zeros if necessary
    let str = bn.toString().padStart(nativePrecision, '0');

    // Insert decimal point
    const idx = str.length - nativePrecision;
    str = `${str.slice(0, idx)}.${str.slice(idx)}`;

    // Trim to desired number of decimal places
    if (decimalPlaces !== undefined) {
      const decimalIdx = str.indexOf('.');
      str = str.slice(0, decimalIdx + decimalPlaces + 1);
    }

    return str;
  };
