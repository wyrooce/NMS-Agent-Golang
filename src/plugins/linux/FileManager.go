package plugins

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
	"strconv"
	"github.com/shirou/gopsutil/disk"
)

//FileManager => a struct define filemanger
type FileManager struct {
	name string
	path string
}

func (fm *FileManager) GetPath() string {
	return fm.path
}

//ansible conf mng
//with tree structure??
func (fm *FileManager) GetFilesInPath(path string) []string {
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


//GetFilesInPathRec => log session vs hamfekr
func (fm *FileManager) GetFilesInPathRec(mypath string) []string {
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

//Search => arry of files found. exactly if set false, return contains string name
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
				goal = "*" + targetFile + "*"
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
	from, err := os.Open(src) //"./sourcefile.txt"
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer from.Close()

	to, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0666) //"./sourcefile.copy.txt"
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (fm *FileManager) Move(srcFile, targetPath string) bool {
	//filename := srcFile[s.LastIndex(srcFile, "/")+1 : len(srcFile)]
	filename := fm.GetName(srcFile)
	fmt.Println(filename)
	err := os.Rename(srcFile, targetPath+"/"+filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("File", filename+": Moved from", srcFile, "to", targetPath+"/"+filename)
	return true

}

func (fm *FileManager) Rename(srcFile, targetPath string) bool {
	// filename := srcFile[s.LastIndex(srcFile, "/")+1 : len(srcFile)]
	filename := fm.GetName(srcFile)
	fmt.Println(filename)
	err := os.Rename(srcFile, targetPath+"/"+filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("File:", "Renamed from ", srcFile, "to", targetPath+"/"+filename)
	return true
}

func (fm *FileManager) GetSize(filepath string) int64 {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// fmt.Println(fi.Sys())
	return fi.Size() // in byte
}

func (fm *FileManager) GetLastModifiedDate(filepath string) time.Time {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return fi.ModTime()
}

func (fm *FileManager) GetName(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return fi.Name()
}
func (fm *FileManager) Remove(filepath string) bool {

	// delete file
	si, e := os.Stat(filepath)
	if os.IsExist(e) && si.IsDir() {
		err := os.RemoveAll(filepath) //if is a full dir
		if err != nil {
			return false
		}
	} else {
		err := os.Remove(filepath) //file only, not dir
		if err != nil {
			return false
		}
	}
	return true
}

//PrintPartitions => all partitions of the host system and details
func (fm *FileManager) PrintPartitions() {
	parts, err := disk.Partitions(false)
	check(err)

	var usage []*disk.UsageStat

	for _, part := range parts {
		u, err := disk.Usage(part.Mountpoint)
		check(err)
		usage = append(usage, u)
		printUsage(u)
	}
}

func printUsage(u *disk.UsageStat) {
	fmt.Println(u.Path + "\t" + strconv.FormatFloat(u.UsedPercent, 'f', 2, 64) + "% full.")
	fmt.Println("Total: " + strconv.FormatUint(u.Total/1024/1024/1024, 10) + " GiB")
	fmt.Println("Free:  " + strconv.FormatUint(u.Free/1024/1024/1024, 10) + " GiB")
	fmt.Println("Used:  " + strconv.FormatUint(u.Used/1024/1024/1024, 10) + " GiB")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//Partitions => disk partitions label, size in GiB and used GiB
func (fm *FileManager) Partitions() ([]string, []string, []string) {
	parts, err := disk.Partitions(false)
	check(err)

	var usage []*disk.UsageStat
	var label []string
	var size []string
	var used []string

	for _, part := range parts {
		u, err := disk.Usage(part.Mountpoint)
		check(err)
		usage = append(usage, u)

		label = append(label, u.Path)
		size = append(size, strconv.FormatUint(u.Total/1024/1024/1024, 10))
		used = append(used, strconv.FormatUint(u.Used/1024/1024/1024, 10))
	}
	return label, size, used
}

