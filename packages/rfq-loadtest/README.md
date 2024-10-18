> :warning: **Warning**: This tool is not intended to be used for anything other than testing with inconsequentially small amounts of ETH on EOAs that were created for the sole & explicit purpose of testing.

# RFQ Load Test Configuration Guide

This guide outlines the steps and requirements for setting up the RFQ load testing tool. The tool is designed to send many ETH bridges in rapid succession.

## Wallet Configuration

- **Create and Fund Wallets:** You are required to create and fund as many wallets as you wish to test with. This tool only supports native ETH bridges on chains that use ETH as the gas currency.

- **Auto-Rebalance:** If you only initially fund ETH on one of the test chains, the tool will automatically rebalance the funds to the other chains before beginning the tests. It will also rebalance as needed while the tests are operating until none of the test chains have enough ETH to effectively rebalance to the others - at which point it will cease the process. The RFQ system is also used for these rebalance actions.


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

4. Define the amount of ETH to be sent in each test transaction under `TEST_BRIDGE_AMOUNT_UNITS`. The default is set to `0.00007`.

## Step 5: Configure Gas and Rebalance Settings

5. Set the `MINIMUM_GAS_UNITS` to the desired threshold for triggering a rebalance of funds across chains. Specify the `REBALANCE_TO_UNITS` to determine the target amount for each chain after rebalancing.

## Step 6: Define Chain Settings

6. In the `CHAINS` section, configure the `FastRouterAddr` and `rpcUrl` settings for each chain involved in your tests. These settings include URLs for reading, simulation, and submitting transactions.

## Step 7: Configure Test Routes

7. Under `TEST_ROUTES`, define the routes for your test transactions, including `fromChainId`, `toChainId`, and `testDistributionPercentage` to control the flow and distribution of transactions between chains.

## Step 8: Save and Run

8. Save your changes to `config-run.yaml`. To start the load test, use the provided startup example, adjusting the path to your configuration file as necessary.
   ```bash
   python3 pyRepeater.py 'node index.js --configFile ../config-run.yaml --pKeyIndex 1'
   ```
