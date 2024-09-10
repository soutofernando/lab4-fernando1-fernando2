import {
	"fmt"
	"io"
	"os"
}

func fileSum(filePath string) ([]int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		retunr nil, err
	}
	defer file.Close()

	var chunks []int64
	buffer := make([]byte, 100)
	for{
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if bytesRead == 0 {
			break
		}
		chunkSum := int64(0)
		for i := 0; i< bytesRead; i++ {
			chunkSum += int64(buffer[i])
		}

		chunks = append(chunks, chunkSum)
	}
	return chunks, nil

}



