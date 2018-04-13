package fs

import (
	"path/filepath"
	"strings"
	"os"
	"archive/zip"
	"io/ioutil"
)

//Zipa um arquivo e o salva na pasta dir.
//O nome do arquivo gerado é o mesmo que path, exceto pela extensão, que é zip.
//Se dir for vazio, salva no mesmo diretório que file.
func ZipFile(path, dir string) error {
	if dir == "" {
		dir = filepath.Dir(path)
	}
	nome := filepath.Base(path)
	pathZip := filepath.Join(dir, nome[:strings.LastIndex(nome, ".")]+".zip")
	z, err := os.OpenFile(pathZip, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0744)
	if err != nil {
		return err
	}
	defer z.Close()
	w := zip.NewWriter(z)
	defer w.Close()
	f, err := w.Create(nome)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	f.Write(data)
	return nil
}
