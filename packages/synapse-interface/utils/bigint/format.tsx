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

export function bigIntToFixed(bigIntValue, numDecimals = 2) {
    const asString = bigIntValue.toString();
    const truncatedString = asString.slice(0, numDecimals);

    return `${truncatedString}.${asString.slice(numDecimals)}`;
  }



  export function powBigInt(base, exponent) {
    let result = 1n;
    for (let i = 0; i < exponent; i++) {
        result *= base;
    }
    return result;
}

  export const formatBigIntToPercentString = (
    bn: bigint,
    nativePrecison: number,
    decimalPlaces = 2
  ) => {
    // Calculate the conversion factor based on the native precision and required decimal places
    const conversionFactor = powBigInt(10n, BigInt(nativePrecison - 2 + decimalPlaces))

    // Convert the bigint to a floating-point number, preserving the requested number of decimal places
    const num = Number(bn) / Number(conversionFactor)

    // Format the number as a percentage string
    return `${num.toFixed(decimalPlaces)}%`
  }

  export function fixNumberToPercentageString(num, numDecimals = 2) {
    return `${num?.toFixed(numDecimals)}%`
  }




  // Some environments have issues with RegEx that contain back-tracking, so we cannot
// use them.
export function commify(value: string | number): string {
  const comps = String(value).split(".");

  if (comps.length > 2 || !comps[0].match(/^-?[0-9]*$/) || (comps[1] && !comps[1].match(/^[0-9]*$/)) || value === "." || value === "-.") {
      console.log("invalid value", "value", value);
  }

  // Make sure we have at least one whole digit (0 if none)
  let whole = comps[0];

  let negative = "";
  if (whole.substring(0, 1) === "-") {
      negative = "-";
      whole = whole.substring(1);
  }

  // Make sure we have at least 1 whole digit with no leading zeros
  while (whole.substring(0, 1) === "0") { whole = whole.substring(1); }
  if (whole === "") { whole = "0"; }

  let suffix = "";
  if (comps.length === 2) { suffix = "." + (comps[1] || "0"); }
  while (suffix.length > 2 && suffix[suffix.length - 1] === "0") {
      suffix = suffix.substring(0, suffix.length - 1);
  }

  const formatted = [];
  while (whole.length) {
      if (whole.length <= 3) {
          formatted.unshift(whole);
          break;
      } else {
          const index = whole.length - 3;
          formatted.unshift(whole.substring(index));
          whole = whole.substring(0, index);
      }
  }

  return negative + formatted.join(",") + suffix;
}



export const commifyBigIntToString = (big: bigint, decimals = 2) => {
  return commify(formatBNToString(big, 18, decimals))
}

export const commifyBigIntWithDefault = (big: bigint, decimals: number) => {
  return big ? commifyBigIntToString(big, decimals) : '0'
}

export const MAX_UINT256 = 115792089237316195423570985008687907853269984665640564039457584007913129639935n;