package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	var i = true
	for i {
		// var for install process for chocolatey
		var checkExecuPolicy = exec.Command("cmd", "/c", "powershell Get-ExecutionPolicy")
		var checkExecuPolicyOut, checkExecuPolicyErr = checkExecuPolicy.CombinedOutput()

		if checkExecuPolicyErr != nil {
			fmt.Println(checkExecuPolicyErr.Error())
		}

		execuPolicy := strings.TrimSpace(string(checkExecuPolicyOut))

		if execuPolicy != "Restricted" {
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
		var wincmd = exec.Command("cmd", "/c", "choco install firefox -y && choco install discord.install -y && choco install vscode -y && choco install steam -y && choco install epicgameslauncher -y && choco install obsidian -y && choco install 7zip -y && choco install git -y && -y && choco install python -y && choco install golang -y && choco install docker-desktop -y && choco install vlc -y")
		var winout, winerr = wincmd.CombinedOutput()

		if winerr != nil {
			fmt.Println(winerr.Error())
			return
		}

		fmt.Println(string(winout))
	}

}
