
# 🏋️‍♂️ Gym Membership API

API untuk mengelola sistem keanggotaan gym, termasuk user, membership, trainer, dan workout session.  
Dibangun menggunakan **Golang (Gin + GORM)** dan **PostgreSQL**, serta siap untuk **deploy di Railway**.

---

## 🚀 Base URL
```
https://gym-management-production-078b.up.railway.app/api
```

---

## 📘 API Endpoints

### 🔑 Auth
| Method | Endpoint | Deskripsi |
|---------|-----------|-----------|
| POST | [/users/register](https://gym-management-production-078b.up.railway.app/api/users/register) | Registrasi user baru |
| POST | [/users/login](https://gym-management-production-078b.up.railway.app/api/users/login) | Login dan mendapatkan JWT token |

---

### 🧾 Membership
| Method | Endpoint | Deskripsi |
|---------|-----------|-----------|
| POST | [/memberships](https://gym-management-production-078b.up.railway.app/api/memberships) | Membuat membership baru |
| GET | [/memberships](https://gym-management-production-078b.up.railway.app/api/memberships) | Mendapatkan semua membership |
| PUT | [/memberships/:id](https://gym-management-production-078b.up.railway.app/api/memberships/1) | Update membership berdasarkan ID |
| DELETE | [/memberships/:id](https://gym-management-production-078b.up.railway.app/api/memberships/1) | Hapus membership berdasarkan ID |

---

### 🏋️ Trainer
| Method | Endpoint | Deskripsi |
|---------|-----------|-----------|
| POST | [/trainers](https://gym-management-production-078b.up.railway.app/api/trainers) | Membuat trainer baru |
| GET | [/trainers](https://gym-management-production-078b.up.railway.app/api/trainers) | Mendapatkan semua trainer |
| PUT | [/trainers/:id](https://gym-management-production-078b.up.railway.app/api/trainers/1) | Update trainer |
| DELETE | [/trainers/:id](https://gym-management-production-078b.up.railway.app/api/trainers/1) | Hapus trainer |

---

### 🕒 Workout Session
| Method | Endpoint | Deskripsi |
|---------|-----------|-----------|
| POST | [/workouts](https://gym-management-production-078b.up.railway.app/api/workouts) | Membuat workout session |
| GET | [/workouts](https://gym-management-production-078b.up.railway.app/api/workouts) | Mendapatkan semua workout session |
| PUT | [/workouts/:id](https://gym-management-production-078b.up.railway.app/api/workouts/1) | Update workout session |
| DELETE | [/workouts/:id](https://gym-management-production-078b.up.railway.app/api/workouts/1) | Hapus workout session |

---

## 🧠 Cara Menjalankan di Lokal

```bash
# 1. Clone repo
git clone https://github.com/username/gym-membership.git
cd gym-membership

# 2. Jalankan
go run main.go
```

---

## ☁️ Deployment

Aplikasi ini siap untuk deploy di **Railway**.  
Railway otomatis menjalankan migrasi database lewat fungsi `InitDB()` saat aplikasi start.
