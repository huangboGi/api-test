package testdata

import (
	"fmt"
	"time"
)

// NewUserOrder 创建用户订单测试数据
func NewUserOrder(userID, productID uint, coinKey, investAmount string) map[string]interface{} {
	return map[string]interface{}{
		"orderNo":      fmt.Sprintf("ORD_%d", time.Now().UnixNano()),
		"userId":       userID,
		"productId":    productID,
		"coinKey":      coinKey,
		"investAmount": investAmount,
		"orderStatus":  0,
	}
}

// NewUserOrderWithStatus 创建指定状态的订单
func NewUserOrderWithStatus(orderNo string, userID, productID uint, coinKey, investAmount string, status int8) map[string]interface{} {
	return map[string]interface{}{
		"orderNo":      orderNo,
		"userId":       userID,
		"productId":    productID,
		"coinKey":      coinKey,
		"investAmount": investAmount,
		"orderStatus":  status,
	}
}
