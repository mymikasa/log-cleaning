package clean

type FileInfos struct {
	ModificationTime int64
	VisitTime        int64
	Createtime       int64
	Size             int64
	Path             string
	Name             string
	Suffix           string
}

type Option func(*FileInfos)

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
