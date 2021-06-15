#!/bin/sh

if [ $(id -u) != 0 ]; then
    echo "This installer must be executed as root because it installs z onto /usr/local/bin/z"
    exit 1
fi

VERSION=${1:-latest}

if [ $VERSION = "latest" ]; then
    echo "Fetching latest version from GitHub."
    VERSION=$(curl -s "https://api.github.com/repos/serramatutu/z/releases/latest" | grep -Po '"tag_name": "\K.*?(?=")')
fi

SYS_NAME=$(uname -s)
SYS_ARCH=$(uname -m)

echo "Installing z version $VERSION"

VERSION_NO_PREFIX=$(echo $VERSION | sed 's/v//g')
TARBALL_FILENAME="z_${VERSION_NO_PREFIX}_${SYS_NAME}_${SYS_ARCH}.tar.gz"

echo "Fetching tarball from GitHub."

if [ -x /usr/bin/curl ]; then
    curl -s -L "https://github.com/serramatutu/z/releases/download/$VERSION/$TARBALL_FILENAME" -o /tmp/$TARBALL_FILENAME
    DOWNLOAD_SUCCESS=$?
else
    wget -q "https://github.com/serramatutu/z/releases/download/$VERSION/$TARBALL_FILENAME" -o /tmp/$TARBALL_FILENAME
    DOWNLOAD_SUCCESS=$?
fi

if [ "$DOWNLOAD_SUCCESS" != 0 ]; then
    echo "Your distribution \"$SYS_NAME ($SYS_ARCH)\" is currently not supported. Please file an issue at https://github.com/serramatutu/z/issues if you think it should be."
    exit 1
fi

echo "Extracting tarball"
tar -xf /tmp/$TARBALL_FILENAME z

echo "Moving executable to /usr/local/bin/z"
mv z /usr/local/bin/z

echo "Cleaning up"
rm /tmp/$TARBALL_FILENAME

echo "Done! z $VERSION was installed! Run \"z\" to get started."
exit 0
