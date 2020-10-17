package fly

import (
	"fmt"
	"log"
	"net/http"
)

//  request handler
type HandlerFunc func(http.ResponseWriter, *http.Request)

// 基于ServeHTTP实现Engine
type Engine struct {
	router map[string]HandlerFunc
}

// 创建fly.Engine
func New() *Engine {
	// 添加.env的变量
	//AddEnv()
	return &Engine{router: make(map[string]HandlerFunc)}
}


// 模仿gin 添加中间件
func Default() *Engine {
	engine := New()
	// todo some thing eg：error log and recovery

	return engine
}

//func addEnv()  {
//	file, err := os.Open(".env")
//	//读取.env文件，并建文件写入到环境中
//	yaml.Marshal()
//	os.Setenv()
//}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler
}

// 定义GET方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// 定义POST方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 运行服务
func (engine *Engine) Run(addr ...string) (err error) {
	// 处理port
	addrNew := resolveAddress(addr)
	return http.ListenAndServe(addrNew, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
