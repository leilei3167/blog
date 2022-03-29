//tag的数据库交互
package model

type Tag struct {
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//获取所有的Tags
func GetTags(name string) (tags []Tag) {
	db.Where("name=?", name).Find(&Tag{})

	return

}
