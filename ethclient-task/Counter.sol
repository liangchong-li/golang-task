// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

contract Counter {
    uint public number = 0;

    function count() public returns (uint) {
        return number++;
    }
}
