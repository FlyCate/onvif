package replay

//查询回放URI
type GetReplayUri struct {
	StreamSetup struct{
		Stream string
		Transport struct{
			Protocol string
		}
	}
	RecordingToken string
}

type GetReplayUriResponse struct {
	Uri string
}