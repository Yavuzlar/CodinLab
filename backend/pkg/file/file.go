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
