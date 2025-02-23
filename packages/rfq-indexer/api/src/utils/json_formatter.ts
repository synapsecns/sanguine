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
          width: 70%;
          background-color: #333;
          color: #ddd;
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
        }
        .indent-1 {
          padding-left: 20px;
        }
        .indent-2 {
          padding-left: 40px;
        }
        .divider {
          height: 15px;
          background-color: #000 !important;
          border:none;
        }
        .copy-icon {
          cursor: pointer;
          margin-left: 5px;
          color: #ddd;
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
      </style>
      <table>
        <tbody>
    `;

    // Function to add rows for each item
    const addRowsForItem = (item: any, index: number) => {
      const itemId = `item-${index}`;
      htmlTable += `<tr><td colspan="2" class="indent-0" style="background-color: #444;">
        <strong>Item #${index}</strong>
        <i class="fas fa-copy copy-icon" onclick="copyItemToClipboard('${itemId}', this)"></i>
      </td></tr>`;
      htmlTable += `<tr id="${itemId}">`;
      Object.entries(item).forEach(([header, content]) => {
        htmlTable += `<tr><td colspan="2" class="indent-1" style="background-color: #555;"><strong>${header}</strong></td></tr>`;
        if (typeof content === 'object' && content !== null) {
          Object.entries(content).forEach(([key, value]) => {
            htmlTable += `<tr><td class="indent-2" style="background-color: #666;">${key}</td><td>${formatValue(key, value)}</td></tr>`;
          });
        }
      });
      htmlTable += `</tr><tr class="divider"><td colspan="2" class='divider'></td></tr>`;
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
      if (key.toLowerCase().includes('timestamp') && typeof value === 'number') {
        transformations.push((val) => {
          const date = new Date(val * 1000);
          const formattedDate = date.toISOString().replace('T', ' ').slice(0, 19);
          const timeAgo = calculateTimeLabel(val);
          return `${formattedDate} utc (${timeAgo}) Unix: ${val}`;
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
          return `${val} <i class="fas fa-external-link-alt" style="cursor: pointer;" onclick="(function() {
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

    // Add rows for each item in the payload
    payload.forEach((item:any, index:any) => {
      addRowsForItem(item, index);
    });

    // Close table
    htmlTable += '</tbody></table>';

    // Add script for copy functionality
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

    // Attach JSON data to each item for copying
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