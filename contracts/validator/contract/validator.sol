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

// File: @axie/contract-library/contracts/math/SafeMath.sol

pragma solidity ^0.5.2;


library SafeMath {
  function add(uint256 a, uint256 b) internal pure returns (uint256 c) {
    c = a + b;
    require(c >= a);
  }

  function sub(uint256 a, uint256 b) internal pure returns (uint256 c) {
    require(b <= a);
    return a - b;
  }

  function mul(uint256 a, uint256 b) internal pure returns (uint256 c) {
    if (a == 0) {
      return 0;
    }

    c = a * b;
    require(c / a == b);
  }

  function div(uint256 a, uint256 b) internal pure returns (uint256 c) {
    // Since Solidity automatically asserts when dividing by 0,
    // but we only need it to revert.
    require(b > 0);
    return a / b;
  }

  function mod(uint256 a, uint256 b) internal pure returns (uint256 c) {
    // Same reason as `div`.
    require(b > 0);
    return a % b;
  }

  function ceilingDiv(uint256 a, uint256 b) internal pure returns (uint256 c) {
    return add(div(a, b), mod(a, b) > 0 ? 1 : 0);
  }

  function subU64(uint64 a, uint64 b) internal pure returns (uint64 c) {
    require(b <= a);
    return a - b;
  }

  function addU8(uint8 a, uint8 b) internal pure returns (uint8 c) {
    c = a + b;
    require(c >= a);
  }
}

// File: contracts/chain/common/IValidator.sol

pragma solidity ^0.5.2;


contract IValidator {
  event ValidatorAdded(uint256 _id, address indexed validator);
  event ValidatorRemoved(uint256 _id, address indexed validator);
  event ThresholdUpdated(
    uint256 _id,
    uint256 indexed numerator,
    uint256 indexed denominator,
    uint256 previousNumerator,
    uint256 previousDenominator
  );

  function isValidator(address _addr) public view returns (bool);

  function getValidators() public view returns (address[] memory _validators);

  function checkThreshold(uint256 _voteCount) public view returns (bool);
}

// File: contracts/chain/common/Validator.sol

pragma solidity ^0.5.2;




contract Validator is IValidator {
  using SafeMath for uint256;

  mapping(address => bool) validatorMap;
  address[] public validators;
  uint256 public validatorCount;
  uint256 public num;
  uint256 public denom;

  constructor(address[] memory _validators, uint256 _num, uint256 _denom) public {
    validators = _validators;
    validatorCount = _validators.length;

    address _validator;

    for (uint256 _i = 0; _i < validatorCount; _i++) {
      _validator = _validators[_i];
      validatorMap[_validator] = true;
    }

    num = _num;
    denom = _denom;
  }

  function isValidator(address _addr) public view returns (bool) {
    return validatorMap[_addr];
  }

  function getValidators() public view returns (address[] memory _validators) {
    _validators = validators;
  }

  function checkThreshold(uint256 _voteCount) public view returns (bool) {
    return _voteCount.mul(denom) > num.mul(validatorCount);
  }

  function _addValidator(uint256 _id, address _validator) internal {
    require(!validatorMap[_validator]);

    validators.push(_validator);
    validatorMap[_validator] = true;
    validatorCount++;

    emit ValidatorAdded(_id, _validator);
  }

  function _removeValidator(uint256 _id, address _validator) internal {
    require(isValidator(_validator));

    uint256 _index;
    address _lastValidator = validators[validatorCount - 1];

    for (uint256 _i = 0; _i < validatorCount; _i++) {
      if (validators[_i] == _validator) {
        _index = _i;
        break;
      }
    }

    validatorMap[_validator] = false;

    validators[_index] = _lastValidator;
    validators.length--;

    validatorCount--;

    emit ValidatorRemoved(_id, _validator);
  }

  function _updateQuorum(uint256 _id, uint256 _numerator, uint256 _denominator) internal {
    require(_numerator < _denominator);
    uint256 _previousNumerator = num;
    uint256 _previousDenominator = denom;

    num = _numerator;
    denom = _denominator;

    emit ThresholdUpdated(_id, _numerator, _denominator, _previousNumerator, _previousDenominator);
  }
}

// File: @axie/contract-library/contracts/access/HasOperators.sol

pragma solidity ^0.5.2;



contract HasOperators is HasAdmin {
  event OperatorAdded(address indexed _operator);
  event OperatorRemoved(address indexed _operator);

  address[] public operators;
  mapping (address => bool) public operator;

  modifier onlyOperator {
    require(operator[msg.sender], "Not operator");
    _;
  }

  function addOperators(address[] memory _addedOperators) public onlyAdmin {
    address _operator;

    for (uint256 i = 0; i < _addedOperators.length; i++) {
      _operator = _addedOperators[i];

      if (!operator[_operator]) {
        operators.push(_operator);
        operator[_operator] = true;
        emit OperatorAdded(_operator);
      }
    }
  }

  function removeOperators(address[] memory _removedOperators) public onlyAdmin {
    address _operator;

    for (uint256 i = 0; i < _removedOperators.length; i++) {
      _operator = _removedOperators[i];

      if (operator[_operator]) {
        operator[_operator] = false;
        emit OperatorRemoved(_operator);
      }
    }

    uint256 i = 0;

    while (i < operators.length) {
      _operator = operators[i];

      if (!operator[_operator]) {
        operators[i] = operators[operators.length - 1];
        delete operators[operators.length - 1];
        operators.length--;
      } else {
        i++;
      }
    }
  }
}

