# ✅ To-Do LIST - Gestor de Tareas en Línea de Comandos

Un gestor de tareas simple y poderoso desde la terminal.  
Permite **agregar, listar, completar y eliminar tareas** con persistencia local usando **JSON**.

> Construido en Go, usando `cobra` para el CLI y `color` para una experiencia más amigable en la terminal. 🐹✨

---

## 🚀 Características

- 📌 **Agregar tareas** con descripción y fecha límite.
- 📋 **Listar tareas** pendientes, completadas o todas.
- ✅ **Marcar tareas como completadas**.
- ❌ **Eliminar tareas**.
- 💾 **Persistencia de datos** usando archivos `.json` o `.csv`.
- 🎨 **Colores en la terminal** para resaltar tareas según su estado.

---

## 🛠️ Tecnologías utilizadas

- **Go** (`fmt`, `os`, `time`, `encoding/json`, etc.)
- **[cobra](https://github.com/spf13/cobra)** – para crear la interfaz de línea de comandos.
- **[fatih/color](https://github.com/fatih/color)** – para salida con colores en la terminal.

---

## 📦 Instalación

```bash
git clone https://github.com/lmirasutn/ToDoList
cd toDoList
go build -o todo
```
---

### 🧪 Uso

```bash
./todo [comando]
```

### 📌 Agregar una tarea

```bash
./todo add [tarea] [dias a vencer] "
```

### 📋 Listar tareas

```bash
./todo list
```

### ✅ Completar una tarea

```bash
./todo complete 1   # Marca la tarea con ID 1 como completada
```

### ❌ Eliminar una tarea

```bash
./todo delete 1              # Elimina la tarea con ID 1
./todo delete --all          # Elimina todas las tareas
./todo delete --completed    # Elimina solo las tareas completadas

```

---

## 🗂️ Estructura del archivo JSON

```json
[
  {
    "id": 1,
    "descripcion": "Aprender Go",
    "completado": false,
    "due_date": "02/07/2025"
  }
]
```

---

## 🌈 Ejemplo visual (en terminal)

```bash
+----+-----------------+------------+--------------+
| ID |   DESCRIPCIÓN   |   ESTADO   | FECHA LÍMITE |
+----+-----------------+------------+--------------+
|  1 | Comprar A       | Pendiente  | 03/07/2025   |
+----+-----------------+------------+--------------+
|  2 | Comprar B       | Finalizada | 02/07/2025   |
+----+-----------------+------------+--------------+
|  3 | Comprar C       | Finalizada | 05/07/2025   |
+----+-----------------+------------+--------------+
|  5 | Comprar D       | Pendiente  | 08/07/2025   |
+----+-----------------+------------+--------------+
```

---

## 🙌 Contribuciones

¡Pull requests, sugerencias o ideas son siempre bienvenidas!

---

## 📄 Licencia

Este proyecto está bajo la licencia MIT.  
Hecho por [Lautaro Miras](https://github.com/lmirasutn)

---


