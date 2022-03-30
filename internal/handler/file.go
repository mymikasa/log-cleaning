package handler

type Option func(*FileInfos)

type FileInfos struct {
	ModificationTime int64
	VisitTime        int64
	Createtime       int64
	Size             int64
	Path             string
	Name             string
	Suffix           string
}

func WithVisitTime(time int64) Option {
	return func(f *FileInfos) {
		f.VisitTime = time
	}
}

func WithModificationTime(time int64) Option {
	return func(f *FileInfos) {
		f.ModificationTime = time
	}
}

func WithCreatetime(time int64) Option {
	return func(f *FileInfos) {
		f.Createtime = time
	}
}

func WithSize(s int64) Option {
	return func(f *FileInfos) {
		f.Size = s
	}
}

func WithPath(p string) Option {
	return func(f *FileInfos) {
		f.Path = p
	}
}

func WithFileName(p string) Option {
	return func(f *FileInfos) {
		f.Name = p
	}
}

func WithSuffix(s string) Option {
	return func(f *FileInfos) {
		f.Suffix = s
	}
}

func New(option ...Option) *FileInfos {
	o := &FileInfos{}
	for _, opt := range option {
		opt(o)
	}
	return o
}
