package easynginx

import (
	"fmt"
	"os/exec"
	"log"
)


func Get_conf_path() (string,error) {
	cmd := exec.Command("nginx", "-V")
	var stderr bytes.Buffer
	cmd.Stderr=&stderr
	cmd.Stdout=nil

	if err := cmd.Run(); err !- nil{
		return "", fmt.Errorf("failed to run nginx -V: %w if you dont have nginx please run sudo apt install nginx -y", err)
	}
	output:=stderr.String()
}
