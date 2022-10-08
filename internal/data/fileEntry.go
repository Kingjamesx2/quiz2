//filename: internal/data/fileEntry.go

package data

import (
	"quiz2/jamesfaber.net/internal/validator"
	"time"
)

type FileEntry struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Level     string    `json:"level"`
	Contact   string    `json:"contact"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email,omitempty"`
	Website   string    `json:"website,omitempty"`
	Address   string    `json:"address"`
	Mode      []string  `json:"mode"`
	Version   int32     `json:"version"`
}

func ValidateFileEntry(v *validator.Validator, fileEntry *FileEntry) {
	// Use the Check() method to execute our validation checks
	v.Check(fileEntry.Name != "", "name", "must be provided")
	v.Check(len(fileEntry.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(fileEntry.Level != "", "level", "must be provided")
	v.Check(len(fileEntry.Level) <= 200, "level", "must not be more than 200 bytes long")

	v.Check(fileEntry.Contact != "", "contact", "must be provided")
	v.Check(len(fileEntry.Contact) <= 200, "contact", "must not be more than 200 bytes long")

	v.Check(fileEntry.Phone != "", "phone", "must be provided")
	v.Check(validator.Matches(fileEntry.Phone, validator.PhoneRX), "phone", "must be a valid phone number")

	v.Check(fileEntry.Email != "", "email", "must be provided")
	v.Check(validator.Matches(fileEntry.Email, validator.EmailRX), "email", "must be a valid email address")

	v.Check(fileEntry.Website != "", "website", "must be provided")
	v.Check(validator.ValidWebsite(fileEntry.Website), "website", "must be a valid URL")

	v.Check(fileEntry.Address != "", "address", "must be provided")
	v.Check(len(fileEntry.Address) <= 500, "address", "must not be more than 500 bytes long")

	v.Check(fileEntry.Mode != nil, "mode", "must be provided")
	v.Check(len(fileEntry.Mode) >= 1, "mode", "must contain at least 1 entry")
	v.Check(len(fileEntry.Mode) <= 5, "mode", "must contain at most 5 entries")
	v.Check(validator.Unique(fileEntry.Mode), "mode", "must not contain duplicate entries")
}
