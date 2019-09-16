# goprettyjson

Simple program which can receive unformatted JSON piped into in and then will format the JSON and write to stdout.

I used this to create a Apple Automator application which gets JSON text from the clipboard and formats the JSON and puts the formatted text back into the clipboard. 

```bash
export GOBIN="/Users/richard/Development/source/go/bin"
export PATH="$PATH:$GOBIN"
[ $# -ge 1 -a -f "$1" ] && input="$1" || input="-"
cat $input | goprettyjson
```

