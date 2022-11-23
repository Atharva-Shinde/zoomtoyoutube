// get list of zoom cloud recordings
package main

type (
	RecordingFile struct {
		DownloadURL string `json:download_url`
		ID          string `json:id`         // id of recording file
		MeetingID   string `json:meeting_id` // meeting number/id
		PlayURL     string `json:play_url`
		Status      string `json:status`
	}

	file struct {
		MeetingID     string `json:meetingid`
		UUID          string `json:uuid`
		IncludeFields string `json:include_fields`
		AccountID     string `json:account_id`
	}
)
