package mysql

type Video struct {
	Id        int
	UniqueStr string
	MediaId   int
	Length    int
}

type Media struct {
	Id     int
	String string
}

type Action struct {
	Id      int
	VideoId string
	Type    ActionType
	Time    int
	Start   int
	End     int
}

type ActionType struct {
	Id     int
	String string
}