// File: contracts/chain/sidechain/Acknowledgement.sol

pragma solidity ^0.5.2;





contract Acknowledgement is HasOperators {
  // Acknowledge status
  enum Status {NotApproved, FirstApproved, AlreadyApproved}
  // Mapping from channel => boolean
  mapping(bytes32 => bool) enabledChannels;
  // Mapping from channel => id => validator => hash entry
  mapping(bytes32 => mapping(uint256 => mapping(address => bytes32))) validatorAck;
  // Mapping from channel => id => hash => ack count
  mapping(bytes32 => mapping(uint256 => mapping(bytes32 => uint256))) ackCount;
  // Mapping from channel => id => hash => ack status
  mapping(bytes32 => mapping(uint256 => mapping(bytes32 => uint8))) ackStatus;

  string public constant DEPOSIT_CHANNEL = "DEPOSIT_CHANNEL";
  string public constant WITHDRAWAL_CHANNEL = "WITHDRAWAL_CHANNEL";
  string public constant VALIDATOR_CHANNEL = "VALIDATOR_CHANNEL";

  Validator public validator;

  constructor (address _validator) public {
    addChannel(DEPOSIT_CHANNEL);
    addChannel(WITHDRAWAL_CHANNEL);
    addChannel(VALIDATOR_CHANNEL);
    validator = Validator(_validator);
  }

  function getChannelHash(string memory _name) public view returns (bytes32 _channel) {
    _channel = _getHash(_name);
    _validChannel(_channel);
  }

  function addChannel(string memory _name) public onlyAdmin {
    bytes32 _channel = _getHash(_name);
    enabledChannels[_channel] = true;
  }

  function removeChannel(string memory _name) public onlyAdmin {
    bytes32 _channel = _getHash(_name);
    _validChannel(_channel);
    delete enabledChannels[_channel];
  }

  function updateValidator(address _validator) public onlyAdmin {
    validator = Validator(_validator);
  }

  function acknowledge(string memory _channelName, uint256 _id, bytes32 _hash, address _validator) public onlyOperator returns (Status) {
    bytes32 _channel = getChannelHash(_channelName);
    require(validatorAck[_channel][_id][_validator] == bytes32(0), "the validator has already acknowledged");

    validatorAck[_channel][_id][_validator] = _hash;
    uint8 _status = ackStatus[_channel][_id][_hash];
    uint256 _count = ackCount[_channel][_id][_hash];

    if (validator.checkThreshold(_count + 1)) {
      if (_status == uint8(Status.NotApproved)) {
        ackStatus[_channel][_id][_hash] = uint8(Status.FirstApproved);
      } else {
        ackStatus[_channel][_id][_hash] = uint8(Status.AlreadyApproved);
      }
    }

    ackCount[_channel][_id][_hash]++;

    return Status(ackStatus[_channel][_id][_hash]);
  }

  function hasValidatorAcknowledged(string memory _channelName, uint256 _id, address _validator) public view returns (bool) {
    bytes32 _channel = _getHash(_channelName);
    return validatorAck[_channel][_id][_validator] != bytes32(0);
  }

  function _getHash(string memory _name) internal pure returns (bytes32 _hash) {
    _hash = keccak256(abi.encode(_name));
  }

  function _validChannel(bytes32 _hash) internal view {
    require(enabledChannels[_hash], "invalid channel");
  }
}

// File: contracts/chain/sidechain/SidechainValidator.sol

pragma solidity ^0.5.2;





/**
 * @title Validator
 * @dev Simple validator contract
 */
contract SidechainValidator is Validator {
  Acknowledgement public acknowledgement;

  modifier onlyValidator() {
    require(isValidator(msg.sender));
    _;
  }

  constructor(
    address _acknowledgement,
    address[] memory _validators,
    uint256 _num,
    uint256 _denom
  ) Validator(_validators, _num, _denom) public {
    acknowledgement = Acknowledgement(_acknowledgement);
  }

  function addValidator(uint256 _id, address _validator) external onlyValidator {
    bytes32 _hash = keccak256(abi.encode("addValidator", _validator));

    Acknowledgement.Status _status = acknowledgement.acknowledge(_getAckChannel(), _id, _hash, msg.sender);
    if (_status == Acknowledgement.Status.FirstApproved) {
      _addValidator(_id, _validator);
    }
  }

  function removeValidator(uint256 _id, address _validator) external onlyValidator {
    require(isValidator(_validator));

    bytes32 _hash = keccak256(abi.encode("removeValidator", _validator));

    Acknowledgement.Status _status = acknowledgement.acknowledge(_getAckChannel(), _id, _hash, msg.sender);
    if (_status == Acknowledgement.Status.FirstApproved) {
      _removeValidator(_id, _validator);
    }
  }

  function updateQuorum(uint256 _id, uint256 _numerator, uint256 _denominator) external onlyValidator {
    bytes32 _hash = keccak256(abi.encode("updateQuorum", _numerator, _denominator));

    Acknowledgement.Status _status = acknowledgement.acknowledge(_getAckChannel(), _id, _hash, msg.sender);
    if (_status == Acknowledgement.Status.FirstApproved) {
      _updateQuorum(_id, _numerator, _denominator);
    }
  }

  function _getAckChannel() internal view returns (string memory) {
    return acknowledgement.VALIDATOR_CHANNEL();
  }
}
