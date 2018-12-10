#!/bin/bash
# TwitchBOT Linux Builder

tb_BINNAME='twitchBOT'

tb_SCRIPTNAME="$tb_BINNAME.go"
tb_PATH=".builds"
tb_STAGE="alpha"
tb_VERSION="0.0.1"
tb_OS="linux"

tb_BIN_TEMP="${tb_BINNAME}_temp"
tb_BIN_LATEST="${tb_PATH}/${tb_STAGE}/${tb_OS}-latest"
tb_BIN_EXEC="${tb_PATH}/${tb_STAGE}/${tb_OS}-${tb_BINNAME}-${tb_VERSION}"

err=$( { go build -o "$tb_BIN_TEMP" "$tb_SCRIPTNAME"; } 2>&1 )

if [ "$err" = '' ]; then
	cp $tb_BIN_TEMP $tb_BIN_LATEST
	cp $tb_BIN_TEMP $tb_BIN_EXEC
	rm $tb_BIN_TEMP
else
	date=$(date)
	output="[FAILED]\n${err} \n"
	echo -e $output
	echo -e "${date} ${output}" > "error.log"
fi

exit