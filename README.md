# etl-client
Developed by Gabriel Cordovado.

## Install Instructions
The following instructions can be used to run the etl client on nix-based operating systems. Testing on Windows operating systems is limited.

### Build
Build the command line client by running the following commands in the repositories working directory. Then add the build folder to your os environment path **/etc/paths**
```shell
mkdir build
go build -o build/
```

### Verification
The version command can be used to verify that the command-line is build and added to the environment path correctly. The command should output the version and current date.
```
etl version
```