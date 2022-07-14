import {expect} from "./chai-setup";
import {setupUsers, setupUser} from './utils';
import {ethers, deployments, getNamedAccounts, getUnnamedAccounts} from 'hardhat';
import { formatEther, parseEther } from "ethers/lib/utils";
const hre = require("hardhat");
const { network } = hre;

async function setup () {
  await deployments.fixture(["BiddingWar"]);

  const contracts = {
    BiddingWar: (await ethers.getContract('BiddingWar')),
  };

  const {owner} = await getNamedAccounts();

  const users = await setupUsers(await getUnnamedAccounts(), contracts);

  return {
    ...contracts,
    users,
    owner: await setupUser(owner, contracts),
  };
}

const defaultBidEpoch = 60 * 60;
const defaultTimeExtension = 10 * 60;
const commission = 5;
const playingPeriod = 60 * 1000;

describe("BiddingWar contract", function() {
  describe("Deployment", function () {

    it("Should set the right owner", async function () {
      const {BiddingWar} = await setup();
      const {owner} = await getNamedAccounts();

      expect(await BiddingWar.owner()).to.equal(owner);
    });
  });

  describe("Check default parameters", function() {
    it ("Should be same with default values", async function() {
      const {BiddingWar, users, owner} = await setup();

      // console.log(await BiddingWar.defaultBidEpoch());
      expect(await BiddingWar.defaultBidEpoch()).to.be.equal(60 * 60);
      expect(await BiddingWar.defaultTimeExtension()).to.be.equal(10 * 60);
      expect(await BiddingWar.commission()).to.be.equal(5);
    })
  })

  describe("test start()", function() {

    it ("owner start()", async function() {
      const {BiddingWar, users, owner} = await setup();

      expect(await BiddingWar.isPlaying()).to.be.equal(false);
      await owner.BiddingWar.startGame(playingPeriod);
      expect(await BiddingWar.isPlaying()).to.be.equal(true);

      let gameInfo = await BiddingWar.gameInfo();
      expect(gameInfo[1]).to.be.equal(0);  // start epoch should be 0
      expect(gameInfo[2]).to.be.equal(defaultTimeExtension);
      expect(gameInfo[3]).to.be.equal(0);
      expect(gameInfo[4]).to.be.equal(0);
      expect(gameInfo[5]).to.be.equal(ethers.constants.AddressZero);
      expect(gameInfo[6]).to.be.equal(playingPeriod);

      // should be reverted
      await expect(users[0].BiddingWar.startGame(playingPeriod)).to.be.revertedWith('Ownable: caller is not the owner');
    });

    it ("Not started Game test", async function() {
      const {BiddingWar, users, owner} = await setup();

      await expect(users[0].BiddingWar.bid({
        value: parseEther("0.1")
      })).to.be.revertedWith("GameNotPlaying");
    })
  })


  describe("bid test", function() {

    it ("test first bid", async function() {
      const {BiddingWar, users, owner} = await setup();
      await owner.BiddingWar.startGame(playingPeriod);

      let bidPrice = 0.1;
      let commision = parseInt(await BiddingWar.commission());

      // first bid
      await users[0].BiddingWar.bid({
        value: parseEther(bidPrice.toString())
      });

      let gameInfo = await BiddingWar.gameInfo();
      expect(gameInfo[1]).to.be.equal(defaultBidEpoch); //bid epoch should be changed
      expect(gameInfo[3]).to.be.equal(parseEther(bidPrice.toString()));
      expect(gameInfo[5]).to.be.equal(users[0].address);
      expect((gameInfo[4])).to.be.equal(parseEther((bidPrice * (100 - commision) / 100).toString()));
    });


    it ("bid epoch test", async function() {
      const {BiddingWar, users, owner} = await setup();
      await owner.BiddingWar.startGame(playingPeriod);

      let bidPrice = 0.1;

      //first bid
      await users[0].BiddingWar.bid({
        value: parseEther(bidPrice.toString())
      });


      //should fail before bidEpoch
      await expect(users[0].BiddingWar.bid({
        value: parseEther((bidPrice + 1).toString())
      })).to.be.revertedWith("NotElapsedToBid");

      await network.provider.send("evm_increaseTime", [60 * 60]);
      await network.provider.send("evm_mine");

      // should be reverted with lowerbid
      await expect(users[1].BiddingWar.bid({
        value: parseEther((bidPrice).toString())
      })).to.be.revertedWith("LowerBid");


      await users[1].BiddingWar.bid({
        value: parseEther((bidPrice + 1).toString())
      });
    })

    it ("several bids test", async function() {
      const {BiddingWar, users, owner} = await setup();
      await owner.BiddingWar.startGame(playingPeriod);

      let firstBidPrice = 1;
      let totalBidPrice = 0;
      let commision = parseInt(await BiddingWar.commission());

      //first bid
      await users[0].BiddingWar.bid({
        value: parseEther(firstBidPrice.toString())
      });
      totalBidPrice += firstBidPrice;

      await network.provider.send("evm_increaseTime", [60 * 60]);
      await network.provider.send("evm_mine");


      //second bid
      let secondBidPrice = 2;
      await users[1].BiddingWar.bid({
        value: parseEther(secondBidPrice.toString())
      });

      totalBidPrice += secondBidPrice;
      let gameInfo = await BiddingWar.gameInfo();
      expect(gameInfo[1]).to.be.equal(defaultBidEpoch + defaultTimeExtension); //bid epoch should be changed
      expect(gameInfo[3]).to.be.equal(parseEther(secondBidPrice.toString()));
      expect(gameInfo[5]).to.be.equal(users[1].address);
      expect((gameInfo[4])).to.be.equal(parseEther((totalBidPrice * (100 - commision) / 100).toString()));


      await network.provider.send("evm_increaseTime", [60 * 70]);
      await network.provider.send("evm_mine");

      let thirdBidPrice = 3;
      await users[2].BiddingWar.bid({
        value: parseEther(thirdBidPrice.toString())
      });
      totalBidPrice += thirdBidPrice;

      gameInfo = await BiddingWar.gameInfo();
      expect(gameInfo[1]).to.be.equal(defaultBidEpoch + defaultTimeExtension * 2); //bid epoch should be changed
      expect(gameInfo[3]).to.be.equal(parseEther(thirdBidPrice.toString()));
      expect(gameInfo[5]).to.be.equal(users[2].address);
      expect((gameInfo[4])).to.be.equal(parseEther((totalBidPrice * (100 - commision) / 100).toString()));
    });

  })

  describe("End test", function() {
    it ("Not Owner End", async function() {
      const {BiddingWar, users, owner} = await setup();
      await owner.BiddingWar.startGame(playingPeriod);

      // should be fail : Not owner
      await expect(users[0].BiddingWar.endGame()).to.be.revertedWith("Ownable: caller is not the owner");
    })

    it ("End not playing game", async function() {
      const {BiddingWar, users, owner} = await setup();
      await expect(owner.BiddingWar.endGame()).to.be.revertedWith("GameNotPlaying");
      
    })

    it ("End game before playingPeriod", async function() {
      const {BiddingWar, users, owner} = await setup();
      await owner.BiddingWar.startGame(playingPeriod);

      await expect(owner.BiddingWar.endGame()).to.be.revertedWith("NotElapsedToEnd");
    })

    it ("Enable to end game", async function() {
      const {BiddingWar, users, owner} = await setup();
      await owner.BiddingWar.startGame(playingPeriod);

      await network.provider.send("evm_increaseTime", [playingPeriod]);
      await network.provider.send("evm_mine");

      await owner.BiddingWar.endGame();
      expect(await BiddingWar.isPlaying()).to.be.equal(false);

    });

    it ("GameInfo after game ended", async function() {
      const {BiddingWar, users, owner} = await setup();
      await owner.BiddingWar.startGame(playingPeriod);

      let firstBidPrice = 1;
      let totalBidPrice = 0;
      let commision = parseInt(await BiddingWar.commission());

      //first bid
      await users[0].BiddingWar.bid({
        value: parseEther(firstBidPrice.toString())
      });
      totalBidPrice += firstBidPrice;

      await network.provider.send("evm_increaseTime", [60 * 60]);
      await network.provider.send("evm_mine");


      //second bid
      let secondBidPrice = 2;
      await users[1].BiddingWar.bid({
        value: parseEther(secondBidPrice.toString())
      });

      totalBidPrice += secondBidPrice;


      await network.provider.send("evm_increaseTime", [60 * 70]);
      await network.provider.send("evm_mine");

      let thirdBidPrice = 3;
      await users[2].BiddingWar.bid({
        value: parseEther(thirdBidPrice.toString())
      });
      totalBidPrice += thirdBidPrice;

      await network.provider.send("evm_increaseTime", [60 * 900]);
      await network.provider.send("evm_mine");

      //Unable to bid
      await expect(users[3].BiddingWar.bid({
        value: parseEther("1000")
      })).to.be.revertedWith("GameEnded");

      let priceForWinner = totalBidPrice * (100 - commision) / 100;
      let bUser2Bal = await ethers.provider.getBalance(users[2].address);
      console.log("Before balance", formatEther(bUser2Bal));
      //End game
      await owner.BiddingWar.endGame();

      let gameInfo = await BiddingWar.gameInfo();
      // winner address
      expect(gameInfo[5]).to.be.equal(users[2].address);
      // winner bidPrice
      expect(gameInfo[3]).to.be.equal(parseEther(thirdBidPrice.toString()));
      // accumulated funds
      expect((gameInfo[4])).to.be.equal(parseEther((totalBidPrice * (100 - commision) / 100).toString()));

      //Transfer transfer accumulated funds to winner
      let aUser2Bal = await ethers.provider.getBalance(users[2].address);
      expect(aUser2Bal.sub(bUser2Bal)).to.be.equal(parseEther(priceForWinner.toString()));

      //Check contract balance
      expect(await ethers.provider.getBalance(BiddingWar.address)).to.be.equal(parseEther((totalBidPrice * commision / 100).toString()));
    });

  })

  describe("Withdraw test", function() {
    it ("Not owner withdraw", async function() {
      const {BiddingWar, users, owner} = await setup();

      let withdrawAddr = users[3].address;
      await expect(users[0].BiddingWar.withdraw(withdrawAddr)).to.be.revertedWith("Ownable: caller is not the owner")
    });

    it("Not game ended", async function() {
      const {BiddingWar, users, owner} = await setup();

      await owner.BiddingWar.startGame(playingPeriod);
      let withdrawAddr = users[3].address;
      await expect(owner.BiddingWar.withdraw(withdrawAddr)).to.be.revertedWith("GameNotEnded")      
    });


    it ("test withdraw balance", async function() {
      const {BiddingWar, users, owner} = await setup();
      await owner.BiddingWar.startGame(playingPeriod);

      let firstBidPrice = 10;
      let totalBidPrice = 0;
      let commision = parseInt(await BiddingWar.commission());
      let withdrawAddr = users[3].address;

      //first bid
      await users[0].BiddingWar.bid({
        value: parseEther(firstBidPrice.toString())
      });
      totalBidPrice += firstBidPrice;
      await network.provider.send("evm_increaseTime", [60 * 60]);
      await network.provider.send("evm_mine");


      //second bid
      let secondBidPrice = 20;
      await users[1].BiddingWar.bid({
        value: parseEther(secondBidPrice.toString())
      });
      totalBidPrice += secondBidPrice;


      await network.provider.send("evm_increaseTime", [60 * 70]);
      await network.provider.send("evm_mine");
      let thirdBidPrice = 30;
      await users[2].BiddingWar.bid({
        value: parseEther(thirdBidPrice.toString())
      });
      totalBidPrice += thirdBidPrice;

      await network.provider.send("evm_increaseTime", [60 * 900]);
      await network.provider.send("evm_mine");
      await owner.BiddingWar.endGame();

      let contractBal = await ethers.provider.getBalance(BiddingWar.address);

      let bWithdrawBal = await ethers.provider.getBalance(withdrawAddr);
      //withdraw
      await owner.BiddingWar.withdraw(withdrawAddr);

      let aWithdrawBal = await ethers.provider.getBalance(withdrawAddr);
      expect(aWithdrawBal.sub(bWithdrawBal)).to.be.equal(contractBal)
    })
  })
});
