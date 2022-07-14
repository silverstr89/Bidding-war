import {HardhatRuntimeEnvironment} from 'hardhat/types'; 
import {DeployFunction} from 'hardhat-deploy/types';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) { 
  const {deployments, getNamedAccounts} = hre;
  const {deploy} = deployments;

  const {owner} = await getNamedAccounts();

  console.log(owner);

  await deploy('BiddingWar', {
    from: owner,
    args: [
      60 * 60,
      60 * 10,
      5
    ],
    log: true,
  });
};
export default func;
func.tags = ['BiddingWar'];