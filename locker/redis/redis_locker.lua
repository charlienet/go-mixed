#!lua name=charlie_locker

-- 安装命令
-- cat redis_locker.lua | redis-cli -x --cluster-only-masters --cluster call 192.168.123.30:6379 FUNCTION LOAD REPLACE

local function lock(keys, args) 
    if redis.call("GET", keys[1]) == args[1] then
        redis.call("SET", keys[1], args[1], "PX", args[2])
        return "OK"
    else
        return redis.call("SET", keys[1], args[1], "NX", "PX", args[2])
    end
end

local function del(keys, args)
    if redis.call("GET", keys[1]) == args[1] then
        return redis.call("DEL", keys[1])
    else
        return '0'
    end
end

local function expire(keys, args)
    if redis.call('get', keys[1]) == args[1] then
        return redis.call('expire', keys[1], args[2])
    else
        return '0'
    end
end

redis.register_function('locker_lock',lock)
redis.register_function('locker_unlock',del)
redis.register_function('locker_expire',expire)
