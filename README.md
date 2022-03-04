# Image Javascript
A simple tool which encodes your Javascript Code into a .png image file

Uses [Auyer's steganography library](https://github.com/auyer/steganography/) to encode data



## Demo
Non-Encoded Image

![Non Encoded Image](https://i.imgur.com/9CjvCAT.png)

SHA256: `3D568164DE25AE2AB8F1EF2E55E7A925BF0A49AF1FD85817C5113FCEA8705750`

Javascript Encoded Image

![Javascript Encoded Image](https://i.imgur.com/DbMcEM0.png)

SHA256: `16633A42672CD271865C99DC3FC9A34E517C2C694AFB822B0EA2C03808C3B093`



## Run Locally

Clone the project

```bash
  git clone https://github.com/yashraj-n/image-js
```

Go to the project directory

```bash
  cd image-js
```

Install dependencies

```bash
 go get github.com/auyer/steganography/
```

Run the main.go file

```bash
  go run .
```


## Usage

### Compile
Compile a javascript code into a image file

```bash
ijs compile [Javascript file] [Image file]
```

Example:
```bash
ijs compile test.js test.png --debug
```

### Run
Decodes a image file into a Javascript file and runs it. (Must have node.js installed)
```bash
ijs run [Image File]
```

Example:
```bash
ijs run test.png
```

### Export
Decodes a image file into a Javascript file
```bash
ijs export [Image File]
```

Example:
```bash
ijs export test.png
```


## License

[MIT](https://choosealicense.com/licenses/mit/)

