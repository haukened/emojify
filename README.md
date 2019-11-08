# emojify
A simple unicode emoji converter for terminal apps, written in Go.

This list currently has about 2100 supported unicode emojis and aliases.  They aren't yet perfect, and contributions are welcome!

## Usage (as a package)
`import "github.com/haukened/emojify"`\
and then\
`fmt.Println(emojify.Render(":beer:"))`\

## Usage (as a useless command line program)

`git clone https://github.com/haukened/emojify.git`\ 
`cd emojify/cmd/emojify`\
`go build`\
`./emojify :beer:`\



