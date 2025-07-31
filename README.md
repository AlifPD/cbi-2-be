# Technical Test CBI

Repo ini berisi Source code untuk Technical Test CBI Nomor 2 Part BE.

## What App is this?
App ini merupakan Simple To Do list.
|Fungsi                |Deskripsi                          |
|----------------|-------------------------------|
|Login|Login dengan username dan password            |
|Register|Register user baru|
|Create item baru|Menambah item baru pada To do list|
|Read to do list|View semua To do list item yang dimiliki seorang user|
|Update item|Mengubah item untuk menandakan item done|
|Delete item|Menghapus item pada To do list|

Dokumentasi API dapat dilihat pada link dibawah ini:
https://documenter.getpostman.com/view/20192847/2sB3B8tE23

## Prerequisite
- Go

## How to Run

1. Add ``` .env ``` file in the project's root folder
```
# Add the JWT_SECRET and ADMIN Credentials
JWT_SECRET=
ADMIN_USERNAME=
ADMIN_PASSWORD=
```
2. Run ``` go run main.go ```