
import { scannerLink } from "./enrichResults";

export function jsonToHtmlTable(payload: any): string {
  try {
    // Ensure payload is an array of objects -- make an array of 1 if not.
    if (!Array.isArray(payload)) {
      payload = [payload];
    } else if (payload.length === 0 || typeof payload[0] !== 'object') {
      payload = [{}];
    }

    // Start table with CSS styling
    let htmlTable = `
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.1/css/all.min.css">
      <style>
        body { background-color: #000;}
        table {
          font-family: 'Consolas', monospace;
          font-size: 12px;
          border-collapse: collapse;
          background-color: #222;
          color: white;
          width: 820px;
        }
        th, td {
          border: 1px solid #555;
          padding: 2px;
          text-align: left;
        }
        th {
          background-color: #444;
        }
        td[colspan="2"] {
          font-weight: bold;
          background-color: #444;
        }
        .indent-0 {
          padding-left: 5px;
          padding-right: 5px;
        }
        .indent-1 {
          padding-left: 20px;
          padding-right: 5px;
        }
        .indent-2 {
          padding-left: 40px;
          padding-right: 5px;
        }
        .divider {
          height: 15px;
          background-color: #000 !important;
          border:none;
        }
        .copy-icon {
          cursor: pointer;
          margin-left: 5px;
          color: white;
        }
        .copy-icon:hover {
          color: #fff;
        }
        .check-icon {
          color: green;
        }
        td:first-child {
          width: 1%;
          white-space: nowrap;
        }
        .valuecell {
          padding-left: 5px;
          background-color: #222 !important;
        }
      </style>
      <table>
        <tbody>
    `;


    
const addRowsForItem = (item: any, index: number) => {
  
  const origChainId = item.Bridge.originChainId;
  const destChainId = item.Bridge.destChainId;

  const itemId = `item-${index}`;
  htmlTable += `<tr><td colspan="3" class="indent-0" style="background-color: #292929;">
    <strong>Item #${index}</strong>
    <i class="fas fa-copy copy-icon" onclick="copyItemToClipboard('${itemId}', this)"></i>
  </td></tr>`;
  htmlTable += `<tr id="${itemId}">`;

  Object.entries(item).forEach(([header, content]: [string, any]) => {
    htmlTable += `<tr><td colspan="3" class="indent-1" style="background-color: #393939;"><strong>${header}</strong></td></tr>`;
    if (typeof content === 'object' && content !== null) {
      const originEntries: [string, any][] = [];
      const destEntries: [string, any][] = [];
      const otherEntries: [string, any][] = [];
      let usdValueDifference = null;
      let amountFormattedDifference = null;
      let timeDifference = null;

      Object.entries(content).forEach(([key, value]) => {
        if (key.startsWith('origin')) {
          originEntries.push([key, value]);
        } else if (key.startsWith('dest')) {
          destEntries.push([key, value]);
        } else {
          otherEntries.push([key, value]);
        }
      });

      if (originEntries.length > 0 || destEntries.length > 0) {
        htmlTable += `<tr><td class="indent-2" style="text-align:center; background-color: #494949;"></td><td style="padding-right: 10px; text-align:right; background-color: #494949;"><strong>Origin</strong></td><td style="padding-left: 10; text-align:left; background-color: #494949;"><strong>Destination</strong></td></tr>`;
        const maxLength = Math.max(originEntries.length, destEntries.length);
        for (let i = 0; i < maxLength; i++) {
          const originEntry = originEntries[i] || ['', ''];
          const destEntry = destEntries[i] || ['', ''];
          const label = originEntry[0].replace(/^origin|dest/, '') || destEntry[0].replace(/^origin|dest/, '');
          htmlTable += `<tr><td class="indent-2" style="text-align:right; background-color: #494949;">${label}</td>`;
          htmlTable += `<td class="valuecell" style="text-align:right;">${formatValue(originEntry[0], originEntry[1])}</td>`;
          htmlTable += `<td class="valuecell" style="text-align:left;">${formatValue(destEntry[0], destEntry[1])}</td></tr>`;

          // Calculate usdvalue differences
          if (label === 'UsdValue') {
            const originUsdValue = originEntry[1] !== null ? parseFloat(originEntry[1]) : null;
            const destUsdValue = destEntry[1] !== null ? parseFloat(destEntry[1]) : null;
            usdValueDifference = (originUsdValue !== null && destUsdValue !== null) ? (originUsdValue - destUsdValue).toFixed(6) : 'unknown';
          }

          // Calculate amount differences
          if (label === 'AmountFormatted') {
            const originTokenSymbol = content['originTokenSymbol'];
            const destTokenSymbol = content['destTokenSymbol'];
            const originAmountFormatted = parseFloat(originEntry[1]) || 0;
            const destAmountFormatted = parseFloat(destEntry[1]) || 0;
            amountFormattedDifference = (originTokenSymbol && destTokenSymbol) ? 
              (originTokenSymbol === destTokenSymbol ? (originAmountFormatted - destAmountFormatted).toFixed(18) : 'n/a') : 'hide';
          }

        }

        // Add amount differences row
        if (amountFormattedDifference !== null && amountFormattedDifference!='hide') {
          htmlTable += `<tr><td class="indent-2" style="text-align:right; background-color: #494949;">Units Diff</td>`;
          htmlTable += `<td colspan="2" class="valuecell" style="text-align:left;">${amountFormattedDifference}</td></tr>`;
        }
        
        // Add usdvalue differences row
        if (usdValueDifference !== null) {
          htmlTable += `<tr><td class="indent-2" style="text-align:right; background-color: #494949;">UsdValue Diff</td>`;
          htmlTable += `<td colspan="2" class="valuecell" style="text-align:left;">${usdValueDifference}</td></tr>`;
        }
      }


      // Calculate relay time
      if (header === 'BridgeRelay') {
        timeDifference = item.BridgeRelay.blockTimestamp - item.BridgeRequest.blockTimestamp;
      }

      // relay time
      if (timeDifference !== null) {
        htmlTable += `<tr><td class="indent-2" style="text-align:right; background-color: #494949;">Relay Time</td>`;
        htmlTable += `<td colspan="2" class="valuecell" style="text-align:left;">${timeDifference} seconds</td></tr>`;
      }

      otherEntries.forEach(([key, value]) => {
        let displayValue = formatValue(key, value);
        const chainId = header == 'BridgeRelay' ? destChainId : origChainId;

        if (['transactionHash', 'sender', 'relayer', 'to'].includes(key) && chainId) {
          displayValue = scannerLink(chainId, value);
        }

        if (key.toLowerCase() === 'deadline' && value !== null && !isNaN(Number(value))) {
          const deadlineTimestamp = Number(value);
          const now = Math.floor(Date.now() / 1000);
          if (deadlineTimestamp > now) {
            const transactionId = item.Bridge.transactionId; // Assuming transactionId is available in the item object
            displayValue += ` <i class="fas fa-undo" title="Clipboard OpBot Cancel Command" style="color: yellow; cursor: pointer;" onclick="copyToClipboard('@OpBot refund ${transactionId}', this)"></i>`;
          }
        }

        htmlTable += `<tr><td class="indent-2" style="text-align:right; background-color: #494949;">${key}</td><td colspan="2" class="valuecell">${displayValue}</td></tr>`;
      });
    }
  });
  htmlTable += `</tr><tr class="divider"><td colspan="3" class='divider'></td></tr>`;
};

    // Function to format value with transformations
    const formatValue = (key: string, value: any) => {
      let displayValue = value;
      const transformations: Array<(val: any) => any> = [];

      // Transformation for long strings
      if (typeof value === 'string' && value.length > 100) {
        transformations.push((val) => `${val.slice(0, 20)}...${val.slice(-20)}`);
      }

      // Transformation for strings starting with "0x"
      if (typeof value === 'string' && value.startsWith('0x')) {
        transformations.push((val) => `
          ${val}
          <i class="fas fa-copy copy-icon" onclick="copyToClipboard('${value}', this)"></i>
        `);
      }
      
      // Transformation for timestamps
      if (value !== null && (key.toLowerCase().includes('timestamp') || key.toLowerCase().includes('deadline') || key.toLowerCase().includes('endtime')) && !isNaN(Number(value))) {
        transformations.push((val) => {
          const numericValue = Number(val);
          const date = new Date(numericValue * 1000);
          const formattedDate = date.toISOString().replace('T', ' ').slice(0, 19);
          const timeAgo = calculateTimeLabel(numericValue);
          return `${formattedDate} utc (${timeAgo}) Unix: ${numericValue}`;
        });
      }

      // Transformation for senderStatus
      if (key === 'senderStatus') {
        transformations.push((val) => {
          if (val === 'OK') {
            return `${val} <i class="fas fa-check-circle" style="color: green;"></i>`;
          } else if (val === 'SCREENED') {
            return `${val} <i class="fas fa-exclamation-circle" style="color: red;"></i>`;
          } else {
            return `${val} <i class="fas fa-exclamation-triangle" style="color: yellow;"></i>`;
          }
        });
      }
      if (key === 'transactionId') {
        transformations.push((val) => {
          return `${val} 
            <i class="fas fa-book" style="cursor: pointer; margin-left: 5px;" title="Clipboard OpBot Trace Command" onclick="(function() {
              const logText = '@OpBot trace transaction_id:${value}';
              navigator.clipboard.writeText(logText).then(() => {
                alert('Log command copied to clipboard');
              }).catch(err => {
                console.error('Error copying log command:', err);
              });
            })()"></i>
            <i class="fas fa-external-link-alt" style="cursor: pointer;" onclick="(function() {
              const currentUrl = document.location;
              const baseUrl = currentUrl.origin + currentUrl.pathname.split('/').slice(0, -2).join('/');
              const extraParams = currentUrl.search;
              const newUrl = baseUrl.concat('/transaction-id/', '${value}', extraParams);
              window.open(newUrl, '_blank');
            })()"></i>`;
        });
      }

      // Apply all transformations
      transformations.forEach(transform => {
        displayValue = transform(displayValue);
      });

      return displayValue;
    };

    const calculateTimeLabel = (timestamp: number) => {
      const now = Math.floor(Date.now() / 1000);
      let diff = now - timestamp;
      const isFuture = diff < 0;
      diff = Math.abs(diff);
      const days = Math.floor(diff / (24 * 3600));
      diff %= 24 * 3600;
      const hours = Math.floor(diff / 3600);
      diff %= 3600;
      const minutes = Math.floor(diff / 60);
      const seconds = diff % 60;

      let timeString = '';
      if (days > 0) timeString += `${days}d `;
      if (hours > 0 || days > 0) timeString += `${hours}h `;
      if (minutes > 0 || hours > 0 || days > 0) timeString += `${minutes}m `;
      timeString += `${seconds}s`;

      return isFuture ? `${timeString} from now` : `${timeString} ago`;
    };

    payload.forEach((item:any, index:any) => {
      addRowsForItem(item, index);
    });

    htmlTable += '</tbody></table>';

    htmlTable += `
      <script>
        function copyToClipboard(text, iconElement) {
          navigator.clipboard.writeText(text).then(() => {
            iconElement.classList.remove('fa-copy');
            iconElement.classList.add('fa-check', 'check-icon');
            iconElement.style.color = 'green'; // Immediately change color to green
            setTimeout(() => {
              iconElement.classList.remove('fa-check', 'check-icon');
              iconElement.classList.add('fa-copy');
              iconElement.style.color = ''; // Reset color
            }, 1000);
          }).catch(err => {
            console.error('Failed to copy text: ', err);
          });
        }

        function copyItemToClipboard(itemId, iconElement) {
          const itemElement = document.getElementById(itemId);
          if (itemElement) {
            const itemData = JSON.parse(itemElement.dataset.json);
            const prettyJson = JSON.stringify(itemData, null, 2);
            copyToClipboard(prettyJson, iconElement);
          }
        }
      </script>
    `;

    payload.forEach((item: any, index: number) => {
      const itemId = `item-${index}`;
      const itemJson = JSON.stringify(item);
      htmlTable = htmlTable.replace(`<tr id="${itemId}">`, `<tr id="${itemId}" data-json='${itemJson}'>`);
    });

    return htmlTable;
  } catch (error) {
    console.error('Error converting JSON to HTML table:', error);
    throw new Error('Failed to convert JSON to HTML table');
  }
}