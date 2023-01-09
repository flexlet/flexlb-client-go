#!/bin/bash

# installation target
BIN="/usr/local/bin"
ETC="/etc/flexlb"

function fn_main() {
    echo "install certs to: ${ETC}"
    if [ ! -d "${ETC}" ]; then mkdir -p ${ETC}; fi
    cp -r etc/certs ${ETC}/

    echo "install binary to: ${BIN}"
    cp bin/flexlb-cli  ${BIN}/
    chmod +x ${BIN}/flexlb-cli
   
    echo "done"
}

fn_main $@
