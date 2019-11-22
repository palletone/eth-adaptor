pragma solidity ^0.4.24;

/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
library LSafeMath {

  /**
  * @dev Multiplies two numbers, throws on overflow.
  */
  function mul(uint256 a, uint256 b) internal pure returns (uint256) {
    if (a == 0) {
      return 0;
    }
    uint256 c = a * b;
    require(c / a == b);
    return c;
  }

  /**
  * @dev Integer division of two numbers, truncating the quotient.
  */
  function div(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b > 0); // Solidity automatically throws when dividing by 0
    uint256 c = a / b;
    // assert(a == b * c + a % b); // There is no case in which this doesn't hold
    return c;
  }

  /**
  * @dev Substracts two numbers, throws on overflow (i.e. if subtrahend is greater than minuend).
  */
  function sub(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b <= a);
    return a - b;
  }

  /**
  * @dev Adds two numbers, throws on overflow.
  */
  function add(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a + b;
    require(c >= a);
    return c;
  }
}

/**
 * @title Token
 * @dev Token interface necessary for working with tokens within the exchange contract.
 */
contract IToken {
  /// @return total amount of tokens
  function totalSupply() public constant returns (uint256 supply) {}

  /// @param _owner The address from which the balance will be retrieved
  /// @return The balance
  function balanceOf(address _owner) public constant returns (uint256 balance) {}

  /// @notice send `_value` token to `_to` from `msg.sender`
  /// @param _to The address of the recipient
  /// @param _value The amount of token to be transferred
  /// @return Whether the transfer was successful or not
  function transfer(address _to, uint256 _value) public returns (bool success) {}

  /// @notice send `_value` token to `_to` from `_from` on the condition it is approved by `_from`
  /// @param _from The address of the sender
  /// @param _to The address of the recipient
  /// @param _value The amount of token to be transferred
  /// @return Whether the transfer was successful or not
  function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {}

  /// @notice `msg.sender` approves `_addr` to spend `_value` tokens
  /// @param _spender The address of the account able to transfer the tokens
  /// @param _value The amount of wei to be approved for transfer
  /// @return Whether the approval was successful or not
  function approve(address _spender, uint256 _value) public returns (bool success) {}

  /// @param _owner The address of the account owning tokens
  /// @param _spender The address of the account able to transfer the tokens
  /// @return Amount of remaining tokens allowed to spent
  function allowance(address _owner, address _spender) public constant returns (uint256 remaining) {}

  event Transfer(address indexed _from, address indexed _to, uint256 _value);
  event Approval(address indexed _owner, address indexed _spender, uint256 _value);

  uint public decimals;
  string public name;
}

