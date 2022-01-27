package main

import(
  "github.com/alecthomas/kong"
)


var CLI struct {
    Rm struct {
    Force     bool `help:"Force removal."`
    Recursive bool `help:"Recursively remove files."`

    Paths []string `arg:"" name:"path" help:"Paths to remove." type:"path"`
  } `cmd:"" help:"Remove files."`
}

func main() {
    ctx := kong.Parse(&CLI)
    switch ctx.Command() {
      case "rm <path>":
      default:
        panic(ctx.Command())
    }
}
