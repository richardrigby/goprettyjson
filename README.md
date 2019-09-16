# goprettyjson

Simple program which can receive unformatted JSON piped into it and then will format the JSON and write to stdout.

I used this to create an Apple Automator application which gets JSON text from the clipboard and formats the JSON and puts the formatted text back into the clipboard. 


![Apple Automator Screen Shot](https://github.com/richardrigby/goprettyjson/blob/master/docs/screen-shot.jpg)

```bash
export GOBIN="/Users/richard/Development/source/go/bin"
export PATH="$PATH:$GOBIN"
[ $# -ge 1 -a -f "$1" ] && input="$1" || input="-"
cat $input | goprettyjson
```

Run `go install` to install into GOBIN
