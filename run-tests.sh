#!/usr/bin/env bash

set -e

export GOPATH=`pwd`
export GOBIN=`pwd`/bin
export GOMAXPROCS=`grep -c ^processor /proc/cpuinfo`

# Remove old coverage file
rm -f coverage.txt
rm -f coverage.*.txt

# Aurora library
declare -A PACKAGES=(
  ["aurora.file.are"]="aurora/file/are"
  ["aurora.file.erf"]="aurora/file/erf"
  ["aurora.file.gff"]="aurora/file/gff"

  ["aurora.tools"]="aurora/tools"
  ["aurora.tools.fileReader"]="aurora/tools/fileReader"
)

# Download all the required packages for the aurora library
go get -t -v aurora/...
# Execute the tests
for fileName in ${!PACKAGES[@]}; do
  go test -race \
    -cover \
    -covermode=atomic \
    -coverprofile=coverage.${fileName}.txt \
    ${PACKAGES[$fileName]}
done

sed -i '/mode: atomic/d' coverage.*.txt
echo "mode: atomic" > coverage.txt
cat coverage.*.txt >> coverage.txt
rm -f coverage.*.txt
