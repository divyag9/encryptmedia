# encryptmedia
Example:

1. Making a call to the server
http://localhost:8080/

POST request body

{
	"data": [8, 1, 18, 4, 116, 101, 115, 116, 26, 4, 116, 101, 115, 116, 34, 4, 116, 101, 115, 116, 42, 4, 116, 101, 115, 116, 50, 4, 116, 101, 115, 116, 61, 0, 0, 128, 63, 69, 0, 0, 0, 64, 74, 4, 116, 101, 115, 116, 82, 4, 116, 101, 115, 116, 90, 4, 116, 101, 115, 116, 98, 4, 116, 101, 115, 116, 106, 4, 116, 101, 115, 116, 114, 4, 116, 101, 115, 116, 122, 4, 116, 101, 115, 116, 130, 1, 4, 116, 101, 115, 116, 138, 1, 4, 116, 101, 115, 116, 146, 1, 4, 1, 2, 3, 4]
}

2. The bytes are of the following protobuf

Media{Version: 1,
		GUID:               "test",
		Client:             "test",
		LoanType:           "test",
		OrderNumber:        "test",
		UserName:           "test",
		Latitude:           1.0,
		Longitude:          2.0,
		DateTaken:          "test",
		DeviceModel:        "test",
		DeviceOS:           "test",
		DeviceOSVersion:    "test",
		FileName:           "test",
		Bytes:              []byte{1, 2, 3, 4},
		MimeType:           "test",
		Application:        "test",
		ApplicationID:      "test",
		ApplicationVersion: "test"}
