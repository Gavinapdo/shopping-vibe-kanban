package product

func MockProducts() []Product {
	return []Product{
		{
			ID:          1,
			Name:        "无线鼠标",
			Description: "静音按键，支持蓝牙与2.4G",
			Price:       89.00,
			Stock:       120,
		},
		{
			ID:          2,
			Name:        "机械键盘",
			Description: "87键红轴，支持热插拔",
			Price:       299.00,
			Stock:       60,
		},
		{
			ID:          3,
			Name:        "27英寸显示器",
			Description: "2K分辨率，75Hz刷新率",
			Price:       1199.00,
			Stock:       25,
		},
	}
}
