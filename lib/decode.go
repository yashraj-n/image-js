package lib

import (
	"bufio"
	"github.com/auyer/steganography"
	"github.com/yashraj-n/ijs/utils"
	"image/png"
	"os"
	"strconv"
	"strings"
)

func Decode(imgPath string, debug bool) {

	if _, err := os.Stat(imgPath); os.IsNotExist(err) {
		utils.Error("The Image file path specified " + imgPath + " does not exist")
		return
	}

	if debug {
		utils.Debug("Image file path: " + imgPath)
		utils.Debug("Checking extension of the image file")
	}

	if !strings.HasSuffix(imgPath, ".png") {
		utils.Error("The Image file path specified " + imgPath + " is not a png file")
		return
	}

	if debug {
		utils.Debug("Extension of the image file is correct")
		utils.Debug("Opening the image file")
	}

	inFile, _ := os.Open(imgPath)
	defer inFile.Close()
	if debug {
		utils.Debug("Reading image file")
	}
	reader := bufio.NewReader(inFile)
	img, _ := png.Decode(reader)
	if debug {
		utils.Debug("Getting message size")
	}
	size := steganography.GetMessageSizeFromImage(img)
	if debug {
		utils.Debug("Message size: " + strconv.Itoa(int(size)))
	}
	if debug {
		utils.Debug("Decoding...")
	}
	msg := steganography.Decode(size, img)
	//Creating new file
	if debug {
		utils.Debug("Creating new file")
	}
	outFile, err := os.Create(strings.TrimSuffix(imgPath, ".png") + ".js")
	if err != nil {
		utils.Error("Error creating new file")
		return
	}

	//Writing message to outFile
	if debug {
		utils.Debug("Writing message to file")
	}
	_, err = outFile.WriteString(string(msg))
	if err != nil {
		return
	}
	if debug {
		utils.Debug("Message written to file")
	}
	outFile.Close()
	if debug {
		utils.Debug("File closed")
	}
	utils.Success("Message decoded successfully to file" + strings.TrimSuffix(imgPath, ".png") + ".js")

}
