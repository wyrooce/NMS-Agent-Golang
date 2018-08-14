package plugins

import "io/ioutil"
import "fmt"
import "log"
import "path/filepath"
import "os"
import s "strings"


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

func (fm *FileManager) Search(targetFolder, targetFile string, exactly bool) []string {
	var searchResult []string

	var findFile = func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
                 fmt.Println(err)
                 return nil
         }

         // get absolute path of the folder that we are searching
         absolute, err := filepath.Abs(path)

         if err != nil {
                 fmt.Println(err)
                 return nil
         }

         if fileInfo.IsDir() {
                 fmt.Println("Searching directory ... ", absolute)

                 // correct permission to scan folder?
                 testDir, err := os.Open(absolute)

                 if err != nil {
                         if os.IsPermission(err) {
                                 fmt.Println("No permission to scan ... ", absolute)
                                 fmt.Println(err)
                         }
                 }
                 testDir.Close()
                 return nil
         } else {
                 // ok, we are dealing with a file
                 // is this the target file?

                 // yes, need to support wildcard search as well
                 // https://www.socketloop.com/tutorials/golang-match-strings-by-wildcard-patterns-with-filepath-match-function
				goal := targetFile
				if !exactly {
				 goal = "*"+targetFile+"*"
				}
                 matched, err := filepath.Match(goal, fileInfo.Name())
                 if err != nil {
                         fmt.Println(err)
                 }

                 if matched {
                         // yes, add into our search result
                         add := absolute
                         searchResult = append(searchResult, add)
                 }
         }

         return nil
 }
 //-------------------------------------------------
 //-- find file usage
 	fmt.Println("Searching for [", targetFile, "]")
         fmt.Println("Starting from directory [", targetFolder, "]")

         // sanity check
         testFile, err := os.Open(targetFolder)
         if err != nil {
                 fmt.Println(err)
                 os.Exit(-1)
         }
         defer testFile.Close()

         testFileInfo, _ := testFile.Stat()
         if !testFileInfo.IsDir() {
                 fmt.Println(targetFolder, " is not a directory!")
                 os.Exit(-1)
         }

         err = filepath.Walk(targetFolder, findFile)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(-1)
         }

         // display our search result

         fmt.Println("\n\nFound ", len(searchResult), " hits!")
         //fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")

        //  for _, v := range searchResult {
        //          fmt.Println(v)
        //  }
	
	return searchResult
}

func (fm *FileManager) Copy(src string, dst string) bool {
	return false
}

func (fm *FileManager) Move(srcFile, targetPath string) bool {
	filename := srcFile[s.LastIndex(srcFile, "/")+1:len(srcFile)]
	fmt.Println(filename)
	err := os.Rename(srcFile, targetPath+"/"+filename)
      if err != nil {
		fmt.Println(err)			  
		return false
	  }
	fmt.Println("File",filename+": Moved from", srcFile,"to",targetPath+"/"+filename)
	return true

}



