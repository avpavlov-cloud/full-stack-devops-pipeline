# Full-Stack DevOps Pipeline (Golang edition)

Комплексный проект по автоматизации развертывания распределенной инфраструктуры веб-приложения с системой мониторинга в локальном облаке Multipass.

---

## 🚀 Архитектура проекта

Проект развертывается на трех независимых узлах Ubuntu (Noble 24.04):

1. **Control Node (10.157.62.20)** — узел управления (Ansible)
2. **App Server (10.157.62.200)** — хост для Go-приложения
3. **Monitor Server (10.157.62.19)** — стек Prometheus + Grafana

---

## 🛠 Технологический стек

* **Язык:** Golang (сервис с метриками Prometheus)
* **Контейнеризация:** Docker (multi-stage сборка для минимизации образов)
* **IaC / Оркестрация:** Ansible (provisioning и deployment)
* **Виртуализация:** Multipass (эмуляция облачной среды)
* **Мониторинг:** Prometheus + Grafana (observability)
* **Метод деплоя:** Air-gapped (передача образов через `.tar` без внешнего реестра)

---

## 📂 Структура репозитория

```bash
full-stack-devops-pipeline/
├── app/
│   ├── main.go          # Go-приложение с экспортом метрик
│   ├── Dockerfile       # Оптимизированная multi-stage сборка
│   └── go.mod
├── ansible/
│   ├── inventory.ini    # IP-адреса серверов
│   ├── setup_nodes.yml  # Подготовка узлов (Docker)
│   └── deploy.yml       # Деплой и конфигурация
└── README.md
```

---

## ⚙️ Как это работает (Workflow)

1. **Provisioning**
   Ansible устанавливает Docker на чистые узлы Multipass

2. **Build**
   Локально собирается Docker-образ Go-приложения

3. **Artifacts**
   Образы (App, Prometheus, Grafana) сохраняются в `.tar`

4. **Transfer**
   Ansible копирует архивы на серверы по SSH

5. **Run**
   Контейнеры запускаются, настраиваются сети и порты

6. **Config**
   Автоматически генерируется конфиг Prometheus для сбора метрик

---

## ▶️ Запуск проекта

### Docker Compose (локальный запуск)

```bash
docker-compose up -d --build
```

---

### Запуск деплоя через Ansible

```bash
ansible-playbook -i ansible/inventory.ini ansible/deploy.yml -K
```

или

```bash
ansible-playbook -i ansible/inventory.ini ansible/deploy.yml --ask-become-pass
```

---

### Просмотр логов Docker

Полезно, когда образы скачиваются в фоне и нет явного вывода:

```bash
journalctl -u docker -f
```

---

## 🌐 Доступ к сервисам

### Локально (Docker Compose)

* **Grafana:** [http://localhost:3000](http://localhost:3000)
* **Go-приложение:** [http://localhost/](http://localhost/)

### Через Multipass

* **Web App:** [http://10.157.62.200:80](http://10.157.62.200:80)
* **Prometheus:** [http://10.157.62.19:9090](http://10.157.62.19:9090)
* **Grafana:** [http://10.157.62.19:3000](http://10.157.62.19:3000)

  * login: `admin`
  * password: `admin`

---

## ☁️ Multipass (эмуляция виртуальных машин)

### Создание виртуалок

```bash
# 1. Управляющий узел (Ansible / CI/CD)
multipass launch --name control --cpus 1 --memory 1G

# 2. Сервер приложения
multipass launch --name app-server --cpus 1 --memory 1G
```

---

### Получение IP-адресов

```bash
multipass list
```

---

### Настройка SSH-доступа

```bash
multipass shell control

ssh-keygen -t rsa -b 4096 -N "" -f ~/.ssh/id_rsa
```

---

### Копирование публичного ключа

```bash
multipass shell monitor-server
```

Добавь содержимое публичного ключа в файл:

```bash
~/.ssh/authorized_keys
```

---

### Проверка доступности серверов

```bash
ansible all -i ansible/inventory.ini -m ping
```

---

## 🐳 Установка Docker на все узлы

```bash
ansible-playbook -i ansible/inventory.ini ansible/setup_nodes.yml
```

---

## 🔄 Пересоздание сервера с увеличенным диском

```bash
multipass delete monitor-server --purge
multipass launch --name monitor-server --cpus 1 --memory 1G --disk 15G
```

---

## 📌 Что демонстрирует проект

* Практику Infrastructure as Code (IaC)
* Автоматизацию развертывания (Ansible)
* Работу с контейнерами (Docker)
* Настройку мониторинга и метрик
* Понимание сетевого взаимодействия сервисов
* Air-gapped deployment (enterprise-практика)
