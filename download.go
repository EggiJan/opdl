package main

import (
  "fmt"
  "net/http"
  "net/url"
  "os"
  "strings"
  "path/filepath"
  "strconv"
  "io"
)

func Download(imageURL string)  {
  fmt.Println("downloading", imageURL)
  chapString := strconv.Itoa(chapter)

  fileURL, err := url.Parse(imageURL)

  if err != nil {
    panic(err)
  }

  path := fileURL.Path
  segments := strings.Split(path, "/")


  fileName := strings.Join(segments, "_")
  fileName = strings.TrimPrefix(fileName, "_")

  os.Mkdir(folder + "/" + chapString, 0777)

  // Create the file
  file, err := os.Create(filepath.Join(folder, chapString, fileName))

  if err != nil {
    fmt.Println(err)
    panic(err)
  }

  defer file.Close()

  check := http.Client{
    CheckRedirect: func (r *http.Request, via []*http.Request) error  {
      r.URL.Opaque = r.URL.Path
      return nil
    },
  }

  resp, err := check.Get(imageURL)

  if err != nil {
    fmt.Println(err)
    panic(err)
  }

  defer resp.Body.Close()


  size, err := io.Copy(file, resp.Body)

  if err != nil {
    panic(err)
  }

  if resp.StatusCode >= 400 {
    fp := filepath.Join(folder, chapString, fileName)
    os.Remove(fp)
  }

  fmt.Printf("%s with %v bytes downloaded\n", fileName, size)
}
