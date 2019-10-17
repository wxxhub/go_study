#!/bin/bash
# echo '$0: '$0
# echo "pwd: "`pwd`
# echo "scriptPath1: "$(cd `dirname $0`; pwd)
# echo $(cd `dirname $0`; pwd)"/start.sh"

export GOPATH=$(cd `dirname $0`; pwd)"/mplay"
code $(cd `dirname $0`; pwd)