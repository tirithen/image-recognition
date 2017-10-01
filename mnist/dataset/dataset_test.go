package dataset_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tirithen/image-recognition/mnist/dataset"
)

func TestNewDataSetFromFilesMnistTestData(test *testing.T) {
	testSet, err := dataset.NewDataSetFromFiles("../data/t10k-images-idx3-ubyte.gz", "../data/t10k-labels-idx1-ubyte.gz")
	assert.NoError(test, err)
	assert.Equal(test, 10000, testSet.Size)
	assert.Equal(test, 28, testSet.Width)
	assert.Equal(test, 28, testSet.Height)
	assert.Equal(test, 10000, len(testSet.Images))
	assert.Equal(test, 10000, len(testSet.Labels))
}

func TestNewDataSetFromFilesMnistTrainData(test *testing.T) {
	testSet, err := dataset.NewDataSetFromFiles("../data/train-images-idx3-ubyte.gz", "../data/train-labels-idx1-ubyte.gz")
	assert.NoError(test, err)
	assert.Equal(test, 60000, testSet.Size)
	assert.Equal(test, 28, testSet.Width)
	assert.Equal(test, 28, testSet.Height)
	assert.Equal(test, 60000, len(testSet.Images))
	assert.Equal(test, 60000, len(testSet.Labels))
}

func TestNewDataSetFromFilesWithBadFile(test *testing.T) {
	_, err := dataset.NewDataSetFromFiles("../data/bad-gz-file", "../data/t10k-labels-idx1-ubyte.gz")
	assert.Error(test, err)
}

func TestNewDataSetFromFilesWithMissingFile(test *testing.T) {
	_, err := dataset.NewDataSetFromFiles("../data/t10k-images-idx3-ubyte.gz", "../data/i-do-not-exist-either")
	assert.Error(test, err)
}
