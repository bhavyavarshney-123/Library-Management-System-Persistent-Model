package Library

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

//GobEncode encodes the given data
func GobEncode(object any) ([]byte, error) {
	//Declare a Buffer and a new Gob encoder
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	gob.Register(PhysicalBook{})
	gob.Register(DigitalBook{})
	//Encode the object into the buffer
	if err := encoder.Encode(object); err != nil {
		return nil, err
	}

	//Return bytes in buffer
	return buffer.Bytes(), nil
}

//GobDecode decodes the bytes into data
func GobDecode(data []byte, object Book) (Book, error) {
	//Declare a new reader from the data and a new gob Decoder
	reader := bytes.NewReader(data)
	decoder := gob.NewDecoder(reader)
	println(decoder)
	//Decode the data into object
	if err := decoder.Decode(&object); err != nil {
		fmt.Print("HERRe\n\n")
		return nil, err
	}
	fmt.Println(object)
	return object, nil
}
