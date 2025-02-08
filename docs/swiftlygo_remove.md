## Remove SwiftlyGo Command

```yaml
sudo swiftlygo removeSwiftlygo
```
Removes the SwiftlyGo app.

*NOTE: **_swiftlygo_** must be run with administrator privileges (sudo)*

### Removing the SwiftlyGo App 
 
This command removes the SwiftlyGo app.

This command only removes the SwiftlyGo app.  
All Swift installations will remain on your system.

For example:
```yaml
sudo swiftlygo removeSwiftlygo
```
```yaml
Are you sure you want to remove SwiftlyGo? [Y/n]: y
Removing swiftlygo ...
removed swiftlygo symlink - /usr/bin/swiftlygo
removed swiftlygo directory - /usr/libexec/swiftlygo
SwiftlyGo removed successfully
```

SwiftlyGo can easily be re-installed by running the install script.
```yaml
curl -sL https://swiftlygo.xyz/install.sh | sudo bash
```
