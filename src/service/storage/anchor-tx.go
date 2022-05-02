package storage

func (d *DataBase) GetLastTxId() int64 {
	anchorTx := AnchorTx{}
	d.db.Model(AnchorTx{}).Order("id desc").First(&anchorTx)
	return anchorTx.ID
}

func (d *DataBase) GetDepositedAmount(tokenAmount string) string {
	anchorTx := AnchorTx{}
	d.db.Model(AnchorTx{}).Where("aust_amount = ? and type = ?", tokenAmount, "deposit_stable").Order("id desc").First(&anchorTx)
	return anchorTx.USTAmount
}

func (d *DataBase) GetRedeemedAmount(tokenAmount string) string {
	anchorTx := AnchorTx{}
	d.db.Model(AnchorTx{}).Where("ust_amount = ? and type = ?", tokenAmount, "redeem_stable").Order("id desc").First(&anchorTx)
	return anchorTx.AUSTAmount
}

func (d *DataBase) SaveAnchorTx(anchorTx *AnchorTx) error {
	if err := d.db.Model(AnchorTx{}).Create(&anchorTx).Error; err != nil {
		return err
	}
	return nil
}
