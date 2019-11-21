package main
import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
// 初始化一个新的redis client
client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "", // redis没密码，没有设置，则留空
		DB:       0,  // 使用默认数据库
	})
	
// 设置一个key，过期时间为0，意思就是永远不过期
err := client.Set("key", "valuevvvv", 0).Err()

// 检测设置是否成功
if err != nil {
	panic(err)
}

// 根据key查询缓存，通过Result函数返回两个值
//  第一个代表key的值，第二个代表查询错误信息
val, err := client.Get("key").Result()

// 检测，查询是否出错
if err != nil {
	panic(err)
}
fmt.Println("key", val)
}