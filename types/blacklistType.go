package types

//黑名单表名 黑名单原始记录表
const (
	BEIJINGCARD    = "b_hmd_card_11" //北京
	TIANJINCARD    = "b_hmd_card_12" //天津
	HEBEICARD      = "b_hmd_card_13" //河北
	SHANXICARD     = "b_hmd_card_14" //山西
	NEIMENGFGUCARD = "b_hmd_card_15" //内蒙古

	LIAONINGCARD     = "b_hmd_card_21" //辽宁
	JILINCARD        = "b_hmd_card_22" //吉林
	HEILONGJIANGCARD = "b_hmd_card_23" //黑龙江

	SHANGHAICARD = "b_hmd_card_31" //上海
	JIANGSUCARD  = "b_hmd_card_32" //江苏
	ZHEJIANGCARD = "b_hmd_card_33" //浙江
	ANHUICARD    = "b_hmd_card_34" //安徽
	FUJIANCARD   = "b_hmd_card_35" //福建
	JIANGXICARD  = "b_hmd_card_36" //江西
	SHANDONGCARD = "b_hmd_card_37" //山东

	HENANCARD     = "b_hmd_card_41" //河南
	HUBEICARD     = "b_hmd_card_42" //湖北
	HUNANCARD     = "b_hmd_card_43" //湖南
	GUANGDONGCARD = "b_hmd_card_44" //广东
	GUANGXICARD   = "b_hmd_card_45" //广西
	HAINANCARD    = "b_hmd_card_46" //海南

	CHONGQINGCARD = "b_hmd_card_50" //重庆
	SICHUANCARD   = "b_hmd_card_51" //四川
	GUIZHOUCARD   = "b_hmd_card_52" //贵州
	YUNNANCARD    = "b_hmd_card_53" //云南
	XIZANGCARD    = "b_hmd_card_54" //西藏

	SHAANXICARD  = "b_hmd_card_61"  //陕西
	GANSUCARD    = "b_hmd_card_62"  //甘肃
	QINGHAICARD  = "b_hmd_card_63"  //青海
	NINGXIACARD  = "b_hmd_card_64"  //宁夏
	XINJIANGCARD = "b_hmd_card_65"  //新疆
	JUNKACARD    = "b_hmd_card_501" //军车卡

	BEIJINGOBU    = "b_hmd_obu_11" //北京
	TIANJINOBU    = "b_hmd_obu_12" //天津
	HEBEIOBU      = "b_hmd_obu_13" //河北
	SHANXIOBU     = "b_hmd_obu_14" //山西
	NEIMENGFGUOBU = "b_hmd_obu_15" //内蒙古

	LIAONINGOBU     = "b_hmd_obu_21" //辽宁
	JILINOBU        = "b_hmd_obu_22" //吉林
	HEILONGJIANGOBU = "b_hmd_obu_23" //黑龙江

	SHANGHAIOBU = "b_hmd_obu_31" //上海
	JIANGSUOBU  = "b_hmd_obu_32" //江苏
	ZHEJIANGOBU = "b_hmd_obu_33" //浙江
	ANHUIOBU    = "b_hmd_obu_34" //安徽
	FUJIANOBU   = "b_hmd_obu_35" //福建
	JIANGXIOBU  = "b_hmd_obu_36" //江西
	SHANDONGOBU = "b_hmd_obu_37" //山东

	HENANOBU     = "b_hmd_obu_41" //河南
	HUBEIOBU     = "b_hmd_obu_42" //湖北
	HUNANOBU     = "b_hmd_obu_43" //湖南
	GUANGDONGOBU = "b_hmd_obu_44" //广东
	GUANGXIOBU   = "b_hmd_obu_45" //广西
	HAINANOBU    = "b_hmd_obu_46" //海南

	CHONGQINGOBU = "b_hmd_obu_50" //重庆
	SICHUANOBU   = "b_hmd_obu_51" //四川
	GUIZHOUOBU   = "b_hmd_obu_52" //贵州
	YUNNANOBU    = "b_hmd_obu_53" //云南
	XIZANGOBU    = "b_hmd_obu_54" //西藏

	SHAANXIOBU  = "b_hmd_obu_61"  //陕西
	GANSUOBU    = "b_hmd_obu_62"  //甘肃
	QINGHAIOBU  = "b_hmd_obu_63"  //青海
	NINGXIAOBU  = "b_hmd_obu_64"  //宁夏
	XINJIANGOBU = "b_hmd_obu_65"  //新疆
	JUNKAOBU    = "b_hmd_obu_501" //军车卡

)

//黑名单表名
var Blacklist = []string{
	BEIJINGCARD, TIANJINCARD, HEBEICARD, SHANXICARD, NEIMENGFGUCARD,
	LIAONINGCARD, JILINCARD, HEILONGJIANGCARD,
	SHANGHAICARD, JIANGSUCARD, ZHEJIANGCARD, ANHUICARD, FUJIANCARD, JIANGXICARD, SHANDONGCARD,
	HENANCARD, HUBEICARD, HUNANCARD, GUANGDONGCARD, GUANGXICARD, HAINANCARD,
	CHONGQINGCARD, SICHUANCARD, GUIZHOUCARD, YUNNANCARD, XIZANGCARD,
	SHAANXICARD, GANSUCARD, QINGHAICARD, NINGXIACARD, XINJIANGCARD, JUNKACARD,
	BEIJINGOBU, TIANJINOBU, HEBEIOBU, SHANXIOBU, NEIMENGFGUOBU,
	LIAONINGOBU, JILINOBU, HEILONGJIANGOBU,
	SHANGHAIOBU, JIANGSUOBU, ZHEJIANGOBU, ANHUIOBU, FUJIANOBU, JIANGXIOBU, SHANDONGOBU,
	HENANOBU, HUBEIOBU, HUNANOBU, GUANGDONGOBU, GUANGXIOBU, HAINANOBU,
	CHONGQINGOBU, SICHUANOBU, GUIZHOUOBU, YUNNANOBU, XIZANGOBU,
	SHAANXIOBU, GANSUOBU, QINGHAIOBU, NINGXIAOBU, XINJIANGOBU, JUNKAOBU,
}
