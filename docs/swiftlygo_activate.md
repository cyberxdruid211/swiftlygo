## Activate Command

```yaml
sudo swiftlygo activate
```

Activates a Swift version

*NOTE: **_swiftlygo_** must be run with administrator privileges (sudo)*

### Activating a Swift Version

*NOTE: To activate a ***Swift*** version it must already be installed locally*

This command activates the Swift version requested.

For example:

- This command will activate Swift version 6.0.3.

```yaml
sudo swiftlygo activate 6.0.3
```
```yaml
Are you sure you want to activate version 6.0.3? [Y/n]: y
Version 6.0.3 is now activated.
```

---

### Command

```yaml
swiftlygo delete [flags]
```

### Options

```yaml
-h, --help             help for delete
```

### Options inherited from parent commands

```yaml
  -y, --yes   Automatically answer 'yes' to all prompts
```

### SEE ALSO

- [swiftlygo](../README.md) - An Installer for the Swift Programming Language
