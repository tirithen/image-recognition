package dataset

import (
	"compress/gzip"
	"encoding/binary"
	"io/ioutil"
	"os"
)

// DataSet represents a data set that can be used for training
type DataSet struct {
	ImagesFilename string
	LabelsFilename string
	Size           int
	Width          int
	Height         int
	Labels         []byte
	Images         [][]byte
}

// ReadGzipFile reads gziped data from a file
func (dataSet *DataSet) ReadGzipFile(filename string) (data []byte, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		return
	}
	defer reader.Close()

	data, _ = ioutil.ReadAll(reader)

	return
}

// NewDataSetFromFiles will read a image data filename and a label data filename, read the contents and return the data in a DataSet struct
func NewDataSetFromFiles(imagesFilename string, labelsFilename string) (result DataSet, err error) {
	result.ImagesFilename = imagesFilename
	result.LabelsFilename = labelsFilename

	imagesData, err := result.ReadGzipFile(result.ImagesFilename)
	if err != nil {
		return
	}

	result.Size = int(binary.BigEndian.Uint32(imagesData[4:8]))
	result.Height = int(binary.BigEndian.Uint32(imagesData[8:12]))
	result.Width = int(binary.BigEndian.Uint32(imagesData[12:16]))

	for index := 0; index < result.Size; index++ {
		result.Images = append(result.Images, imagesData[16+result.Width*index:16+result.Width*(index+1)])
	}

	labelsData, err := result.ReadGzipFile(result.LabelsFilename)
	if err != nil {
		return
	}

	result.Size = int(binary.BigEndian.Uint32(labelsData[4:8]))

	for index := 0; index < result.Size; index++ {
		result.Labels = append(result.Labels, labelsData[8+index])
	}

	return
}
