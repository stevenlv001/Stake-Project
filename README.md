# Stake-Project - 去中心化质押挖矿系统

一个完整的去中心化质押挖矿解决方案，包含 Solidity 智能合约和 Go 后端服务。

## 项目概述

本项目实现了一个基于区块链的质押挖矿系统，用户可以质押 ERC20 代币并获得奖励代币作为收益。系统采用可升级合约架构，支持批量事件索引、防穿透缓存、链重组处理等企业级特性。

### 核心功能

- **质押挖矿**：用户质押代币获得奖励
- **收益领取**：随时领取累积的挖矿奖励
- **黑名单管理**：管理员可限制特定地址参与
- **合约升级**：支持 UUPS 模式的可升级合约
- **事件索引**：高效的链上事件同步与存储
- **缓存优化**：Redis 多级缓存，防穿透防击穿
- **API 限流**：防止恶意请求

---

## 项目结构

```
Stake-Project/
├── StakeBackend/                 # Go 后端服务
│   ├── cmd/
│   │   └── main.go              # 服务入口
│   ├── internal/
│   │   ├── abi/                 # 合约绑定代码（abigen 生成）
│   │   ├── api/                 # API 控制器
│   │   ├── config/              # 配置管理
│   │   ├── contract/            # 合约交互层
│   │   ├── db/                  # 数据库操作
│   │   ├── indexer/             # 事件索引器
│   │   ├── middleware/          # 中间件（JWT、限流）
│   │   ├── model/               # 数据模型
│   │   ├── pkg/logger/          # 日志封装
│   │   ├── redis/               # Redis 缓存
│   │   └── utils/               # 工具函数
│   ├── configs/
│   │   └── config.yaml          # 配置文件
│   ├── go.mod                   # Go 模块依赖
│   └── go.sum
│
├── StakeContracts/              # Solidity 智能合约
│   ├── src/
│   │   └── StakeMining.sol      # 质押挖矿合约
│   ├── test/
│   │   └── StakeMining.t.sol    # 合约测试
│   ├── foundry.toml             # Foundry 配置
│   └── remappings.txt           # 依赖映射
│
└── README.md
```

---

## 快速开始

### 环境要求

- **Go**: 1.24.0+
- **Node.js**: 18+ (可选，用于合约开发)
- **Foundry**: 最新稳定版
- **MySQL**: 8.0+
- **Redis**: 6.0+

### 1. 安装依赖

#### 后端依赖

```bash
cd StakeBackend
go mod tidy
```

#### 合约依赖

```bash
cd StakeContracts
forge install OpenZeppelin/openzeppelin-contracts@v5.2.0
forge install OpenZeppelin/openzeppelin-contracts-upgradeable@v5.2.0
```

### 2. 配置服务

创建配置文件 `StakeBackend/configs/config.yaml`：

```yaml
app:
  name: stake-mining
  port: "8080"
  mode: debug

mysql:
  dsn: "root:your_password@tcp(127.0.0.1:3306)/stake_db?charset=utf8mb4&parseTime=True&loc=Local"
  max_idle: 10
  max_open: 100

redis:
  addr: "127.0.0.1:6379"
  password: ""
  db: 0

jwt:
  secret: "your-secret-key-change-in-production"
  expire: 86400

chain:
  rpc: "http://localhost:8545"
  stake_token: "0x..."        # 质押代币地址
  reward_token: "0x..."       # 奖励代币地址
  mining_proxy: "0x..."       # 代理合约地址
  start_block: 0

indexer:
  batch_size: 100
  batch_wait_seconds: 10

rate_limit:
  max_requests: 100
  window_seconds: 60
```

### 3. 部署智能合约

```bash
cd StakeContracts

# 编译合约
forge build

# 运行测试
forge test -vvv

# 部署到本地测试网（示例）
forge script script/Deploy.s.sol --rpc-url http://localhost:8545 --broadcast
```

### 4. 生成合约绑定代码

```bash
cd StakeBackend

# 从合约 ABI 生成 Go 绑定
abigen --abi=../StakeContracts/out/StakeMining.sol/StakeMining.json \
       --pkg=abi \
       --type=StakeMining \
       --out=internal/abi/stakemining.go
```

### 5. 启动后端服务

```bash
cd StakeBackend

# 确保 MySQL 和 Redis 已启动
# 然后运行服务
go run cmd/main.go
```

服务启动后，访问 `http://localhost:8080`

---

## API 文档

### 公共接口

#### 用户登录

```http
GET /api/login?address={wallet_address}
```

**参数：**
- `address`: 钱包地址（必填）

**响应：**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 认证接口

所有认证接口需要在 Header 中携带 JWT Token：

```
Authorization: Bearer {token}
```

#### 查询质押信息

```http
GET /api/stake/info
```

**响应：**
```json
{
  "data": {
    "id": 1,
    "user_address": "0x123...",
    "stake_amount": "1000000000000000000000",
    "update_time": 1699999999
  }
}
```

#### 查询待领取收益

```http
GET /api/stake/reward
```

**响应：**
```json
{
  "data": {
    "pending_reward": "50000000000000000000"
  }
}
```

---

## 🔧 技术栈

