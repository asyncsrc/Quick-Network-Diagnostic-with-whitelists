package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

// used from -> https://github.com/donovanhide/btctop/blob/master/monitor/layout.go
func drawStringCustom(x, y int, str string, fg, bg termbox.Attribute) int {
	for _, s := range str {
		termbox.SetCell(x, y, s, fg, bg)
		x++
	}
	return x
}

func drawString(x, y int, str string) int {
	for _, s := range str {
		termbox.SetCell(x, y, s, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}
	return x
}

func downloadFile(url string) (bytesWritten int64, err error) {
	out, err := os.OpenFile("largeFile.txt", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return
	}

	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bytesWritten, err = io.Copy(out, resp.Body)
	if err != nil {
		return
	}
	return
}

func completionPercentage(sites map[string][]string, progress *chan int, row *int) {
	totalSites := 0
	currentProgress := 0

	for key, _ := range sites {
		totalSites += len(sites[key])
	}

	for {
		currentProgress = <-*progress
		drawString(1, *row, fmt.Sprintf("Sites scanned: %v of %v -- %.2f%%", currentProgress, totalSites, (float64(currentProgress)/float64(totalSites))*100))
		termbox.Flush()
	}
}

func downloadingAnimation(endLocation int, row *int, finishCheck *chan bool) {
	loopCounter := 0
	checking := [...]string{".", "..", "..."}

	for {
		select {
		case <-*finishCheck:
			drawString(endLocation, *row, "                ")
			return
		default:
			if loopCounter > 2 {
				drawString(endLocation, *row, "               ")
				loopCounter = 0
			}
			drawString(endLocation, *row, fmt.Sprintf("Downloading%s", checking[loopCounter]))
			time.Sleep(350 * time.Millisecond)
			termbox.Flush()
			loopCounter += 1
		}
	}
}

func main() {
	finishCheck := make(chan bool)
	largeFileURL := "http://url_to_large_file_here"

	currentProgress := 0
	start := time.Now()
	row := 1

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	drawStringCustom(1, row, "Network Diagnostics...", termbox.ColorWhite, termbox.ColorBlue)
	termbox.Flush()

	reportFileName := "pm_dotcom_report.txt"
	reportFile, err := os.OpenFile(reportFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Unable to create report file: %s in current directory.  Error: %s", reportFileName, err)
		os.Exit(-1)
	}

	defer reportFile.Close()
	reportFile.WriteString(fmt.Sprintf("Report Performed on %v\r\n\r\n", time.Now()))

	sites := populateSites()
	progress := make(chan int, 5)
	done := make(chan bool)

	progress <- currentProgress

	go completionPercentage(sites, &progress, &row)

	row += 2
	drawStringCustom(1, row, "Scanning sites now... Please wait...", termbox.ColorBlue, termbox.ColorWhite)
	termbox.Flush()

	row += 2
	for _, groups := range sites {
		for _, url := range groups {
			go func(url string) {
				_, err := http.Get(url)
				if err != nil {
					reportFile.WriteString(fmt.Sprintf("%s - error: %s\r\n", url, err))
				} else {
					reportFile.WriteString(fmt.Sprintf("%s - successful.\r\n", url))
				}
				currentProgress += 1
				progress <- currentProgress
				done <- true
			}(url)
		}
	}

	for _, groups := range sites {
		for _, _ = range groups {
			<-done
		}
	}

	row += 2
	elapsed := time.Since(start)
	drawString(1, row, fmt.Sprintf("Done scanning sites in %s.", elapsed))
	row += 2
	reportFile.WriteString("\r\nNetwork performance\r\n")

	drawStringCustom(1, row, fmt.Sprintf("Performing speed test..."), termbox.ColorBlue, termbox.ColorWhite)
	row += 2
	drawString(1, row, fmt.Sprintf("Attempt #1: "))
	termbox.Flush()

	beginDownload := time.Now()
	go downloadingAnimation(15, &row, &finishCheck)
	bytesWritten, err := downloadFile(largeFileURL)
	finishCheck <- true
	if err != nil {
		reportFile.WriteString(fmt.Sprintf("Attempt #1 - Error grabbing %s.  Reason: %s\r\n", largeFileURL, err))
		drawString(15, row, "Error downloading file.")
		row += 2
	} else {
		elapsedDownload := time.Since(beginDownload)
		drawString(15, row, fmt.Sprintf("Downloaded %d bytes with approximate rate: %.2f KB/sec", bytesWritten, (float64(bytesWritten)/float64(elapsedDownload.Seconds()))/1024.0))
		reportFile.WriteString(fmt.Sprintf("Attempt #1 - Successfully downloaded %d bytes in %s seconds. Estimated rate: %f KB/sec\r\n", bytesWritten, elapsedDownload, (float64(bytesWritten)/float64(elapsedDownload.Seconds()))/1024.0))
		row += 2
	}

	drawString(1, row, fmt.Sprintf("Attempt #2: "))
	termbox.Flush()

	beginDownload = time.Now()
	go downloadingAnimation(15, &row, &finishCheck)
	bytesWritten, err = downloadFile(largeFileURL)
	finishCheck <- true
	if err != nil {
		reportFile.WriteString(fmt.Sprintf("Attempt #2 - Error grabbing %s.  Reason: %s\r\n", largeFileURL, err))
		drawString(15, row, "Error downloading file")
		row += 2
	} else {
		elapsedDownload := time.Since(beginDownload)
		drawString(15, row, fmt.Sprintf("Downloaded %d bytes with approximate rate: %.2f KB/sec", bytesWritten, (float64(bytesWritten)/float64(elapsedDownload.Seconds()))/1024.0))
		reportFile.WriteString(fmt.Sprintf("Attempt #2 - Successfully downloaded %d bytes in %s seconds. Estimated rate: %f KB/sec\r\n", bytesWritten, elapsedDownload, (float64(bytesWritten)/float64(elapsedDownload.Seconds()))/1024.0))
		row += 2
	}
	drawString(1, row, fmt.Sprintf("Please send file: %s to Support.  Thank you!", reportFileName))
	row += 2
	drawString(1, row, "Press any key to exit.")
	termbox.Flush()
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			break loop
		}
	}
}
