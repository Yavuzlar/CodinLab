package file

import "os"

func CreateDir(dir string) (err error) {
	err = os.Mkdir(dir, 0755)
	if err != nil {
		return err
	}

	return nil
}

func CheckDir(dir string) (err error) {
	fileInfo, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return err
	}

	if !fileInfo.IsDir() {
		return os.ErrNotExist
	}

	return nil
}

func CreateFileAndWrite(filePath, content string) (err error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	if content != "" {
		file.WriteString(content)
	}

	defer file.Close()

	return nil
}

func CheckFile(filePath string) (err error) {
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		return err
	}
	return nil
}
