#!/bin/bash

# API 测试脚本
# 使用前请确保服务已启动：go run cmd/main.go

BASE_URL="http://localhost:8080"

echo "=========================================="
echo "Stake Mining API 测试脚本"
echo "=========================================="
echo ""

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 测试计数器
PASS=0
FAIL=0

# 测试函数
test_api() {
    local name=$1
    local method=$2
    local url=$3
    local data=$4
    
    echo -e "${YELLOW}测试: ${name}${NC}"
    
    if [ "$method" == "GET" ]; then
        response=$(curl -s -w "\n%{http_code}" "$url")
    else
        response=$(curl -s -w "\n%{http_code}" -X "$method" -H "Content-Type: application/json" -d "$data" "$url")
    fi
    
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | sed '$d')
    
    if [ "$http_code" == "200" ]; then
        echo -e "${GREEN}✓ 通过 (HTTP $http_code)${NC}"
        PASS=$((PASS + 1))
    else
        echo -e "${RED}✗ 失败 (HTTP $http_code)${NC}"
        FAIL=$((FAIL + 1))
    fi
    
    echo "响应: $body"
    echo ""
}

# 测试管理员登录获取 Token
echo -e "${YELLOW}步骤 1: 获取管理员 Token${NC}"
ADMIN_RESPONSE=$(curl -s "$BASE_URL/api/admin/login?admin_id=admin001&role=super_admin")
ADMIN_TOKEN=$(echo $ADMIN_RESPONSE | grep -o '"token":"[^"]*"' | cut -d'"' -f4)

if [ -z "$ADMIN_TOKEN" ]; then
    echo -e "${RED}✗ 获取管理员 Token 失败${NC}"
    exit 1
fi

echo -e "${GREEN}✓ 管理员 Token 获取成功${NC}"
echo ""

# 设置 Authorization Header
AUTH_HEADER="Authorization: Bearer $ADMIN_TOKEN"

# 测试 Swagger 文档
echo -e "${YELLOW}步骤 2: 测试 Swagger 文档${NC}"
SWAGGER_CODE=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/swagger/index.html")
if [ "$SWAGGER_CODE" == "200" ]; then
    echo -e "${GREEN}✓ Swagger 文档可访问 (http://localhost:8080/swagger/index.html)${NC}"
    PASS=$((PASS + 1))
else
    echo -e "${RED}✗ Swagger 文档访问失败${NC}"
    FAIL=$((FAIL + 1))
fi
echo ""

# 测试管理员 API
echo -e "${YELLOW}步骤 3: 测试管理员 API${NC}"
echo ""

# 测试添加黑名单
test_api "添加黑名单" \
    "POST" \
    "$BASE_URL/api/admin/blacklist/add" \
    '{"address":"0x1234567890123456789012345678901234567890"}'

# 测试移除黑名单
test_api "移除黑名单" \
    "POST" \
    "$BASE_URL/api/admin/blacklist/remove" \
    '{"address":"0x1234567890123456789012345678901234567890"}'

# 测试暂停合约
test_api "暂停合约" \
    "POST" \
    "$BASE_URL/api/admin/pause" \
    ""

# 测试恢复合约
test_api "恢复合约" \
    "POST" \
    "$BASE_URL/api/admin/unpause" \
    ""

# 测试更新质押限额
test_api "更新质押限额" \
    "POST" \
    "$BASE_URL/api/admin/limits/update" \
    '{"min_amount":"100000000000000000","max_amount":"10000000000000000000"}'

# 测试更新收益费率
test_api "更新收益费率" \
    "POST" \
    "$BASE_URL/api/admin/rate/update" \
    '{"rate":"100"}'

# 测试查询合约状态
test_api "查询合约状态" \
    "GET" \
    "$BASE_URL/api/admin/status" \
    ""

# 测试无效 Token
echo -e "${YELLOW}步骤 4: 测试权限控制${NC}"
INVALID_RESPONSE=$(curl -s -w "\n%{http_code}" \
    -X POST \
    -H "Authorization: Bearer invalid_token" \
    "$BASE_URL/api/admin/pause")

INVALID_CODE=$(echo "$INVALID_RESPONSE" | tail -n1)
if [ "$INVALID_CODE" == "401" ]; then
    echo -e "${GREEN}✓ 无效 Token 被正确拒绝 (HTTP 401)${NC}"
    PASS=$((PASS + 1))
else
    echo -e "${RED}✗ 权限控制失效 (HTTP $INVALID_CODE)${NC}"
    FAIL=$((FAIL + 1))
fi
echo ""

# 测试结果汇总
echo "=========================================="
echo "测试结果汇总"
echo "=========================================="
echo -e "通过: ${GREEN}$PASS${NC}"
echo -e "失败: ${RED}$FAIL${NC}"
echo "总计: $((PASS + FAIL))"
echo ""

if [ $FAIL -eq 0 ]; then
    echo -e "${GREEN}🎉 所有测试通过！${NC}"
    exit 0
else
    echo -e "${RED}⚠️  部分测试失败，请检查日志${NC}"
    exit 1
fi
