import {HardhatUserConfig} from 'hardhat/types';
import 'hardhat-deploy';
import 'hardhat-deploy-ethers';
import {node_url, accounts} from './utils/network';

const config: HardhatUserConfig = {
  solidity: {
    compilers: [
      {
        version: "0.8.6",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200
          }
        }
      },
    ]
  },
  defaultNetwork: "hardhat",
  networks: {
    hardhat: {
      forking: {
        // url: "https://eth-rinkeby.alchemyapi.io/v2/7yKOjZzEgFwVijCpEriJDu2ihLlmpG7X",
        url: `https://nodeapi.test.energi.network/v1/jsonrpc`,
        // blockNumber: 10380183
        blockNumber: 1216262
      }
    },
    rinkeby: {
      url: 'https://eth-rinkeby.alchemyapi.io/v2/T76f9ewGuXiIe0l3mVwh1t75zarg_SDz',
      accounts: accounts('rinkeby'),
    },
    testEnergi: {
      url: `https://nodeapi.test.energi.network/v1/jsonrpc`,
      accounts: accounts('testenergi')
    }
  },
  namedAccounts: {
    owner: 0,
    user1: 1,
    buyer: 2,
  },
  paths: {
    sources: 'src',
  },
  verify: {
    etherscan: {
      apiKey: '1ZRP5UIMBSWXNKPZBDXNRNVAP8TAPF2294'
    }
  }
};
export default config;
