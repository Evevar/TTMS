package main

//	func main() {
//		nums := []int{200, 1, 3, 4}
//		fmt.Println(nums)
//	}
//func main() {
//	ctx := context.Background()
//
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6378",
//		Password: consts.RedisPassword, // no password set
//		DB:       0,                    // use default DB
//	})
//
//	err := rdb.Set(ctx, "key", "value11", 0).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	val, err := rdb.Get(ctx, "key").Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("key", val)
//
//	val2, err := rdb.Get(ctx, "key2").Result()
//	if err == redis.Nil {
//		fmt.Println("key2 does not exist")
//	} else if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("key2", val2)
//	}
//}
