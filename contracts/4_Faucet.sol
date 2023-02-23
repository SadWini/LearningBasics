contract Faucet{
    // Give otu ether to anything who asks
    function withdraw(uint withdraw_amount) public{
        
        address owner;
        //Initialize Faucet contract: set owner
        constructor(){
            owner = msg.sender;
        } 
        
        //Limit withdraw amount
        require(withdraw_amount<=100000000000000000);

        //Send the amount to the address that requested it
        payable(msg.sender).transfer(withdraw_amount);
    }

    fallback() external payable{}
}