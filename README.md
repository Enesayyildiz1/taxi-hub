# 🚖 Taxi Hub API Gateway & Driver Microservice

Taxi Hub, merkezi bir API geçidi ve sürücü mikroservisi sunar. JWT tabanlı kimlik doğrulama, API anahtarı desteği, sürücü yönetimi, istatistik toplama ve sistem sağlık kontrolü gibi temel işlevleri içerir.

## 🏗 Teknoloji Stack
- Go (Golang) 1.21
- Gin-Gonic Web Framework
- MongoDB
- Swagger 2.0
- RESTful API
- JWT Authentication
- Docker & Docker Compose
- Logrus
- Rate Limiting

---

## 📁 API Endpoint Listesi ve Açıklamaları

### 🔐 Auth (Kimlik Doğrulama)

#### `POST /auth/login`
Kullanıcı adı ve şifre ile giriş yapar, JWT token döner.

**Request:**
```json
{
  "username": "admin",
  "password": "password"
}
```
**Response:**
```json
{
  "token": "<JWT_TOKEN>"
}
```

---

### 🚗 Drivers (Sürücü İşlemleri)

#### `POST /api/v1/drivers`
Yeni sürücü ekler.

**Request:**
```json
{
  "name": "Ali Veli",
  "location": {
    "lat": 41.015137,
    "lng": 28.979530
  },
  "car": "Renault Clio"
}
```
**Response:**
```json
{
  "id": "driver_id",
  "name": "Ali Veli",
  "location": {...},
  "car": "Renault Clio"
}
```

#### `GET /api/v1/drivers`
Tüm sürücüleri listeler.

**Response:**
```json
[
  {
    "id": "driver_id",
    "name": "Ali Veli",
    "location": {...},
    "car": "Renault Clio"
  },
  ...
]
```

#### `GET /api/v1/drivers/{id}`
Belirli bir sürücüyü getirir.

#### `PUT /api/v1/drivers/{id}`
Sürücü bilgilerini günceller.

#### `DELETE /api/v1/drivers/{id}`
Sürücüyü siler.

#### `GET /api/v1/drivers/nearby?lat={LAT}&lng={LNG}&radius={RADIUS}`
Belirtilen konuma yakın sürücüleri listeler.

**Örnek:**
```
GET /api/v1/drivers/nearby?lat=41.015137&lng=28.979530&radius=5
```
**Response:**
```json
[
  {
    "id": "driver_id",
    "name": "Ali Veli",
    "distance": 1.2
  },
  ...
]
```

---

### 🛠 Admin (Yönetici İstatistikleri)

#### `GET /api/v1/admin/stats`
Sistem istatistiklerini döner. JWT ve API Key gerektirir.

**Headers:**
```
Authorization: Bearer <TOKEN>
x-api-key: YOUR_KEY
```
**Response:**
```json
{
  "totalDrivers": 120,
  "activeDrivers": 80,
  "totalRequests": 5000
}
```

---

### ❤️ Health (Sağlık Kontrolü)

#### `GET /health`
Servisin çalışıp çalışmadığını kontrol eder.

**Response:**
```json
{
  "status": "ok"
}
```

---

## 🔒 Güvenlik

Tüm korumalı endpointler için aşağıdaki header'ları eklemelisiniz:
```
Authorization: Bearer <TOKEN>
x-api-key: YOUR_KEY
```

---

## ▶ Çalıştırma

```bash
go mod tidy
go run main.go
```
veya Docker ile:
```bash
docker-compose up --build
```

---

## 📄 Lisans
Apache 2.0 License

---

## 📚 Swagger Dökümantasyonu

Swagger arayüzü ile API endpointlerini test edebilirsiniz.  
`/docs` veya `/swagger` endpointinden erişebilirsiniz.

---

## 📝 Notlar

- Rate limiting ve hata loglama otomatik olarak uygulanır.
- Tüm endpointler RESTful standartlarına uygundur.
- API Gateway, mikroservisler arasında reverse proxy görevi görür.

Sorularınız için: [github.com/enesayyildiz](https://github.com/enesayyildiz)