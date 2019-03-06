package bilibili

import (
	"fmt"
	"github.com/Yesterday17/bili-archive/utils"
	"io"
	"os"
	"time"
)

type VideoURL struct {
	URL  string `json:"url"`
	Size int64  `json:"size"`
	Ext  string `json:"ext"`
}

type VideoStream struct {
	// [VideoURL: {VideoURL, Size, Ext}, ...]
	// Some video files have multiple fragments
	// and support for downloading multiple image files at once
	URLs    []VideoURL `json:"urls"`
	Quality string     `json:"quality"`
	// total size of all urls
	Size int64 `json:"size"`

	// name used in sortedStreams
	name string
}

type VideoData struct {
	Site  string `json:"site"`
	Title string `json:"title"`
	Type  string `json:"type"`
	// each stream has it's own URLs and Quality
	Streams       map[string]VideoStream `json:"streams"`
	sortedStreams []VideoStream

	// Err is used to record whether an error occurred when extracting data.
	// It is used to record the error information corresponding to each url when extracting the list data.
	// NOTE(iawia002): err is only used in VideoData list
	Err error `json:"-"`
	// VideoURL is used to record the address of this download
	URL string `json:"url"`
}

func EmptyVideoData(url string, err error) VideoData {
	return VideoData{
		URL: url,
		Err: err,
	}
}

func writeFile(url, cookies string, header map[string]string, file *os.File) (int64, error) {
	res, err := utils.Request("GET", url, cookies, nil, header)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	writer := io.MultiWriter(file)
	// Note that io.Copy reads 32kb(maximum) from input and writes them to output, then repeats.
	// So don't worry about memory.
	written, copyErr := io.Copy(writer, res.Body)
	if copyErr != nil {
		return written, fmt.Errorf("file copy error: %s", copyErr)
	}
	return written, nil
}

func SaveFile(urlData VideoURL, fileName, filePath, cookies string, callback func(status string)) error {
	filePath, err := utils.FilePath(filePath, fileName, urlData.Ext, false)
	if err != nil {
		return err
	}
	fileSize, exists, err := utils.FileSize(fileName)
	if err != nil {
		return err
	}
	if exists && fileSize == urlData.Size {
		// 跳过重复文件
		return nil
	}

	tempFilePath := filePath + ".download"
	tempFileSize, _, err := utils.FileSize(tempFilePath)
	if err != nil {
		return err
	}
	headers := map[string]string{}
	var (
		file      *os.File
		fileError error
	)
	if tempFileSize > 0 {
		// range start from 0, 0-1023 means the first 1024 bytes of the file
		headers["Range"] = fmt.Sprintf("bytes=%d-", tempFileSize)
		file, fileError = os.OpenFile(tempFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	} else {
		file, fileError = os.Create(tempFilePath)
	}
	if fileError != nil {
		return fileError
	}
	var start, end, chunkSize int64
	chunkSize = 5 * 1024 * 1024
	remainingSize := urlData.Size
	if tempFileSize > 0 {
		start = tempFileSize
		remainingSize -= tempFileSize
	}
	chunk := remainingSize / chunkSize
	if remainingSize%chunkSize != 0 {
		chunk++
	}
	var i int64 = 1
	for ; i <= chunk; i++ {
		end = start + chunkSize - 1
		headers["Range"] = fmt.Sprintf("bytes=%d-%d", start, end)
		temp := start
		for i := 0; ; i++ {
			written, err := writeFile(urlData.URL, cookies, headers, file)
			if err == nil {
				break
			} else if i+1 >= 10 {
				// 10 为重试次数
				return err
			}
			temp += written
			headers["Range"] = fmt.Sprintf("bytes=%d-%d", temp, end)
			time.Sleep(1 * time.Second)
		}
		start = end + 1
	}

	// 结束时重命名下载文件
	defer func() {
		// 先关闭再修改
		// 否则会提示有进程占用
		file.Close()
		if err == nil {
			os.Rename(tempFilePath, filePath)
		}
	}()
	return nil
}

func DownloadVideo(v VideoData, vData DownloadVideoRequest, basePath, cookies string, callback func(status string)) error {
	var err error
	v.genSortedStreams()
	stream := v.sortedStreams[0].name

	data, ok := v.Streams[stream]
	if !ok {
		return fmt.Errorf("no stream named %s", stream)
	}

	// 生成文件名
	title := fmt.Sprintf("%s - %s", vData.Title, vData.Page.PageName)

	// 跳过已合并文件
	mergedFilePath, err := utils.FilePath(basePath, title, "mp4", false)
	if err != nil {
		return err
	}
	_, mergedFileExists, err := utils.FileSize(mergedFilePath)
	if err != nil {
		return err
	}
	// After the merge, the file size has changed, so we do not check whether the size matches
	if mergedFileExists {
		fmt.Printf("%s: file already exists, skipping\n", mergedFilePath)
		return nil
	}
	// 单线程下载
	parts := make([]string, len(data.URLs))
	for index, url := range data.URLs {
		var fileName, partFilePath string
		if len(data.URLs) == 1 {
			fileName = title
		} else {
			fileName = fmt.Sprintf("%s[%d]", title, index)
			partFilePath, err = utils.FilePath(basePath, fileName, url.Ext, false)
			if err != nil {
				return err
			}
			parts[index] = partFilePath
		}
		err := SaveFile(url, fileName, basePath, cookies, callback)
		if err != nil {
			return err
		}
	}
	if v.Type != "video" {
		return nil
	}
	// 合并多段文件
	if len(data.URLs) != 1 {
		fmt.Printf("Merging video parts into %s\n", mergedFilePath)
		// TODO: 合并文件
		//if err := utils.MergeToMP4(parts, mergedFilePath, title); err != nil {
		//	return err
		//}
	}
	return nil
}
