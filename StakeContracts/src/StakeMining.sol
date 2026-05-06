// SPDX-Lincense-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/**
 * @title 可升级质押挖矿合约
 */

contract StakeMining is Initializable, OwnableUpgradeable, UUPSUpgradeable, PausableUpgradeable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    // ==== 状态变量 ====
    IERC20 public stakeToken;  // 质押代币
    IERC20 public rewardToken; // 收益代币
    uint256 public rewardRate; // 每秒收益速率

    uint256 public stakeMinAmount;  // 最小质押金额
    uint256 public stakeMaxAmount;  // 最大质押金额

    mapping(address => bool) public blacklist; // 黑名单地址

    // 用户质押信息
    struct StakeInfo {
        uint256 amount;      // 质押金额
        uint256 rewardDebt;  // 已经领取的收益
        uint256 lastUpdate;  // 上次更新质押信息的时间戳
    }
    mapping(address => StakeInfo) public userStakes;

    // ==== 事件 ====
    event Staked(address indexed user, uint256 amount, uint256 time);
    event Unstaked(address indexed user, uint256 amount, uint256 time);
    event RewardClaimed(address indexed user, uint256 reward, uint256 time);
    event BlacklistAdded(address indexed account);
    event BlacklistRemoved(address indexed account);
    event StakeLimitsUpdated(uint256 min, uint256 max);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    // ==== 初始化函数 ====
    function initialize(
        address _stakeToken,
        address _rewardToken,
        uint256 _rewardRate,
        uint256 _stakeMinAmount,
        uint256 _stakeMaxAmount
    ) public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
        __Pausable_init();

        stakeToken = IERC20(_stakeToken);
        rewardToken = IERC20(_rewardToken);
        rewardRate = _rewardRate;
        stakeMinAmount = _stakeMinAmount;
        stakeMaxAmount = _stakeMaxAmount;
    }

    // 仅所有者可升级
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    // 紧急暂停
    function pause() external onlyOwner {
        _pause();
    }
    function unpause() external onlyOwner {
        _unpause();
    }

    // 添加/移除黑名单
    function addBlacklist(address _account) external onlyOwner {
        blacklist[_account] = true;
        emit BlacklistAdded(_account);
    }
    function removeBlacklist(address _account) external onlyOwner {
        blacklist[_account] = false;
        emit BlacklistRemoved(_account);
    }

    // 黑名单校验修饰器
    modifier notBlacklisted() {
        require(!blacklist[msg.sender], "account in blacklist");
        _;
    }

    // 更新质押限制
    function updateStakeLimits(uint256 _min, uint256 _max) external onlyOwner {
        require(_min < _max, "invalid limits");
        stakeMinAmount = _min;
        stakeMaxAmount = _max;
        emit StakeLimitsUpdated(_min, _max);
    }


    // ==================== 核心业务逻辑（全防护） ====================
    // 质押函数
    function stake(uint256 _amount) external whenNotPaused notBlacklisted nonReentrant {
        require(_amount > 0, "stake amount must be greater than zero");
        require(_amount >= stakeMinAmount && _amount <= stakeMaxAmount, "stake amount out of limits");
        StakeInfo storage stakeInfo = userStakes[msg.sender];
        _updateReward(msg.sender);

        stakeToken.safeTransferFrom(msg.sender, address(this), _amount);
        stakeInfo.amount += _amount;
        stakeInfo.lastUpdate = block.timestamp;

        emit Staked(msg.sender, _amount, block.timestamp);
    }

    // 赎回函数
    function unstake(uint256 _amount) external whenNotPaused notBlacklisted nonReentrant {
        StakeInfo storage stakeInfo = userStakes[msg.sender];
        require(stakeInfo.amount >= _amount, "unstake amount exceeds staked amount");

        _updateReward(msg.sender);
        stakeInfo.amount -= _amount;
        stakeToken.safeTransfer(msg.sender, _amount);
        stakeInfo.lastUpdate = block.timestamp;

        emit Unstaked(msg.sender, _amount, block.timestamp);
    }

    // 领取收益函数
    function claimReward() external whenNotPaused notBlacklisted nonReentrant {
        _updateReward(msg.sender);
        StakeInfo storage stakeInfo = userStakes[msg.sender];
        uint256 reward = stakeInfo.rewardDebt;
        require(reward > 0, "no rewards to claim");

        stakeInfo.rewardDebt = 0;
        rewardToken.safeTransfer(msg.sender, reward);

        emit RewardClaimed(msg.sender, reward, block.timestamp);
    }

    // ==================== 内部工具函数 ====================
    function _updateReward(address _user) internal {
        StakeInfo storage stakeInfo = userStakes[_user];
        if (stakeInfo.amount == 0) return;

        uint256 timeDiff = block.timestamp - stakeInfo.lastUpdate;
        uint256 reward = stakeInfo.amount * rewardRate * timeDiff;
        stakeInfo.rewardDebt += reward;
        stakeInfo.lastUpdate = block.timestamp;
    }

    function getPendingReward(address _user) external view returns (uint256) {
        StakeInfo storage stakeInfo = userStakes[_user];
        uint256 timeDiff = block.timestamp - stakeInfo.lastUpdate;
        return stakeInfo.rewardDebt + (stakeInfo.amount * rewardRate * timeDiff);
    }

    function setRewardRate(uint256 _rate) external onlyOwner {
        rewardRate = _rate;
    }
}
