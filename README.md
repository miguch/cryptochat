# CryptoChat

本仓库为 CryptoChat 项目的合约部分，合约代码位于contracts 文件夹下的CryptoChat.sol 文件。

[Web 前端仓库](https://github.com/miguch/cryptochat-web)

[CryptoChat (支持 Ropsten / Rinkeby测试网络)](https://cryptochat.miguch.com)

## CryptoChat 简介

CryptoChat 是一个利用以太坊平台实现的去中心化应用，目的是实现去中心化的安全简讯通信，避免在通信过程中发生的聊天数据泄露，数据被监听、篡改等情况。通过前端的配合将加密后的聊天数据发送至区块链上，利用 RSA 等非对称加密算法使得只有知道私钥的用户才可对聊天信息进行解密，实现安全的简讯通信。

## 合约结构简介

合约中使用如下结构体记录每个用户的数据：

```
struct User {
        string username;
        address addr;
        // Messages stores with a JSON format, containing the RSA cipher
        // text and the digital sign from the sender.
        string[] recv_messages;
        // Since the cipher for receiver and sender is different, 
        // a message has to be stored in two ciphers seperately.
        string[] sent_messages;
        string pubKey;
        bool exists;
        //signature is used to verify the user is created by the right address.
        // web3.personal.sign is used to generate the signature
        string signature;
        //keySig is used to verify the public key, generated by RSA signing.
        string keySig;
    }
```

以下两个映射用于存储用户地址到用户结构体以及用户名到用户地址的映射：

```
mapping(address => User) users;
mapping(string => address) name_addr;
```

合约中的函数介绍如下：

```
function getRecvMsgCount() public view returns (uint)
```

获取调用函数用户已收到的消息数量。

```
function getUserRecvMsg(uint index) public view returns (string)
```

获取用户收到的对应索引的消息。

```
function getSentMsgCount() public view returns (uint)
```

获取用户已发送的消息数

```
function getUserSentMsg(uint index) public view returns (string)
```

获取用户已发送的对应索引的消息

```
function sendMessage(address targetUser, string memory recv_msg, string memory send_msg) public
```

发送消息至目标用户，需要`recv_msg`和`send_msg`两个参数的原因是聊天数据经加密后仅有持有公钥对应私钥的用户才可解密，意味着发送者将无法解密曾经发送的消息，因此需要对消息使用发送者的公钥也进行一次加密后存入区块链才能令发送者和接收者都能重新获取消息。

```
function getCurrentAddr() public view returns (address)
```

获取调用者的地址

```
function hasUser(address target) public view returns (bool)
```

查询地址是否已存在注册用户

```
function checkNewUser(string name, string pubKey) public view returns (uint)
```

检查是否符合注册用户条件，根据情况返回如下错误码：

- 0：OK，符合注册条件
- 1：地址已存在注册用户
- 2：用户名过长
- 3：公钥长度不正确
- 4：用户名已存在

```
function addUser(string name, string pubKey, string sign, string keySig) public
```

注册新用户，传入用户名，公钥，公钥的签名以及使用 Web3对完整用户信息的签名。

```
function getAddrFromName(string username) public view returns (address target, bool status)
```

根据传入的用户名返回对应用户地址

```
function sendEther(address target) public payable
```

向指定用户地址发送以太币，用于实现红包功能

```
function getUsername(address target) public view returns (string memory)
```

返回地址对应的用户名

```
function getUserPublicKey(address target) public view returns (string memory) 
```

返回地址对应用户的公钥

```
function getUserSignature(address target) public view returns (string memory)
```

返回完整用户信息的签名

```
function getUserKeySig(address target) public view returns (string memory)
```

返回用户公钥的签名



## 前端配合

由于 CryptoChat 项目关键的加密解密以及密钥对生成的部分都是运算量较大的任务，不适合也难以在区块链上进行，因此合约的实现逻辑较为简单，而前端则需要承担正确地对数据进行加密、解密、签名以及验证的任务，对于无法解密或者签名验证失败的聊天数据不需要反馈给用户。