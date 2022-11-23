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

	Meeting struct {
		AccountID      string          `json:account_id`
		HostID         string          `json:host_id`
		ID             string          `json:id`
		IncludeFields  string          `json:include_fields`
		Password       string          `json:password` // cloud recording password
		RecordingCount string          `json:recording_count`
		RecordingFiles []RecordingFile `json:recording_files`
		Topic          string          `json:topic`
		UUID           string          `json:uuid`
	}
)
