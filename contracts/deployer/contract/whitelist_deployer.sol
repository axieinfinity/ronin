
// File: @axie/contract-library/contracts/access/HasAdmin.sol

pragma solidity ^0.5.2;


contract HasAdmin {
  event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin);
  event AdminRemoved(address indexed _oldAdmin);

  address public admin;

  modifier onlyAdmin {
    require(msg.sender == admin);
    _;
  }

  constructor() internal {
    admin = msg.sender;
    emit AdminChanged(address(0), admin);
  }

  function changeAdmin(address _newAdmin) external onlyAdmin {
    require(_newAdmin != address(0));
    emit AdminChanged(admin, _newAdmin);
    admin = _newAdmin;
  }

  function removeAdmin() external onlyAdmin {
    emit AdminRemoved(admin);
    admin = address(0);
  }
}

// File: ethereum-contracts/contracts/others/WhitelistDeployer.sol

pragma solidity ^0.5.2;


contract WhitelistDeployer is HasAdmin {
  event AddressWhitelisted(address indexed _address, bool indexed _status);
  event WhitelistAllChange(bool indexed _status);

  mapping (address => bool) public whitelisted;
  bool public whitelistAll;

  constructor() public {}

  function whitelist(address _address, bool _status) external onlyAdmin {
    whitelisted[_address] = _status;
    emit AddressWhitelisted(_address, _status);
  }

  function whitelistAllAddresses(bool _status) external onlyAdmin {
    whitelistAll = _status;
    emit WhitelistAllChange(_status);
  }

  function isWhitelisted(address _address) external view returns (bool) {
    return whitelistAll || whitelisted[_address];
  }
}
