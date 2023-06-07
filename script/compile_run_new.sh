#bin/bash

# get the dir of current file
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
# go to the root dir of the project
cd $DIR/../
pwd

# check whether app "ConfBackend" is running. If so, kill it
if ps -ax|grep "ConfBackend"; then
    echo "ConfBackend is already running. Stopping current instance..."
    pkill ConfBackend
fi

# compile new proj
echo "Compiling new project..."
go build -o ConfBackend main.go


# run new proj in the background
echo "Running new project in the background..."
nohup ./ConfBackend > /dev/null 2>&1 &
