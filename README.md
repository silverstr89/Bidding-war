# Game Mechanics

Users try to out bid each other. The instructions for each game are as below:

- The bidding starts with 0 NRG and 60 minutes(Default value) countdown.
- Every next bid has to be higher than the previous one.
- For each bid, the time extends for 10 minutes(Default value).
- When the time runs out, the winner is the last bidder.
- The reward is all the funds that were accumulated minus the service fee(Game commissions).

*Each successful bid will take 5% commission in order to fund the game operations.*

# Task
- Serves endpoints for an external web app.
- Logs all events related to the smart contract.
- Pays out the winners at the end of each game.
- Closes the latest game when the timer has expired.


# Contract
## How To Test
- Clone the repo.
- Install node.js and yarn
- Run `yarn hardhat test`

## Directory
- test script
  /BiddingWar/test
- contract abi and address
  /BiddingWar/deployments/testEnergi/BiddingWar.json

# API
## How To Install
- Clone the repo.
- Run `go install` and `npm i` to install all app dependencies.
- Update env variables or leave as is to run in prod by default(At least the operator private key must be set).
- Run `go run .` or `npm run start` and `npm run dev` to run the app in prod and dev respectively.

# Routes


- Clone the repo.
- Run `go install` and `npm i` to install all app dependencies.
- Update env variables or leave as is to run in prod by default(At least the operator private key must be set).
- Run `go run .` or `npm run start` and `npm run dev` to run the app in prod and dev respectively.
