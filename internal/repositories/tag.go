package repositories

type TagEntity struct {
	Title string
	Color string
}

type TagRepository interface {
	AddTag()
	UpdateTag()
	DeleteTag()
	GetTags()
}

type TagRepo struct{}

func (t *TagRepo) AddTag() {}

func (t *TagRepo) UpdateTag() {}

func (t *TagRepo) DeleteTag() {}

func (t *TagRepo) GetTags() {}
