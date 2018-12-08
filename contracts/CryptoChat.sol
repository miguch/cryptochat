pragma solidity ^0.4.24;



contract CryptoChat {
    
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
        string signature;
    }
    mapping(address => User) users;
    mapping(string => address) name_addr;

    function getRecvMsgCount(address userAddr) public view returns (uint){
        require(users[userAddr].exists);
        return users[userAddr].recv_messages.length;
    }
    
    function getUserRecvMsg(address userAddr, uint index) public view returns (string) {
        require(users[userAddr].exists && users[userAddr].recv_messages.length > index);
        return users[userAddr].recv_messages[index];
    }
    
    function getSentMsgCount() public view returns (uint){
        require(users[msg.sender].exists);
        return users[msg.sender].sent_messages.length;
    }
    
    function getUserSentMsg(uint index) public view returns (string) {
        require(users[msg.sender].exists && users[msg.sender].sent_messages.length > index);
        return users[msg.sender].sent_messages[index];
    }
    
    function sendMessage(address targetUser, string memory recv_msg, string memory send_msg) public{
        require(users[targetUser].exists && users[msg.sender].exists);
        users[targetUser].recv_messages.push(recv_msg);
        users[msg.sender].sent_messages.push(send_msg);
    }
    
    function getCurrentAddr() public view returns (address) {
        return msg.sender;
    }
    
    function hasUser(address target) public view returns (bool) {
        return users[target].exists;
    }
    
    // Only add user when this function returns 0.
    function checkNewUser(string name, string pubKey) public view returns (uint) {
        if (users[msg.sender].exists) {
            //1 - represent user already exists.
            return 1;
        }
        if (bytes(name).length > 20) {
            //2 - represent username too long
            return 2;
        }
        if (bytes(pubKey).length > 300) {
            //3 - represent the pub
            return 3;
        }
        if (name_addr[name] != 0x0) {
            //4 - username already exists.
            return 4;
        }
        //OK
        return 0;
    }
    
    //Returns ERROR code to represent what kinds of error were encountered.
    function addUser(string name, string pubKey, string sign) public  {
        require(!users[msg.sender].exists && bytes(name).length <= 20 && bytes(pubKey).length <= 300 && name_addr[name] == 0x0);
        users[msg.sender] = User({
            username: name,
            addr: msg.sender,
            pubKey: pubKey,
            recv_messages: new string[](0),
            sent_messages: new string[](0),
            exists: true,
            signature: sign
        });
        name_addr[name] = msg.sender;
    }

    function getAddrFromName(string username) public view returns (address) {
        return name_addr[username];
    }
    
    function sendEther(address target) public payable {
        require(users[target].exists && users[msg.sender].exists);
        target.transfer(msg.value);
    }
    
    function getUsername(address target) public view returns (string memory) {
        require(users[target].exists);
        return users[target].username;
    }
    

}