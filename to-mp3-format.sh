#!/bin/bash

convert-audio-files() {
	count=0
	inc=1
	for f in *.aac *.opus *.mpeg .m4a; do
		ffmpeg -i "$f" -b:a 192K "${f%.*}.mp3" 
		count=$(($count + $inc))
	done
	echo "converted $count files on folder $(pwd)"
}

remove-old-files() {
	for f in *.aac *.opus *.mpeg .m4a; do
		rm -rf "$f"
	done
}

convert-folder() {
	if [ "$1" ]; then 
		cd "$1"
	fi
	
	for fe in *; do
		if [[ -d $fe ]]; then
			convert-folder "$fe"
		fi
	done
	
	convert-audio-files
	remove-old-files


	if [ "$1" ]; then 
		cd "-1"
	fi
}

if [ "$1" ]; then
	cd "$1"
fi
convert-folder "."

#loop-folders() {
#	if [ "$1" ]; then 
#			cd "$1"
#	fi
#	
#	convert-folder "$(pwd)"
#}
#
#if [ "$1" ]; then 
#	loop-folders "$1"
#fi
#loop-folders "$(pwd)"

