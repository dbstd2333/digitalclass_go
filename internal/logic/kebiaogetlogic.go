package logic

import (
	"context"
	"encoding/json"
	"time"

	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KebiaogetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}
type KebiaoJson struct {
	Class     int    `json:"class"`
	Weekly    string `json:"weekly"`
	Lession1  string `json:"lession1"`
	Src1      string `json:"src1"`
	Lession2  string `json:"lession2"`
	Src2      string `json:"src2"`
	Lession3  string `json:"lession3"`
	Src3      string `json:"src3"`
	Lession4  string `json:"lession4"`
	Src4      string `json:"src4"`
	Lession5  string `json:"lession5"`
	Src5      string `json:"src5"`
	Lession6  string `json:"lession6"`
	Src6      string `json:"src6"`
	Lession7  string `json:"lession7"`
	Src7      string `json:"src7"`
	Lession8  string `json:"lession8"`
	Src8      string `json:"src8"`
	Lession9  string `json:"lession9"`
	Src9      string `json:"src9"`
	Lession10 string `json:"lession10"`
	Src10     string `json:"src10"`
	Lession11 string `json:"lession11"`
	Src11     string `json:"src11"`
	Lession12 string `json:"lession12"`
	Src12     string `json:"src12"`
	Lession13 string `json:"lession13"`
	Src13     string `json:"src13"`
}

func NewKebiaogetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KebiaogetLogic {
	return &KebiaogetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KebiaogetLogic) Kebiaoget(req *types.Kebiaoreq) (resp *types.Kebiaores, err error) {
	if req.Weekly == 0 {

		rsbj, err := l.svcCtx.Redis.Get("class" + string(req.Class) + "subject").Result() //查redis
		var Sbj []KebiaoJson
		errj := json.Unmarshal([]byte(rsbj), &Sbj)

		if errj != nil || err != nil { //如果redis查不到数据或json无法对应，则查mysql后回填redis，return

			logs := svc.Log{ //写入mysql log
				Type: "info",
				Logs: "redis缓存未命中/json解析失败",
				Time: time.Now(),
			}
			l.svcCtx.Mysql.Create(&logs)

			var sbj1 svc.Subject //初始化查询结构体
			var sbj2 svc.Subject
			var sbj3 svc.Subject
			var sbj4 svc.Subject
			var sbj5 svc.Subject
			var sbj6 svc.Subject
			var sbj7 svc.Subject
			var sbj8 svc.Subject
			var sbj9 svc.Subject
			var sbj10 svc.Subject
			var sbj11 svc.Subject
			var sbj12 svc.Subject
			var sbj13 svc.Subject
			//开始查mysql
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 1, time.Now().Weekday()).First(&sbj1)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 2, time.Now().Weekday()).First(&sbj2)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 3, time.Now().Weekday()).First(&sbj3)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 4, time.Now().Weekday()).First(&sbj4)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 5, time.Now().Weekday()).First(&sbj5)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 6, time.Now().Weekday()).First(&sbj6)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 7, time.Now().Weekday()).First(&sbj7)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 8, time.Now().Weekday()).First(&sbj8)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 9, time.Now().Weekday()).First(&sbj9)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 10, time.Now().Weekday()).First(&sbj10)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 11, time.Now().Weekday()).First(&sbj11)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 12, time.Now().Weekday()).First(&sbj12)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 13, time.Now().Weekday()).First(&sbj13)
			mkebiao := &KebiaoJson{ //创建map，统一写入，避免多次写redis
				Class:     int(req.Class),
				Weekly:    time.Now().Weekday().String(),
				Lession1:  sbj1.Sbjname,
				Src1:      sbj1.Src,
				Lession2:  sbj2.Sbjname,
				Src2:      sbj2.Src,
				Lession3:  sbj3.Sbjname,
				Src3:      sbj3.Src,
				Lession4:  sbj4.Sbjname,
				Src4:      sbj4.Src,
				Lession5:  sbj5.Sbjname,
				Src5:      sbj5.Src,
				Lession6:  sbj6.Sbjname,
				Src6:      sbj6.Src,
				Lession7:  sbj7.Sbjname,
				Src7:      sbj7.Src,
				Lession8:  sbj8.Sbjname,
				Src8:      sbj8.Src,
				Lession9:  sbj9.Sbjname,
				Src9:      sbj9.Src,
				Lession10: sbj10.Sbjname,
				Src10:     sbj10.Src,
				Lession11: sbj11.Sbjname,
				Src11:     sbj11.Src,
				Lession12: sbj12.Sbjname,
				Src12:     sbj12.Src,
				Lession13: sbj13.Sbjname,
				Src13:     sbj13.Src,
			}
			rdkebiao, _ := json.Marshal(mkebiao)        //map转json
			rwerr := l.svcCtx.Redis.Set("class"+string(rune(req.Class))+"subject", rdkebiao, 28800).Err() //mysql数据写入redis
			if rwerr != nil {
				logs := svc.Log{ //写入mysql log
					Type: "warning",
					Logs: "redis写入失败 kebiaologic.go 138",
					Time: time.Now(),
				}
				l.svcCtx.Mysql.Create(&logs)
			}
			return &types.Kebiaores{
				Weekly:    time.Now().Weekday().String(),
				Lession1:  sbj1.Sbjname,
				Src1:      sbj1.Src,
				Lession2:  sbj2.Sbjname,
				Src2:      sbj2.Src,
				Lession3:  sbj3.Sbjname,
				Src3:      sbj3.Src,
				Lession4:  sbj4.Sbjname,
				Src4:      sbj4.Src,
				Lession5:  sbj5.Sbjname,
				Src5:      sbj5.Src,
				Lession6:  sbj6.Sbjname,
				Src6:      sbj6.Src,
				Lession7:  sbj7.Sbjname,
				Src7:      sbj7.Src,
				Lession8:  sbj8.Sbjname,
				Src8:      sbj8.Src,
				Lession9:  sbj9.Sbjname,
				Src9:      sbj9.Src,
				Lession10: sbj10.Sbjname,
				Src10:     sbj10.Src,
				Lession11: sbj11.Sbjname,
				Src11:     sbj11.Src,
				Lession12: sbj12.Sbjname,
				Src12:     sbj12.Src,
				Lession13: sbj13.Sbjname,
				Src13:     sbj13.Src,
			}, nil
		} else {
			var sbj1 svc.Subject //初始化查询结构体
			var sbj2 svc.Subject
			var sbj3 svc.Subject
			var sbj4 svc.Subject
			var sbj5 svc.Subject
			var sbj6 svc.Subject
			var sbj7 svc.Subject
			var sbj8 svc.Subject
			var sbj9 svc.Subject
			var sbj10 svc.Subject
			var sbj11 svc.Subject
			var sbj12 svc.Subject
			var sbj13 svc.Subject
			//开始查mysql
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 1, time.Now().Weekday()+1).First(&sbj1)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 2, time.Now().Weekday()+1).First(&sbj2)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 3, time.Now().Weekday()+1).First(&sbj3)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 4, time.Now().Weekday()+1).First(&sbj4)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 5, time.Now().Weekday()+1).First(&sbj5)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 6, time.Now().Weekday()+1).First(&sbj6)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 7, time.Now().Weekday()+1).First(&sbj7)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 8, time.Now().Weekday()+1).First(&sbj8)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 9, time.Now().Weekday()+1).First(&sbj9)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 10, time.Now().Weekday()+1).First(&sbj10)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 11, time.Now().Weekday()+1).First(&sbj11)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 12, time.Now().Weekday()+1).First(&sbj12)
			l.svcCtx.Mysql.Where("class = ? AND th = ? AND weekly = ?", req.Class, 13, time.Now().Weekday()+1).First(&sbj13)
			return &types.Kebiaores{
				Weekly:    "Tomorrow",
				Lession1:  sbj1.Sbjname,
				Src1:      sbj1.Src,
				Lession2:  sbj2.Sbjname,
				Src2:      sbj2.Src,
				Lession3:  sbj3.Sbjname,
				Src3:      sbj3.Src,
				Lession4:  sbj4.Sbjname,
				Src4:      sbj4.Src,
				Lession5:  sbj5.Sbjname,
				Src5:      sbj5.Src,
				Lession6:  sbj6.Sbjname,
				Src6:      sbj6.Src,
				Lession7:  sbj7.Sbjname,
				Src7:      sbj7.Src,
				Lession8:  sbj8.Sbjname,
				Src8:      sbj8.Src,
				Lession9:  sbj9.Sbjname,
				Src9:      sbj9.Src,
				Lession10: sbj10.Sbjname,
				Src10:     sbj10.Src,
				Lession11: sbj11.Sbjname,
				Src11:     sbj11.Src,
				Lession12: sbj12.Sbjname,
				Src12:     sbj12.Src,
				Lession13: sbj13.Sbjname,
				Src13:     sbj13.Src,
			}, nil
		}

	} else {

		return &types.Kebiaores{}, nil
	}
}
