## SwiftlyGo Development Guide

### Build Host Requirements
The only requirement for the build host is that it can run Go version 1.20.0 or later.

For details on installing Go and setting up a Go build environment please see [Go Install](https://go.dev/doc/install)

### Target System Requirements 
The target system needs to be a Ubuntu or Debian based distribution that supports Swift.

To check if the target system is compatible, copy and run the `create-env.sh` script from the source directory.
```bash
chmod +x create-env.sh
./create-env.sh

Checking your OS system for Swift compatibility ...

Your OS System is compatible with Swift for Ubuntu 22.04 / x86_64

Creating swiftlygo.env file
```

</br>

## How to Build and Test SwiftlyGo on a Swift Compatible Host

### Step 1. Setup Go Environment
Install set up a Go build environment as detailed here - [Go Install](https://go.dev/doc/install)

```bash
# check go version is at least version 1.20.0
go version
go version go1.24.0 linux/arm64
```

### Step 2. Clone the SwiftlyGo Repository
Clone the SwiftlyGo repository and cd to the `source` directory

```bash
git clone https://github.com/cyberxdruid211/swiftlygo.git
#......
#......
cd source
ls
LICENSE  README.md  cmd  create_env.sh  go.mod  go.sum  logo.txt  main.go  swiftlygo
```

### Step 3. Build SwiftlyGo
Build `swiftlygo` and create the `swiftlygo.env` file.  
You will need to make `create-env.sh` executable.
```bash
# build swiftlygo
go build

# create swiftlygo.env
chmod +x create-env.sh
./create-env.sh
```

### Step 4. Test SwiftlyGo
You can now run and test SwiftlyGo using the `./swiftlygo` command
```bash
./swiftlygo --help

 ____          _  __ _   _        ____       
/ ___|_      _(_)/ _| |_| |_   _ / ___| ___  
\___ \ \ /\ / / | |_| __| | | | | |  _ / _ \ 
 ___) \ V  V /| |  _| |_| | |_| | |_| | (_) |
|____/ \_/\_/ |_|_|  \__|_|\__, |\____|\___/ 
                           |___/  v1.0      

An Installer for the Swift Programming Language.

You can install and delete any Swift version that is available for your platform.
For example:

swiftlygo install 6.0.3
This command will install Swift version 6.0.3.

swiftlygo delete 6.0.3
This command will delete Swift version 6.0.3.

The activate command will activate a locally installed Swift version.
For example:

swiftlygo activate 6.0.3
This command will activate Swift version 6.0.3.

Usage:
  swiftlygo [command]

Available Commands:
  activate        Activates a Swift version
  delete          Deletes a Swift version
  help            Help about any command
  install         Installs a Swift version
  list            Lists the Swift versions available
  removeSwiftlygo Removes the SwiftlyGo app

Flags:
  -h, --help   help for swiftlygo
  -y, --yes    Automatically answer 'yes' to all prompts

Use "swiftlygo [command] --help" for more information about a command.
```

</br>

## How to Install SwiftlyGo on the Build Host

SwiftlyGo is installed in the `/usr/libexec/swiftlygo/bin` directory with a symlink to the `/usr/bin/` directory.

```bash
# swiftlygo installation location
# /usr/libexec/swiftlygo/bin/
# ├── swiftlygo -> symlink to /usr/bin/swiftlygo
# └── swiftlygo.env
#
```

```bash
# create swiftlygo/bin directory
sudo mkdir -p /usr/libexec/swiftlygo/bin

#  copy swiftlygo and swiftlygo.env to the directory
sudo cp swiftlygo /usr/libexec/swiftlygo/bin/swiftlygo
sudo cp swiftlygo.env /usr/libexec/swiftlygo/bin/swiftlygo.env

# create symlink
sudo ln -sf /usr/libexec/swiftlygo/bin/swiftlygo /usr/bin/swiftlygo
```
You can now run and use SwiftlyGo using the `swiftlygo` command

</br>

## How to Build and Install SwiftlyGo on a Target System

### Step 1. Setup Go Environment
Install set up a Go build environment as detailed here - [Go Install](https://go.dev/doc/install)

```bash
# check go version is at least version 1.20.0
go version
go version go1.24.0 linux/arm64
```

### Step 2. Clone the SwiftyGo Repository
Clone the SwiftlyGo repository and cd to the `source` directory
```bash
git clone https://github.com/cyberxdruid211/swiftlygo.git
#......
#......
cd source
ls
LICENSE  README.md  cmd  create_env.sh  go.mod  go.sum  logo.txt  main.go  swiftlygo
```

### Step 3. Build SwiftlyGo
#### On the build machine
To build `swiftlygo` for a target you will need to know the cpu arch-type of the target.  
The arch-types supported are x86_64/amd64 or aarch64/arm64.  
NOTE: It does not matter what the arch-type of the build host is, only the target.

```bash
# build swiftlygo
# for installation on x86_64
GOOS=linux GOARCH=amd64 go build -o swiftlygo-x86_64 

# for  installation on arm64 (aarch64)
GOOS=linux GOARCH=arm64 go build -o swiftlygo-aarch64
```
### Step 4 Copy Files to Target System
Copy the `swiftlygo` binary that matches the arch-type and the `create-env.sh` file to a temp location on the target machine.  
e.g. `$USER/swiftlygo-setup`

### Step 5. Install SwiftlyGo on the Target System
#### On the target machine
Go to the temp location of the SwiftlyGo files

```bash
cd swiftlygo-setup
# check contents of directory
ls
create-env.sh swiftlygo-x86_64 
```
Create the `swiftlygo.env` file by running the `create-env.sh` script.
```bash
# make create-env.sh executable
chmod +x create-env.sh
# run script
./create-env.sh
# we should now have 3 files
ls
create-env.sh swiftlygo.env swiftlygo-x86_64
```
Create the SwiftlyGo install directory
```bash
sudo mkdir -p /usr/libexec/swiftlygo/bin/
```
Move SwiftlyGo and .env file to install directory
```bash
# move swiftlygo
sudo mv swiftlygo-x86_64 /usr/libexec/swiftlygo/bin/swiftlygo
# make executable
sudo chmod +x /usr/libexec/swiftlygo/bin/swiftlygo
# move .env
sudo mv swiftlygo.env /usr/libexec/swiftlygo/bin/swiftlygo.env
```
Add symlink for `swiftlygo`
```bash
# create symlink
sudo ln -sf /usr/libexec/swiftlygo/bin/swiftlygo /usr/bin/swiftlygo
```

SwiftlyGo is now installed and ready to use.  
You can remove the `swiftlygo-setup` directory as it is no longer needed.










