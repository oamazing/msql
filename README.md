# msql
## mysql 封装

[![Build Status](https://github.com/oamazing/msql/actions/workflows/go.yml/badge.svg)](https://github.com/oamazing/msql/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/oamazing/msql/badge.svg?branch=master)](https://coveralls.io/github/oamazing/msql)
[![Go Report Card](https://goreportcard.com/badge/github.com/oamazing/msql)](https://goreportcard.com/report/github.com/oamazing/msql)
[![Documentation](https://pkg.go.dev/badge/github.com/oamazing/msql)](https://pkg.go.dev/github.com/oamazing/msql@v0.0.3)

---
## Example:
```go
db, err = Open(source)
if err != nil {
    Panic(err)
}
db.SetTimeOut(time.Minute)

var ts struct {
    Id        int64     `json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
}
if err := db.Query(&ts, `SELECT id,name,created_at FROM tests WHERE id = 1`); err != nil {
    return err
}
```
---

> scan basic
- [x] int
- [x] int64 
- [x] bool
- [x] time
- [x] json
- [ ] float
> other
- [x] struct
- [ ] 事务
- [x] slice

---

Todo:

- [x] 获取结构体字段
- [ ] 结构体值获取


