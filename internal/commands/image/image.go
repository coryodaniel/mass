package image

type PushImageInput struct {
	ImageName          string
	Location           string
	OrganizationID     string
	Tag                string
	ArtifactID         string
	Dockerfile         string
	DockerBuildContext string
	TargetPlatform     string
}

type ErrorLine struct {
	Error       string      `json:"error"`
	ErrorDetail ErrorDetail `json:"errorDetail"`
}

type ErrorDetail struct {
	Message string `json:"message"`
}
