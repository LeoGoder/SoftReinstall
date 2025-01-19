# SoftReinstall
This program as for goal to help install all the software need, especially on a fresh install of Windows 

## Installation
Download on the release page the exe

## build from source
Go to src and type 
```
go build
go install
```


## How to use
First of all check if the program work correctly by taping 

```
.\softInstall.exe version
```

if you don't have chocolatey not install type this command 
```
.\softInstall.exe choco
```

To see all possible soft to install 
```
.\softInstall.exe list
```

To launch an install 
```
.\softInstall.exe install <the software>
```
## Roadmap

- Additional OS support

- Improvement of soft selection


