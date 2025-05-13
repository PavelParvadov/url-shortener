package stat

import (
	"gorm.io/datatypes"
	"time"
	"url/pkg/db"
)

type StatRepository struct {
	Db *db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{Db: db}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repo.Db.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)
	if stat.ID == 0 {
		repo.Db.Create(&Stat{
			LinkId: linkId,
			Date:   currentDate,
			Clicks: 1,
		})
	} else {
		stat.Clicks += 1
		repo.Db.Save(&stat)
	}
}

func (repo *StatRepository) GetStats(by string, from time.Time, to time.Time) []GetStatsResponse {
	var stats []GetStatsResponse
	var selectQuery string
	switch by {
	case FilterByDay:
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks)"
	case FilterByMonth:
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks)"
	}

	repo.Db.Table("stats").Select(selectQuery).Where("date between ? and ?", from, to).Group("period").Order("period").Scan(&stats)
	return stats
}
