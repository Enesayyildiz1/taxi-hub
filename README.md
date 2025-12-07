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
  "firstName": "Mehmet",
  "lastName": "Kaya",
  "plate": "34XYZ789",
  "taksiType": "turuncu",
  "carBrand": "Honda",
  "carModel": "Civic",
  "lat": 41.0082,
  "lon": 28.9784
}
```
**Response:**
```json
{
  "id": "driver_id",
}
```

#### `GET /api/v1/drivers`
Tüm sürücüleri listeler.

**Response:**
```json
[
  {
    "id": "6934112a0a3d041839246dcf",
    "firstName": "Mehmet",
    "lastName": "Kaya",
    "plate": "34XYZ789",
    "taksiType": "turuncu",
    "carBrand": "Honda",
    "carModel": "Civic",
    "lat": 41.0082,
    "lon": 28.9784,
    "createdAt": "2025-12-06T11:19:06Z"
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

#### `GET /api/v1/drivers/nearby?lat=41.0082&lon=28.9784&taksiType=sari`
Belirtilen konuma yakın sürücüleri listeler.

**Response:**
```json
[
  {
    "id": "6931d559734d982a29d7ef99",
    "firstName": "Efe2",
    "lastName": "Ayyildiz",
    "plate": "34FFJ850",
    "taksiType": "sari",
    "carBrand": "Fiat",
    "carModel": "Egea",
    "distanceKm": 0,
    "lat": 41.0082,
    "lon": 28.9784
  }
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
  "total_requests": 1000,
  "active_users": 50
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