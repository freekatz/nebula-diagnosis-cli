# nebula-diagnosis-cli

Nebula Diagnosis CLI Tool is an information diagnosis cli tool for the nebula service and the node to which the service belongs.

## Usage

### Build

```shell
make clean && make build
```

### Commands

```shell
nebula-diag-cli -h --help
nebula-diag-cli -v --version
```

Then type and enter `nebula-diag-cli [sub-commands] -C --config [filepath]` to run commands.

```shell
nebula-diag-cli [sub-commands]: info, diag, pack
```

### Info

```shell
./nebula-diag-cli info --config <config filepath>
```

### diag

```shell
nebula-diag-cli diag  -I --input_dir_path [dir_path] -O --output_dir_path [dir_path] --option partition,others
```

### Pack

Use the `-C --config [filepath]` to set the upload ssh config and cli tool will upload the Package automatically.

```shell
nebula-diag-cli pack -C --config [filepath] -I --input_dir_path [dir_path] -O --output_dir_path [dir_path]
```
