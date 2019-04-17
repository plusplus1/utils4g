# utils4g Golang utils

---

```bash

# Redis config example

# prefix, all keys with "prefix:"
prefix: xxx
# 连接机器
host: 127.0.0.1
# 连接端口
port: 6379
# 认证信息
password: ""
# 默认连接的db
db: 1
# 最大连接数
max_connections: 1000

---

# Mongo db config example 
# The seed servers must be provided in the following format:
#
#     [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
#
# Relevant documentation:
#
#     http://docs.mongodb.org/manual/reference/connection-string/
#
connect_string: "mongodb://localhost/result"

# 连接池大小，控制最大并发，最大4096
pool_limit: 1

---

# Mysql config example
host: 127.0.0.1
port: 3306
user: root
password: root
db: test

params:
  charset: utf8
  
```