pragma solidity >=0.4.26 <0.9.0;

pragma experimental ABIEncoderV2;

// Should fix yet —> OpenZeppelin -> ^0.8.0
// import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/access/Ownable.sol";
// import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/Context.sol";

// import "https://gist.githubusercontent.com/ageyev/779797061490f5be64fb02e978feb6ac/raw/dbd08a23531b784186bc4bfffcadefcb23db88e9/StringsAndBytes.sol";

contract Auther {
    mapping(address => bool) nodes;

    address private chairman;

    constructor() public {
        chairman = msg.sender;
        nodes[chairman] = true;
    }

    modifier onlyChairman() {
        if (msg.sender != chairman) revert();
        _;
    }

    modifier onlyRegisted() {
        if (!nodes[msg.sender]) revert();
        _;
    }

    function register(address addr) public onlyChairman {
        require(!nodes[addr], "The node has been registed!");

        nodes[addr] = true;
    }
}

contract DDoS is Auther {
    event msgConn(uint eventID, address addr, string typeName, string name);

    struct Rconn {
        address commiterAddr;
        string targetIP;
        uint speed;
        string timestamp;
        bool isDDoS;

        bool isUsed;
    }

    mapping(uint => Rconn) rconns;

    function insertRconn(uint eventID, address addr, uint sp, string[] memory rconn) public onlyRegisted {
        require(!rconns[eventID].isUsed , "This eventID has been used , can't insert into rconns");

        rconns[eventID].commiterAddr = addr;
        rconns[eventID].targetIP = rconn[0];
        rconns[eventID].speed = sp;
        rconns[eventID].timestamp = rconn[1];

        rconns[eventID].isDDoS = false;
        rconns[eventID].isUsed = true;

        emit msgConn(eventID, addr, "insert", "Rconn");
    }

    function indexRconn(uint eventID) public view onlyRegisted returns(Rconn memory) {
        return rconns[eventID];
    }

    mapping(uint => Rconn) rddoses;
    uint[] public events;

    function insertRddos(uint eventID) onlyRegisted public {
        require(!rconns[eventID].isUsed, "This eventID has been used , can't insert into rddoses!");

        events.push(eventID);
        rddoses[eventID] = rconns[eventID];
        delete rconns[eventID];

        emit msgConn(eventID, rddoses[eventID].commiterAddr, "insert", "Rddos");
    }
}
