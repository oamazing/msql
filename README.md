# msql
## mysql 封装

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