### 后端技术

| 技术 | 版本 | 用途 |
|------|------|------|
| Go | 1.24.0 | 主要编程语言 |
| Gin | v1.10.0 | Web 框架 |
| GORM | v1.31.1 | ORM 框架 |
| Redis | v8.11.5 | 缓存层 |
| Zap | v1.28.0 | 日志框架 |
| Viper | v1.21.0 | 配置管理 |
| go-ethereum | v1.17.2 | 区块链交互 |
| JWT | v4.5.2 | 认证授权 |

### 合约技术

| 技术 | 版本 | 用途 |
|------|------|------|
| Solidity | 0.8.20+ | 智能合约语言 |
| Foundry | 最新版 | 合约开发工具链 |
| OpenZeppelin | v5.2.0 | 安全合约库 |
| UUPS Proxy | - | 可升级合约模式 |

---

## 核心特性

### 1. 批量事件索引器

- **批量处理**：支持按数量或时间自动批量入库
- **链重组处理**：自动检测并处理区块链重组
- **优雅关机**：服务关闭前强制刷新队列，防止数据丢失
- **并发安全**：使用互斥锁保证队列操作线程安全

```go
// 批量入库触发条件
// 1. 达到批量数量阈值
if len(eventQueue) >= batchSize {
    flushBatch()
}

// 2. 超时自动提交
if time.Since(lastFlushTime) > batchWaitTime {
    flushBatch()
}
```

### 2. Redis 缓存优化

- **防穿透**：空值也缓存，避免恶意查询
- **防击穿**：使用分布式锁防止并发击穿
- **自动过期**：设置合理的 TTL

```go
// 缓存查询流程
cacheData, err := redis.GetCache(cacheKey)
if err == nil {
    if cacheData == "null" {
        // 空值缓存，防止穿透
        return emptyResult
    }
    return cacheData
}

// 加锁防击穿
if !redis.TryLock(cacheKey) {
    return busyError
}
defer redis.Unlock(cacheKey)
```

### 3. 可升级合约

采用 **TransparentUpgradeableProxy** 模式：

```solidity
// 部署代理合约
proxy = new TransparentUpgradeableProxy(
    address(implementation),
    address(proxyAdmin),
    abi.encodeWithSelector(
        StakeMining.initialize.selector,
        stakeToken,
        rewardToken,
        rewardRate,
        stakeMinAmount,
        stakeMaxAmount
    )
);
```

### 4. 安全特性

- 使用 SafeERC20 进行代币转账
- 重入攻击防护（ReentrancyGuard）
- 暂停机制（Pausable）
- 黑名单管理
- 质押限额控制

---

## 🧪 测试

### 合约测试

```bash
cd StakeContracts

# 运行所有测试
forge test

# 详细输出
forge test -vvv

# 运行特定测试
forge test --match-test testStake
```

**测试覆盖率：**
- 初始化测试
- 质押功能测试
- 赎回功能测试
- 收益领取测试
- 黑名单功能测试
- 暂停功能测试
- 限额更新测试
- 边界条件测试

### 后端测试

```bash
cd StakeBackend

# 运行单元测试
go test ./...

# 带覆盖率
go test -cover ./...
```

---

## 数据库设计

### UserStake（用户质押表）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | 主键 |
| user_address | VARCHAR(42) | 用户地址 |
| stake_amount | VARCHAR(78) | 质押数量（Wei） |
| update_time | BIGINT | 更新时间戳 |

### ChainEvent（链上事件表）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | 主键 |
| block_number | BIGINT | 区块号 |
| block_hash | VARCHAR(66) | 区块哈希 |
| tx_hash | VARCHAR(66) | 交易哈希 |
| event_type | VARCHAR(20) | 事件类型 |
| user | VARCHAR(42) | 用户地址 |
| amount | VARCHAR(78) | 金额（Wei） |
| event_time | BIGINT | 事件时间戳 |

### BlockSync（区块同步表）

| 字段 | 类型 | 说明 |
|------|------|------|
| block_number | BIGINT | 区块号 |
| block_hash | VARCHAR(66) | 区块哈希 |
| parent_hash | VARCHAR(66) | 父区块哈希 |

---

## 安全建议

### 生产环境部署

1. **配置安全**
   - 修改 JWT Secret 为强随机字符串
   - 数据库密码使用环境变量
   - Redis 设置访问密码

2. **网络隔离**
   - 使用内网访问数据库和 Redis
   - 配置防火墙规则
   - 使用反向代理（Nginx）

3. **监控告警**
   - 接入日志收集系统（ELK）
   - 配置服务监控（Prometheus + Grafana）
   - 设置异常告警

---

## 📝 开发指南

### 添加新的合约事件

1. 在 Solidity 合约中添加事件定义
2. 重新编译合约：`forge build`
3. 重新生成绑定代码：`abigen --abi=...`
4. 在 `events.go` 中添加事件 ID 常量
5. 在 `indexer.go` 中添加事件解析逻辑

### 添加新的 API 接口

1. 在 `controller.go` 中实现 Handler 函数
2. 在 `main.go` 中注册路由
3. 添加必要的中间件（JWT、限流）
4. 编写单元测试

---