package copier

import "github.com/google/go-github/v41/github"

// SetClient allows the client ot be overridden so we can take advantage of mocking.
func (r *ReleaseCopier) SetClient(client *github.Client) {
	r.client = client
}

// SetSourceOwner allows the source owner to be overridden so we can take advantage of mocking.
func (r *ReleaseCopier) SetSourceOwner(sourceOwner string) {
	r.sourceOwner = sourceOwner
}

// SetSourceRepo allows the source repo to be overridden so we can take advantage of mocking.
func (r *ReleaseCopier) SetSourceRepo(sourceRepo string) {
	r.sourceRepo = sourceRepo
}

// SetTargetOwner allows the target owner to be overridden so we can take advantage of mocking.
func (r *ReleaseCopier) SetTargetOwner(targetOwner string) {
	r.targetOwner = targetOwner
}

// SetTargetRepo allows the target repo to be overridden so we can take advantage of mocking.
func (r *ReleaseCopier) SetTargetRepo(targetRepo string) {
	r.targetRepo = targetRepo
}
