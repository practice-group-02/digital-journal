# 🗃️ Scholar.KZ – Database Structure

## 👩‍💻 Роль: Database & Data Engineer (PostgreSQL)

Этот раздел проекта отвечает за проектирование, реализацию и наполнение базы данных для цифрового журнала международных программ Scholar.KZ. Ниже описана структура таблиц, связи, объяснение логики и шаги подключения к backend.

---

## 📐 1. ERD (Entity-Relationship Diagram)

![ERD Diagram](./erd/ScholarKZ_ERD.png)

### Связи:
- `programs.country_id → countries.id` (многие к одному)
- `programs.language_id → languages.id`
- `programs.university_id → universities.id`
- `program_tags.program_id → programs.id`
- `program_tags.tag_id → tags.id`
- `eligibility.program_id → programs.id`
- `documents.program_id → programs.id`

---

## 🗄 2. Структура таблиц

### `countries`
| Поле | Тип |
|------|-----|
| id | SERIAL |
| name | TEXT |

### `languages`
| Поле | Тип |
|------|-----|
| id | SERIAL |
| name | TEXT |

### `universities`
| Поле | Тип |
|------|-----|
| id | SERIAL |
| name | TEXT |
| website | TEXT |

### `tags`
| Поле | Тип |
|------|-----|
| id | SERIAL |
| name | TEXT |

### `programs`
| Поле | Тип |
|------|-----|
| id | SERIAL |
| name | TEXT |
| description | TEXT |
| deadline | DATE |
| duration | TEXT |
| degree | TEXT |
| funding | TEXT |
| url | TEXT |
| country_id | INT (FK) |
| language_id | INT (FK) |
| university_id | INT (FK) |

### `program_tags`
Связующая таблица (многие ко многим):
- program_id → programs.id
- tag_id → tags.id

### `eligibility`
| Поле | Тип |
|------|-----|
| id | SERIAL |
| program_id | INT |
| citizenship | TEXT |
| age_range | TEXT |
| education_level | TEXT |
| experience_required | TEXT |

### `documents`
| Поле | Тип |
|------|-----|
| id | SERIAL |
| program_id | INT |
| name | TEXT |

---

## ⚙️ 3. SQL-инициализация

📄 Файл `init.sql` находится в `db/migrations/init.sql`  
Он создаёт все таблицы и связи между ними.

---

## 🔌 4. Подключение к Go Backend

- Используется библиотека GORM
- Подключение через `.env`:
  ```env
  DB_USER=postgres
  DB_PASSWORD=пароль
  DB_NAME=scholarkz
  DB_HOST=localhost
  DB_PORT=5432
