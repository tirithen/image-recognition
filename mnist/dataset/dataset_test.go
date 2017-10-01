package dataset_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tirithen/image-recognition/mnist/dataset"
)

func TestNewDataSetFromFilesMnistTestData(test *testing.T) {
	dataSet, err := dataset.NewDataSetFromFiles("../data/t10k-images-idx3-ubyte.gz", "../data/t10k-labels-idx1-ubyte.gz")
	assert.NoError(test, err)
	assert.Equal(test, 10000, dataSet.Size)
	assert.Equal(test, 28, dataSet.Width)
	assert.Equal(test, 28, dataSet.Height)
	assert.Equal(test, 10000, len(dataSet.Images))
	assert.Equal(test, 10000, len(dataSet.Labels))
}

func TestNewDataSetFromFilesMnistTrainData(test *testing.T) {
	dataSet, err := dataset.NewDataSetFromFiles("../data/train-images-idx3-ubyte.gz", "../data/train-labels-idx1-ubyte.gz")
	assert.NoError(test, err)
	assert.Equal(test, 60000, dataSet.Size)
	assert.Equal(test, 28, dataSet.Width)
	assert.Equal(test, 28, dataSet.Height)
	assert.Equal(test, 60000, len(dataSet.Images))
	assert.Equal(test, 60000, len(dataSet.Labels))
}

func TestNewDataSetFromFilesWithBadFile(test *testing.T) {
	_, err := dataset.NewDataSetFromFiles("../data/bad-gz-file", "../data/t10k-labels-idx1-ubyte.gz")
	assert.Error(test, err)
}

func TestNewDataSetFromFilesWithMissingFile(test *testing.T) {
	_, err := dataset.NewDataSetFromFiles("../data/t10k-images-idx3-ubyte.gz", "../data/i-do-not-exist-either")
	assert.Error(test, err)
}
