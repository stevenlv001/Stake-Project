# Stake Mining Project

一个基于区块链的质押挖矿系统，包含智能合约和 Go 后端服务。

## 📋 项目概览

本项目实现了一个完整的去中心化质押挖矿系统，支持用户质押代币获取收益，并提供 RESTful API 接口供前端调用。

### ✨ 功能特性

**智能合约层**
- ✅ 可升级质押挖矿合约（UUPS 模式）
- ✅ 质押/解质押/收益领取
- ✅ 黑名单机制
- ✅ 紧急暂停功能
- ✅ 质押限额控制
- ✅ 实时收益计算

**后端服务层**
- ✅ RESTful API 接口
- ✅ Swagger API 文档
- ✅ JWT 身份认证（用户 + 管理员）
- ✅ Redis 缓存优化
- ✅ MySQL 数据持久化
- ✅ 交易状态追踪
- ✅ 事件批量索引
- ✅ 请求限流保护
- ✅ 管理员后台 API

---

## 🛠️ 技术栈

### 智能合约
| 技术 | 版本 | 说明 |
|------|------|------|
| Solidity | ^0.8.20 | 合约开发语言 |
| OpenZeppelin | ^5.0 | 安全合约库 |
| Foundry | ^0.2 | 智能合约测试框架 |

### 后端服务
| 技术 | 版本 | 说明 |
|------|------|------|
| Go | 1.21+ | 后端开发语言 |
| Gin | ^1.9 | Web 框架 |
| GORM | ^1.25 | ORM 框架 |
| MySQL | 8.0+ | 关系型数据库 |
| Redis | 7.0+ | 缓存服务 |
| go-ethereum | ^1.12 | Ethereum 客户端库 |
| Zap | ^1.27 | 日志框架 |

---

## 🏗️ 项目结构

```
Stake-Project/
├── StakeContracts/           # 智能合约模块
│   ├── src/                  # 合约源代码
│   │   └── StakeMining.sol   # 质押挖矿合约
│   ├── test/                 # 合约测试
│   │   └── StakeMining.t.sol # Foundry 测试文件
│   ├── foundry.toml          # Foundry 配置
│   └── README.md             # 合约模块说明
├── StakeBackend/             # Go 后端服务
│   ├── cmd/                  # 入口文件
│   │   └── main.go           # 主函数
│   ├── configs/              # 配置文件
│   │   └── config.yaml       # 应用配置
│   ├── internal/             # 内部模块
│   │   ├── api/              # API 控制器
│   │   ├── abi/              # ABI 定义
│   │   ├── config/           # 配置管理
│   │   ├── contract/         # 合约交互
│   │   ├── db/               # 数据库连接
│   │   ├── indexer/          # 事件索引器
│   │   ├── middleware/       # 中间件
│   │   ├── model/            # 数据模型
│   │   ├── pkg/              # 工具包
│   │   │   ├── logger/       # 日志组件
│   │   │   ├── request/      # 请求处理
│   │   │   ├── response/     # 响应封装
│   │   │   └── txtracker/    # 交易追踪
│   │   ├── redis/            # Redis 客户端
│   │   └── utils/            # 工具函数
│   ├── go.mod                # Go 依赖
│   └── go.sum                # 依赖校验
└── README.md                 # 项目说明（本文件）
```

---

## 🚀 快速开始

### 前置条件

- Go 1.21+
- Node.js 18+（用于 Foundry）
- MySQL 8.0+
- Redis 7.0+

### 1. 克隆项目

```bash
git clone https://github.com/your-username/Stake-Project.git
cd Stake-Project
```

### 2. 配置智能合约

```bash
cd StakeContracts
# 安装依赖
forge install
# 编译合约
forge build
# 运行测试
forge test -v
```

### 3. 配置后端服务

```bash
cd ../StakeBackend

# 安装依赖
go mod tidy

# 修改配置文件
cp configs/config.yaml configs/config.yaml.local
# 编辑 config.yaml.local，配置数据库、Redis、链RPC等
```

### 4. 启动服务

```bash
go run cmd/main.go
```

服务将在 `http://localhost:8080` 启动。

---

## 📡 API 接口文档

### 🌐 Swagger 在线文档

启动服务后访问：`http://localhost:8080/swagger/index.html`

- ✅ 完整的 API 接口文档
- ✅ 在线测试 API 接口
- ✅ 自动生成请求/响应示例

详细使用说明请查看 [API_GUIDE.md](./StakeBackend/API_GUIDE.md)

---

### 认证接口

#### 登录获取 Token

```
GET /api/login?address=<钱包地址>
```

**请求示例：**
```bash
curl "http://localhost:8080/api/login?address=0x1234567890abcdef1234567890abcdef12345678"
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

---

### 质押接口

#### 查询质押信息

```
GET /api/stake/info
Authorization: Bearer <token>
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "user_address": "0x1234...",
    "stake_amount": "1000000000000000000",
    "reward_debt": "50000000000000000",
    "created_at": "2024-01-01T12:00:00Z"
  }
}
```

#### 查询待领取收益

```
GET /api/stake/reward
Authorization: Bearer <token>
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "pending_reward": "1234567890123456789"
  }
}
```

#### 质押代币（获取交易数据）

```
POST /api/stake/do
Authorization: Bearer <token>
Content-Type: application/json

{
  "amount": "1000000000000000000"
}
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "to": "0x合约地址",
    "data": "0xa9059cbb...",
    "value": "0",
    "gas_limit": 300000
  }
}
```

#### 解质押

```
POST /api/stake/unstake
Authorization: Bearer <token>
Content-Type: application/json

