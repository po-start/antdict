#! /usr/bin/env bash

scriptPath=$(cd `dirname $0`; pwd)
goPath=$(cd ${scriptPath}/../; pwd)

sed -i '' '/export GOPATH/d' ~/.bash_profile
echo "export GOPATH=${goPath}" >> ~/.bash_profile

source ~/.bash_profile 

printf "\x1B[32m"
echo -e "运行以下命令，设置GOPATH路径:"
printf "\x1B[0m"
echo "source ~/.bash_profile"
