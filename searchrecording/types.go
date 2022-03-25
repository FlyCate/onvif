package searchrecording

import "encoding/xml"

//查询回放概要信息
type GetRecordingSummary struct {
	XMLName xml.Name `xml:"GetRecordingSummary"`
}

type GetRecordingSummaryResponse struct {
	Summary struct{
		DataFrom string
		DataUntil string
		NumberRecordings int
	}
}

//查询回放token
type FindRecordings struct {
	Scope string
	KeepaliveTime string
}

type FindRecordingsResponse struct {
	SearchToken string
}

//查询录像记录
type GetRecordingSearchResults struct {
	SearchToken string
	MinResults int
	MaxResults int
	WaitTime string
}

type GetRecordingSearchResultsResponse struct {
	ResultList struct{
		SearchState string
		RecordingInformation []struct{
			RecordingToken string
			Source struct{
				SourceId string
				Name string
				Location string
				Description string
				Address string
			}
			EarliestRecording string
			LatestRecording string
			Content string
			Track []struct{
				TrackToken string
				TrackType string
				Description string
				DataFrom string
				DataTo string
			}
			RecordingStatus string
		}
	}
}