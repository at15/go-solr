#!/usr/bin/env bash

# get the script path http://stackoverflow.com/questions/4774054/reliable-way-for-a-bash-script-to-get-the-full-path-to-itself
pushd `dirname $0` > /dev/null
SCRIPT_PATH=`pwd -P`
popd > /dev/null
ORIGINAL_WD=${PWD}
cd ${SCRIPT_PATH}

SOLR_VERSION=6.6.0

if [ -f solr-${SOLR_VERSION}.tgz ]; then
    echo "Solr ${SOLR_VERSION} already downloaded "
else
    echo "Downloading Solr ${SOLR_VERSION}"
    wget http://www-us.apache.org/dist/lucene/solr/6.6.0/solr-6.6.0.tgz
    echo "Download finished for Solr ${SOLR_VERSION}"
fi

if [ -d "solr" ]; then
    echo "Solr ${SOLR_VERSION} already installed"
else
    echo "Extracting tarball"
    tar zxf solr-${SOLR_VERSION}.tgz
    mv solr-${SOLR_VERSION} solr
    echo "Install finished for Solr ${SOLR_VERSION}"
fi

if hash java 2>/dev/null; then
    echo "Java is installed"
else
    echo "Warn: Java not installed"
fi

cd ${ORIGINAL_WD}