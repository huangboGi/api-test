# API 接口 Curl 命令手册

> 直接复制命令到终端执行即可

---

## 用户端接口

### 1. 产品列表
```bash
curl --location 'http://localhost:8080/api/v1/wth/product/page' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiSDN1Vko5TGtGRWhwOXpOUk1GcURIbkJXQzY1UmhZemlnSG5rM3JqeWc2UTciLCJjaGFpbiI6NTAxLCJleHAiOjE3NzAyNzQ5MDAsImlhdCI6MTc3MDE4ODUwMCwib3JpZ19pYXQiOjE3NzAxODg1MDAsInVzZXJJZCI6MTcyOTUxfQ.wT0a8tFdc2PUmy-qtIGairx9dP45VjxYM4AU3cUpOns' \
--header 'Content-Type: application/json'
```

### 2. 订单列表
```bash
curl --location 'http://localhost:8080/api/v1/wth/order/list' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiSDN1Vko5TGtGRWhwOXpOUk1GcURIbkJXQzY1UmhZemlnSG5rM3JqeWc2UTciLCJjaGFpbiI6NTAxLCJleHAiOjE3NzAyNzQ5MDAsImlhdCI6MTc3MDE4ODUwMCwib3JpZ19pYXQiOjE3NzAxODg1MDAsInVzZXJJZCI6MTcyOTUxfQ.wT0a8tFdc2PUmy-qtIGairx9dP45VjxYM4AU3cUpOns' \
--header 'Content-Type: application/json' \
--data '{
    "page": 1,
    "size": 10
}'
```

### 3. 订单详情
```bash
curl --location 'http://localhost:8080/api/v1/wth/order/detail' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiSDN1Vko5TGtGRWhwOXpOUk1GcURIbkJXQzY1UmhZemlnSG5rM3JqeWc2UTciLCJjaGFpbiI6NTAxLCJleHAiOjE3NzAyNzQ5MDAsImlhdCI6MTc3MDE4ODUwMCwib3JpZ19pYXQiOjE3NzAxODg1MDAsInVzZXJJZCI6MTcyOTUxfQ.wT0a8tFdc2PUmy-qtIGairx9dP45VjxYM4AU3cUpOns' \
--header 'Content-Type: application/json' \
--data '{
    "orderNo": "ORDER123456"
}'
```

### 4. 申购
```bash
curl --location 'http://localhost:8080/api/v1/wth/order/subscribe' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiSDN1Vko5TGtGRWhwOXpOUk1GcURIbkJXQzY1UmhZemlnSG5rM3JqeWc2UTciLCJjaGFpbiI6NTAxLCJleHAiOjE3NzAyNzQ5MDAsImlhdCI6MTc3MDE4ODUwMCwib3JpZ19pYXQiOjE3NzAxODg1MDAsInVzZXJJZCI6MTcyOTUxfQ.wT0a8tFdc2PUmy-qtIGairx9dP45VjxYM4AU3cUpOns' \
--header 'Content-Type: application/json' \
--data '{
    "productId": 1,
    "volume": "1000"
}'
```

### 5. 赎回
```bash
curl --location 'http://localhost:8080/api/v1/wth/order/redeem' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiSDN1Vko5TGtGRWhwOXpOUk1GcURIbkJXQzY1UmhZemlnSG5rM3JqeWc2UTciLCJjaGFpbiI6NTAxLCJleHAiOjE3NzAyNzQ5MDAsImlhdCI6MTc3MDE4ODUwMCwib3JpZ19pYXQiOjE3NzAxODg1MDAsInVzZXJJZCI6MTcyOTUxfQ.wT0a8tFdc2PUmy-qtIGairx9dP45VjxYM4AU3cUpOns' \
--header 'Content-Type: application/json' \
--data '{
    "orderNo": "ORDER123456",
    "volume": "500"
}'
```

### 6. 持仓查询
```bash
curl --location 'http://localhost:8080/api/v1/wth/order/holdPosition' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiSDN1Vko5TGtGRWhwOXpOUk1GcURIbkJXQzY1UmhZemlnSG5rM3JqeWc2UTciLCJjaGFpbiI6NTAxLCJleHAiOjE3NzAyNzQ5MDAsImlhdCI6MTc3MDE4ODUwMCwib3JpZ19pYXQiOjE3NzAxODg1MDAsInVzZXJJZCI6MTcyOTUxfQ.wT0a8tFdc2PUmy-qtIGairx9dP45VjxYM4AU3cUpOns' \
--header 'Content-Type: application/json'
```

