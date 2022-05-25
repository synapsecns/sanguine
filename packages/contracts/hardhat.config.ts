import { HardhatUserConfig } from 'hardhat/types'
import 'solidity-coverage'
import * as dotenv from 'dotenv'
import './tasks/flatten'


const config: HardhatUserConfig = {
    networks: {
      hardhat: {

      },
    },
}
