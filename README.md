# gitx

`gitx` is a Go-based command-line tool providing enhanced git workflows.


## Commands

### `reclone`
Deletes and re-clones a repository.

**Usage:**
```bash
gitx reclone [path] [flags]
```
**Flags:**
- `--force`, `-f`: Skip all confirmations.
- `--depth <n>`: Perform a shallow clone.
- `--single-branch`: Clone only the current branch.
- `--no-tags`: Do not fetch tags.
- `--verbose`, `-v`: Print detailed logs.

**Windows Note:**
Due to file locking in Windows, if you are currently inside the repository directory you wish to reclone, the deletion might fail. It is recommended to run the command from the parent directory:
```powershell
gitx reclone ./my-repo
```

### `master`
Switches to the `master` branch and runs `git pull --all`.

**Usage:**
```bash
gitx master
```
Shows all git output and a summary. Fails if the branch does not exist.

### `main`
Switches to the `main` branch and runs `git pull --all`.

**Usage:**
```bash
gitx main
```
Shows all git output and a summary. Fails if the branch does not exist.


## Build and Deploy

### Build
To build with a specific version:
```powershell
./build.ps1 v0.1.0
# or just ./build.ps1 for dev build
```

### Deploy
To copy the binary to your Go bin directory (or specify a target):
```powershell
./deploy.ps1
# or ./deploy.ps1 C:\some\other\dir
```

## Versioning Best Practices
1. **Semantic Versioning (SemVer):** Use `vMAJOR.MINOR.PATCH` (e.g., `v0.1.0`).
2. **Build-Time Injection:** Use `-ldflags` to inject the version from your CI/CD or Git tags. This ensures the binary always knows exactly which version it is.
3. **Git Tags:** Always tag your releases in Git:
   ```bash
   git tag -a v0.1.0 -m "Release version 0.1.0"
   git push origin v0.1.0
   ```

## Project Structure

- `cmd/gitx/`: Entry point and root command.
- `internal/reclone/`: Logic for the `reclone` command.
- `internal/git/`: Git utility functions (including branch switching and pulling).
- `internal/fs/`: File system utilities (including safe delete).
- `pkg/prompt/`: Interactive prompt utilities.
- `build.ps1`: PowerShell build script.
- `deploy.ps1`: PowerShell deploy script.
