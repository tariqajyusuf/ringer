# Package Definitions

Each file in this directory defines a package that Ringer can install. Files must be named `<package-name>.package.yaml`, where `<package-name>` is the identifier used with `ringer add` and `ringer remove` and in guise files.

## Schema

```yaml
name: "Display Name"
description: "Brief description of the package."
platforms:
  homebrew:
    name: "formula-or-cask-name"
  windows:
    name: "Publisher.PackageId"
```

### Fields

| Field | Required | Description |
|---|---|---|
| `name` | Yes | Human-readable display name shown in CLI output. |
| `description` | Yes | One-sentence description of what the package does. |
| `platforms` | Yes | Map of platform keys to platform-specific configuration. At least one platform entry is required. |

### Platform keys

| Key | Package manager | Supported OS |
|---|---|---|
| `homebrew` | [Homebrew](https://brew.sh/) | macOS, Linux |
| `windows` | [Winget](https://learn.microsoft.com/en-us/windows/package-manager/winget/) | Windows |

### Per-platform fields

| Field | Required | Description |
|---|---|---|
| `name` | Yes | The package identifier used by that platform's package manager (e.g. a Homebrew formula/cask name, or a winget package ID). |

## Examples

**Cross-platform package** — available on all supported platforms:

```yaml
name: "Git"
description: "A distributed version control system for tracking changes in source code."
platforms:
  homebrew:
    name: "git"
  windows:
    name: "Git.Git"
```

**macOS-only package** — omit platforms that don't apply:

```yaml
name: "iTerm"
description: "A terminal emulator for macOS."
platforms:
  homebrew:
    name: "iterm2"
```

## Notes

- A package that omits a platform key simply cannot be installed on that platform. Ringer will report an error if a user tries to install it there.
- Homebrew formula names and cask names are both valid under the `homebrew` key. Ringer passes the name directly to `brew install`.
- Winget package IDs follow the `Publisher.PackageName` convention. Find them with `winget search <name>`.
