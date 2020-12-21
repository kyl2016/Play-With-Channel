# redis

## 订阅 Keyspace 的 key expired

## key 失效与 set 并发

业务需求：会收到同一个 key 的数据，收到后，expire 延长 15 秒，15 秒内没收到新数据，则移除。
支持并发，如果 new 数据和过期同时发生，则不要移除数据。

每次新增 key，不设置过期时间；同时，增加一个`key_expire`的 key 来处理过期。
当`key_expire`过期时，需要判断`key`中的`lastTime`与当前时间间隔
- 如果>=15秒，则过期时没有新数据到来，进行业务处理，然后移除`key`
- 如果<15 秒，则过期时又收到了新数据，则不处理。

## Config

- K     Keyspace events, published with __keyspace@<db>__ prefix.
- E     Keyevent events, published with __keyevent@<db>__ prefix.
- g     Generic commands (non-type specific) like DEL, EXPIRE, RENAME, ...
- $     String commands
- l     List commands
- s     Set commands
- h     Hash commands
- z     Sorted set commands
- t     Stream commands
- x     Expired events (events generated every time a key expires)
- e     Evicted events (events generated when a key is evicted for maxmemory)
- m     Key miss events (events generated when a key that doesn't exist is accessed)
- A     Alias for "g$lshztxe", so that the "AKE" string means all the events except "m".