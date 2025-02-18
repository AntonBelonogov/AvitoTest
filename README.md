# AvitoTest
### Тестовое задание для стажёра Backend-направления (зимняя волна 2025)

---

## Запуск проекта
**Через Docker:**
1. Убедитесь, что у вас установлен Docker и docker-compose.
2. Для запуска проекта должен быть `.env` файл. Для запуска используется данный файл:
```
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASSWORD=password
DATABASE_NAME=shop
DATABASE_HOST=db

SERVER_PORT=8080

SECRET_KEY=SECRETSECRETKEY
```
3. В корне проекта выполните:
```sh
docker-compose up --build
```

4. Сервис поднимется на http://localhost:8080.