import { HardhatUserConfig } from 'hardhat/types'
import 'solidity-coverage'
import * as dotenv from 'dotenv'



const config: HardhatUserConfig = {
    networks: {
      hardhat: {
        live: false,
        saveDeployments: false,
        tags: ['local'],
      },

    }
}