{
  "amount": "500000000000000000"
}
```

#### 领取收益

```
POST /api/stake/claim
Authorization: Bearer <token>
```

---

### 交易追踪接口

#### 查询交易状态

```
GET /api/tx/status?tx_hash=<交易哈希>
Authorization: Bearer <token>
```

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "tx_hash": "0xabc123...",
    "status": "confirmed",
    "confirmations": 12,
    "block_number": 18500000
  }
}
```

#### 等待交易确认

```
GET /api/tx/wait?tx_hash=<交易哈希>
Authorization: Bearer <token>
```

---

### 管理员接口

> 所有管理员接口需要使用管理员 Token 进行认证

#### 管理员登录

```
GET /api/admin/login?admin_id=<管理员ID>&role=<角色>
```

**角色类型：**
- `admin` - 普通管理员
- `super_admin` - 超级管理员

**响应示例：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### 黑名单管理

**添加黑名单：**
```
POST /api/admin/blacklist/add
Authorization: Bearer <admin_token>
Content-Type: application/json

{
  "address": "0x不良地址"
}
```

**移除黑名单：**
```
POST /api/admin/blacklist/remove
Authorization: Bearer <admin_token>
Content-Type: application/json

{
  "address": "0x地址"
}
```

#### 合约控制

**暂停合约：**
```
POST /api/admin/pause
Authorization: Bearer <admin_token>
```

**恢复合约：**
```
POST /api/admin/unpause
Authorization: Bearer <admin_token>
```

#### 参数调整

**更新质押限额：**
```
POST /api/admin/limits/update
Authorization: Bearer <admin_token>
Content-Type: application/json

{
  "min_amount": "100000000000000000",
  "max_amount": "10000000000000000000"
}
```

**更新收益费率：**
```
POST /api/admin/rate/update
Authorization: Bearer <admin_token>
Content-Type: application/json

{
  "rate": "100"
}
```

#### 状态查询

**查询合约状态：**
```
GET /api/admin/status
Authorization: Bearer <admin_token>
```

---

## 📄 智能合约接口

### 主要函数

| 函数名 | 功能 | 调用权限 |
|--------|------|----------|
| `stake(uint256 amount)` | 质押代币 | 普通用户 |
| `unstake(uint256 amount)` | 解质押 | 普通用户 |
| `claimReward()` | 领取收益 | 普通用户 |
| `getPendingReward(address)` | 查询待领取收益 | 公开 |
| `pause()` | 暂停合约 | 管理员 |
| `unpause()` | 恢复合约 | 管理员 |
| `addBlacklist(address)` | 添加黑名单 | 管理员 |
| `removeBlacklist(address)` | 移除黑名单 | 管理员 |

### 事件日志

```solidity
event Staked(address indexed user, uint256 amount, uint256 time);
event Unstaked(address indexed user, uint256 amount, uint256 time);
event RewardClaimed(address indexed user, uint256 reward, uint256 time);
event BlacklistAdded(address indexed account);
event BlacklistRemoved(address indexed account);
event StakeLimitsUpdated(uint256 min, uint256 max);
```

---

## 🗂️ 数据库模型

### UserStake（用户质押记录）

| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | uint | 主键 |
| user_address | string | 用户钱包地址 |
| stake_amount | string | 质押金额 |
| reward_debt | string | 已领取收益 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

### ChainEvent（链上事件记录）

| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | uint | 主键 |
| event_type | string | 事件类型 |
| tx_hash | string | 交易哈希 |
| block_number | uint | 区块号 |
| data | json | 事件数据 |
| processed | bool | 是否已处理 |

---

## ⚙️ 配置说明

### 配置文件结构

```yaml
app:
  name: "Stake-Mining-Backend"
  port: 8080
  mode: "prod"

mysql:
  dsn: "root:password@tcp(127.0.0.1:3306)/stake_mining?charset=utf8mb4&parseTime=True&loc=Local"
  max_idle: 10
  max_open: 100

redis:
  addr: "127.0.0.1:6379"
  password: ""
  db: 0

jwt:
  secret: "your-jwt-secret"
  expire: 86400

chain:
  rpc: "https://sepolia.infura.io/v3/your-infura-key"
  stake_token: "0x质押代币地址"
  reward_token: "0x收益代币地址"
  mining_proxy: "0x代理合约地址"
  start_block: 0

indexer:
  batch_size: 50
  batch_wait_seconds: 10

rate_limit:
  max_requests: 60
  window_seconds: 60
```

---

## 🧪 测试

### 智能合约测试

```bash
cd StakeContracts
forge test -v
forge coverage
```

### 后端测试

```bash
cd StakeBackend
go test ./...
```

### API 测试

**方式一：使用 Swagger UI（推荐）**

1. 启动服务：`go run cmd/main.go`
2. 访问：`http://localhost:8080/swagger/index.html`
3. 在线测试所有 API 接口

**方式二：使用测试脚本**

```bash
cd StakeBackend
./test_api.sh
```

**方式三：使用 cURL**

详细示例请查看 [API_GUIDE.md](./StakeBackend/API_GUIDE.md)

---

## 📝 部署

### 测试网部署

```bash
# 部署到 Sepolia 测试网
forge script script/Deploy.s.sol:Deploy --rpc-url $SEPOLIA_RPC --private-key $PRIVATE_KEY --broadcast
```

### 生产环境部署

```bash
# 使用 Docker Compose
docker-compose up -d

# 或使用 systemd
sudo systemctl daemon-reload
sudo systemctl enable stake-backend
sudo systemctl start stake-backend
```

---

## 🔒 安全注意事项

1. **合约安全**：部署前请进行专业安全审计
2. **私钥管理**：不要将私钥硬编码到代码中
3. **RPC节点**：生产环境建议使用专用节点服务
4. **数据备份**：定期备份数据库和关键配置
5. **HTTPS**：生产环境务必启用 HTTPS

---

## 📄 许可证

MIT License

---
