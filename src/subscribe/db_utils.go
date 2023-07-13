package subscribe

import (
	"github.com/pkg/errors"
	"github.com/setcy/spider/src/rule"
	"gorm.io/gorm"
	"strings"
)

func SaveSubscribeInDatabase(db *gorm.DB, StaffID string, Rule *rule.RuleItem) error {
	data := make([]Subscribe, 0)
	tx := db.Begin()
	tx.Model(&Subscribe{}).Where("staff_id = ? AND tag LIKE ? ", StaffID, Rule.Tag+"%").Find(&data)
	if len(data) == 0 {
		tx.Create(&Subscribe{StaffID: StaffID, Tag: Rule.Tag, Name: Rule.Site + Rule.Section})
		tx.Commit()
		return nil
	}
	if len(data) > 1 || len(data) == 1 && strings.Contains(data[0].Tag, Rule.Tag) {
		err := DeleteSubscribeInDatabase(db, StaffID, Rule)
		tx.Create(&Subscribe{StaffID: StaffID, Tag: Rule.Tag, Name: Rule.Site + "-" + Rule.Section})
		if err != nil {
			tx.Callback()
			return err
		}
		tx.Commit()
	}
	if len(data) == 1 && data[0].Tag == Rule.Tag {
		tx.Callback()
		return errors.New("already subscribed")
	}
	tx.Callback()
	return nil
}

func DeleteSubscribeInDatabase(db *gorm.DB, StaffID string, Rule *rule.RuleItem) error {
	data := make([]Subscribe, 0)
	tx := db.Begin()
	tx.Model(&Subscribe{}).Where(" staff_id = ? AND tag LIKE ? ", StaffID, Rule.Tag+"%").Find(&data)
	if len(data) != 0 {
		tx.Delete(data)
		tx.Commit()
		return nil
	}
	tx.Callback()
	return errors.New("no this subscribed")
}

func GetSubscribesByStaffID(db *gorm.DB, StaffID string) []*Subscribe {
	data := make([]*Subscribe, 0)
	db.Where(&Subscribe{StaffID: StaffID}).Find(&data)
	return data
}

func GetSubscribesByRule(db *gorm.DB, Rule *rule.RuleItem) []*Subscribe {
	data := make([]*Subscribe, 0)
	db.Where("TAG = ?", Rule.Tag).Find(&data)
	return data
}
