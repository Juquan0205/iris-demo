package errcode

import (
	"maizuo.com/back-end/iris-demo/src/server/entity"
	"errors"
)
var(
	SQL_EMPTY_DATA=errors.New("数据库不存在数据")
)

//返回结果错误码定义
var (
	SUCCESS = &ErrCode{0, "ok"}
	REQUEST_ILLEGAL_ERROR = &ErrCode{-2, "非法请求"}
	SYSTEM_ERROR = &ErrCode{-1, "系统异常"}
	PARAM_ERROR = &ErrCode{1001, "参数错误"}
	DATA_ERROR = &ErrCode{1016, "业务调用失败"}
	RATE_ERROR = &ErrCode{1014, "请求过于频繁"}
	PARAM_PARSE_ERROR = &ErrCode{1002, "参数解析错误"}
	PRODUCT_UNDERSTOCK = &ErrCode{800003, "商品库存不足"}
	PRODUCT_INFO_CHANGE = &ErrCode{800006, "商品信息发生变化"}
	PRODUCT_NOT_FIND = &ErrCode{800007, "商品信息不存在"}
	ORDER_VERIFY_FAILURE = &ErrCode{200006, "价格校验失败"}
	ORDER_SAVE_FAILURE = &ErrCode{200012, "订单保存失败"}
	SCORE_UNDERSTOCK = &ErrCode{200010, "积分不足"}
	EPAY_CARD_FAILURE = &ErrCode{200011, "卖座卡抵扣失败"}
	EPAY_COUPON_FAILURE = &ErrCode{200012, "现金券抵扣失败"}
	EPAY_BALANCE_FAILURE = &ErrCode{200013, "余额抵扣失败"}
	EPAY_SCRORE_FAILURE = &ErrCode{200014, "积分抵扣失败"}
	ORDER_CHNAGESTATUS_FAILURE = &ErrCode{200015, "支付成功,标记订单状态失败,需要人工处理"}
	ORDER_NOT_FIND = &ErrCode{200016, "订单不存在"}
	ORDER_STATUS_ERR = &ErrCode{200017, "当前操作的订单状态错误"}
	ORDER_INVENTORY_REDUCE_FAILURE = &ErrCode{200018, "库存扣减失败"}
	ORDER_INVENTORY_ROLLBACK_FAILURE = &ErrCode{200019, "库存回滚失败"}
	ORDER_PASSWORD_ERR = &ErrCode{200020, "安全密码错误"}
	PAY_CHECK_STATUS_ERR = &ErrCode{200021, "查询支付状态错误"}
	PAY_INFO_CHANGE_ERR = &ErrCode{200021, "已确认支付方式后,不允许更换哦"}




)

type ErrCode struct {
	Code int
	Msg  string
}

func (e *ErrCode)Result() *entity.Result {
	return &entity.Result{
		Status:e.Code,
		Msg:e.Msg,
		Data:"",
	}
}

func (e *ErrCode)ResultWithData(data interface{}) *entity.Result {
	return &entity.Result{
		Status:e.Code,
		Msg:e.Msg,
		Data:data,
	}
}

func (e *ErrCode)ResultWithMsg(msg string) *entity.Result {
	return &entity.Result{
		Status:e.Code,
		Msg:msg,
		Data:"",
	}
}

func  (e *ErrCode)ReplaceMsg(msg string) *ErrCode {
	return &ErrCode{
		Code:e.Code,
		Msg:msg,
	}
}



