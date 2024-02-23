package parser

import (
	"compress/gzip"
	"encoding/xml"
	"io"
	"log"
	"os"
)

type Manual struct {
	Value string `xml:"Value,attr"`
}

type Tempo struct {
	Manual Manual `xml:"Manual"`
}

type Mixer struct {
	Tempo Tempo `xml:"Tempo"`
}

type DeviceChain struct {
	Mixer Mixer `xml:"Mixer"`
}

type RootNote struct {
	Value string `xml:"Value,attr"`
}

type Name struct {
	Value string `xml:"Value,attr"`
}

type ScaleInformation struct {
	RootNote RootNote `xml:"RootNote"`
	Name     Name     `xml:"Name"`
}

type MainTrack struct {
	DeviceChain DeviceChain `xml:"DeviceChain"`
}

type LiveSet struct {
	MainTrack MainTrack        `xml:"MainTrack"`
	ScaleInfo ScaleInformation `xml:"ScaleInformation"`
}

type Ableton struct {
	LiveSet LiveSet `xml:"LiveSet"`
}

// ALSData represents the combined data of Tempo and ScaleInformation
type ALSData struct {
	Tempo     string
	ScaleInfo ScaleInformation
}

func ExtractALS(sourcePath string) (*ALSData, error) {
	reader, err := os.Open(sourcePath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	decoder := xml.NewDecoder(gzipReader)

	var ableton Ableton
	var currentElement string

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Error reading token:", err)
			return nil, err
		}

		switch t := token.(type) {
		case xml.StartElement:
			currentElement = t.Name.Local

			// Check if we are within the <Ableton> element
			if currentElement == "Ableton" {
				if err := decoder.DecodeElement(&ableton, &t); err != nil {
					log.Println("Error decoding Ableton:", err)
					return nil, err
				}

				// Access the Tempo and ScaleInformation from the parsed data
				tempoValue := ableton.LiveSet.MainTrack.DeviceChain.Mixer.Tempo.Manual.Value
				scaleInfo := ableton.LiveSet.ScaleInfo

				alsData := &ALSData{
					Tempo:     tempoValue,
					ScaleInfo: scaleInfo,
				}

				log.Printf("Tempo: %s BPM\n", alsData.Tempo)
				log.Printf("Scale Information: %+v\n", alsData.ScaleInfo)

				return alsData, nil
			}
		}
	}

	log.Println("Tempo and ScaleInformation not found in the XML")
	return nil, nil
}
