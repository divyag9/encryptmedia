package encryptMedia

// Media represents the data provided for a media coming in
type Media struct {
	Version            int64
	GUID               string
	Client             string
	LoanType           string
	OrderNumber        string
	UserName           string
	Latitude           float32
	Longitude          float32
	DateTaken          string
	DeviceModel        string
	DeviceOS           string
	DeviceOSVersion    string
	FileName           string
	MimeType           string
	Application        string
	ApplicationID      string
	ApplicationVersion string
	Bytes              []byte
}

// MediaEncrypted represents encrypted media bytes and keys, along with other properties from Media
type MediaEncrypted struct {
	Version            int64
	GUID               string
	Client             string
	LoanType           string
	OrderNumber        string
	UserName           string
	Latitude           float32
	Longitude          float32
	DateTaken          string
	DeviceModel        string
	DeviceOS           string
	DeviceOSVersion    string
	FileName           string
	MimeType           string
	Application        string
	ApplicationID      string
	ApplicationVersion string
	EncryptedBytes     []byte
	EncryptedKey       []byte
	PublicKey          []byte
}

// MediaService is service for managing Media
type MediaService interface {
	SaveMediaEncrypted(mediaEncryptedBytes []byte, fileName string) error
}
