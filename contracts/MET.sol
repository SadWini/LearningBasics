pragma solidity >=0.7.0 <0.9.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract METoken is ERC20 {
    string private constant NAME = "Mastering Ethereum Token";
    string private constant SYMBOL = "MET";
    uint8 private constant DECIMALS = 2;
    uint256 private constant INITIAL_SUPPLY = 2100000000;

    constructor() public ERC20(NAME, SYMBOL){
        _mint(_msgSender(), INITIAL_SUPPLY);
    }
    
        function decimals() override public view returns (uint8) {
        return DECIMALS;
    }
    
}