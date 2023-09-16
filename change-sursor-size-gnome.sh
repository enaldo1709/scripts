#!/bin/bash

# change cursor size command:
#   gsettings set org.gnome.desktop.interface cursor-size 32

current_size=$(gsettings get org.gnome.desktop.interface cursor-size)
echo "current cursor size: $current_size"

default_size=24
re='^[0-9]+$'

main() {
    size=$default_size
    if [ "$1" ]; then 
        if ! [[ $1 =~ $re ]] ; then
            echo "error: not a number" >&2; exit 1
        fi
        if [[ $1 == $current_size ]] ; then
            echo "cursor size input is the same than current size: $current_size..."
            echo "exiting with nothing to change..."
            return 0
        else
            size=$1
        fi
        echo "changing cursor size from $current_size to $size"
        gsettings set org.gnome.desktop.interface cursor-size $size
    fi
    
    echo "ok!"
}


main $1

