#!/usr/bin/env bash

export T_GO_VERSION=1.24.4
export T_OS=linux
export T_ARCH=amd64
export T_ZIP_NAME="go${T_GO_VERSION}.${T_OS}-${T_ARCH}.tar.gz"
export T_ZIP_PATH="/usr/local/${T_ZIP_NAME}"
export T_DOWNLOAD_URL="https://go.dev/dl/${T_ZIP_NAME}"
echo "T_ZIP_NAME: [${T_ZIP_NAME}]"
echo "T_ZIP_PATH: [${T_ZIP_PATH}]"
echo "T_DOWNLOAD_URL: [${T_DOWNLOAD_URL}]"

sleep 5

wget -O ${T_ZIP_PATH} ${T_DOWNLOAD_URL} && rm -rf /usr/local/go && tar -zxvf ${T_ZIP_PATH} -C /usr/local && rm -rf ${T_ZIP_PATH}
