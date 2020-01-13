package tool

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

func MakeDir(path string) error {
	_, err := os.Stat(path)
	// 目录已存在
	if err == nil {
		return nil
	} else {
		// 创建文件夹
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("mkdir failed![%v]\n", err)
		} else {
			return nil
		}
	}
}

func Home() (string, error) {
	path, err := user.Current()
	if nil == err {
		return path.HomeDir, nil
	}
	// cross compile support
	if "windows" == runtime.GOOS {
		return homeWindows()
	}
	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}
	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}
	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}
	return home, nil
}
