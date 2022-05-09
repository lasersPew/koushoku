package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	. "koushoku/config"
	"koushoku/server"
	"koushoku/services"
)

func main() {
	server.Init()

	server.GET("/archive/:id/:slug/download", download)
	server.GET("/data/:id/:pageNum", serve)
	server.GET("/data/:id/:pageNum/*width", serve)

	server.NoRoute(func(c *server.Context) {
		c.Redirect(http.StatusFound, Config.Meta.BaseURL)
	})

	server.Start(Config.Server.DataPort)
}

func download(c *server.Context) {
	id, err := c.ParamInt64("id")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	fp, err := services.GetArchiveSymlink(int(id))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	} else if len(fp) == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	c.FileAttachment(fp, filepath.Base(fp))
}

func createThumbnail(c *server.Context, f io.Reader, fp string, w int) (ok bool) {
	tmp, err := os.CreateTemp("", "tmp-")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer func() {
		tmp.Close()
		os.Remove(tmp.Name())
	}()

	if _, err := io.Copy(tmp, f); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	opts := services.ResizeOptions{Width: w, Height: w * 3 / 2}
	if err := services.ResizeImage(tmp.Name(), fp, opts); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	return true
}

func serve(c *server.Context) {
	id, err := c.ParamInt("id")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	pageNum := services.GetPageNum(c.Param("pageNum"))
	if pageNum <= 0 {
		c.Status(http.StatusBadRequest)
		return
	}

	str := strings.TrimPrefix(c.Param("width"), "/")
	width, _ := strconv.Atoi(strings.TrimSuffix(str, filepath.Ext(str)))

	path, err := services.GetArchiveSymlink(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	} else if len(path) == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	var fp string
	if (pageNum == 1 && (width == 288 || width == 896)) || width == 320 {
		fp = filepath.Join(Config.Directories.Thumbnails, fmt.Sprintf("%d-%d.%d.webp", id, pageNum, width))
		if _, err := os.Stat(fp); err == nil {
			c.ServeFile(fp)
			return
		}
	}

	zf, err := zip.OpenReader(path)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer zf.Close()

	var files []*zip.File
	for _, f := range zf.File {
		stat := f.FileInfo()
		name := stat.Name()

		if stat.IsDir() || !services.IsImage(name) {
			continue
		}
		files = append(files, f)
	}

	index := pageNum - 1
	if index > len(files) {
		c.Status(http.StatusNotFound)
		return
	}

	sort.SliceStable(files, func(i, j int) bool {
		return services.GetPageNum(filepath.Base(files[i].Name)) < services.GetPageNum(filepath.Base(files[j].Name))
	})

	file := files[index]
	stat := file.FileInfo()

	f, err := file.Open()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if len(fp) > 0 {
		if createThumbnail(c, f, fp, width) {
			c.ServeFile(fp)
		}
	} else {
		buf, err := io.ReadAll(f)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.ServeData(stat, bytes.NewReader(buf))
	}
}
