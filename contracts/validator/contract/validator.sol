pragma solidity ^0.5.2;

contract SkyMavisValidator {
  address[] validators;
  uint256 validatorCount;

  constructor() public {
  }

  function addValidator(address _validator) external {
    validators.push(_validator);
    validatorCount++;
  }

  function removeValidator(uint256 _index) external {
    require(_index < validatorCount);
    for (uint256 _i = _index; _i + 1 < validatorCount; _i++) {
      validators[_i] = validators[_i + 1];
    }

    validatorCount--;
    validators.length--;
  }

  function getValidators() public view returns(address[] memory _validators) {
    _validators = validators;
  }
}
