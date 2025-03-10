#!/bin/bash
set -e

# Navigate to the project root directory
cd "$(dirname "$0")/.."

echo "Compiling TypeScript..."
npx tsc --skipLibCheck --esModuleInterop scripts/adjustBridgemap.ts

echo "Running BridgeMap adjustment script..."
node scripts/adjustBridgemap.js

echo "Cleaning up temporary files..."
rm scripts/adjustBridgemap.js
rm constants/bridgeMap.js

echo "Done! BridgeMap has been successfully adjusted."
