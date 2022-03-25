package recording

import "encoding/xml"

//查询回放记录
type GetRecordings struct {
	XMLName xml.Name `xml:"GetRecordings"`
}

//查询回放记录回复
type GetRecordingsResponse struct {
	RecordingItem []RecordingItem
}

type RecordingItem struct {
	RecordingToken string
	Configuration struct{
		Source struct{
			SourceId string
			Name string
			Location string
			Description string
			Address string
		}
		Content string
		MaximumRetentionTime string
	}
	Tracks []Track
}

type Track struct {
	TrackToken string
	Configuration struct{
		TrackType string
		Description string
	}
}