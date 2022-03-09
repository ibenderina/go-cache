# go-cache

Simple in-memory cache

##Getting Cache

`go get github.com/ibenderina/go-cache`

##Import 
`import github.com/ibenderina/go-cache`

##Usage
```
func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)
}
```
