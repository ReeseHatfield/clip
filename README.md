# clip
A clipboard-like command line utility. Not exactly an alternative to xclip,
but just a simple took I wanted to have. It essentially just acts as a pipable 
buffer for any data you would like

## Usage:
Pip some input into clip and it will save it. Run `clip` and it will print the data to stdout
```bash
echo "hello" | clip
clip # -> hello
```

```bash
find -name "main.go" | clip
# ... sometime later
clip | xargs cat # -> contents of main.go
```

# Installation: To install `clip`, run the following commands:
```bash
git clone https://github.com/ReeseHatfield/clip.git
cd clip
chmod u+x install.sh
./install.sh
```