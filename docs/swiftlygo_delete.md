## Delete Command

```yaml
sudo swiftlygo delete
```

Deletes a Swift version

*NOTE: **_swiftlygo_** must be run with administrator privileges (sudo)*

### Deleting Swift

This command deletes the Swift version requested.

For example:

- This command will delete Swift version 5.10.1.

```yaml
sudo swiftlygo delete 5.10.1
```
```yaml
Are you sure you want to delete version 5.10.1? [Y/n]: y
Deleting ...
Version 5.10.1 deleted.
```

The delete command will also deactivate this version if it is currently active.

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

