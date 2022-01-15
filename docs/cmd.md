diag-cli -h --help

diag-cli -v --version

diag-cli [sub-commands]: info, diag, pack, unpack

diag-cli [sub-commands] -C --config [filepath]

diag-cli diag  -I --input_dirPath [dirPath] -O --output_dirPath [dirPath] --option [", , , , "]

// -C: upload ssh config, -I input files, -O output tar files
diag-cli pack  -C --config [filepath] -I --input_dirPath [dirPath] -O --output_dirPath [dirPath]

config.yaml:
