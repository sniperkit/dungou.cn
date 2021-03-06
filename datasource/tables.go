package datasource

type Dungouset struct {
	Project   string `gorm:"size:512;column:project"`
	Path      string `gorm:"size:512;column:path"`
	Section   string `gorm:"size:512;column:section"`
	Dungou    string `gorm:"size:512;column:dungou"`
	Type      string `gorm:"size:512;column:type"`
	Company1  string `gorm:"size:512;column:company1"`
	Company2  string `gorm:"size:512;column:company2"`
	Client    string `gorm:"size:512;column:client"`
	Datano    string `gorm:"size:512;column:datano"`
	Pressures int    `gorm:"column:pressures"`
	Jack      int    `gorm:"column:jack"`
	Ringnum   int    `gorm:"column:ringnum"`
	Lon       string `gorm:"size:512;column:longitude"`
	Lat       string `gorm:"size:512;column:latitude"`
	City      string `gorm:"size:512;column:city"`
	Status    string `gorm:"size:512;column:status"`
	Schedule  string `gorm:"size:512;column:schedule"`
	Persent   string `gorm:"size:512;column:persent"`
	Own       string `gorm:"size:512;column:own"`
}
type Risk struct {
	Section string `gorm:"size:512;column:section"`
	Risks   string `gorm:"size:512;column:risks"`
	Ringnum string `gorm:"size:512;column:ringnum"`
}
type Rtinfo struct {
	Line string `gorm:"size:512;column:line"`
	Name string `gorm:"size:512;column:name"`
}

type Seclonlat struct {
	Section string `gorm:"size:512;column:section"`
	Lon     string `gorm:"size:512;column:lon"`
	Lat     string `gorm:"size:512;column:lat"`
}

type Prolonlat struct {
	Section string `gorm:"size:512;column:section"`
	Lon     string `gorm:"size:512;column:lon"`
	Lat     string `gorm:"size:512;column:lat"`
}
type Profile struct {
	Section string `gorm:"size:512;column:section"`
	Url     string `gorm:"size:512;column:url"`
}

type Daopan struct {
	Dungou      string `gorm:"size:512;column:dungou"`
	Jilutime    string `gorm:"size:512;column:Jilutime"`
	Shike       string `gorm:"size:512;column:shike"`
	No1         string `gorm:"size:512;column:no1"`
	No2         string `gorm:"size:512;column:no2"`
	No3         string `gorm:"size:512;column:no3"`
	No4         string `gorm:"size:512;column:no4"`
	Zongniuli   string `gorm:"size:512;column:zongniuli"`
	Waizhou     string `gorm:"size:512;column:waizhou"`
	Neizhou     string `gorm:"size:512;column:neizhou"`
	Zuozhuan    string `gorm:"size:512;column:zuozhuan"`
	Youzhuan    string `gorm:"size:512;column:youzhuan"`
	Chaowali    string `gorm:"size:512;column:chaowali"`
	Huixuansudu string `gorm:"size:512;column:huixuansudu"`
	Ringnum     int    `gorm:"column:ringnum"`
	Batch       int    `gorm:"column:batch"`
}

type Jingbao struct {
	Dungou   string `gorm:"siz:512;column:dungou"`
	Jilutime string `gorm:"size:512;column:Jilutime"`
	Shike    string `gorm:"size:512;column:shike"`
	Dyfx     string `gorm:"size:512;column:dyfx"`
	Dpcfh    string `gorm:"size:512;column:dpcfh"`
	Pzjhz    string `gorm:"size:512;column:pzjhz"`
	No1lx    string `gorm:"size:512;column:no1lx"`
	No2lx    string `gorm:"size:512;column:no2lx"`
	PzjJPU   string `gorm:"size:512;column:pzjJPU"`
	Zy       string `gorm:"size:512;column:zy"`
	Sb       string `gorm:"size:512;column:sb"`
	Jn       string `gorm:"size:512;column:jn"`
	No1zz    string `gorm:"size:512;column:no1zz"`
	No2zz    string `gorm:"size:512;column:no2zz"`
	Yl       string `gorm:"size:512;column:yl"`
	Xh       string `gorm:"size:512;column:xh"`
	Pdcfh    string `gorm:"size:512;column:pdcfh"`
	Zzdy     string `gorm:"size:512;column:zzdy"`
	Pddy     string `gorm:"size:512;column:pddy"`
	Dpdy     string `gorm:"size:512;column:dpdy"`
	Zzhl     string `gorm:"size:512;column:zzhl"`
	Yxhl     string `gorm:"size:512;column:yxhl"`
	Dpcnl    string `gorm:"size:512;column:dpcnl"`
	Hzj      string `gorm:"size:512;column:hzj"`
	Mf       string `gorm:"size:512;column:mf"`
	Dpss     string `gorm:"size:512;column:dpss"`
	Batch    int    `gorm:"column:batch"`
}

type Jiaojie struct {
	Dungou   string `gorm:"size:512;column:dungou"`
	Jilutime string `gorm:"size:512;column:Jilutime"`
	Shike    string `gorm:"size:512;column:shike"`
	Xc1      string `gorm:"size:512;column:xc1"`
	Xc2      string `gorm:"size:512;column:xc2"`
	Xc3      string `gorm:"size:512;column:xc3"`
	Xc4      string `gorm:"size:512;column:xc4"`
	Yl1      string `gorm:"size:512;column:yl1"`
	Yl2      string `gorm:"size:512;column:yl2"`
	Yl3      string `gorm:"size:512;column:yl3"`
	Yl4      string `gorm:"size:512;column:yl4"`
	Jdsx     string `gorm:"size:512;column:jdsx"`
	Jdzy     string `gorm:"size:512;column:jdzy"`
	Batch    int    `gorm:"column:batch"`
}

type Luoxuanji struct {
	Dungou   string `gorm:"size:512;column:dungou"`
	Jilutime string `gorm:"size:512;column:Jilutime"`
	Shike    string `gorm:"size:512;column:shike"`
	Hzm      string `gorm:"size:512;column:hzm"`
	Yl       string `gorm:"size:512;column:yl"`
	Zt       string `gorm:"size:512;column:zt"`
	Sd       string `gorm:"size:512;column:sd"`
	Batch    int    `gorm:"column:batch"`
}

type Juejin struct {
	Dungou   string `gorm:"size:512;column:dungou"`
	Jilutime string `gorm:"size:512;column:Jilutime"`
	Shike    string `gorm:"size:512;column:shike"`
	Fy       string `gorm:"size:512;column:fy"`
	Hz       string `gorm:"size:512;column:hz"`
	Spq      string `gorm:"size:512;column:spq"`
	CZq      string `gorm:"size:512;column:cZq"`
	Sph      string `gorm:"size:512;column:sph"`
	Czh      string `gorm:"size:512;column:czh"`
	Fw       string `gorm:"size:512;column:fw"`
	Zjdqh    string `gorm:"size:512;column:zjdqh"`
	Pmhhy    string `gorm:"size:512;column:pmhhy"`
	Jhw      string `gorm:"size:512;column:jhw"`
	Dp       string `gorm:"size:512;column:dp"`
	HBW      string `gorm:"size:512;column:hBW"`
	Ep2      string `gorm:"size:512;column:ep2"`
	Zjzs     string `gorm:"size:512;column:zjzs"`
	Zjys     string `gorm:"size:512;column:zjys"`
	Zjyx     string `gorm:"size:512;column:zjyx"`
	Zjzx     string `gorm:"size:512;column:zjzx"`
	Jjjxz    string `gorm:"size:512;column:jjjxz"`
	Jjms     string `gorm:"size:512;column:jjms"`
	Zzms     string `gorm:"size:512;column:zzms"`
	Batch    int    `gorm:"column:batch"`
}

