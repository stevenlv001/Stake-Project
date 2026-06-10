#!/bin/bash

# 项目改进验证脚本
# 用于测试所有改进功能是否正常工作

BASE_URL="http://localhost:8080"
echo "=========================================="
echo "  Stake Project 改进功能验证"
echo "=========================================="
echo ""

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 测试结果统计
PASS=0
FAIL=0

# 测试函数
test_api() {
    local name=$1
    local method=$2
    local url=$3
    local headers=$4
    local data=$5
    
    echo -e "${YELLOW}测试: $name${NC}"
    
    if [ "$method" == "GET" ]; then
        response=$(curl -s -w "\n%{http_code}" -H "$headers" "$BASE_URL$url")
    else
        response=$(curl -s -w "\n%{http_code}" -X POST -H "$headers" -H "Content-Type: application/json" -d "$data" "$BASE_URL$url")
    fi
    
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | sed '$d')
    
    if [ "$http_code" -ge 200 ] && [ "$http_code" -lt 300 ]; then
        echo -e "${GREEN}✓ 通过 (HTTP $http_code)${NC}"
        PASS=$((PASS+1))
    else
        echo -e "${RED}✗ 失败 (HTTP $http_code)${NC}"
        echo "响应: $body"
        FAIL=$((FAIL+1))
    fi
    echo ""
}

echo "1. 检查服务是否运行..."
health_response=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/swagger/index.html")
if [ "$health_response" != "200" ]; then
    echo -e "${RED}错误: 服务未运行或Swagger不可访问${NC}"
    echo "请先启动服务: go run cmd/main.go"
    exit 1
fi
echo -e "${GREEN}✓ 服务运行正常${NC}"
echo ""

echo "2. 测试环境变量配置..."
echo "检查JWT_SECRET是否从环境变量读取..."
# 这个需要查看启动日志，这里只做提示
echo -e "${YELLOW}提示: 请检查启动日志中是否有 '使用环境变量 JWT_SECRET' 的提示${NC}"
echo ""

echo "3. 测试管理员登录（数据库验证）..."
# 使用默认管理员地址测试
test_api "管理员登录" "GET" "/api/admin/login?admin_id=0x0000000000000000000000000000000000000001" "" ""

echo "4. 测试合约状态查询（带缓存）..."
ADMIN_TOKEN=$(curl -s "$BASE_URL/api/admin/login?admin_id=0x0000000000000000000000000000000000000001" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
if [ -n "$ADMIN_TOKEN" ]; then
    test_api "查询合约状态" "GET" "/api/admin/status" "Authorization: Bearer $ADMIN_TOKEN" ""
else
    echo -e "${RED}无法获取管理员Token，跳过后续测试${NC}"
    exit 1
fi

echo "5. 测试用户登录..."
USER_ADDRESS="0x1234567890123456789012345678901234567890"
test_api "用户登录" "GET" "/api/login?address=$USER_ADDRESS" "" ""

USER_TOKEN=$(curl -s "$BASE_URL/api/login?address=$USER_ADDRESS" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)

if [ -n "$USER_TOKEN" ]; then
    echo "6. 测试质押历史查询（新接口）..."
    test_api "查询质押历史" "GET" "/api/stake/history?page=1&size=10" "Authorization: Bearer $USER_TOKEN" ""
    
    echo "7. 测试交易状态查询（统一格式化）..."
    test_api "查询交易状态" "GET" "/api/tx/status?tx_hash=0x0000000000000000000000000000000000000000000000000000000000000000" "Authorization: Bearer $USER_TOKEN" ""
else
    echo -e "${YELLOW}警告: 无法获取用户Token，跳过部分测试${NC}"
fi

echo "8. 测试错误处理（统一错误码）..."
# 故意发送错误请求测试错误处理
response=$(curl -s -X POST -H "Authorization: Bearer $USER_TOKEN" -H "Content-Type: application/json" -d '{"amount":"invalid"}' "$BASE_URL/api/stake/do")
if echo "$response" | grep -q "code"; then
    echo -e "${GREEN}✓ 错误处理正常（返回错误码）${NC}"
    PASS=$((PASS+1))
else
    echo -e "${RED}✗ 错误处理异常${NC}"
    FAIL=$((FAIL+1))
fi
echo ""

echo "=========================================="
echo "  测试结果汇总"
echo "=========================================="
echo -e "通过: ${GREEN}$PASS${NC}"
echo -e "失败: ${RED}$FAIL${NC}"
echo ""

if [ $FAIL -eq 0 ]; then
    echo -e "${GREEN}🎉 所有测试通过！${NC}"
    exit 0
else
    echo -e "${YELLOW}⚠️  部分测试失败，请检查日志${NC}"
    exit 1
fi
