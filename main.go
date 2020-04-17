package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
	"strings"
	"strconv"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	var replacer = strings.NewReplacer("AM", "", "PM", "")	
	hhStr, mmStr, ssStr := "", "", ""
	hh, mm, ss := 0, 0, 0	

	tempStr := replacer.Replace(s)
	timeSplit := strings.Split(tempStr, ":") 
	hh, _ = strconv.Atoi(timeSplit[0])
	mm, _ = strconv.Atoi(timeSplit[1]) 
	ss, _ = strconv.Atoi(timeSplit[2])
	if strings.Contains(s, "PM") {		
		if hh > 1 && ( hh <= 11 && mm <= 59) {
			hh += 12 
		}
	} else {		
		if hh == 12 && mm <= 59 {
			hh -= 12			
		}
	}

	hhStr = fmt.Sprintf("%02d", hh)
	mmStr = fmt.Sprintf("%02d", mm)
	ssStr = fmt.Sprintf("%02d", ss)

	return hhStr + ":" + mmStr + ":" + ssStr
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer outputFile.Close()

    writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
