package repo

type repo struct {
	firebaseClient *db.Config
}

type RepoInterface interface {
	SavePost(contex.Context, string, string, string, string, time.Time, bool) error
}

func NewRepo(firebaseClient *db.Client) *repo {
	return repo{
		firebaseClient: firebaseClient
	}
}
