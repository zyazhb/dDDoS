pragma solidity >=0.4.26 <0.9.0;

pragma experimental ABIEncoderV2;

contract Auther {
    /*
    **  address: Fringe node -> Account address
    **  bool: isRegist
    */
    mapping(address => bool) nodes;
    address[] public nodesAddr;
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
        nodesAddr.push(addr);
    }
}

contract DDoS is Auther {
    event msgConn(uint eventID, address addr, string typeName, string name);

    struct Rconn {
        address commiterAddr;
        string targetIP;
        uint speed;
        string timestamp;
        string descrptions;

        bool isUsed;
    }

    // Rconn -> Actions

    mapping(uint => Rconn) rconns;

    function insertRconn(uint eventID, address addr, uint sp, string[] memory rconn) public onlyRegisted {
        require(!rconns[eventID].isUsed , "This eventID has been used , can't insert into rconns");

        rconns[eventID].commiterAddr = addr;
        rconns[eventID].speed        = sp;
        rconns[eventID].targetIP     = rconn[0];
        rconns[eventID].timestamp    = rconn[1];
        rconns[eventID].descrptions  = rconn[2];

        rconns[eventID].isUsed = true;

        emit msgConn(eventID, addr, "insert", "Rconn");
    }

    function indexRconn(uint eventID) public view onlyRegisted returns(uint) {
        return rconns[eventID].speed;
    }

    // Rddos -> Actions

    mapping(uint => Rconn) rddoses;
    uint[] public events;

    function insertRddos(uint eventID) onlyRegisted public {
        require(!rconns[eventID].isUsed, "This eventID has been used , can't insert into rddoses!");

        events.push(eventID);
        rddoses[eventID] = rconns[eventID];
        delete rconns[eventID];

        emit msgConn(eventID, rddoses[eventID].commiterAddr, "insert", "Rddos");
    }

    // Valid Checker

    mapping(address => bool) valids;

    constructor() public {
        for (uint i = 0; i < nodesAddr.length; i++) {
            valids[nodesAddr[i]] = false;
        }
    }

    function reCheckDDos(uint rconnSpeed, uint threshold) onlyRegisted public {
        require(nodes[msg.sender], "The node's address is not registed!");

        if(rconnSpeed >= threshold) {
            valids[msg.sender] = true;
        }
    }

    modifier onlyValidAttack() {
        for (uint i = 0; i < nodesAddr.length; i++) {
            if (valids[nodesAddr[i]] == false) {
                revert();
            }
        }
        _;
    }
}