type Tuya struct {
	Dungou   string `gorm:"size:512;column:dungou"`
	Jilutime string `gorm:"size:512;column:Jilutime"`
	Shike    string `gorm:"size:512;column:shike"`
	Qjdque   string `gorm:"size:512;column:qjdque"`
	Tuyaque  string `gorm:"size:512;column:tuyaque"`
	Tls      string `gorm:"size:512;column:tls"`
	Xcs      string `gorm:"size:512;column:xcs"`
	Sds      string `gorm:"size:512;column:sds"`
	Tlx      string `gorm:"size:512;column:tlx"`
	Xcx      string `gorm:"size:512;column:xcx"`
	Sdx      string `gorm:"size:512;column:sdx"`
	Tlz      string `gorm:"size:512;column:tlz"`
	Xcz      string `gorm:"size:512;column:xcz"`
	Sdz      string `gorm:"size:512;column:sdz"`
	Tly      string `gorm:"size:512;column:tly"`
	Xcy      string `gorm:"size:512;column:xcy"`
	Sdy      string `gorm:"size:512;column:sdy"`
	Tuya1    string `gorm:"size:512;column:tuya1"`
	Tuya2    string `gorm:"size:512;column:tuya2"`
	Tuya3    string `gorm:"size:512;column:tuya3"`
	Tuya4    string `gorm:"size:512;column:tuya4"`
	Tuya5    string `gorm:"size:512;column:tuya5"`
	Ztl      string `gorm:"size:512;column:ztl"`
	Batch    int    `gorm:"column:batch"`
}
type Commum struct {
	Dungou   string `gorm:"size:512;column:dungou"`
	Jilutime string `gorm:"size:512;column:Jilutime"`
	Shike    string `gorm:"size:512;column:shike"`
	Batch    int    `gorm:"column:batch"`
}

type Sediment struct {
	Dungou    string `gorm:"size:512;column:dungou"`
	Line      string `gorm:"size:512;column:line"`
	Project   string `gorm:"size:512;column:project"`
	Loopnum   string `gorm:"size:512;column:loopnum"`
	Loop      string `gorm:"size:512;column:loop"`
	Groundlr1 string `gorm:"size:512;column:groundlr1"`
	Loop1     string `gorm:"size:512;column:loop1"`
	Groundlr2 string `gorm:"size:512;column:groundlr2"`
	Batch     int    `gorm:"column:batch"`
}
type User struct {
	Id        int    `json:"size:32;column:id;auto_increment"`
	Username  string `json:"size:512;column:username"`
	Password  string `json:"size:512;column:password"`
	Role      string `json:"size:512;column:role"`
	Grade     string `json:"size:512;column:grade"`
	Companyid string `json:"size:512;column:companyid"`
	T1        string `json:"size:512;column:t1"`
	T2        string `json:"size:512;column:t2"`
}

type Message struct {
	Id        int    `json:"size:32;column:id;auto_increment"`
	Username  string `json:"size:512;column:username"`
	Line      string `json:"size:512;column:line"`
	Dungou    string `json:"size:512;column:dungou"`
	Type      string `json:"size:512;column:type"`
	Ringnum   string `json:"size:512;column:ringnum"`
	Schedule  string `json:"size:512;column:schedule"`
	Companyid string `json:"size:512;column:companyid"`
	Date      string `json:"size:512;column:date"`
	Text      string `json:"size:1024;column:text"`
	Img       string `json:"size:64;column:img"`
	Section   string `json:"size:64;column:section"`
}
type Remark struct {
	Id        int    `json:"size:32;column:id;auto_increment"`
	Username  string `json:"size:512;column:username"`
	Companyid string `json:"size:512;column:companyid"`
	Date      string `json:"size:512;column:date"`
	Text      string `json:"size:1024;column:text"`
	Messageid int    `json:"size:64;column:messageid"`
}

type Video struct {
	DeviceSerial string `json:"size:512;column:deviceSerial"`
	Dungou   string `json:"size:512;column:dungou"`
	Section 	 string `json:"size:512;column:section"`
	ChannelNo    string `json:"size:512;column:channelNo"`
	LiveAddress  string `json:"size:512;column:liveAddress"`
	HdAddress    string `json:"size:512;column:hdAddress"`
	Rtmp         string `json:"size:512;column:rtmp"`
	RtmpHd       string `json:"size:512;column:rtmpHd"`
	Status       string `json:"size:512;column:status"`
	Exception    string `json:"size:512;column:exception"`
	BeginTime    string `json:"size:512;column:beginTime"`
	EndTime      string `json:"size:512;column:endTime"`
}
