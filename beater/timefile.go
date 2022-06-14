package beater

import(
	"fmt"
	"io/ioutil"
	"log"
)

type Archivos struct{
	Nombre string
	Fecha_crea string
}

func getTime(path string)[]Archivos{
	archivos, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	longitud_archivos := len(archivos)
	info := []Archivos{}
	for k := 0; k <  longitud_archivos; k++{
		archivo := archivos[k]
		nombre := archivo.Name()
		fecha := archivo.ModTime()
		datos := Archivos{nombre, fecha.Format("15:04:05")}
		info = append(info,datos)
	}
	fmt.Println("info: ",info)
	return info
}
