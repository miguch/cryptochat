var Migrations = artifacts.require("./Migrations.sol");
const CryptoChat = artifacts.require("./CryptoChat.sol");

module.exports = function(deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(CryptoChat);
};
