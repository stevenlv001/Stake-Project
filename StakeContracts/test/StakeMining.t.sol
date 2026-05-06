// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "../src/StakeMining.sol";

contract TestToken is ERC20 {
    constructor(string memory name, string memory symbol) ERC20(name, symbol) {}

    function mint(address to, uint256 amount) public {
        _mint(to, amount);
    }
}

contract StakeMiningTest is Test {
    TestToken public stakeToken;
    TestToken public rewardToken;
    StakeMining public stakeMining;
    ProxyAdmin public proxyAdmin;
    TransparentUpgradeableProxy public proxy;

    address public owner = address(0x1234);
    address public user1 = address(0x5678);
    address public user2 = address(0xABCD);
    address public blacklistedUser = address(0xEF0123456789ABCDEF);

    uint256 public constant INITIAL_SUPPLY = 1000000 * 10**18;
    uint256 public constant REWARD_SUPPLY = 100000000 * 10**18; // 增加奖励代币供应
    uint256 public constant REWARD_RATE = 100;
    uint256 public constant STAKE_MIN = 100 * 10**18;
    uint256 public constant STAKE_MAX = 10000 * 10**18;

    function setUp() public {
        // 部署测试代币
        stakeToken = new TestToken("Stake Token", "STK");
        rewardToken = new TestToken("Reward Token", "RWD");

        // 铸造代币给测试用户
        stakeToken.mint(user1, INITIAL_SUPPLY);
        stakeToken.mint(user2, INITIAL_SUPPLY);
        stakeToken.mint(blacklistedUser, INITIAL_SUPPLY);
        rewardToken.mint(address(this), REWARD_SUPPLY);

        // 部署质押挖矿合约实现
        StakeMining implementation = new StakeMining();
        
        // 部署 ProxyAdmin
        proxyAdmin = new ProxyAdmin(address(this));
        
        // 编码初始化调用数据
        bytes memory initData = abi.encodeWithSelector(
            StakeMining.initialize.selector,
            address(stakeToken),
            address(rewardToken),
            REWARD_RATE,
            STAKE_MIN,
            STAKE_MAX
        );
        
        // 部署透明代理
        proxy = new TransparentUpgradeableProxy(
            address(implementation),
            address(proxyAdmin),
            initData
        );
        
        // 获取代理的合约实例
        stakeMining = StakeMining(address(proxy));

        // 授权合约转移奖励代币
        rewardToken.transfer(address(stakeMining), REWARD_SUPPLY);

        // 用户授权质押代币
        vm.prank(user1);
        stakeToken.approve(address(stakeMining), INITIAL_SUPPLY);

        vm.prank(user2);
        stakeToken.approve(address(stakeMining), INITIAL_SUPPLY);

        vm.prank(blacklistedUser);
        stakeToken.approve(address(stakeMining), INITIAL_SUPPLY);
    }

    // ==================== 测试初始化 ====================
    function testInitialization() public view {
        assertEq(address(stakeMining.stakeToken()), address(stakeToken));
        assertEq(address(stakeMining.rewardToken()), address(rewardToken));
        assertEq(stakeMining.rewardRate(), REWARD_RATE);
        assertEq(stakeMining.stakeMinAmount(), STAKE_MIN);
        assertEq(stakeMining.stakeMaxAmount(), STAKE_MAX);
    }

    // ==================== 测试质押功能 ====================
    function testStake() public {
        uint256 stakeAmount = 500 * 10**18;

        vm.prank(user1);
        stakeMining.stake(stakeAmount);

        (uint256 amount, uint256 rewardDebt, uint256 lastUpdate) = stakeMining.userStakes(user1);
        assertEq(amount, stakeAmount);
        assertEq(rewardDebt, 0);
        assertGt(lastUpdate, 0);
    }

    function testStakeBelowMin() public {
        uint256 stakeAmount = 50 * 10**18;

        vm.prank(user1);
        vm.expectRevert("stake amount out of limits");
        stakeMining.stake(stakeAmount);
    }

    function testStakeAboveMax() public {
        uint256 stakeAmount = 20000 * 10**18;

        vm.prank(user1);
        vm.expectRevert("stake amount out of limits");
        stakeMining.stake(stakeAmount);
    }

    function testStakeZeroAmount() public {
        vm.prank(user1);
        vm.expectRevert("stake amount must be greater than zero");
        stakeMining.stake(0);
    }

    // ==================== 测试赎回功能 ====================
    function testUnstake() public {
        uint256 stakeAmount = 500 * 10**18;
        uint256 unstakeAmount = 200 * 10**18;

        vm.prank(user1);
        stakeMining.stake(stakeAmount);

        vm.prank(user1);
        stakeMining.unstake(unstakeAmount);

        (uint256 amount, , ) = stakeMining.userStakes(user1);
        assertEq(amount, stakeAmount - unstakeAmount);
    }

    function testUnstakeMoreThanStaked() public {
        uint256 stakeAmount = 500 * 10**18;
        uint256 unstakeAmount = 600 * 10**18;

        vm.prank(user1);
        stakeMining.stake(stakeAmount);

        vm.prank(user1);
        vm.expectRevert("unstake amount exceeds staked amount");
        stakeMining.unstake(unstakeAmount);
    }

    // ==================== 测试收益计算和领取 ====================
    function testClaimReward() public {
        uint256 stakeAmount = 1000 * 10**18;
        uint256 timePassed = 1000;

        vm.prank(user1);
        stakeMining.stake(stakeAmount);

        vm.warp(block.timestamp + timePassed);

        vm.prank(user1);
        stakeMining.claimReward();

        uint256 expectedReward = stakeAmount * REWARD_RATE * timePassed;
        (, uint256 rewardDebt, ) = stakeMining.userStakes(user1);
        assertEq(rewardDebt, 0);
        assertEq(rewardToken.balanceOf(user1), expectedReward);
    }

    function testClaimRewardWithoutStake() public {
        vm.prank(user1);
        vm.expectRevert("no rewards to claim");
        stakeMining.claimReward();
    }

    // ==================== 测试黑名单功能 ====================
    function testBlacklist() public {
        uint256 stakeAmount = 500 * 10**18;

        stakeMining.addBlacklist(blacklistedUser);

        vm.prank(blacklistedUser);
        vm.expectRevert("account in blacklist");
        stakeMining.stake(stakeAmount);

        stakeMining.removeBlacklist(blacklistedUser);

        vm.prank(blacklistedUser);
        stakeMining.stake(stakeAmount);
    }

    // ==================== 测试暂停功能 ====================
    function testPause() public {
        uint256 stakeAmount = 500 * 10**18;

        stakeMining.pause();

        vm.prank(user1);
        vm.expectRevert("EnforcedPause()");
        stakeMining.stake(stakeAmount);

        stakeMining.unpause();

        vm.prank(user1);
        stakeMining.stake(stakeAmount);
    }

    // ==================== 测试质押限额更新 ====================
    function testUpdateStakeLimits() public {
        uint256 newMin = 200 * 10**18;
        uint256 newMax = 5000 * 10**18;

        stakeMining.updateStakeLimits(newMin, newMax);

        assertEq(stakeMining.stakeMinAmount(), newMin);
        assertEq(stakeMining.stakeMaxAmount(), newMax);
    }

    function testUpdateStakeLimitsInvalid() public {
        vm.expectRevert("invalid limits");
        stakeMining.updateStakeLimits(STAKE_MAX, STAKE_MIN);
    }

    // ==================== 测试批量质押和收益累加 ====================
    function testMultipleStakes() public {
        uint256 stakeAmount1 = 500 * 10**18;
        uint256 stakeAmount2 = 300 * 10**18;
        uint256 timePassed = 100;

        vm.prank(user1);
        stakeMining.stake(stakeAmount1);

        vm.warp(block.timestamp + timePassed);

        vm.prank(user1);
        stakeMining.stake(stakeAmount2);

        (uint256 amount, uint256 rewardDebt, ) = stakeMining.userStakes(user1);
        assertEq(amount, stakeAmount1 + stakeAmount2);
        assertGt(rewardDebt, 0);
    }
}