### 7. 收益列表
```bash
curl --location 'http://localhost:8080/api/v1/wth/order/interest/page' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiSDN1Vko5TGtGRWhwOXpOUk1GcURIbkJXQzY1UmhZemlnSG5rM3JqeWc2UTciLCJjaGFpbiI6NTAxLCJleHAiOjE3NzAyNzQ5MDAsImlhdCI6MTc3MDE4ODUwMCwib3JpZ19pYXQiOjE3NzAxODg1MDAsInVzZXJJZCI6MTcyOTUxfQ.wT0a8tFdc2PUmy-qtIGairx9dP45VjxYM4AU3cUpOns' \
--header 'Content-Type: application/json' \
--data '{
    "page": 1,
    "size": 10
}'
```

### 8. 历史订单
```bash
curl --location 'http://localhost:8080/api/v1/wth/order/his' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiSDN1Vko5TGtGRWhwOXpOUk1GcURIbkJXQzY1UmhZemlnSG5rM3JqeWc2UTciLCJjaGFpbiI6NTAxLCJleHAiOjE3NzAyNzQ5MDAsImlhdCI6MTc3MDE4ODUwMCwib3JpZ19pYXQiOjE3NzAxODg1MDAsInVzZXJJZCI6MTcyOTUxfQ.wT0a8tFdc2PUmy-qtIGairx9dP45VjxYM4AU3cUpOns' \
--header 'Content-Type: application/json' \
--data '{
    "page": 1,
    "size": 10
}'
```

### 9. 定期详情
```bash
curl --location 'http://localhost:8080/api/v1/wth/order/period/detail' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiSDN1Vko5TGtGRWhwOXpOUk1GcURIbkJXQzY1UmhZemlnSG5rM3JqeWc2UTciLCJjaGFpbiI6NTAxLCJleHAiOjE3NzAyNzQ5MDAsImlhdCI6MTc3MDE4ODUwMCwib3JpZ19pYXQiOjE3NzAxODg1MDAsInVzZXJJZCI6MTcyOTUxfQ.wT0a8tFdc2PUmy-qtIGairx9dP45VjxYM4AU3cUpOns' \
--header 'Content-Type: application/json' \
--data '{
    "productId": 1
}'
```

### 10. 开放申购
```bash
curl --location 'http://localhost:8080/api/v1/wth/order/openSub' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiSDN1Vko5TGtGRWhwOXpOUk1GcURIbkJXQzY1UmhZemlnSG5rM3JqeWc2UTciLCJjaGFpbiI6NTAxLCJleHAiOjE3NzAyNzQ5MDAsImlhdCI6MTc3MDE4ODUwMCwib3JpZ19pYXQiOjE3NzAxODg1MDAsInVzZXJJZCI6MTcyOTUxfQ.wT0a8tFdc2PUmy-qtIGairx9dP45VjxYM4AU3cUpOns' \
--header 'Content-Type: application/json'
```

---

## 管理端接口

### 1. 添加币种
```bash
curl --location 'http://localhost:8080/api/v1/admin/wth/coin/add' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ5MjQyMDE1NTYsImlhdCI6MTc3MDYwMTU1NiwibG9naW5UeXBlIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTc3MDYwMTU1NiwicGFzc3dvcmRDaGFuZ2VBdCI6MTc2NjU1OTE2MywicHdkIjoiMVEydzNlNHIiLCJyb2xla2V5IjoiYWRtaW4iLCJzZXJ2ZXJTdGFydFRpbWUiOjE3NzA2MDEzNzYsInVzZXJJZCI6MSwidXNlck5hbWUiOiJhZG1pbiJ9.BATEs3Qtjj1h3KEZ_tZNZlq3JUduoGcdY6w3YDrf44g' \
--header 'Content-Type: application/json' \
--data '{
    "coin": "USDT",
    "name": "Tether USD",
    "icon": "https://example.com/usdt.png"
}'
```

### 2. 币种上下架
```bash
curl --location 'http://localhost:8080/api/v1/admin/wth/coin/updateShelves' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ5MjQyMDE1NTYsImlhdCI6MTc3MDYwMTU1NiwibG9naW5UeXBlIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTc3MDYwMTU1NiwicGFzc3dvcmRDaGFuZ2VBdCI6MTc2NjU1OTE2MywicHdkIjoiMVEydzNlNHIiLCJyb2xla2V5IjoiYWRtaW4iLCJzZXJ2ZXJTdGFydFRpbWUiOjE3NzA2MDEzNzYsInVzZXJJZCI6MSwidXNlck5hbWUiOiJhZG1pbiJ9.BATEs3Qtjj1h3KEZ_tZNZlq3JUduoGcdY6w3YDrf44g' \
--header 'Content-Type: application/json' \
--data '{
    "id": 1,
    "shelvesStatus": 1
}'
```

