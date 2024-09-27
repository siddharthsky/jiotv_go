# Output the target information
# This prints the type of build being generated
echo "GEN BIN"
echo "LINUX-X64"

# Set environment variables for Linux x64 build
# GOOS is set to 'linux' to specify the target operating system
# GOARCH is set to 'amd64' to specify the architecture (x64)
# CGO_ENABLED is set to '1' to enable Cgo, allowing the inclusion of C code
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"

# Get the current date in YYYYMMDD format to append to the binary name
$Date = Get-Date -Format "yyyyMMdd"

# Define the binary output name, appending the current date
# $BinaryName = "jstartlinux64"
$BinaryName = "vCustom_linux64"

# Run the Go build command to compile the code into a Linux x64 binary
# The '-o' flag specifies the output file name
go build -o $BinaryName .

# Check if the build succeeded
# $? contains the exit status of the last command, and it's true if the build was successful
if ($?) {
    echo "Build succeeded! Binary created: $BinaryName"
} else {
    echo "Build failed."
}
