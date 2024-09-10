#!/bin/zsh

if [ "$#" -ne 4 ]; then
  echo "Usage: $0 <FILTER> <COUNT> <OUT_OLD> <OUT_NEW>"
  echo "Example 1: $0 . 10 old.txt new.txt"
  echo "Example 2: $0 ascii-7 10 old.txt new.txt"
  exit 1
fi

FILTER=$1
COUNT=$2
OUT_OLD=$3
OUT_NEW=$4

FUNC_OLD=std
FUNC_NEW=opt

echo "Benchmarking standard to $FUNC_OLD"
go test \
  -test.timeout=0 \
  -bench ToLower/$FILTER \
  -benchmem \
  -func $FUNC_OLD \
  -count $COUNT \
  | tee $OUT_OLD
clear


echo "Benchmarking optimized $FUNC_NEW"
go test \
  -test.timeout=0 \
  -bench ToLower/$FILTER \
  -benchmem \
  -func $FUNC_NEW \
  -count $COUNT \
  | tee $OUT_NEW
clear

benchstat $OUT_OLD $OUT_NEW
