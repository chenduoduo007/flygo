# flygo
flygo is an open-source, easy to use web framework for the Go .

a simple web framework
   
go get -u github.com/chenduoduo007/flygo

go mod
    
    require (
    	github.com/chenduoduo007/flygo v0.0.2
    )

main.go

    package main
    
    import "github.com/chenduoduo007/flygo"
    
    func main()  {
        r := flygo.Default()
        r.GET("/", func(c *flygo.Context) {
            c.String(http.StatusOK, "Hello Flygo")
        })
    
        r.GET("/api", func(c *flygo.Context) {
            data := map[string]interface{}{
                "code": 200,
                "msg": "success",
                "data": map[string]interface{}{},
            }
            c.JSON(http.StatusOK, data)
        })
        endPoint := ":8080"
        log.Printf("[info] start http server listening %s", endPoint)
        err := r.Run(endPoint)
        if err != nil {
            log.Printf("[error] %s", err)
        }
    }