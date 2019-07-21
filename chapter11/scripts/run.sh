#!/usr/bin/env bash

:<<EOF
scripts for this project

EOF


echo $(go version)

command=${1}


function Run() {
    echo "Run Function"
}

function Build() {
    echo "Build Function"
}

function Deploy() {
    echo "Deploy Function"
}

commands=("run" "build" "deploy")

case ${command} in
${commands[0]}) Run;;
${commands[1]}) Build;;
${commands[2]}) Deploy;;
esac


