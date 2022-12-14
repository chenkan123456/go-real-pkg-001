package e

const (
	DB_REGION             = "DB_REGION"
	DB_SOURCE             = "DB_SOURCE"
	DB_REGION_TATAL_PRICE = "DB_REGION_TATAL_PRICE"
	DB_REGION_PRICE       = "DB_REGION_PRICE"
	DB_BUILD              = "DB_BUILD"
	DB_CITY               = "DB_CITY"
	DB_RING               = "DB_RING"
	DB_FIVE_LOCATION      = "DB_FIVE_LOCATION"
	DB_NINE_LOCATION      = "DB_NINE_LOCATION"
	DB_BANK               = "DB_BANK"
	DB_AREA               = "DB_AREA"
	DB_PROPERTY           = "DB_PROPERTY"
	DB_TOTAL_PRICE        = "DB_REGION_TATAL_PRICE"
	DB_UNIT_PRICE         = "DB_REGION_PRICE"
	DB_DEAL_TYPE          = "DB_DEAL_TYPE"

	DB_BUILD_KEY            = "695ec522-bc71-11ec-8ae3-0050568cfd0b"
	DB_AREA_KEY             = "69664d75-bc71-11ec-8ae3-0050568cfd0b"
	DB_TOTAL_PRICE_KEY      = "5bf4bdc4-bc77-11ec-8ae3-0050568cfd0b"
	DB_UNIT_PRICE_KEY       = "5bf50b24-bc77-11ec-8ae3-0050568cfd0b"
	DB_ROOT_KEY             = "74ab3565-bc6d-11ec-8ae3-0050568cfd0b"
	DB_PROPERTY_KEY         = "9f9a5c12-bc6e-11ec-8ae3-0050568cfd0b"
	DB_LAND_BUILD_AREA_KEY  = "f8e737c5-33f9-11ed-96b3-0050568cfd0b"
	DB_LAND_FLOOR_PRICE_KEY = "27931297-3412-11ed-96b3-0050568cfd0b"

	DB_SUMRANGE_SALE_COUNT_KEY      = "fcc6dac1-2905-11ed-9627-0050568cfd0b"
	DB_SUMRANGE_SALE_AREA_KEY       = "9d56d66d-2908-11ed-9627-0050568cfd0b"
	DB_SUMRANGE_SALE_AVGPRICE_KEY   = "539b61b2-2909-11ed-9627-0050568cfd0b"
	DB_SUMRANGE_SALE_TOTALPRICE_KEY = "487154d5-290a-11ed-9627-0050568cfd0b"
	DB_SUMRANGE_ADD_COUNT_KEY       = "d9bfbdc5-290a-11ed-9627-0050568cfd0b"
	DB_SUMRANGE_ADD_AREA_KEY        = "48a21569-290b-11ed-9627-0050568cfd0b"
	DB_SUMRANGE_STOCK_COUNT_KEY     = "bc72548c-2926-11ed-9627-0050568cfd0b"
	DB_SUMRANGE_STOCK_AREA_KEY      = "cc1cf8d9-2928-11ed-9627-0050568cfd0b"
	DB_SUMRANGE_BACK_COUNT_KEY      = "f1e1fc5b-292a-11ed-9627-0050568cfd0b"
	DB_SUMRANGE_BACK_AREA_KEY       = "6fdb3051-292b-11ed-9627-0050568cfd0b"

	DB_REGION_KEY     = "22ac8c56-fe5f-40b8-8165-33339a088b91"
	DB_BUILD_WORK_KEY = "5648823f-0ca5-11ed-8d90-0050568cfd0b"

	DB_GROUP_BUILD_KEY       = "BuildingType"
	DB_GROUP_AREA_KEY        = "AreaGroup"
	DB_GROUP_UNIT_PRICE_KEY  = "UnitPriceType"
	DB_GROUP_TOTAL_PRICE_KEY = "TotalPriceType"
	DB_SOURCE_RG_KEY         = "_a" //??????
	DB_SOURCE_BA_KEY         = "_b" //??????
	STATE_IS_SELL            = "??????"
	STATE_SELLING            = "??????"
	STATE_SELL               = "??????"

	YES_TEXT   = "???"
	NOT_TEXT   = "???"
	EMPTY_TEXT = ""

	UNIT1 = "???"
	UNIT2 = "???"
	UNIT3 = "???/???"
	UNIT4 = "%"

	TITLE_TYPE_SALE_AREA      = "????????????"
	TITLE_TYPE_ADD_AREA       = "????????????"
	TITLE_TYPE_SALE_NUMBER    = "????????????"
	TITLE_TYPE_ADD_NUMBER     = "????????????"
	TITLE_TYPE_SALE_AVG_PRICE = "????????????"

	COUNT_SALE_KEY  = "sale"
	COUNT_ADD_KEY   = "add"
	COUNT_STOCK_KEY = "stock"
	COUNT_UN_KEY    = "unsale"
	COUNT_BACK_KEY  = "back"

	TIME_WEEK_KEY  = "Week"
	TIME_MONTH_KEY = "Month"

	FILTER_BUILD = 0
	FILTER_AREA  = 1

	FILTER_TIME_YEAR      = 1
	FILTER_TIME_HALF_YEAR = 2
	FILTER_TIME_QUARTER   = 3
	FILTER_TIME_MONTH     = 4
	FILTER_TIME_WEEK      = 5

	RADIO_NULL_KEY = 0
	RADIO_YES_KEY  = 1
	RADIO_NO_KEY   = 2

	DB_CODE_RING          = "DB_RING"
	DB_CODE_FIVE_LOCATION = "DB_FIVE_LOCATION"
	DB_CODE_NINE_LOCATION = "DB_NINE_LOCATION"
	DB_CODE_BANK          = "DB_BANK"
	DB_CODE_REGION        = "DB_REGION"

	DB_CODE_SUMRANGE_SALE_COUNT      = "DB_SUMRANGE_SALE_COUNT"
	DB_CODE_SUMRANGE_SALE_AREA       = "DB_SUMRANGE_SALE_AREA"
	DB_CODE_SUMRANGE_SALE_AVGPRICE   = "DB_SUMRANGE_SALE_AVGPRICE"
	DB_CODE_SUMRANGE_SALE_TOTALPRICE = "DB_SUMRANGE_SALE_TOTALPRICE"
	DB_CODE_SUMRANGE_ADD_TOTAL       = "DB_SUMRANGE_ADD_TOTAL"
	DB_CODE_SUMRANGE_ADD_AREA        = "DB_SUMRANGE_ADD_AREA"
	DB_CODE_SUMRANGE_STOCK_COUNT     = "DB_SUMRANGE_STOCK_COUNT"
	DB_CODE_SUMRANGE_STOCK_AREA      = "DB_SUMRANGE_STOCK_AREA"
	DB_CODE_SUMRANGE_BACK_COUNT      = "DB_SUMRANGE_BACK_COUNT"
	DB_CODE_SUMRANGE_BACK_AREA       = "DB_SUMRANGE_BACK_AREA"
	DB_CODE_SUMRANGE_ALL             = "DB_SUMRANGE_ALL"

	TABLE_LAND_TRAN       = "rap_land_transaction"
	TABLE_LAND_BASE       = "rap_land_transfer_a"
	TABLE_LAND_EXTEND     = "rap_land_transfer_b"
	TABLE_LAND_DEVLOPER   = "rap_land_estate"
	TABLE_LAND_MAP_ESTATE = "rap_estate_land"

	TABLE_SALE   = "rap_sale"
	TABLE_STOCK  = "rap_surplus"
	TABLE_UNSALE = "rap_unsale"
	TABLE_ADD    = "rap_newsupply"
	TABLE_BACK   = "rap_back"

	TABLE_DW_SALE   = "rap_dw_sale"
	TABLE_DW_STOCK  = "rap_dw_surplus"
	TABLE_DW_UNSALE = "rap_dw_unsale"
	TABLE_DW_ADD    = "rap_dw_newsupply"
	TABLE_DW_BACK   = "rap_dw_back"

	TABLE_DW_SALE_OFFICE   = "rap_dw_sale_office"
	TABLE_DW_STOCK_OFFICE  = "rap_dw_surplus_office"
	TABLE_DW_UNSALE_OFFICE = "rap_dw_unsale_office"
	TABLE_DW_ADD_OFFICE    = "rap_dw_newsupply_office"
	TABLE_DW_BACK_OFFICE   = "rap_dw_back_office"

	TABLE_PRESALE       = "rap_presale"
	TABLE_ESTATE_BASE   = "rap_estate"
	TABLE_ESTATE_BRANCH = "rap_estate_branch"
	TABLE_BUILD         = "rap_build"

	CENTA_ESTAE_OPEN_FILE_ROOT = "cent_estate_open_files/"
	CENTA_ESTAE_PIC_ROOT       = "centa_estate_pics/"
	CENTA_DIAGRAM_PIC_ROOT     = "cent_diagram/"
	CENTA_LAND_PIC_ROOT        = "rap_land_pics/"
	CENTA_LAND_FILE_ROOT       = "rap_land_attachment/"
	PIC_BIG_SIZE               = "_big"
	PIC_MID_SIZE               = "_mid"
	PIC_ICO_SIZE               = "_cover"

	LOG_ESTATE_OPEN_REPORT_KEY = "ESTATE_OPEN_REPORT"
	LOG_ESTATE_MAP_KEY         = "ESTATE_MAP"

	SHAPE_POLYGON  = 0
	SHAPE_CIRCULAR = 1

	POINT_TYPE_DEFAULT = 0
	POINT_TYPE_CENTER  = 1
)
