# Save this script as build-windows-x64.ps1

# Output the target information
# This prints the type of build being generated
echo "GEN BIN"
echo "WINDOWS-X64"

# Set environment variables for Windows x64 build
# GOOS is set to 'windows' to specify the target operating system
# GOARCH is set to 'amd64' to specify the architecture (x64)
# CGO_ENABLED is set to '1' to enable Cgo, allowing the inclusion of C code
$env:GOOS = "windows"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "1"

# Get the current date in YYYYMMDD format to append to the binary name
$Date = Get-Date -Format "yyyyMMdd"

# Define the binary output name, appending the current date and '.exe' for Windows executables
# $BinaryName = "jstartwin64.exe"
$BinaryName = "vCustom_win64.exe"

# Run the Go build command to compile the code into a Windows x64 binary
# The '-o' flag specifies the output file name
go build -o $BinaryName .

# Check if the build succeeded
# $? contains the exit status of the last command, and it's true if the build was successful
if ($?) {
    echo "Build succeeded! Binary created: $BinaryName"
} else {
    echo "Build failed."
}