contract PalletOne {
  using LSafeMath for uint;

  event Deposit(address token, address user, uint amount, bytes redeem);
  event Withdraw(address token, address user, bytes redeem, address recver, uint amount, uint confirmvalue, string state);

  address public admin; //the admin address

  struct Multisig {
      uint balance;
      uint nonece;
  }
  
  mapping (address => mapping (bytes32 => Multisig)) public tokens; //mapping of token addresses to mapping of account balances (token=0 means Ether)

  /// This is a modifier for functions to check if the sending user address is the same as the admin user address.
  modifier isAdmin() {
    require(msg.sender == admin);
    _;
  }

  constructor(address admin_) public {
    admin = admin_;
  }

  function splitaddress(address[] owners, bytes redeem) private pure {
    address addra;
    address addrb;
    address addr1;
    
    if (redeem.length < 60) {
      return ;
    }

    assembly {
      addra := mload(add(redeem, 20))
      addrb := mload(add(redeem, 40))
      addr1 := mload(add(redeem, 60))
    }
    owners[0] = addra;
    owners[1] = addrb;
    owners[2] = addr1;
  }

  function setoneconfirm(uint8[] addrconfirms, address addr, address[] owners) private pure {
      for (uint8 i=0; i < 3;i++) {
          if (addr != owners[i]) {
              continue;
          }
          addrconfirms[i] = 1;
          break;
      }
  }

  function setallconfirms(uint8[] addrconfirms, bytes32 tranhash, address[] owners, bytes sigstr1, bytes sigstr2) private pure {
    address addr = 0;
    if (sigstr1.length != 0) {
        addr = getaddr(tranhash, sigstr1);
        setoneconfirm(addrconfirms, addr, owners);
    }
    if (sigstr2.length != 0) {
        addr = getaddr(tranhash, sigstr2);
        setoneconfirm(addrconfirms, addr, owners);
    }
  }

  function calconfirm(uint8[] addrconfirms) private pure returns (uint8) {
    uint8[] memory weights = new uint8[](3);
    weights[0] = 1;
    weights[1] = 1;
    weights[2] = 1;

    uint8 confirms = 0;
    for (uint8 i=0;i<3;i++) {
        confirms += addrconfirms[i]*weights[i];
    }
    return confirms;
  }

  function getconfirm(address[] owners, bytes32 tranhash, bytes sigstr1, bytes sigstr2) private pure returns (uint8)  {
    uint8[] memory addrconfirms = new uint8[](3);
    
    setallconfirms(addrconfirms, tranhash, owners, sigstr1, sigstr2);
    
    uint8 confirms = 0;
    confirms = calconfirm(addrconfirms);
    return confirms;  
  }

  function withdraw(bytes redeem, address recver, uint amount, uint nonece, bytes sigstr1, bytes sigstr2) public {
    bytes32 hash = keccak256(abi.encodePacked(redeem));
    require(tokens[0][hash].balance >= amount);
    require(tokens[0][hash].nonece == nonece.sub(1));
    
    address[] memory owners= new address[](6);
    splitaddress(owners, redeem);
    
    uint8 confirms = 0;
    bytes32 tranhash = keccak256(abi.encodePacked(redeem, recver, address(this), amount, nonece));
    confirms = getconfirm(owners, tranhash, sigstr1, sigstr2);
    
    require(confirms >= 2);
    tokens[0][hash].balance = tokens[0][hash].balance.sub(amount);
    tokens[0][hash].nonece = tokens[0][hash].nonece.add(1);
    recver.transfer(amount);
    emit Withdraw(0, msg.sender, redeem, recver, amount, confirms, "withdraw");
  }

  function subamount(address token, bytes32 hash, uint amount) private {
    tokens[token][hash].balance = tokens[token][hash].balance.sub(amount);
    tokens[token][hash].nonece = tokens[token][hash].nonece.add(1); 
  }

  function withdrawtoken(address token, bytes redeem, address recver, uint amount, uint nonece, bytes sigstr1, bytes sigstr2) public {
    bytes32 hash = keccak256(abi.encodePacked(redeem));
    require(tokens[token][hash].balance >= amount);
    require(tokens[token][hash].nonece == nonece.sub(1));
    
    address[] memory owners= new address[](6);
    splitaddress(owners, redeem);

    uint8 confirms = 0;
    bytes32 tranhash = keccak256(abi.encodePacked(token, redeem, recver, address(this), amount, nonece));
    confirms = getconfirm(owners, tranhash, sigstr1, sigstr2);

    require(confirms >= 2);
    subamount(token, hash, amount);
    require(IToken(token).transfer(recver, amount));
    emit Withdraw(token, msg.sender, redeem, recver, amount, confirms, "withdrawtoken");
  }

  function getmultisig(address addr, bytes redeem) public view returns(uint, uint) {
    bytes32 hash = keccak256(abi.encodePacked(redeem));
    return (tokens[addr][hash].balance, tokens[addr][hash].nonece);
  }

  function my_eth_bal() public view returns(uint) {
      return address(this).balance;
  }

  function recover(bytes32 hash, bytes sig) private pure returns (address) {
    bytes32 r;
    bytes32 s;
    uint8 v;

    // Check the signature length
    if (sig.length != 65) {
      return (address(0));
    }

    // Divide the signature in r, s and v variables
    // ecrecover takes the signature parameters, and the only way to get them
    // currently is to use assembly.
    // solium-disable-next-line security/no-inline-assembly
    assembly {
      r := mload(add(sig, 32))
      s := mload(add(sig, 64))
      v := byte(0, mload(add(sig, 96)))
    }

    // Version of signature should be 27 or 28, but 0 and 1 are also possible versions
    if (v < 27) {
      v += 27;
    }

    // If the version is correct return the signer address
    if (v != 27 && v != 28) {
      return (address(0));
    } else {
      // solium-disable-next-line arg-overflow
      return ecrecover(hash, v, r, s);
    }
  }
  function getaddr(bytes32 tranhash, bytes sigstr) private pure returns (address) {
    return recover(tranhash, sigstr);
  }

  
  function getMapPtnAddr(address addr) external view returns (string){
    if (addrmap[addr] == address(0)) {
        return "";
    }
    return encodeBase58(addrmap[addr]);
  }
  function getMapEthAddr(address ptnAddr) external view returns (address){
    return addrmapPTN[ptnAddr];
  }


  function bytesConcat(bytes _b) internal returns (string){
    string memory ret = new string(2 + _b.length);
    bytes memory bret = bytes(ret);
    uint k = 0;
    bret[k++] = byte('P');
    bret[k++] = byte('1');
    for (uint8 i = 0; i < _b.length; i++) bret[k++] = _b[i];
    return string(ret);
  }
  function toBytes(uint160 x) internal returns (bytes) {
    bytes20 a = bytes20(x);
    bytes memory b = new bytes(25);
    b[0] = byte(0);
    for (uint8 i=0; i < 20; i++) {
        b[1+i] = a[i];
    }
    bytes32 cksum = sha256(sha256(0,a));
    for (uint8 j=0; j < 4; j++) {
        b[21+j] = cksum[j];
    }
    return b;
  }
  function encodeBase58(address addrHex) constant returns (string) {
    uint160 a = uint160(addrHex);
    string memory result = bytesConcat(toBase58(toBytes(a)));
    return result;
  }
  function toBase58(bytes source) internal returns (bytes) {
    if (source.length == 0) {
        return "";
    }
    uint8[] memory digits = new uint8[](40); //TODO: figure out exactly how much is needed
    digits[0] = 0;
    uint8 digitlength = 1;
    for (uint8 i = 0; i<source.length; ++i) {
        uint carry = uint8(source[i]);
        for (uint8 j = 0; j<digitlength; ++j) {
            carry += uint(digits[j]) * 256;
            digits[j] = uint8(carry % 58);
            carry = carry / 58;
        }
        
        while (carry > 0) {
            digits[digitlength] = uint8(carry % 58);
            digitlength++;
            carry = carry / 58;
        }
    }

    return toAlphabet(reverse(truncate(digits, digitlength)));
  }
  bytes constant ALPHABET = '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz';
  function toAlphabet(uint8[] indices) internal returns (bytes) {
    bytes memory output = new bytes(indices.length);
    for (var i = 0; i<indices.length; i++) {
        output[i] = ALPHABET[indices[i]];
    }
    return output;
  }
  function reverse(uint8[] input) internal returns (uint8[]) {
    uint8[] memory output = new uint8[](input.length);
    for (var i = 0; i<input.length; i++) {
        output[i] = input[input.length-1-i];
    }
    return output;
  }
  function truncate(uint8[] array, uint8 length) internal returns (uint8[]) {
    uint8[] memory output = new uint8[](length);
    for (var i = 0; i<length; i++) {
        output[i] = array[i];
    }
    return output;
  }


  function suicideto(address addr) public isAdmin {
      selfdestruct(addr);
  }

  function () payable {
      // can receive eth
  }
}