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
#
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
                *noble* | *oracular* | *plucky*)
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
                *)
                echo
                echo "Sorry you have an unsupported OS version"
                exit 1
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
echo "Your OS System is compatible with Swift for $dist_ver / $arch_type"
echo
#
echo "Creating swiftlygo.env file"
echo
# Save environment variables to .env file
cat << EOF > swiftlygo.env
ARCH=$arch
ARCH_TYPE=$arch_type
DIST_VER=$dist_ver
ID=$id
OS_DIR=$os_dir
OS_FILE_NAME=$os_file_name
URL_BASE=https://download.swift.org
EOF
#
echo "Finished!"
echo
# swiftlygo installation location
# /usr/libexec/swiftlygo/bin/
# ├── swiftlygo -> symlink to /usr/bin/swiftlygo
# └── swiftlygo.env
#
