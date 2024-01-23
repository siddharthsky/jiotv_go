try {
    # Check if running with admin privileges
    $isAdmin = ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
    
     # If user wants to access from anywhere, add to PATH
    # (Note: This section assumes that you have the user's consent to modify the system environment variable)
    $accessFromAnywhere = $null
    while ($accessFromAnywhere -eq $null) {
        $accessFromAnywhere = Read-Host "Do you want to access jiotv_go from anywhere? (yes/no)"
        if ($accessFromAnywhere -notin @("yes", "no")) {
            Write-Host "Invalid choice. Please enter 'yes' or 'no'."
            $accessFromAnywhere = $null
        }
    }

    if ($accessFromAnywhere -eq "yes") {
        if (-not $isAdmin) {
            Write-Host "Requesting admin privileges..."
            
            # Relaunch the script with admin privileges
            Start-Process -FilePath PowerShell.exe -Verb Runas -ArgumentList "-File `"$($MyInvocation.MyCommand.Path)`"  `"$($MyInvocation.MyCommand.UnboundArguments)`""
            exit
        }
    }

    # Identify operating system architecture
    $architecture = (Get-WmiObject Win32_OperatingSystem).OSArchitecture
    switch ($architecture) {
        "64-bit" {
            $arch = "amd64"
            break
        }
        "32-bit" {
            $arch = "386"
            break
        }
        "ARM64" {
            $arch = "arm64"
            break
        }
        default {
            throw "Unsupported architecture: $architecture"
        }
    }

    Write-Host "Detected architecture: $arch"

    # Determine the user's home directory
    $homeDirectory = [System.IO.Path]::Combine($env:USERPROFILE, ".jiotv_go")

    # Create the directory if it doesn't exist
    if (-not (Test-Path $homeDirectory -PathType Container)) {
        New-Item -ItemType Directory -Force -Path $homeDirectory
    }

    # Change to the home directory
    Set-Location -Path $homeDirectory

    # If the binary already exists, delete it
    if (Test-Path jiotv_go.exe) {
        Write-Host "Deleting existing binary"
        Remove-Item jiotv_go.exe
    }

    # Fetch the latest binary
    $binaryUrl = "https://api.github.com/repos/rabilrbl/jiotv_go/releases/latest/download/jiotv_go-windows-$arch.exe"
    Write-Host "Fetching the latest binary from $binaryUrl"
    Invoke-WebRequest -Uri $binaryUrl -OutFile jiotv_go.exe -UseBasicParsing

    if ($accessFromAnywhere -eq "yes") {
        # Add the directory to PATH in the current session
        $env:Path = "$env:Path;$homeDirectory"
        
        # Modify system environment variable to persist
        [System.Environment]::SetEnvironmentVariable("Path", [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine) + ";$homeDirectory", [System.EnvironmentVariableTarget]::Machine)
        
        Write-Host "JioTV Go has successfully downloaded and added to PATH. Start by running jiotv_go help"
    } else {
        Write-Host "JioTV Go has successfully downloaded. You can run it from the current folder."
    }
}
catch {
    Write-Host "Error: $_"
    exit 1
}
