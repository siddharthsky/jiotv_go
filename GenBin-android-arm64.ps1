# Save this script as build-android-arm64.ps1

# Output the target information
echo "GEN BIN"
echo "ANDROID-ARM8"

# Set environment variables for Android ARMv7 build
$env:GOOS = "android"
$env:GOARCH = "arm64"
$env:CC = "C:\android-ndk-r27\toolchains\llvm\prebuilt\windows-x86_64\bin\aarch64-linux-android32-clang"
$env:CGO_ENABLED = "1"

# Get the current date in YYYYMMDD format
$Date = Get-Date -Format "yyyyMMdd"

# Define the binary output name with the date included
$BinaryName = "vCustom-android-arm64"

# Run the Go build command
go build -o $BinaryName .

# Check if build succeeded
if ($?) {
    echo "Build succeeded! Binary created: $BinaryName"
} else {
    echo "Build failed."
}
