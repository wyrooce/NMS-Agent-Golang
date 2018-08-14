package plugins

import "io/ioutil"
import "fmt"
import "log"
import "path/filepath"
import "os"

type FileManager struct{
	name string
	path string

}

func (fm *FileManager) GetName() string{
	return fm.name
}

func (fm *FileManager) GetPath() string{
	return fm.path
}
//ansible conf mng
//with tree structure??
func (fm *FileManager) GetFilesInPath(path string) []string{
	files, err := ioutil.ReadDir(path)
	    if err != nil {
			fmt.Println("GetFilesInPath, err:", err)
			return nil
    }
	var result []string
    for _, file := range files {
		fmt.Println(file.Name())
		result = append(result, file.Name())
    }
	return result
}
//log session vs hamfekr
func (fm *FileManager) GetFilesInPathRec(mypath string) []string{
	var result []string
	//use this for search rec, fantastic!!!
	err := filepath.Walk(mypath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		result = append(result, path)
		return nil
		})
	if err != nil {
		log.Println(err)
	}
	return result
}



func (fm *FileManager) Search(mypath string) []string {
	var result []string
	//use this for search rec, fantastic!!!
	err := filepath.Walk(mypath, func(path string, info os.FileInfo, err error) error {
									if err != nil {
										return err
									}
									
									result = append(result, path)
									return nil
									})
	if err != nil {
		log.Println(err)
	}
	return result
}

func (fm *FileManager) copy(src string, dst string) bool {
	return false
}

func (fm *FileManager) move()