#!/bin/bash
BASE=$(pwd)
cd ${BASE}/cmd/server && ./build.sh
cd ${BASE}/cmd/wasm && ./build.sh
cd ${BASE}

