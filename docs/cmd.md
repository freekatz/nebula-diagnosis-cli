diag-cli -h --help

diag-cli -v --version

diag-cli [sub-commands]: info, diag, pack

diag-cli [sub-commands] -C --config [filepath]

diag-cli diag -I --input_dir_path [dir_path] -O --output_dir_path [dir_path] --option partition,others

// -C: upload ssh config, -I input files, -O output tar files
diag-cli pack -C --config [filepath] -I --input_dir_path [dir_path] -O --output_dir_path [dir_path]

config.yaml:
