package lib

import (
	"bufio"
	"bytes"
	"github.com/auyer/steganography"
	"github.com/yashraj-n/ijs/utils"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

func Compile(filePath string, imgPath string, debug bool) {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		utils.Error("The Javascript file path specified " + filePath + " does not exist")
		return
	}

	if debug {
		utils.Debug("Got Javascript file path: " + filePath)
		utils.Debug("Got Image file path: " + imgPath)
	}

	if _, err := os.Stat(imgPath); os.IsNotExist(err) {
		utils.Error("The Image file path specified " + imgPath + " does not exist")
		return
	}
	if debug {
		utils.Debug("Checking Suffix of files")
	}

	if !strings.HasSuffix(filePath, ".js") {
		utils.Error("The Javascript file path specified " + filePath + " does not have a .js extension")
		return
	}

	if !strings.HasSuffix(imgPath, ".png") {
		utils.Error("The Image file path specified " + imgPath + " does not have a .png extension")
		return
	}
	if debug {
		utils.Debug("Reading Javascript file")
	}

	jsFileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		utils.Error("Error while reading the Javascript file " + filePath)
		return
	}

	if debug {
		utils.Debug("Reading Image file")
	}

	// Opening File
	imgFile, err := os.Open(imgPath)

	if err != nil {
		utils.Error("Error while opening the Image file " + imgPath)
		return
	}

	if debug {
		utils.Debug("Reading Image file buffer")
	}

	//Reading the buffer
	imgFileBuffer := bufio.NewReader(imgFile)

	if debug {
		utils.Debug("Reading Image file buffer into golang png ")
	}

	img, decodeErr := png.Decode(imgFileBuffer)
	if decodeErr != nil {
		utils.Error("Error while decoding the Image file " + imgPath + " " + decodeErr.Error())
		return
	}

	//Resultant Buffer
	w := new(bytes.Buffer)

	if debug {
		utils.Debug("Encoding Javascrip")
	}

	//Encoding
	encodeErr := steganography.Encode(w, img, jsFileBytes)
	if encodeErr != nil {
		utils.Error("Error while encoding the Image file " + imgPath + " " + encodeErr.Error())
		return
	}

	if debug {
		utils.Debug("Saving image file...")
	}

	//Writing the resultant buffer to a file
	out, e := os.Create(imgPath + "-ijs" + ".png")
	if e != nil {
		utils.Error("Error while creating the output file " + imgPath + "-ijs.png")
		return
	}
	w.WriteTo(out)
	out.Close()
	utils.Success("Successfully encoded the Javascript file " + filePath + " to the Image file " + imgPath + " . Created new image file " + imgPath + "-ijs.png")
}