### 3. 添加规格
```bash
curl --location 'http://localhost:8080/api/v1/admin/wth/spec/add' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ5MjQyMDE1NTYsImlhdCI6MTc3MDYwMTU1NiwibG9naW5UeXBlIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTc3MDYwMTU1NiwicGFzc3dvcmRDaGFuZ2VBdCI6MTc2NjU1OTE2MywicHdkIjoiMVEydzNlNHIiLCJyb2xla2V5IjoiYWRtaW4iLCJzZXJ2ZXJTdGFydFRpbWUiOjE3NzA2MDEzNzYsInVzZXJJZCI6MSwidXNlck5hbWUiOiJhZG1pbiJ9.BATEs3Qtjj1h3KEZ_tZNZlq3JUduoGcdY6w3YDrf44g' \
--header 'Content-Type: application/json' \
--data '{
    "specValue": -1,
    "deadlineType": 0,
    "interestRate": "0.05"
}'
```

### 4. 规格上下架
```bash
curl --location 'http://localhost:8080/api/v1/admin/wth/spec/updateShelves' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ5MjQyMDE1NTYsImlhdCI6MTc3MDYwMTU1NiwibG9naW5UeXBlIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTc3MDYwMTU1NiwicGFzc3dvcmRDaGFuZ2VBdCI6MTc2NjU1OTE2MywicHdkIjoiMVEydzNlNHIiLCJyb2xla2V5IjoiYWRtaW4iLCJzZXJ2ZXJTdGFydFRpbWUiOjE3NzA2MDEzNzYsInVzZXJJZCI6MSwidXNlck5hbWUiOiJhZG1pbiJ9.BATEs3Qtjj1h3KEZ_tZNZlq3JUduoGcdY6w3YDrf44g' \
--header 'Content-Type: application/json' \
--data '{
    "id": 1,
    "shelvesStatus": 1
}'
```

### 5. 添加产品
```bash
curl --location 'http://localhost:8080/api/v1/admin/wth/product/add' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ5MjQyMDE1NTYsImlhdCI6MTc3MDYwMTU1NiwibG9naW5UeXBlIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTc3MDYwMTU1NiwicGFzc3dvcmRDaGFuZ2VBdCI6MTc2NjU1OTE2MywicHdkIjoiMVEydzNlNHIiLCJyb2xla2V5IjoiYWRtaW4iLCJzZXJ2ZXJTdGFydFRpbWUiOjE3NzA2MDEzNzYsInVzZXJJZCI6MSwidXNlck5hbWUiOiJhZG1pbiJ9.BATEs3Qtjj1h3KEZ_tZNZlq3JUduoGcdY6w3YDrf44g' \
--header 'Content-Type: application/json' \
--data '{
    "coinId": 1,
    "specId": 1,
    "minVol": "100",
    "useQuotaTotal": "10000"
}'
```

### 6. 产品上下架
```bash
curl --location 'http://localhost:8080/api/v1/admin/wth/product/updateShelves' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ5MjQyMDE1NTYsImlhdCI6MTc3MDYwMTU1NiwibG9naW5UeXBlIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTc3MDYwMTU1NiwicGFzc3dvcmRDaGFuZ2VBdCI6MTc2NjU1OTE2MywicHdkIjoiMVEydzNlNHIiLCJyb2xla2V5IjoiYWRtaW4iLCJzZXJ2ZXJTdGFydFRpbWUiOjE3NzA2MDEzNzYsInVzZXJJZCI6MSwidXNlck5hbWUiOiJhZG1pbiJ9.BATEs3Qtjj1h3KEZ_tZNZlq3JUduoGcdY6w3YDrf44g' \
--header 'Content-Type: application/json' \
--data '{
    "id": 1,
    "shelvesStatus": 1
}'
```

### 7. 更新产品
```bash
curl --location 'http://localhost:8080/api/v1/admin/wth/product/update' \
--header 'Accept: application/json' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ5MjQyMDE1NTYsImlhdCI6MTc3MDYwMTU1NiwibG9naW5UeXBlIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTc3MDYwMTU1NiwicGFzc3dvcmRDaGFuZ2VBdCI6MTc2NjU1OTE2MywicHdkIjoiMVEydzNlNHIiLCJyb2xla2V5IjoiYWRtaW4iLCJzZXJ2ZXJTdGFydFRpbWUiOjE3NzA2MDEzNzYsInVzZXJJZCI6MSwidXNlck5hbWUiOiJhZG1pbiJ9.BATEs3Qtjj1h3KEZ_tZNZlq3JUduoGcdY6w3YDrf44g' \
--header 'Content-Type: application/json' \
--data '{
    "id": 1,
    "minVol": "100",
    "useQuotaTotal": "50000"
}'
```

---

> **注意**: Token 可能会过期，请从 `.env` 文件获取最新的 Token
