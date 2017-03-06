#! /bin/sh

readonly PROJECT="cliper"
readonly BIN_PATH="/usr/local/bin"
readonly GH_AUTHOR="hackliff"
readonly LATEST_VERSION="0.1.1"

function download_binary() {
  local version="$1"
  local platform=darwin-amd64

  echo "downloading binary (${PROJECT}-${platform})..."
  curl \
    -ksL \
    -o ${BIN_PATH}/${PROJECT} \
    https://github.com/${GH_AUTHOR}/${PROJECT}/releases/download/v${version}/${PROJECT}-${platform}

  echo "make it executable (${BIN_PATH}/${PROJECT})"
  chmod +x ${BIN_PATH}/${PROJECT}
}

download_binary ${CLIPER_VERSION:-"${LATEST_VERSION}"}

