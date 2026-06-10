# Stake Mining - DeFi 质押挖矿系统

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![Solidity](https://img.shields.io/badge/Solidity-0.8.20-orange.svg)](https://soliditylang.org/)
[![Foundry](https://img.shields.io/badge/Foundry-latest-red.svg)](https://getfoundry.sh/)

一个基于以太坊的 DeFi 质押挖矿系统，包含智能合约和后端 API 服务。用户可以通过质押代币获得收益，管理员可以管理合约参数和用户黑名单。

## 📋 目录

- [项目概述](#项目概述)
- [技术栈](#技术栈)
- [项目结构](#项目结构)
- [功能特性](#功能特性)
- [快速开始](#快速开始)
- [API 文档](#api-文档)
- [智能合约](#智能合约)
- [安全特性](#安全特性)
- [开发指南](#开发指南)
- [测试](#测试)
- [部署](#部署)
- [贡献](#贡献)
- [许可证](#许可证)

## 🎯 项目概述

Stake Mining 是一个完整的去中心化金融（DeFi）质押挖矿解决方案，包括：

- **智能合约**：基于 Solidity 的可升级质押合约，支持 UUPS 代理模式
- **后端服务**：Go 语言实现的 RESTful API，提供用户和管理员接口
- **事件索引器**：链上事件监听和数据同步
- **交易追踪**：实时交易状态监控
- **安全管理**：JWT 认证、速率限制、黑名单机制

## 🛠️ 技术栈

### 智能合约
- **Solidity** ^0.8.20
- **OpenZeppelin Contracts** - 安全的智能合约库
- **Foundry** - 智能合约开发和测试框架
- **UUPS 代理模式** - 可升级合约架构

### 后端服务
- **Go** 1.21+
- **Gin** - Web 框架
- **GORM** - ORM 数据库操作
- **MySQL** 8.0+ - 数据存储
- **Redis** 7.0+ - 缓存和会话管理
- **Swagger** - API 文档
- **JWT** - 身份认证
- **Web3.go** - 以太坊交互

## 📁 项目结构

```
Stake-Project/
├── StakeContracts/           # 智能合约
│   ├── src/
│   │   └── StakeMining.sol   # 质押挖矿主合约
│   ├── test/
│   │   └── StakeMining.t.sol # 合约测试
│   ├── script/               # 部署脚本
│   └── foundry.toml          # Foundry 配置
│
├── StakeBackend/             # 后端服务
│   ├── cmd/
│   │   └── main.go           # 应用入口
│   ├── internal/
│   │   ├── api/              # API 控制器
│   │   ├── config/           # 配置管理
│   │   ├── contract/         # 合约交互
│   │   ├── db/               # 数据库连接
│   │   ├── indexer/          # 事件索引器
│   │   ├── middleware/       # 中间件
│   │   ├── model/            # 数据模型
│   │   ├── pkg/              # 工具包
│   │   ├── redis/            # Redis 客户端
│   │   └── utils/            # 工具函数
│   ├── configs/              # 配置文件
│   ├── docs/                 # Swagger 文档
│   └── scripts/              # 初始化脚本
│
└── README.md
```

## ✨ 功能特性

### 用户功能
- ✅ 质押代币到合约
- ✅ 解除质押并提取代币
- ✅ 领取累积的收益
- ✅ 查询质押信息和待领取收益
- ✅ 查看质押历史记录
- ✅ 交易状态追踪

### 管理员功能
- ✅ 添加/移除用户黑名单
- ✅ 暂停/恢复合约（紧急情况）
- ✅ 调整最小/最大质押限额
- ✅ 更新收益费率
- ✅ 查询合约状态
- ✅ 多级权限管理（普通管理员/超级管理员）

### 系统特性
- ✅ JWT 身份认证
- ✅ 速率限制保护
- ✅ 批量事件索引（高性能）
- ✅ 优雅关闭和数据持久化
- ✅ 完整的 API 文档（Swagger）
- ✅ 结构化日志记录
- ✅ 交易状态追踪器

## 🚀 快速开始

### 前置要求

- **Go** 1.21+
- **Node.js** 18+ (用于 Foundry)
- **MySQL** 8.0+
- **Redis** 7.0+
- **Foundry** (智能合约开发工具)

### 安装步骤

#### 1. 克隆项目

```bash
git clone https://github.com/your-username/Stake-Project.git
cd Stake-Project
```

#### 2. 智能合约设置

```bash
cd StakeContracts

# 安装依赖
forge install

# 编译合约
forge build

# 运行测试
forge test

# 格式化代码
forge fmt
```

#### 3. 后端服务设置

```bash
cd ../StakeBackend

# 安装 Go 依赖
go mod tidy

# 复制配置文件
cp configs/config.yaml configs/config.yaml.local

# 编辑配置文件，设置数据库、Redis、合约地址等
vim configs/config.yaml.local
```

#### 4. 初始化数据库

```bash
# 创建数据库（根据配置文件中的设置）
mysql -u root -p -e "CREATE DATABASE stake_mining CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# 运行管理员初始化脚本（可选）
mysql -u root -p stake_mining < scripts/init_admins.sql
```

#### 5. 启动服务

```bash
# 启动后端服务
go run cmd/main.go
```

服务将在 `http://localhost:8080` 启动。

#### 6. 访问 Swagger 文档

打开浏览器访问：`http://localhost:8080/swagger/index.html`

## 📖 API 文档

### 公开接口

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/api/login` | 用户登录获取 JWT Token |
| GET | `/api/admin/login` | 管理员登录 |

### 用户接口（需要 JWT 认证）

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | `/api/stake/info` | 查询用户质押信息 |
| GET | `/api/stake/reward` | 查询待领取收益 |
| GET | `/api/stake/history` | 查询质押历史 |
| POST | `/api/stake/do` | 执行质押 |
| POST | `/api/stake/unstake` | 执行解质押 |
| POST | `/api/stake/claim` | 领取收益 |
| GET | `/api/tx/status` | 查询交易状态 |
| GET | `/api/tx/wait` | 等待交易确认 |

### 管理员接口（需要管理员 JWT 认证）

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | `/api/admin/blacklist/add` | 添加黑名单 |
| POST | `/api/admin/blacklist/remove` | 移除黑名单 |
| POST | `/api/admin/pause` | 暂停合约 |
| POST | `/api/admin/unpause` | 恢复合约 |
| POST | `/api/admin/limits/update` | 更新质押限额 |
| POST | `/api/admin/rate/update` | 更新收益费率 |
| GET | `/api/admin/status` | 查询合约状态 |

详细的 API 使用说明请参考 [API_GUIDE.md](StakeBackend/API_GUIDE.md)。

## 🔐 智能合约

### 合约架构

`StakeMining` 合约采用可升级架构：

- **UUPS 代理模式**：允许合约逻辑升级而不改变合约地址
- **Ownable**：所有权管理
- **Pausable**：紧急暂停功能
- **ReentrancyGuard**：防止重入攻击

### 核心功能

```solidity
// 质押代币
function stake(uint256 _amount) external;

// 解除质押
function unstake(uint256 _amount) external;

// 领取收益
function claimReward() external;

// 查询待领取收益
function getPendingReward(address _user) external view returns (uint256);
```

### 管理员功能

```solidity
// 添加/移除黑名单
function addBlacklist(address _account) external onlyOwner;
function removeBlacklist(address _account) external onlyOwner;

// 暂停/恢复合约
function pause() external onlyOwner;
function unpause() external onlyOwner;

// 更新参数
function updateStakeLimits(uint256 _min, uint256 _max) external onlyOwner;
function setRewardRate(uint256 _rate) external onlyOwner;
```

### 事件

合约发出以下事件供后端索引：

- `Staked` - 质押事件
- `Unstaked` - 解质押事件
- `RewardClaimed` - 收益领取事件
- `BlacklistAdded/Removed` - 黑名单变更事件
- `StakeLimitsUpdated` - 限额更新事件
- `RewardRateUpdated` - 费率更新事件

## 🛡️ 安全特性

### 智能合约安全

- ✅ 使用 OpenZeppelin 经过审计的合约库
- ✅ 重入攻击防护（ReentrancyGuard）
- ✅ 安全的代币传输（SafeERC20）
- ✅ 溢出检查（Solidity 0.8+ 内置）
- ✅ 访问控制（OnlyOwner）
- ✅ 紧急暂停机制
- ✅ 黑名单系统

### 后端安全

- ✅ JWT 身份认证
- ✅ 管理员权限分级
- ✅ API 速率限制
- ✅ 输入验证
- ✅ SQL 注入防护（GORM ORM）
- ✅ 结构化日志记录
- ✅ 优雅的错误处理

### 生产环境建议

1. **多重签名钱包**：合约所有者应使用多签钱包
2. **时间锁**：关键参数变更应设置时间锁
3. **审计**：生产部署前进行专业安全审计
4. **监控**：实时监控合约活动和异常行为
5. **备份**：定期备份数据库和配置文件
6. **HTTPS**：生产环境使用 HTTPS
7. **IP 白名单**：限制管理员接口访问 IP

## 💻 开发指南

### 重新生成 Swagger 文档

修改 API 注释后：

```bash
cd StakeBackend
swag init -g cmd/main.go -o docs
```

### 添加新的 API 接口

1. 在 `internal/api/` 中添加处理函数
2. 添加 Swagger 注释（godoc 格式）
3. 在 `cmd/main.go` 中注册路由
4. 重新生成 Swagger 文档

### Swagger 注释示例

```go
// GetUserProfile godoc
// @Summary 获取用户资料
// @Description 根据用户ID获取详细信息
// @Tags 用户
// @Security BearerAuth
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response{data=model.User}
// @Router /api/user/{id} [get]
func GetUserProfile(c *gin.Context) {
    // ...
}
```

### 合约开发

```bash
cd StakeContracts

# 编写新测试
forge test --match-test testYourFunction -vvv

# Gas 分析
forge snapshot

# 本地测试网部署
forge script script/Deploy.s.sol --rpc-url http://localhost:8545 --private-key $PRIVATE_KEY
```

## 🧪 测试

### 智能合约测试

```bash
cd StakeContracts

# 运行所有测试
forge test

# 详细输出
forge test -vvv

# 运行特定测试
forge test --match-test testStake -vvv

# Gas 报告
forge test --gas-report
```

### 后端测试

```bash
cd StakeBackend

# 运行测试脚本
./test_api.sh
./test_improvements.sh

# 或使用 curl 手动测试
curl "http://localhost:8080/api/login?address=0xYourAddress"
```

## 🌐 部署

### 智能合约部署

```bash
cd StakeContracts

# 部署到测试网
forge script script/Deploy.s.sol \
  --rpc-url https://sepolia.infura.io/v3/YOUR_API_KEY \
  --private-key $PRIVATE_KEY \
  --broadcast \
  --verify

# 部署到主网（谨慎操作）
forge script script/Deploy.s.sol \
  --rpc-url https://mainnet.infura.io/v3/YOUR_API_KEY \
  --private-key $PRIVATE_KEY \
  --broadcast \
  --verify
```

### 后端部署

#### Docker 部署（推荐）

创建 `Dockerfile`：

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/configs/config.yaml.local ./configs/config.yaml
EXPOSE 8080
CMD ["./main"]
```

构建和运行：

```bash
docker build -t stake-backend .
docker run -p 8080:8080 stake-backend
```

#### 直接部署

```bash
# 编译
go build -o stake-backend cmd/main.go

# 后台运行
nohup ./stake-backend > app.log 2>&1 &

# 或使用 systemd 管理服务
sudo cp stake-backend.service /etc/systemd/system/
sudo systemctl enable stake-backend
sudo systemctl start stake-backend
```

### 环境变量配置

生产环境建议使用环境变量：

```bash
export DB_HOST=your-db-host
export DB_USER=your-db-user
export DB_PASSWORD=your-db-password
export REDIS_HOST=your-redis-host
export CONTRACT_ADDRESS=0x...
export JWT_SECRET=your-secret-key
```

## 🤝 贡献

欢迎贡献代码、报告问题或提出改进建议！

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

### 贡献指南

- 遵循现有的代码风格
- 添加适当的测试
- 更新文档
- 确保所有测试通过

## 📄 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 📞 联系方式

- 项目主页：[GitHub Repository](https://github.com/your-username/Stake-Project)
- 问题反馈：[Issues](https://github.com/your-username/Stake-Project/issues)
- 邮箱：support@example.com

## ⚠️ 免责声明

本软件仅供学习和研究使用。在生产环境中使用前，请：

1. 进行完整的安全审计
2. 在测试网充分测试
3. 咨询专业的区块链安全专家
4. 了解并承担使用风险

开发者不对因使用本软件造成的任何损失负责。

---

**⭐ 如果这个项目对你有帮助，请给个 Star！**
