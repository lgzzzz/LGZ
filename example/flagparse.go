package example

import (
	"flag"
	"fmt"
	"os"
)

type MyFlagSet struct {
	*flag.FlagSet
	cmdComment string // 二级子命令本身的注释
}

func flagParse() {
	flag.Parse()
	// docker ps
	psCmd := &MyFlagSet{
		FlagSet:    flag.NewFlagSet("ps", flag.ExitOnError),
		cmdComment: "List containers",
	}
	psCmd.Bool("a", false, "Show all containers (default shows just running)")
	psCmd.Bool("s", false, "Display total file sizes")

	// docker run
	runCmd := &MyFlagSet{
		FlagSet:    flag.NewFlagSet("run", flag.ExitOnError),
		cmdComment: "Run a command in a new container",
	}
	runCmd.Int("c", 1, "CPU shares (relative weight)")
	runCmd.String("name", "", "Assign a name to the container")

	// 用 map 保存所有的二级子命令，方便快速查找
	subcommands := map[string]*MyFlagSet{
		psCmd.Name():  psCmd,
		runCmd.Name(): runCmd,
	}

	useage := func() { // 整个命令行的帮助信息
		fmt.Printf("Usage: docker COMMAND\n\n")
		for _, v := range subcommands {
			fmt.Printf("%s %s\n", v.Name(), v.cmdComment)
			v.PrintDefaults() // 使用 flag 库自带的格式输出子命令的选项帮助信息
			fmt.Println()
		}
		os.Exit(2)
	}

	if len(os.Args) < 2 { // 即没有输入子命令
		useage()
	}

	cmd := subcommands[os.Args[1]] // 第二个参数必须是我们支持的子命令
	if cmd == nil {
		useage()
	}

	cmd.Parse(os.Args[2:]) // 注意这里是 cmd.Parse 不是 flag.Parse，且值是 Args[2:]

	// 输出解析后的结果
	fmt.Println("command name is:", cmd.Name())
	cmd.Visit(func(f *flag.Flag) {
		fmt.Printf("option %s, value is %s\n", f.Name, f.Value)
	})
}
