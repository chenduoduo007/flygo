package fly

import "os"

func resolveAddress(addr []string) string {
	println(os.Getenv("HOME"),"port")
	switch len(addr) {
	case 0:
		if port := os.Getenv("PORT"); port != "" {
			return ":" + port
		}
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}