// SPDX-License-Identifier: MIT

pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";
import "hardhat/console.sol";

error NotElapsedToBid(uint256 nextEpoch, uint256 currentTime);
error LowerBid(uint256 highestBid, uint256 currentBid);
error GameNotEnded();
error GameNotPlaying();
error GameEnded();
error NotElapsedToEnd(uint256 endTime, uint256 currentTime);
error TransferFailed(address sender, uint256 amount);

contract BiddingWar is Ownable {
  uint256 public defaultBidEpoch;
  uint256 public defaultTimeExtension;
  uint256 public commission;

  struct GameInformation {
    uint256 startedAt;
    uint256 bidEpoch;
    uint256 timeExtension;
    uint256 highestBidPrice;    
    uint256 accumulatedFund;
    address highestBidder;
    uint256 playingPeriod;
    bool isPlaying;
    uint256 roundEndAt;
  }

  GameInformation private _gameInfo;

  event Bid(address indexed bidder, uint256 bidPrice);
  event StartGame(uint256 startedAt);
  event EndGame(address winner, uint256 prizeAmount);
  event Transfer(address sender, uint256 amount);

  constructor(
    uint256 bidEpoch,
    uint256 timeExtension,
    uint256 commission_
  ) {
    defaultBidEpoch = bidEpoch;
    defaultTimeExtension = timeExtension;
    commission = commission_;
  }

  function _startGame(
    uint256 playingPeriod
  ) internal {
    _gameInfo.isPlaying = true;
    _gameInfo.startedAt = block.timestamp;
    _gameInfo.playingPeriod = playingPeriod;
    _gameInfo.bidEpoch = 0;
    _gameInfo.timeExtension = defaultTimeExtension;
    emit StartGame(block.timestamp);
  }


  function gameInfo() external view 
    returns(
      uint256,
      uint256,
      uint256,
      uint256,
      uint256,
      address,
      uint256,
      bool,
      uint256
    )
  {
    return (
      _gameInfo.startedAt,
      _gameInfo.bidEpoch,
      _gameInfo.timeExtension,
      _gameInfo.highestBidPrice,
      _gameInfo.accumulatedFund,
      _gameInfo.highestBidder,
      _gameInfo.playingPeriod,
      _gameInfo.isPlaying,
      _gameInfo.roundEndAt
    );
  }


  function _endGame() internal {
    _gameInfo.isPlaying = false;

    (bool s1, ) = (_gameInfo.highestBidder).call{value: (_gameInfo.accumulatedFund)}("");
    if (!s1) revert TransferFailed(_gameInfo.highestBidder, _gameInfo.accumulatedFund);

    emit EndGame(_gameInfo.highestBidder, _gameInfo.highestBidPrice);
    emit Transfer(_gameInfo.highestBidder, _gameInfo.accumulatedFund);
  }


  function _updateGameInfo(address bidder, uint256 bidPrice) internal {
    if (_gameInfo.highestBidder == address(0)) {
      _gameInfo.bidEpoch = defaultBidEpoch;
    } else {
      _gameInfo.bidEpoch += _gameInfo.timeExtension;
    }

    _gameInfo.roundEndAt = block.timestamp + _gameInfo.bidEpoch;
    _gameInfo.highestBidder = bidder;
    _gameInfo.highestBidPrice = bidPrice;
    _gameInfo.accumulatedFund += ((100 - commission) * bidPrice) / 100;
  }

  function setDefaultBidEpoch(uint256 bidEpoch) external onlyOwner {
    defaultBidEpoch = bidEpoch;
  }

  function setDefaultTimeExtension(uint256 timeExtension) external onlyOwner {
    defaultTimeExtension = timeExtension;
  }

  
  function startGame(uint256 playingPeriod) external onlyOwner {
    if(isPlaying()) revert GameNotEnded();

    _startGame(playingPeriod);
  }

  function endGame() external onlyOwner {
    if(!isPlaying()) revert GameNotPlaying();
    if(!isTimeElapsedToEnd()) revert NotElapsedToEnd(_gameInfo.startedAt + _gameInfo.playingPeriod, block.timestamp);   

    _endGame();
  }

  function bid() external payable {
    if(!isPlaying()) revert GameNotPlaying();
    if (!isTimeElapsedToBid()) revert NotElapsedToBid(_gameInfo.roundEndAt, block.timestamp);
    if(isTimeElapsedToEnd()) revert GameEnded();
    else if (!(msg.value > _gameInfo.highestBidPrice)) revert LowerBid(_gameInfo.highestBidPrice, msg.value);

    _updateGameInfo(_msgSender(), msg.value);

    emit Bid(_msgSender(), msg.value);
  }

  function isTimeElapsedToBid() public view returns(bool) {
    return _gameInfo.roundEndAt < block.timestamp;
  }

  function isTimeElapsedToEnd() public view returns(bool) {
    return _gameInfo.startedAt + _gameInfo.playingPeriod < block.timestamp;
  }

  function withdraw(address sender) external onlyOwner {
    if(isPlaying()) revert GameNotEnded();

    (bool s1, ) = sender.call{value: (address(this).balance)}("");
    if (!s1) revert TransferFailed(sender, address(this).balance);

    emit Transfer(sender, address(this).balance);
  }

  function isPlaying() public view returns(bool) {
    return _gameInfo.isPlaying;
  }
}