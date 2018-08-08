package plugins

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

