package storage

func (d *DataBase) UpsertFarm(farm *Farm) error {
	oldFarm, _ := d.GetFarm(farm.ID)

	if oldFarm.ID == "" {
		return d.saveFarm(farm)
	} else {
		return d.updateFarm(farm)
	}
}

func (d *DataBase) GetFarm(id string) (farm Farm, err error) {
	err = d.db.Model(Farm{}).Where("id = ?", id).First(&farm).Error
	if err != nil {
		return farm, err
	}
	return farm, nil
}

func (d *DataBase) updateFarm(farm *Farm) error {
	if err := d.db.Model(Farm{}).Where("id = ?", farm.ID).Update(&farm).Error; err != nil {
		return err
	}
	return nil
}

func (d *DataBase) saveFarm(farm *Farm) error {
	if err := d.db.Model(Farm{}).Create(&farm).Error; err != nil {
		return err
	}
	return nil
}
