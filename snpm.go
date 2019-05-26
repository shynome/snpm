package snpm

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

// Snpm ee
type Snpm struct {
}

func checkIsLifecycleScript(cmd string) bool {
	return strings.HasPrefix(cmd, "pre") || strings.HasPrefix(cmd, "post")
}

// Exec npm script
func Exec(stage string, args []string, pkg Package) (err error) {

	cmd := pkg.Scripts[stage]
	isLifecycleScript := checkIsLifecycleScript(stage)
	if cmd == "" {
		return fmt.Errorf("can't find script: %v", stage)
	}
	env := pkg.getEnv()
	env = append(env, fmt.Sprintf("npm_lifecycle_script=%v", stage))

	preScript := fmt.Sprintf("pre%v", stage)
	if !isLifecycleScript && pkg.Scripts[preScript] != "" {
		if err = Exec(preScript, []string{}, pkg); err != nil {
			return
		}
	}

	fmt.Printf(
		"> %v %v %v \n",
		pkg.ENV["npm_package_name"],
		stage,
		pkg.DIR,
	)
	printArgsByte, err := json.Marshal(args)
	if err != nil {
		return
	}
	fmt.Printf("> %v \"%v\" \n", cmd, string(printArgsByte))
	if err = runCmd(cmd, args, env, pkg); err != nil {
		return err
	}

	postScript := fmt.Sprintf("post%v", stage)
	if !isLifecycleScript && pkg.Scripts[postScript] != "" {
		if err = Exec(postScript, []string{}, pkg); err != nil {
			return
		}
	}

	return nil

}

func runCmd(cmd string, args []string, env []string, pkg Package) error {

	args = append([]string{"-c", cmd + " " + strings.Join(args, " ")})

	proc := exec.Command("sh", args...)
	proc.Env = append(os.Environ(), env...)
	proc.Dir = pkg.DIR
	proc.Stdin = os.Stdin
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs,
		syscall.SIGHUP,  //终端控制进程结束(终端连接断开)
		syscall.SIGINT,  //用户发送INTR字符(Ctrl+C)触发
		syscall.SIGQUIT, //用户发送QUIT字符(Ctrl+/)触发
		syscall.SIGABRT, //调用abort函数触发
		syscall.SIGKILL, //无条件结束程序(不能被捕获、阻塞或忽略)
		syscall.SIGPIPE, //消息管道损坏(FIFO/Socket通信时，管道未打开而进行写操作)
		syscall.SIGTERM, //结束程序(可以被捕获、阻塞或忽略)
	)
	handleSIGTERM := func() {
		for {
			if proc.ProcessState != nil && proc.ProcessState.Exited() {
				break
			}
			sig := <-sigs
			proc.Process.Signal(sig)
		}
	}
	go handleSIGTERM()

	return proc.Run()

}
