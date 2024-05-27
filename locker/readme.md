同步锁

EmptyLocker， 空锁
RWLocker, 读写锁
SpinLocker, 旋转锁


锁可以添加一个外部存储成为分布式锁。WithRedis, WithZookeeper

单例锁


资源锁


分布式锁
在锁的基础上添加分布式存储升级为分布式锁 

locker.WithRedis()
locker.WithZookeeper()


