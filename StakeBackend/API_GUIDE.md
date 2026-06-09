# API 使用指南

## 📚 Swagger API 文档

项目已集成 Swagger UI，启动服务后可通过浏览器访问交互式 API 文档。

### 访问地址

```
http://localhost:8080/swagger/index.html
```

### 功能特性

- ✅ 完整的 API 接口文档
- ✅ 在线测试 API 接口
- ✅ 自动生成请求/响应示例
- ✅ 支持 JWT 认证测试

---

## 🔐 认证说明

### 1. 普通用户登录

**接口：** `GET /api/login`

**请求示例：**
```bash
curl "http://localhost:8080/api/login?address=0x1234567890abcdef1234567890abcdef12345678"
```

**响应：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**使用 Token：**
在后续请求的 Header 中添加：
```
Authorization: Bearer <token>
```

---

### 2. 管理员登录

**接口：** `GET /api/admin/login`

**请求示例：**
```bash
# 普通管理员
curl "http://localhost:8080/api/admin/login?admin_id=admin001&role=admin"

# 超级管理员
curl "http://localhost:8080/api/admin/login?admin_id=superadmin&role=super_admin"
```

**响应：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

> ⚠️ **注意：** 当前为演示实现，生产环境应使用签名验证或数据库查询来验证管理员身份。

---

## 👨‍💼 管理员 API

所有管理员接口都需要使用管理员 Token 进行认证。

### 1. 黑名单管理

#### 添加黑名单

**接口：** `POST /api/admin/blacklist/add`

**请求体：**
```json
{
  "address": "0xBadActor1234567890abcdef1234567890abcd"
}
```

**响应：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "to": "0x合约地址",
    "data": "0xa9059cbb...",
    "value": "0",
    "gas_limit": 100000
  }
}
```

#### 移除黑名单

**接口：** `POST /api/admin/blacklist/remove`

**请求体：**
```json
{
  "address": "0xBadActor1234567890abcdef1234567890abcd"
}
```

---

### 2. 合约控制

#### 暂停合约

**接口：** `POST /api/admin/pause`

**说明：** 紧急情况下暂停合约，禁止所有质押操作。

**响应：** 返回交易数据，需要管理员钱包签名并提交。

#### 恢复合约

**接口：** `POST /api/admin/unpause`

**说明：** 恢复已暂停的合约。

---

### 3. 参数调整

#### 更新质押限额

**接口：** `POST /api/admin/limits/update`

**请求体：**
```json
{
  "min_amount": "100000000000000000",
  "max_amount": "10000000000000000000"
}
```

**说明：** 
- 金额单位为 wei（1 ETH = 10^18 wei）
- min_amount 必须小于 max_amount

#### 更新收益费率

**接口：** `POST /api/admin/rate/update`

**请求体：**
```json
{
  "rate": "100"
}
```

**说明：** 设置每秒收益速率（单位：wei/秒）

---

### 4. 状态查询

#### 查询合约状态

**接口：** `GET /api/admin/status`

**响应：**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "message": "功能开发中"
  }
}
```

> 📝 TODO: 实现完整的合约状态查询功能

---

## 🧪 测试示例

### 使用 cURL 测试管理员 API

```bash
# 1. 获取管理员 Token
ADMIN_TOKEN=$(curl -s "http://localhost:8080/api/admin/login?admin_id=admin001&role=super_admin" | jq -r '.data.token')

# 2. 添加黑名单
curl -X POST "http://localhost:8080/api/admin/blacklist/add" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"address":"0xBadActor1234567890abcdef1234567890abcd"}'

# 3. 暂停合约
curl -X POST "http://localhost:8080/api/admin/pause" \
  -H "Authorization: Bearer $ADMIN_TOKEN"

# 4. 更新质押限额
curl -X POST "http://localhost:8080/api/admin/limits/update" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"min_amount":"100000000000000000","max_amount":"10000000000000000000"}'
```

### 使用 Swagger UI 测试

1. 访问 `http://localhost:8080/swagger/index.html`
2. 点击任意管理员接口
3. 点击 "Try it out"
4. 在 Authorization 字段输入：`Bearer <your_admin_token>`
5. 填写请求参数
6. 点击 "Execute" 执行请求

---

## 🔒 安全建议

### 生产环境部署

1. **管理员身份验证**
   - 实现基于签名的身份验证
   - 或使用数据库存储管理员凭证
   - 支持多因素认证（MFA）

2. **权限分离**
   - 普通管理员：只能查看状态、管理黑名单
   - 超级管理员：可以暂停合约、调整参数

3. **操作审计**
   - 记录所有管理员操作日志
   - 关键操作需要多重签名确认

4. **IP 白名单**
   - 限制管理员接口的访问 IP
   - 仅允许内部网络或特定 IP 访问

5. **速率限制**
   - 对管理员接口实施更严格的限流
   - 防止暴力攻击

---

## 📖 API 端点总览

### 公开接口
- `GET /api/login` - 用户登录
- `GET /api/admin/login` - 管理员登录
- `GET /swagger/*any` - Swagger 文档

### 用户接口（需要 JWT）
- `GET /api/stake/info` - 查询质押信息
- `GET /api/stake/reward` - 查询待领取收益
- `POST /api/stake/do` - 质押代币
- `POST /api/stake/unstake` - 解质押
- `POST /api/stake/claim` - 领取收益
- `GET /api/tx/status` - 查询交易状态
- `GET /api/tx/wait` - 等待交易确认

### 管理员接口（需要 Admin JWT）
- `POST /api/admin/blacklist/add` - 添加黑名单
- `POST /api/admin/blacklist/remove` - 移除黑名单
- `POST /api/admin/pause` - 暂停合约
- `POST /api/admin/unpause` - 恢复合约
- `POST /api/admin/limits/update` - 更新质押限额
- `POST /api/admin/rate/update` - 更新收益费率
- `GET /api/admin/status` - 查询合约状态

---

## 🛠️ 开发说明

### 重新生成 Swagger 文档

修改 API 注释后，运行：
```bash
cd StakeBackend
swag init -g cmd/main.go -o docs
```

### 添加新的 API 接口

1. 在 controller 中添加处理函数
2. 添加 Swagger 注释（godoc 格式）
3. 在 main.go 中注册路由
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

更多注释语法参考：[Swag 官方文档](https://github.com/swaggo/swag)
