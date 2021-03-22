pragma solidity ^0.4.26;

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

    function register(address addr) public onlyChairman returns(bool) {
        require(!nodes[addr], "The node has been registed!");

        nodes[addr] = true;
    }
}

contract DDoS is Auther {
    event msgConn(bytes32 eventID, address addr, string typeName, string name);

    struct Rconn {
        address commiterAddr;
        string fromIP;
        string targetIP;
        uint speed;
        string timestamp;
        string description;
        bool isDDoS;

        bool isUsed;
    }

    mapping(bytes32 => Rconn) rconns;

    function insertRconn(bytes32 eventID, address addr, uint sp, bool ddos, string[] rconn) public onlyRegisted {
        require(!rconns[eventID].isUsed , "This eventID has been used , can't insert into rconns");

        rconns[eventID].commiterAddr = addr;
        rconns[eventID].fromIP = rconn[0];
        rconns[eventID].targetIP = rconn[1];
        rconns[eventID].speed = sp;
        rconns[eventID].timestamp = rconn[3];
        rconns[eventID].description = rconn[4];

        rconns[eventID].isDDoS = ddos;

        emit msgConn(eventID, addr, "insert", "Rconn");
    }

    function indexRconn(bytes32 eventID) public view returns(Rconn) {
        return rconns[eventID];
    }

    mapping(bytes32 => Rconn) rddoses;
    bytes32[] public events;

    function insertRddos(bytes32 eventID) public {
        require(!rconns[eventID].isUsed, "This eventID has been used , can't insert into rddoses!");

        events.push(eventID);
        rddoses[eventID] = rconns[eventID];
        delete rconns[eventID];
    }
}
