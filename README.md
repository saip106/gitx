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

## Build and Install

### Local Build
To create a binary in the current folder:
```powershell
go build -o gitx.exe ./cmd/gitx
# You can now run it as:
.\gitx reclone
```

### Global Install
To make `gitx` available everywhere as a command:
```powershell
go install ./cmd/gitx
# You can now run it anywhere as:
gitx reclone
```

## Project Structure

- `cmd/gitx/`: Entry point and root command.
- `internal/reclone/`: Logic for the `reclone` command.
- `internal/git/`: Git utility functions.
- `internal/fs/`: File system utilities (including safe delete).
- `pkg/prompt/`: Interactive prompt utilities.
