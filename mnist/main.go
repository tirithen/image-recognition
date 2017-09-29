package main

import (
	"compress/gzip"
	"fmt"
	"log"
	"os"
)

// DataSet represents a data set that can be used for training
type DataSet struct {
	ImagesFilename string
	LabelsFilename string
	Size           int32
	MagicNumber    int32
	Width          int32
	Height         int32
	Labels         []byte
	Images         []byte
}

// ReadGzipFile reads gziped data from a file
func (dataSet *DataSet) ReadGzipFile(filename string) (data []byte, err error) {
	file, err := os.Open(filename)
	defer file.Close()
	reader, err := gzip.NewReader(file)
	defer reader.Close()
	return
}

func newDataSetFromFiles(imagesFilename string, labelsFilename string) (result DataSet, err error) {
	result.ImagesFilename = imagesFilename
	result.LabelsFilename = labelsFilename

	imagesData, err := result.ReadGzipFile(result.ImagesFilename)

	result.Images = imagesData

	return
}

func main() {
	testSet, err := newDataSetFromFiles("data/t10k-images-idx3-ubyte.gz", "data/t10k-labels-idx1-ubyte.gz")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(testSet.Size)
}
