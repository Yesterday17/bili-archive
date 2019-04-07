// Edited from annie
package bilibili

import (
	"fmt"
	"github.com/Yesterday17/bili-archive/utils"
	"io"
	"os"
	"sort"
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

func (data *VideoStream) calculateTotalSize() {
	var size int64
	for _, urlData := range data.URLs {
		size += urlData.Size
	}
	data.Size = size
}

func (v *VideoData) genSortedStreams() {
	for k, data := range v.Streams {
		if data.Size == 0 {
			data.calculateTotalSize()
		}
		data.name = k
		v.Streams[k] = data
		v.sortedStreams = append(v.sortedStreams, data)
	}
	if len(v.Streams) > 1 {
		sort.Slice(
			v.sortedStreams, func(i, j int) bool { return v.sortedStreams[i].Size > v.sortedStreams[j].Size },
		)
	}
}

func EmptyVideoData(url string, err error) VideoData {
	return VideoData{
		URL: url,
		Err: err,
	}
}

func writeFile(url, cookies string, header map[string]string, file *os.File, pg *utils.Progress) (int64, error) {
	res, err := utils.Request("GET", url, cookies, nil, header)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	writer := io.MultiWriter(file, pg)
	// Note that io.Copy reads 32kb(maximum) from input and writes them to output, then repeats.
	// So don't worry about memory.
	written, copyErr := io.Copy(writer, res.Body)
	// TODO: Check whether the content of res.Body is HTML
	if copyErr != nil {
		return written, fmt.Errorf("file copy error: %s", copyErr)
	}
	return written, nil
}

func SaveFile(urlData VideoURL, fName, fPath, cookies string, pg *utils.Progress) error {
	filePath := utils.FilePath(fPath, fName, urlData.Ext)
	fileSize, exists, err := utils.FileSize(fName)
	if err != nil {
		pg.Error(err)
		pg.Finish()
		return err
	}
	if exists && fileSize == urlData.Size {
		// 跳过重复文件
		pg.Finish()
		return nil
	}

	tempFilePath := filePath + ".download"
	tempFileSize, _, err := utils.FileSize(tempFilePath)
	if err != nil {
		pg.Error(err)
		pg.Finish()
		return err
	}
	headers := map[string]string{}
	var (
		file      *os.File
		fileError error
	)
	if tempFileSize > 0 {
		// range start from 0, 0-1023 means the first 1024 bytes of the file
		pg.Add(tempFileSize)
		headers["Range"] = fmt.Sprintf("bytes=%d-", tempFileSize)
		file, fileError = os.OpenFile(tempFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	} else {
		file, fileError = os.Create(tempFilePath)
	}
	if fileError != nil {
		pg.Error(fileError)
		pg.Finish()
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
			written, err := writeFile(urlData.URL, cookies, headers, file, pg)
			if err == nil {
				break
			} else if i+1 >= 10 {
				// 10 为重试次数
				pg.Error(err)
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

func DownloadVideo(videoData VideoData, vData DownloadVideoRequest, basePath, cookies string, callback func(pg *utils.Progress)) error {
	videoData.genSortedStreams()
	if len(videoData.sortedStreams) == 0 {
		return fmt.Errorf("error getting video stream of av%s", vData.Aid)
	}
	stream := videoData.sortedStreams[0].name

	data, ok := videoData.Streams[stream]
	if !ok {
		return fmt.Errorf("no stream named %s", stream)
	}

	// 生成文件名
	title := vData.Page.CID

	// 最终文件路径
	finalFilePath := utils.FilePath(basePath, title, "mp4")

	// 下载视频（目前为单线程）
	bar := utils.NewProgress(title, data.Size, callback)
	parts := make([]string, len(data.URLs))
	for index, url := range data.URLs {
		fileName := fmt.Sprintf("%s[%d]", title, index)
		partFilePath := utils.FilePath(basePath, fileName, url.Ext)
		parts[index] = partFilePath

		if err := SaveFile(url, fileName, basePath, cookies, bar); err != nil {
			return err
		}
	}
	if videoData.Type != "video" {
		return nil
	}
	// 合并多段文件
	return utils.MergeToMP4(parts, finalFilePath, title)
}
