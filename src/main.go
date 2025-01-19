package main

import (
	"fmt"
	"os/exec"
	"strings"

	//"os/exec"
	//"strings"

	"github.com/spf13/cobra"
)

func main() {
	var getVersion = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of the app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Version 0.1")
		},
	}

	var getListSoftware = &cobra.Command{
		Use: "list",
		Run: listOfSoftware,
	}

	var installSoftware = &cobra.Command{
		Use: "install",
		Run: installSoftware,
	}

	var chocolateyInstall = &cobra.Command{
		Use: "choco",
		Run: chocolateyInstall,
	}

	var rootCmd = &cobra.Command{
		Use:   "reinstall",
		Short: "select and install software on your computer",
	}
	rootCmd.AddCommand(getVersion, getListSoftware, installSoftware, chocolateyInstall)

	installSoftware.Flags().StringP("software", "s", "", "Software to install")
	rootCmd.Execute()
}

func chocolateyInstall(cmd *cobra.Command, args []string) {
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
		var chocoinstall = exec.Command("cmd", "/c", "powershell Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))")
		var chocoinstallOut, chocoinstallErr = chocoinstall.CombinedOutput()

		if chocoinstallErr != nil {
			fmt.Println(chocoinstallErr.Error())
		}
		fmt.Println(string(chocoinstallOut))

	}

}

func listOfSoftware(cmd *cobra.Command, args []string) {
	fmt.Println("List of Software")
	fmt.Println("1. Firefox: firefox")
	fmt.Println("2. Discord: discord.install")
	fmt.Println("3. Visual Studio Code: vscode")
	fmt.Println("4. Steam: steam")
	fmt.Println("5. Epic Games Launcher: epicgameslauncher")
	fmt.Println("6. Obsidian: obsidian")
	fmt.Println("7. 7zip: 7zip")
	fmt.Println("8. Git: git")
	fmt.Println("9. Python: python")
	fmt.Println("10. Golang: golang")
	fmt.Println("11. Docker Desktop: docker-desktop")
	fmt.Println("12. VLC: vlc")
	fmt.Println("13. Notion: notion")
	fmt.Println("14. Vmware: vmwareworkstation")
	fmt.Println("15. VirtualBox: virtualbox")
	fmt.Println("16. DS4: ds4windows")
	fmt.Println("17. Everything: everything")
}

func installSoftware(cmd *cobra.Command, args []string) {
	var commandArgs = append([]string{"/c", "choco", "install"}, args...)
	commandArgs = append(commandArgs, "-y")
	var wincmd = exec.Command("cmd", commandArgs...)
	var winout, winerr = wincmd.CombinedOutput()

	if winerr != nil {
		fmt.Println(winerr.Error())
		return
	}

	fmt.Println(string(winout))
}
