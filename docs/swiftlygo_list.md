## List Command

```yaml
sudo swiftlygo list
```

Lists the Swift versions available

*NOTE: ***swiftlygo*** must be run with administrator privileges (sudo)*

### Listing Available Swift Versions

This command lists the Swift versions available that can be installed on your platform.  
It will also show the Swift versions that are already downloaded on your system
and the Swift version currently active.

For example:

```yaml
sudo swiftlygo list
```
```yaml
Swift versions available for installation on this system:-
5.10.1, 6.0, 6.0.1, 6.0.2, 6.0.3

Locally available versions:-
5.10.3, 6.0.3

The Swift version currently active is - 6.0.3
```

---

### Command

```yaml
swiftlygo list [flags]
```

### Options

```yaml
  -h, --help   help for list
```

### SEE ALSO

- [swiftlygo](../README.md) - An Installer for the Swift Programming Language

