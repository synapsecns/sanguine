import { JsonRpcProvider } from '@ethersproject/providers';
import { SynapseSDK } from '@synapsecns/sdk-router';
import { BigNumber } from '@ethersproject/bignumber';
import express from 'express';


//Setting up RPC providers:
//@simon Need a list of all providers (Can be public) from somewhere else (sdk or interface)
const arbitrumProvider = new JsonRpcProvider('https://arb1.arbitrum.io/rpc');
const avalancheProvider = new JsonRpcProvider('https://api.avax.network/ext/bc/C/rpc');

const app = express();
const port = process.env.PORT || 3000;

//Intro Message for UI
app.get('/', (req, res) => {
  res.send('Welcome to the Synapse Rest API for swap and bridge quotes')
});

//Setting up arguments
//@simon Need a list of all chain ids... maybe can use supportedchainIds from the sdk? was just having package issues.
const chainIds = [42161,43114];
const providers = [ arbitrumProvider, avalancheProvider];

//Set up a SynapseSDK Instance
const Synapse = new SynapseSDK(chainIds, providers);

//Swap get request
app.get('/swap/:chain/:fromToken/:toToken/:amount', async(req,res) => {
  const chain = req.params.chain;
  //@simon Need logic here that takes in the chain and the token symbol and returns the token address for that chain (for both the to and From tokens) @simon
  const fromToken = req.params.fromToken;
  const toToken = req.params.toToken;
  // @simon Need logic here that takes in the amount and multiplies it by the decimals for that token on its respective chain
  const amount = req.params.amount;

  const resp = await Synapse.swapQuote(chain, fromToken, toToken, BigNumber.from(amount));

  //Hardcoded implementation for testing purposes only
  // const resp = await Synapse.swapQuote(42161, '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8', '0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9', BigNumber.from(100000000));


  res.json(resp);
});


//Bridge get request
app.get('/bridge/:fromChain/:toChain/:fromToken/:toToken/:amount', async(req,res) => {
  const fromChain = req.params.fromChain;
  const toChain = req.params.toChain;
  //@simon Need logic here that takes in the chain and the token symbol and returns the token address for that chain (for both the to and From tokens) @simon
  const fromToken = req.params.fromToken;
  const toToken = req.params.toToken;
  //@simon Need logic here that takes in the amount and multiplies it by the decimals for that token on its respective chain @simon
  const amount = req.params.amount;

  const resp = await Synapse.bridgeQuote(fromChain,toChain, fromToken, toToken, BigNumber.from(amount));

  //Hardcoded implementation for testing purposes only
  // const resp = await Synapse.swapQuote(42161, '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8', '0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9', BigNumber.from(100000000));


  res.json(resp);
});


app.listen(port, () => {
  console.log('Server listening at ${port}')
});


// export const allChains = [
//   1,
//   10,
//   56,
//   250,
//   288,
//   1284,
//   1285,
//   137,
//   43144,
//   53935,
//   42161,
//   1313161554,
//   1666600000,
//   25,
//   1088,
//   8217,
//   2000,
//   7700
// ]

// export const rpcProviders = [
//   new JsonRpcProvider('https://rpc.ankr.com/eth'),
//   new JsonRpcProvider('https://rpc.ankr.com/optimism'),
//   new JsonRpcProvider('https://bsc-dataseed1.ninicoin.io/'),
//   new JsonRpcProvider('https://rpc.ftm.tools'),
//   new JsonRpcProvider('https://replica-oolong.boba.network/'),
//   new JsonRpcProvider('https://rpc.api.moonbeam.network'),
//   new JsonRpcProvider('https://rpc.api.moonriver.moonbeam.network'),
//   new JsonRpcProvider('https://rpc-mainnet.matic.quiknode.pro'),
//   new JsonRpcProvider('https://api.avax.network/ext/bc/C/rpc'),
//   new JsonRpcProvider('https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc'),
//   new JsonRpcProvider('https://arb1.arbitrum.io/rpc'),
//   new JsonRpcProvider('https://mainnet.aurora.dev'),
//   new JsonRpcProvider('https://harmony-mainnet.chainstacklabs.com'),
//   new JsonRpcProvider('https://evm-cronos.crypto.org'),
//   new JsonRpcProvider('https://andromeda.metis.io/?owner=1088'),
//   new JsonRpcProvider('https://klaytn.blockpi.network/v1/rpc/public'),
//   new JsonRpcProvider('https://rpc.ankr.com/dogechain'),
//   new JsonRpcProvider('https://mainnode.plexnode.org:8545'),
// ]
