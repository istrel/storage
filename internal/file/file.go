package file

import "github.com/google/uuid"

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	fileID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &File{
		ID:   fileID,
		Name: filename,
		Data: blob,
	}, nil
}
