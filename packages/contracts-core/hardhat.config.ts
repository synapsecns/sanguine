import { HardhatUserConfig } from 'hardhat/types'
import 'solidity-coverage'
import * as dotenv from 'dotenv'


const config: HardhatUserConfig = {
    networks: {
      hardhat: {

      },
    },
  solidity: {
      compilers: [
        {
          version: "0.8.17",
          settings: {
            optimizer: {
              enabled: true,
              runs: 5000000, // see: https://github.com/ethereum/solidity/issues/5394#issue-379536332
            },
          },
        },
      ]
  }
}


module.exports = config;
