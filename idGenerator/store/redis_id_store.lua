#!lua name=charlie_id_generator

-- 安装命令
-- cat redis_id_store.lua | redis-cli -x --cluster-only-masters --cluster call 192.168.123.30:6379 FUNCTION LOAD REPLACE

-- 标识分配的redis函数，
-- updateMachineCode，更新客户端机器码
-- allocateSerial，分配序列段。在分配序列段之前更新机器码

-- 默认机器码有效期秒数
local machineExpires = 60

local function _updateMachineCode(key, code, token, max)

    local machineKey = key..":allocated:"..tostring(code)
    if redis.call("GET", machineKey) == token then
        redis.call("EXPIRE", machineKey, machineExpires)
        return code
    end

    for i = 0, tostring(max), 1 do
        machineKey = key..":allocated:"..tostring(i)

        if redis.call("EXISTS", machineKey) == 0 then 
            redis.call("SET", machineKey, token, "EX", machineExpires)
            return i
        end
    end

    return -1
end

-- 请求参数：机器码当前值，机器标识，机器码最大值
-- 响应参数：分配的机器码。分配失败时返回-1
-- 机器标识在客户端创建时生成，分辨不同的客户。
-- FCALL updateMachineCode 1 "bbcc" -1 "aaaaa" 9
local function updateMachineCode(keys, args)
    local key = keys[1]
    local code = tonumber(args[1])
    local token = args[2]
    local max = tonumber(args[3])

    return _updateMachineCode(key, code, token, max)
end

-- 请求参数：机器码，机器标识，步长，序列最小值，序列最大值，机器码最大值
-- 响应参数：分配成功的机器码，序列起始值和结束值{machineCode, begin, finish}
local function allocateSerial(keys, args)
    local key = keys[1]

    local code = tonumber(args[1])
    local token = args[2]
    local step = tonumber(args[3])
    local min = tonumber(args[4])
    local max = tonumber(args[5])    
    local maxCode = tonumber(args[6])

    code = _updateMachineCode(key, code, token, maxCode)
    if code == -1 then
        -- 刷新机器码失败，响应错误信息。
        return {code, 0, 0, 0, "machine code allocation failed"}
    end
    
    if step > max then 
        step = max
    end
    
    key = key..":sequence"
    if redis.call("HEXISTS", key, code) == 0 then
        redis.call("HSET", key, code, step)
        
        return {code, min, step, 0, "success"}
    end

    local begin = tonumber(redis.call("HGET", key, code))
    local finish = redis.call("HINCRBY", key, code, step)
    local reback = 0

    -- 计算后的起始值超过最大值，从序列段起点重新开始
    if begin >= max then 
        begin = min
        finish = step
        redis.call("HSET", key, code, step)

        -- 检查上次绕回时间，判断是否需要检查时间段冲突
        reback = 1
    end

    -- 计算后结束值超过最大值
    if finish > max then
        finish = max
        redis.call("HSET", key, code, finish)
    end

    return {code, begin, finish, reback, "success"}
end

redis.register_function('updateMachineCode',updateMachineCode)
redis.register_function('allocateSerial',allocateSerial)
