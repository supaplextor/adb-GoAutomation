#!/bin/env bash

find * -name \*.go | while read f
do	echo -e "\n### ${f}\n<pre>"
	grep ^func "${f}" | sed -e s/..$// | awk '{print "<code>"$ff"</code>"}'
	echo -e "</pre>"
done | tee md/all-functions.md 
