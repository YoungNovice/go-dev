#! /bin/sh

kill -9 $(pgrep webserver)
cd ~
git pull git@github.com:YoungNovice/go-dev.git
cd go-dev
./webserver &
