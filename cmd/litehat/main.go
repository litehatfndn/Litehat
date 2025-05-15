package main

import (
    "fmt"
    "github.com/litehat/litehat/pkg/core"
)

func main() {
    fmt.Println("Litehat MCP CLI")
    core.DeployMCPServer("test", 2, 4096)
}
