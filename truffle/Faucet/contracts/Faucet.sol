pragma solidity >=0.7.0 <0.9.0;

contract owned {

    address owner;

    // Contract constructor: set owner
    constructor() {
        owner = msg.sender;
    }
    
    // Access control modifier
    modifier onlyOwner {
        require(msg.sender == owner, 
                "Only the contract owner can call this function");
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
    event Withdrawal(address indexed to, uint amount);
    event Deposit(address indexed from, uint amount);
    // Give out ether to anyone who asks
    function withdraw(uint withdraw_amount) public{
        //Limit withdraw amount
        require(withdraw_amount<=100000000000000000);
        require(address(this).balance >= withdraw_amount, 
                "Insufficient balance in faucet for withdrawl request");
        //Send the amount to the address that requested it
        payable(msg.sender).transfer(withdraw_amount);
        emit Withdrawal(msg.sender, withdraw_amount);
    }
    //Accept any incoming amount
    fallback() external payable{
        emit Deposit(msg.sender, msg.value);
    }
}
