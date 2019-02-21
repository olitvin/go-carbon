#!/bin/sh

cd `dirname $0`
ROOT=`pwd`

PACKAGE="github.com/olitvin/go-carbon"

cd _vendor/src/${PACKAGE}/
GOPATH=$ROOT/_vendor dep $@
