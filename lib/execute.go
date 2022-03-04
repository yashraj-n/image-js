package lib

import (
	"bufio"
	"github.com/auyer/steganography"
	"github.com/yashraj-n/ijs/utils"
	"image/png"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Execute(imgPath string, debug bool) {

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
	//Create temp file using ioutil.TempFile
	if debug {
		utils.Debug("Creating temp file")
	}

	outFile, err := ioutil.TempFile("", "ijs*.js")
	if err != nil {
		utils.Error("Error creating new file")
		return
	}
	if debug {
		utils.Debug("New file created " + outFile.Name())
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
	//Executing the file
	if debug {
		utils.Debug("Executing the file")
	}
	stdout, err := exec.Command("node", outFile.Name()).Output()
	if err != nil {
		utils.Error("Error executing the file")
		return
	}
	if debug {
		utils.Debug("File executed")
	}
	//Printing the output
	if debug {
		utils.Debug("Printing the output")
	}
	utils.Info(string(stdout))
	//Deleting file
	if debug {
		utils.Debug("Deleting the file")
	}
	os.Remove(outFile.Name())
}
