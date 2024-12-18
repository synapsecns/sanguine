> :warning: **Warning**: This tool is not intended to be used for anything other than testing with inconsequentially small amounts of assets on EOAs that were created for the sole & explicit purpose of testing.

# RFQ Load Test Configuration Guide

This guide outlines the steps and requirements for setting up the RFQ load testing tool. The tool is designed to send large varieties/volumes of actual bridges.

## Wallet Configuration

- **Create and Fund Wallets:** You are required to create and fund as many wallets as you wish to test with. They will need sufficient gas and whatever other tokens you wish to test.


# Script Configuration

Follow these steps to configure your load testing environment for blockchain transactions.

## Step 1: Prepare Configuration File

1. Start by copying the `config-template.yaml` file from the `packages/rfq-loadtest` directory. This will serve as the basis for your test configuration.
   ```bash
   cp packages/rfq-loadtest/config-template.yaml packages/rfq-loadtest/config-run.yaml
   ```

## Step 2: Configure Private Keys

2. Open the `config-run.yaml` file and locate the `PRIVATE_KEY_X` entries. Replace the placeholder values with your actual private keys. These keys will be used to execute transactions during the tests.
   ```yaml
   PRIVATE_KEY_1: 'your_private_key_here'
   PRIVATE_KEY_2: 'your_private_key_here'
   PRIVATE_KEY_3: 'your_private_key_here'
   PRIVATE_KEY_4: 'your_private_key_here'
   PRIVATE_KEY_5: 'your_private_key_here'
   ```

## Step 3: Set Transaction Pace

3. Adjust the settings under `##### TEST PACE` to control the pace of transactions. You can modify the `VOLLEY_MILLISECONDS_BETWEEN`, `VOLLEY_MIN_COUNT`, and `VOLLEY_MAX_COUNT` to fit your testing requirements.

## Step 4: Specify Test Bridge Amount

4. Define each asset to be sent in test transactions under the `ASSETS` section. 

## Step 6: Define Chain Settings

6. In the `CHAINS` section, configure the `Label`, Router addresses, and `rpcUrl` settings for each chain involved in your tests. These settings include URLs for reading, simulation, and submitting transactions.

## Step 8: Save and Run

8. Save your changes to `config-run.yaml`. To start the load test, use the provided startup example, adjusting the path to your configuration file as necessary.
   ```bash
   python3 pyRepeater.py 'node index.js --configFile ../config-run.yaml --pKeyIndex 1'
   ```
