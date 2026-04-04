Запуск проекта

```bash
docker-compose up -d --build
```
Запуск деплоя
```bash
ansible-playbook -i ansible/inventory.ini ansible/deploy.yml -K 

или

ansible-playbook -i ansible/inventory.ini ansible/deploy.yml --ask-become-pass

```
Посмотреть логи работы докера, потмоу что во время скачивания и устновки образов не всегда видно что там в фоне происходит
```bash
journalctl -u docker -f
```

Grafana
```bash
http://localhost:3000
```

Приложение Go
```bash
http://localhost/
```


## Запуск multipass для эмуляции вируалок
```bash
# 1. Управляющий узел (здесь будет Ansible и запуск CI/CD)
multipass launch --name control --cpus 1 --memory 1G

# 2. Сервер приложения (Production)
multipass launch --name app-server --cpus 1 --memory 1G
```

Получение списка IP-адресов
```bash
multipass list
```
```bash
multipass shell control
ssh-keygen -t rsa -b 4096 -N "" -f ~/.ssh/id_rsa

```

Копирование публичного ключа на сервера
```bash
multipass shell monitor-server
```
Копировать публичный ключ сюда
```bash
~/.ssh/authorized_keys
```
Проверка доступности серверов с основной машины для ansible
```bash
ansible all -i ansible/inventory.ini -m ping
```

## Установить docker на все сервера
```bash
ansible-playbook -i ansible/inventory.ini ansible/setup_nodes.yml
```

## Пересоздание сервера с большим размером диска
```bash
multipass delete monitor-server --purge
multipass launch --name monitor-server --cpus 1 --memory 1G --disk 15G
```

