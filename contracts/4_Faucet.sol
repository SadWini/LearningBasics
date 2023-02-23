pragma solidity >=0.7.0 <0.9.0;

contract owned {

    address owner;

    // Contract constructor: set owner
    constructor() {
        owner = msg.sender;
    }
    
    // Access control modifier
    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }
}

contract mortal is owned{
    //Contract destructor
    function destroy() public onlyOwner{
        selfdestruct(payable(owner));
    }
}

contract Faucet is mortal{
    // Give otu ether to anything who asks
    function withdraw(uint withdraw_amount) public{
        
        //Limit withdraw amount
        require(withdraw_amount<=100000000000000000);

        //Send the amount to the address that requested it
        payable(msg.sender).transfer(withdraw_amount);
    }

    fallback() external payable{}
}