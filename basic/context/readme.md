## Context
- Sebuah data yang membawa value, sinyal cancel, sinyal timeout, dan sinyal deadline
- Context biasanya dibuat per request (misal tiap ada request masuk ke server web melalui http request)
- Context digunakan untuk mempermudah kita meneruskan value, dan sinyal antar process
- Context menganut konsep parent child (1 parent bisa bnyak child, tp 1 child cuman bisa 1 parent)
- Context merupakan object yg immutable (setelah dibuat maka data tidak dapat diubah, hanya bisa read saja)
- Jika value diubah maka akan membuat context baru

```go
type Context interface {
  Deadline()(deadline time.Time, ok bool)
  Done() <- chan struct{}
  Err() error
  Value(key interface{}) interface{} 
}
```

context.Background() -> context kosong. biasa digunakan di main func ato unit test
context.TODO() -> context kosong biasa digunakan ketika belum jelas context apa yang ingin digunakan