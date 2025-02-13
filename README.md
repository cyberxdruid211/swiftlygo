[![Build SwiftlyGo](https://github.com/cyberxdruid211/swiftlygo/actions/workflows/build-swiftlygo.yaml/badge.svg)](https://github.com/cyberxdruid211/swiftlygo/actions/workflows/build-swiftlygo.yaml)

# SwiftlyGo

An Installer and Toolchain Manager for the **Swift** Programming Language on Linux

**SwiftlyGo** v1.0 is compatible with x86_64 and aarch64 Linux distributions based on *Debian* and *Ubuntu*.  
This includes popular releases such as *LinuxMint*, *popOS*, *RaspberryPi OS*, *Zorin*, *Elementary*, etc.

With **SwiftlyGo** you can quickly and easily install any **Swift** toolchain version that is available for your platform. 

The **Swift** toolchain packages are official releases from *[swift.org](https://www.swift.org/)* and are sourced directly from the *[swift.org downloads](https://www.swift.org/install/linux/#platforms)*.

Multiple versions of **Swift** can be installed and you can switch between active versions in seconds.

### Installation
Installation of **SwiftlyGo** is quick and simple using the `install` script.
```yaml
curl -L https://swiftlygo.xyz/install.sh | bash
```
```yaml
Checking your OS system for Swift compatibility ...

Success! Your OS System is compatible with Swift for Debian 12 / aarch64

Installing SwiftlyGo ...
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
100 7296k  100 7296k    0     0  5220k      0  0:00:01  0:00:01 --:--:-- 11.4M

Congratulations! Swiftlygo has been successfully installed.
 ____          _  __ _   _        ____       
/ ___|_      _(_)/ _| |_| |_   _ / ___| ___  
\___ \ \ /\ / / | |_| __| | | | | |  _ / _ \ 
 ___) \ V  V /| |  _| |_| | |_| | |_| | (_) |
|____/ \_/\_/ |_|_|  \__|_|\__, |\____|\___/ 
                           |___/  v1.0           

For help run 'swiftlygo -h'
```
```yaml
swiftlygo -h
```
```yaml
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

### Usage

Easy to use commands
 * `swiftlygo list`     - lists the **Swift** versions available for your system and versions locally installed.
 * `swiftlygo install` - installs a **Swift** version (also installs Swift dependencies if needed).
 * `swiftlygo activate` - activates a **Swift** version.
 * `swiftlygo delete` - deletes a **Swift** version.
 

### Quick Start
If you just want to install the latest **Swift** version available: -
```yaml
sudo swiftlygo install latest
```
If you want to install a specific **Swift** version: -
```yaml
sudo swiftlygo install 6.0.1
```
If you need to install the dependencies required for **Swift**: -
```yaml
sudo swiftlygo install depends
```

### How to uninstall SwiftlyGo
If you want to uninstall **SwiftlyGo** go can use the removeSwiftlygo command: -
```yaml
sudo swiftlygo removeSwiftlygo
```
### For more details on each command

* [swiftlygo list](docs/swiftlygo_list.md)
* [swiftlygo install](docs/swiftlygo_install.md)
* [swiftlygo activate](docs/swiftlygo_activate.md)
* [swiftlygo delete](docs/swiftlygo_delete.md)
* [swiftlygo removeSwiftlygo](docs/swiftlygo_remove.md)
