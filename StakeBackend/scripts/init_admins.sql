-- 初始化默认管理员数据
-- 注意：生产环境应该通过更安全的方式添加管理员

INSERT INTO admins (admin_id, role, is_active, created_at, updated_at) 
VALUES 
('0x0000000000000000000000000000000000000001', 'super_admin', true, NOW(), NOW()),
('0x0000000000000000000000000000000000000002', 'admin', true, NOW(), NOW())
ON DUPLICATE KEY UPDATE 
    role = VALUES(role),
    is_active = VALUES(is_active),
    updated_at = NOW();
