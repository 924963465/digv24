package image

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
)

//按指定的图片的长边，生成缩略图
func ConvertByLong(orig,dest string,long int) error {
	longstr := strconv.Itoa(long)
	cmd := "/usr/bin/convert -size "+longstr+"x"+longstr+" -resize "+longstr+"x"+longstr+" +profile '*' "+orig+" "+dest+" 2>&1"
	out,err := execShell(cmd)
	fmt.Println(out)
	fmt.Println(err)
	if (err != nil) {
		return err
	}
	if (out != "") {
		return errors.New(out)
	}
	return nil
}

//执行shell命令
func execShell(s string) (string, error){
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	//checkErr(err)
	return out.String(), err
}
