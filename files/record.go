package files

type Record struct {
	Source string
	Target string
}

func NewRecord(Source, Target string) *Record {
	return &Record{Source, Target}
}

func (record *Record) TargetExists() bool {
	return exists(record.Target)
}

func (record *Record) CopyRecord() {
	CopyDirectory(record.Source, record.Target)
}
