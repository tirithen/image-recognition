package main

import (
	"fmt"
	"log"

	"github.com/tirithen/image-recognition/mnist/dataset"
)

func main() {
	testSet, err := dataset.NewDataSetFromFiles("data/t10k-images-idx3-ubyte.gz", "data/t10k-labels-idx1-ubyte.gz")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("loaded testSet with the size of:", testSet.Size)
}
