package facade

/*
	外观模式在客户端和现有系统之间加入一个外观对象, 为子系统提供一个统一的接入接口, 类似与委托
*/

type Music struct {
	Name string
}

func (m *Music) GetMusic() string {
	return m.Name
}

type Video struct {
	Id int64
}

func (v *Video) GetVideoId() int64 {
	return v.Id
}

type Count struct {
	Comment int64
	Praise  int64
	Collect int64
}

func (c *Count) GetComment() int64 {
	return c.Comment
}

type Facade struct {
	music Music
	count Count
	video Video
}

func (f *Facade) PrintServerInfo() {
	f.music.GetMusic()
	f.video.GetVideoId()
	f.count.GetComment()
}

func NewFacade(music Music, count Count, video Video) *Facade {
	return &Facade{
		music: music,
		video: video,
		count: count,
	}
}
