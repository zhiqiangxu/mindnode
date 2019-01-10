#!/bin/bash

set -e
set -x

echo "Doing build..."

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
BUILD_DIR="$DIR/build"
mkdir -p $BUILD_DIR
pushd $BUILD_DIR

CC="$DIR/remote/build-scripts/gcc-fat.sh" ../remote/configure --prefix=$BUILD_DIR; make -j 4; make install
popd
