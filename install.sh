if ! go build -o clip *.go; then
    echo "Could not build clip, is go installed properly?"
else
    sudo mv clip /bin
fi