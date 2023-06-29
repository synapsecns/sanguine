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


export const formatBigIntToString = (
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


  export const formatBigIntToPercentString = (
    bn: bigint,
    nativePrecison: number,
    decimalPlaces = 2
  ) => {
    // Calculate the conversion factor based on the native precision and required decimal places
    const conversionFactor = 10n ** BigInt(nativePrecison - 2 + decimalPlaces);

    // Convert the bigint to a floating-point number, preserving the requested number of decimal places
    const num = Number(bn) / Number(conversionFactor)

    // Format the number as a percentage string
    return `${num.toFixed(decimalPlaces)}%`
  }
