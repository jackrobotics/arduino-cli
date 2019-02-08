package arduino

import (
	"log"

	semver "go.bug.st/relaxed-semver"
)

type VersionCodec struct{}

func (cb *VersionCodec) Marshal(v interface{}) ([]byte, error) {
	p, ok := v.(*semver.Version)
	if !ok {
		log.Fatalf("Invalid type of struct: %+v", v)
	}
	return []byte(p.String()), nil
}

func (cb *VersionCodec) Unmarshal(data []byte, v interface{}) error {
	// p, ok := v.(*semver.Version)
	// if !ok {
	// 	log.Fatalf("Invalid type of struct: %+v", v)
	// }
	// p1, err := semver.Parse(string(data))
	// if err != nil {
	// 	return err
	// }
	// p1.CopyTo(p)
	// return err
	return nil
}

func (cb *VersionCodec) String() string {
	return "arduino.Version"
}
