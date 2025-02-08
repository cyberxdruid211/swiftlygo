## Install Command
```yaml
sudo swiftlygo install
```
Installs Swift and Swift dependencies

*NOTE: ***swiftlygo*** must be run with administrator privileges (sudo)*

### Installing Swift 
 
This command installs the Swift version requested.

For example:
- This command will install the latest Swift version available.
```yaml
sudo swiftlygo install latest
```
```yaml
Do you want to install Swift version 6.0.3? [Y/n]: y
Installing ...
2025/02/07 23:58:45 Downloading from: https://download.swift.org/swift-6.0.3-release/debian12-aarch64/swift-6.0.3-RELEASE/swift-6.0.3-RELEASE-debian12-aarch64.tar.gz
Downloading Swift ... -
File successfully downloaded.
Extracting Swift to: /usr/libexec/swift/6.0.3
Extracting Swift file ... -
Swift 6.0.3 has been installed successfully.

swift --version
Swift version 6.0.3 (swift-6.0.3-RELEASE)
Target: aarch64-unknown-linux-gnu
```
- This command will install Swift version 5.10.1.

```yaml
sudo swiftlygo install 5.10.1
```
- If the Swift version requested is available locally the install command ask if you wish to activate this version.

#### Installing Swift Dependencies

When you use the install command it will automatically check that you have the required dependencies for Swift.  
During the installation process you will be prompted to install dependencies if needed.
```yaml
sudo swiftlygo install latest

Do you want to install Swift version 6.0.3? [Y/n]: 
Installing ...
2025/02/07 15:56:41 Downloading from: https://download.swift.org/swift-6.0.3-release/ubuntu2404/swift-6.0.3-RELEASE/swift-6.0.3-RELEASE-ubuntu24.04.tar.gz
Downloading Swift ... -
File successfully downloaded.
Extracting Swift to: /usr/libexec/swift/6.0.3
Extracting Swift file ... -
Swift 6.0.3 has been installed successfully.

Some dependencies required for Swift are missing. Do you want to install them? [Y/n]:
```

- You can also manually install the dependencies with this command.
```yaml
sudo swiftlygo install depends
```
```yaml

Do you want to install the dependencies required for Swift? [Y/n]: y
Installing ...
Installing dependencies: build-essential git gnupg2 libcurl4 libedit2 libncurses-dev libpython3-dev libxml2 libz3-dev pkg-config python3 tzdata unzip zlib1g-dev
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done
build-essential is already the newest version (12.9).
git is already the newest version (1:2.39.2-1.1).
gnupg2 is already the newest version (2.2.40-1.1).
libcurl4 is already the newest version (7.88.1-10+deb12u6).
libedit2 is already the newest version (3.1-20221030-2).
libncurses-dev is already the newest version (6.4-4).
libpython3-dev is already the newest version (3.11.2-1+b1).
libxml2 is already the newest version (2.9.14+dfsg-1.3~deb12u1).
libz3-dev is already the newest version (4.8.12-3.1).
pkg-config is already the newest version (1.8.1-1).
python3 is already the newest version (3.11.2-1+b1).
tzdata is already the newest version (2024a-0+deb12u1).
unzip is already the newest version (6.0-28).
zlib1g-dev is already the newest version (1:1.2.13.dfsg-1).
0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
Swift dependency installation finished.

```
---

### Command
```yaml
sudo swiftlygo install [flags]
```

### Options

```yaml
  -h, --help             help for install
```

### Options inherited from parent commands

```yaml
  -y, --yes   Automatically answer 'yes' to all prompts
```

### SEE ALSO

* [swiftlygo](../README.md)	 - An Installer for the Swift Programming Language
