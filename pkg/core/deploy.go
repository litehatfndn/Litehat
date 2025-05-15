package core

import "fmt"

func DeployMCPServer(env string, cpu int, memory int) {
    fmt.Printf("Deploying MCP server in %s with %d CPU and %dMB memory\n", env, cpu, memory)
}
