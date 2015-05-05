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
