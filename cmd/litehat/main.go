package main

import (
    "fmt"
    "os"
    "github.com/litehat/litehat/pkg/core"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Expected 'deploy', 'monitor' or 'governance' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "deploy":
        core.DeployMCPServer("test", 2, 4096)
    case "monitor":
        fmt.Println("Monitoring server...")
        // Add monitoring logic here
    case "governance":
        fmt.Println("Governance command...")
        // Add governance logic here
    default:
        fmt.Println("Unknown command:", os.Args[1])
    }
}
