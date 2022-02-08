package controller

//
//import (
//	"fmt"
//	"os"
//	"path/filepath"
//	"strconv"
//	"strings"
//)
//
//type State struct {
//	Opened   bool `json:"opened"`
//	Selected bool `json:"selected"`
//}
//
//type Data struct {
//	Data []Inner `json:"data"`
//}
//
//type Core struct {
//	Core    Data     `json:"framework"`
//	Plugins []string `json:"plugins"`
//}
//
//type Inner struct {
//	Id      string `json:"id"`
//	Parent  string `json:"parent"`
//	Text    string `json:"text"`
//	Summary string `json:"summary"`
//	Size    string `json:"size"`
//}
//
//var directoryLocation string
//
//func Initialize() {
//	directoryLocation = os.Getenv("SRVLOCATION")
//	fmt.Println("Initialized", directoryLocation)
//}
//
//type Message struct {
//	Message string `json:"message"`
//	Result  string `json:"result"`
//}
//
//func (c App) GetFile(filename string) revel.Result {
//
//	p := returnFilePath(filename)
//	if strings.Contains(p, "error") {
//		fmt.Println("p: error: ", p)
//		m := Message{"nothing here", p}
//		return c.RenderJson(m)
//	}
//	file, err := os.Open(p) // For read access.
//	if err != nil {
//		fmt.Println("error'd out")
//	}
//	return c.RenderFile(file, "attachment") //inline would try and display it (apparently).
//}
//
//func returnFilePath(filename string) string {
//
//	var p string
//	p = "error - p not found"
//	filepath.Walk(directoryLocation, func(path string, info os.FileInfo, err error) error {
//		if strings.Contains(path, filename) {
//
//			p = path
//		}
//		return nil
//	})
//
//	fmt.Println("returning p: ", p)
//	return p
//}
//
//func getFolderStructure() []Inner {
//
//	dataArray := make([]Inner, 0)
//	currentDir := directoryLocation
//	root := strings.Split(directoryLocation, "/")
//	filepath.Walk(directoryLocation, func(path string, info os.FileInfo, err error) error {
//
//		if info.Mode().IsDir() {
//			//damned bitsync
//			if strings.Contains(path, ".sync") || strings.Contains(path, "Archive") {
//				return nil
//			}
//			var inner1 Inner
//			//dont show path directory as an explict directory
//			if path != directoryLocation { //i.e its a sub directory
//				absFileName := strings.Replace(path, currentDir, "", -1)
//				//removes .md off the file name to display
//				relFileName := strings.Split(absFileName, "/")
//				//fmt.Println("rel file name: ", relFileName[len(relFileName)-1])
//
//				/*
//
//					THIS NEXT LINE IS ONLY CHECKING FOR MARKDOWN FILES - THIS IS UNECESSARY!
//
//				*/
//				filename := strings.Replace(relFileName[len(relFileName)-1], ".md", "", -1)
//				//sorting out the parent directory:
//				parent := strings.Split(path, "/")
//				if parent[len(parent)-2] != root[len(root)-1] { //if its parent is not the root
//					inner1 = Inner{filename, parent[len(parent)-2], filename, "directory", strconv.FormatInt(info.Size(), 10)}
//				} else {
//					inner1 = Inner{filename, "#", filename, "directory", strconv.FormatInt(info.Size(), 10)}
//				}
//				dataArray = append(dataArray, inner1)
//			}
//
//		} else { // This is a markdown file
//
//			if strings.Contains(path, ".DS_Store") || strings.Contains(path, ".sync") || strings.Contains(path, "Archive") {
//				return nil
//			}
//			absFileName := strings.Replace(path, currentDir, "", -1)
//			//removes .md off the file name to display
//			relFileName := strings.Split(absFileName, "/")
//			filename := relFileName[len(relFileName)-1]
//
//			var inner1 Inner
//			parent := strings.Split(path, "/")
//			if parent[len(parent)-2] != root[len(root)-1] {
//				inner1 = Inner{filename, parent[len(parent)-2], strings.Split(filename, ".")[0], filename, strconv.FormatInt(info.Size(), 10)}
//			} else {
//				inner1 = Inner{filename, "#", strings.Split(filename, ".")[0], filename, strconv.FormatInt(info.Size(), 10)}
//			}
//			dataArray = append(dataArray, inner1)
//
//		}
//
//		return nil
//
//	}) // end of filepath.Walk
//
//	return dataArray
//}
//
//////////////////////////////////////////////////////////////////////////
//
//func (c App) Json() revel.Result {
//
//	return c.RenderJson(getFolderStructure())
//}
//
//func (c App) Explorer() revel.Result {
//	return c.Render()
//}
//
//func (c App) index() revel.Result {
//	Initialize()
//	m := Message{"nothing to see here", "NULL"}
//	return c.RenderJson(m)
//}
