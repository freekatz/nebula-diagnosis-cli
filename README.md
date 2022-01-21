# nebula-diagnose

Nebula Diagnose CLI Tool is an information diagnose cli tool for the nebula service and the node to which the service belongs.

nebula diagnose supports the following functions:

- Collect physical information of nodes 
- Collect nebula service metrics, flags and status
- Download the log under nebula service running log directory
- Diagnose based on collected information (TODO)
- Package files or folders for upload 

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

Then type and enter `nebula-diag-cli <sub-commands> -C <config filepath>` to run commands.

```shell
nebula-diag-cli <sub-commands>: info, diag, pack
```

### Info

```shell
./nebula-diag-cli info -C <config filepath>
```

### diag

```shell
nebula-diag-cli diag -O <output dir path> -I <input data dir path> --option <options, such as: partition,others>
```

### Pack

Use the `-C <config filepath>` to set the upload ssh config and cli tool will upload the Package automatically.

```shell
nebula-diag-cli pack -C <config filepath> -O <output dir path> -I <input tar filepath> -N <output tar filename, will output into output dir path>
```
