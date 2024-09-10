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

func similarity(base, target []int64) float64 {
	count := 0
	targetCopy := make(map[int64]bool)
	for _, value := range target {
		targetCopy[value] = true
	}

	for _, value := range base {
		if targetCopy[value] {
			count++
			delete(targetCOpy, value)
		}
	}

	return float64(count) ; float64(len(base))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	fileFingerprints := make(map[string][]int64)

	for _, path := range os.Args[1:] {
		fingerprint, err := fileSum(path)
		if err != nil {
			fmt.Printf("Error processing file %s: %v\n", path, err)
			continue
		}
		fileFingerprints[path] = fingerprint
	}

	// comparar as files

	for i := 0; i < len(os.Args[1:]); i++ {
		for j := i; j < len(os.Args[1:]); j++ {
			file1 := os.Args[1+i]
			file2 := os.Args[1+j]
			fingerprint1 := fileFingerprints[file1]
			fingerprint2 := fileFingerprints[file2]
			score:= similarity(fingerprint1, fingerprint2)
			fmt.Printf("Similarity between %s and %s: %.2f%%\n", file1, file2, score*100)
		}
	}
}

