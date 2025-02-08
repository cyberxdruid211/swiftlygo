#! /bin/bash
#
# Install script for SwiftlyGo
#

# detect compatible os and run install script
arch=$(uname -m)
arch_type=$(uname -m)
id=""
dist_ver=""
os_dir=""
os_file_name=""
swgo_ver="v1.0"

case $arch in
    *x86_64*)
        arch=""
        ;;
    *aarch64*)
        arch="-aarch64"
        ;;
    *)
        echo "Sorry, you have an unsupported architecture"
        exit 1
        ;;
esac
echo
echo "Checking your OS system for Swift compatibility ..."
if test -r /etc/os-release; then
. /etc/os-release

        case $ID in
            *elementary* | *pop* | *linuxmint* | *bianbu* | *zorin*)
                echo "$PRETTY_NAME is compatible with ubuntu/$UBUNTU_CODENAME"
                echo
                os_version=$UBUNTU_CODENAME
            ;;
            *ubuntu*)
                 os_version=$UBUNTU_CODENAME
            ;;
            *debian*)
                os_version=$VERSION_CODENAME
            ;;
            *)
                echo "sorry you have an unsupported operating system"; exit 1
            ;;
        esac
        # add cases for focal, jammy, noble, bookworm
        case $os_version in
                *focal*)
                id="ubuntu"
                dist_ver="Ubuntu 20.04"
                os_dir="ubuntu2004"
                os_file_name="ubuntu20.04"
                ;;
                *jammy*)
                id="ubuntu"
                dist_ver="Ubuntu 22.04"
                os_dir="ubuntu2204"
                os_file_name="ubuntu22.04"
                ;;
                *noble*)
                id="ubuntu"
                dist_ver="Ubuntu 24.04"
                os_dir="ubuntu2404"
                os_file_name="ubuntu24.04"
                ;;
                *bookworm*)
                id="debian"
                dist_ver="Debian 12"
                os_dir="debian12"
                os_file_name="debian12"
                ;;
        esac
else
    echo
    echo "sorry you have an unsupported system"
    echo
    exit 1
fi
# OS System is compatible with Swift for $dist_ver / $arch_type
echo
echo "Success! Your OS System is compatible with Swift for $dist_ver / $arch_type"
echo
echo "Installing SwiftlyGo ..."
#
# check if user is root and set SUDO accordingly
SUDO=$(if [ "$(id -u)" -ne 0 ]; then echo sudo; else echo ""; fi)
#
DOWNLOAD_URL="https://github.com/cyberxdruid211/swiftlygo/releases/download/$swgo_ver/swiftlygo-$arch_type"
DEST_DIR="/usr/libexec/swiftlygo/bin/"
EXECUTABLE_NAME="swiftlygo"
#
$SUDO mkdir -p "$DEST_DIR"
# Save environment variables to .env file
$SUDO bash -c "cat << EOF > /usr/libexec/swiftlygo/bin/swiftlygo.env
ARCH=$arch
ARCH_TYPE=$arch_type
DIST_VER=$dist_ver
ID=$id
OS_DIR=$os_dir
OS_FILE_NAME=$os_file_name
URL_BASE=https://download.swift.org
EOF"
#
install_swiftlygo() {
    # Remove existing executable if it exists
    $SUDO rm "$DEST_DIR/$EXECUTABLE_NAME" 2>/dev/null
    
    # Download, install, and make executable
    if $SUDO curl -sL "$DOWNLOAD_URL" -o "$DEST_DIR/$EXECUTABLE_NAME" --fail; then
        $SUDO chmod +x "$DEST_DIR/$EXECUTABLE_NAME"
        # Create a symbolic link
        $SUDO ln -sf "$DEST_DIR/$EXECUTABLE_NAME" /usr/bin/swiftlygo
        echo
        echo "Congratulations! Swiftlygo has been successfully installed."
        logo
        echo
        echo "For help run 'swiftlygo -h'"
        echo
    else
        echo "Failed to download SwiftlyGo. Please check if the URL is correct or if the file exists at the given URL."
        echo "URL: $DOWNLOAD_URL"
        echo
        return 1
    fi
}
logo() {
cat << EOF
 ____          _  __ _   _        ____       
/ ___|_      _(_)/ _| |_| |_   _ / ___| ___  
\___ \ \ /\ / / | |_| __| | | | | |  _ / _ \ 
 ___) \ V  V /| |  _| |_| | |_| | |_| | (_) |
|____/ \_/\_/ |_|_|  \__|_|\__, |\____|\___/ 
                           |___/  $swgo_ver           
EOF
}
#
install_swiftlygo

# swiftlygo installation location
# /usr/libexec/swiftlygo/bin/
# ├── swiftlygo -> symlink to /usr/bin/swiftlygo
# └── swiftlygo.env
#
