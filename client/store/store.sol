// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

contract IPFSStorage{
    string hash;
    
    mapping (address => string) private entries;
    
    function setEntry(string memory _hash) public {
        entries[msg.sender] = _hash;
    }
    
    function getEntry(address _address) public view returns(string memory entry) {
        return entries[_address];
    }
    
}