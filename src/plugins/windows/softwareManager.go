package windows

import (
	wapi "github.com/iamacarpet/go-win64api"
	"fmt"
	"strings"
)

//Software a software detail
type Software struct {
	Name string `json:"name"`
	Publisher string `json:"publisher"`
	Version string `json:"version"`
	InstallDate string `json:"installDate"`
	Size string `json:"size"`
	Architecture string `json:"architecture"`
}

//SoftwareList returns list of softwares installed on windows
func SoftwareList() []Software{
	var result []Software
	swList, err := wapi.InstalledSoftwareList();
	 if err != nil {
        fmt.Printf("%s\r\n", err.Error())
    }

    for _, s := range swList {
		if strings.Contains(s.Name(), "Update"){
			continue
		}
		result = append(result, Software{Name: s.Name(), 
										 Version: s.Version(), 
										 Architecture: s.Architecture()})

        //fmt.Printf("%-100s - %s - %s\r\n", s.Name(), s.Architecture(), s.Version())
	}
	
	return result
}