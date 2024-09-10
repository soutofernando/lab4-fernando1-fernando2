#!/bin/bash

args=`find dataset -type f | xargs`

time bash go/serial/run.sh $args
time bash serial-partial/run.sh $args
time bash concurrent-0/run.sh $args