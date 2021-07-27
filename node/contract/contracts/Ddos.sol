pragma solidity >=0.5.0 <0.9.0;
pragma experimental ABIEncoderV2;

contract trafficStation {
    struct upchainTrafficInfo {
        uint    trafficID;
        address sourceAddr;
        string  trafficInfo;
    }

    struct voteInfo {
        address sourceAddr;
        address voteAddr;
        uint    trafficID;
        bool    state;
    }

    event trafficTrans(uint trafficID, address sourceAddr, string trafficInfo);
    event voteTrans(address sourceAddr, address voteAddr, uint trafficID, bool state);

    //  send from: address, traffic id: uint, vote number: uint
    mapping(address => mapping(uint => uint)) voting;
    //  send from: address, traffic id: uint, traffic info: upchainTrafficInfo
    mapping(address => mapping(uint => upchainTrafficInfo)) traffics;
    /*

                node1  -------------traffic info----------------> node~
                node1  <-------------vote info---------------- node~
                node1  -------------traffic info----------------> node~

    */

    function emitTrafficTrans(upchainTrafficInfo memory uti) public {
        require(address(0x0) != uti.sourceAddr);
        voting[uti.sourceAddr][uti.trafficID] = 0;
        emit trafficTrans(uti.trafficID, uti.sourceAddr, uti.trafficInfo);
    }

    function emitVoteTrans(voteInfo memory vi) public {
        require(vi.voteAddr != address(0x0));
        if (vi.state == true) {
            voting[vi.sourceAddr][vi.trafficID] += 1;
        }
        emit voteTrans(vi.sourceAddr, vi.voteAddr, vi.trafficID, vi.state);
    }

    /*
                traffic sender source: send traffic info, get vote number -> judge, resend ddos traffic id, get price: ERC20
    */

    function getVoteInfoByID(uint trafficID) public view returns(uint) {
        return voting[msg.sender][trafficID];
    }

    function getTrafficInfoByID(uint trafficID) public view returns (upchainTrafficInfo memory) {
        return traffics[msg.sender][trafficID];
    }
}