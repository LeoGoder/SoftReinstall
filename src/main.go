package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// var for install process for chocolatey
	var checkExecuPolicy = exec.Command("cmd", "/c", "powershell Get-ExecutionPolicy")
	var checkExecuPolicyOut, checkExecuPolicyErr = checkExecuPolicy.CombinedOutput()

	if checkExecuPolicyErr != nil {
		fmt.Println(checkExecuPolicyErr.Error())
	}

	execuPolicy := strings.TrimSpace(string(checkExecuPolicyOut))

	if execuPolicy != "restricted" {
		fmt.Println("Execution Policy is not restricted")

		var chocoinstall = exec.Command("cmd", "/c", "powershell Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))")
		var chocoinstallOut, chocoinstallErr = chocoinstall.CombinedOutput()

		if chocoinstallErr != nil {
			fmt.Println(chocoinstallErr.Error())
		}
		fmt.Println(string(chocoinstallOut))
	} else {
		fmt.Println("Execution Policy is restricted")

		var setExecuPolicy = exec.Command("cmd", "/c", "powershell Set-ExecutionPolicy AllSigned")
		var setExecuPolicyOut, setExecuPolicyErr = setExecuPolicy.CombinedOutput()

		if setExecuPolicyErr != nil {
			fmt.Println(setExecuPolicyErr.Error())
		}

		fmt.Println(string(setExecuPolicyOut))

	}

	// var for Install process with choco
	var wincmd = exec.Command("cmd", "/c", "choco install googlechrome -y")
	var winout, winerr = wincmd.CombinedOutput()

	if winerr != nil {
		fmt.Println(winerr.Error())
		return
	}

	fmt.Println(string(winout))
}
