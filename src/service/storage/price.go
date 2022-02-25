package storage

func (d *DataBase) SavePriceInformation(priceLogs []*PriceLog) {
	for _, priceLog := range priceLogs {
		previousLog, _ := d.GetPriceLog(priceLog.Name)
		if previousLog.UpdateTime == 0 {
			if e := d.savePrice(priceLog); e != nil {
				continue
			}
		} else if priceLog.UpdateTime > previousLog.UpdateTime {
			if e := d.updatePrice(priceLog); e != nil {
				continue
			}
		}
	}
}

func (d *DataBase) GetPriceLog(name string) (priceLog PriceLog, err error) {
	err = d.db.Model(PriceLog{}).Where("name = ?", name).Order("update_time desc").First(&priceLog).Error
	if err != nil {
		return priceLog, err
	}
	return priceLog, nil
}

func (d *DataBase) updatePrice(priceLog *PriceLog) error {
	if err := d.db.Model(PriceLog{}).Where("name = ?", priceLog.Name).Update(&priceLog).Error; err != nil {
		return err
	}
	return nil
}

func (d *DataBase) savePrice(priceLog *PriceLog) error {
	if err := d.db.Model(PriceLog{}).Create(&priceLog).Error; err != nil {
		return err
	}
	return nil
}
