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
multipass shell control
ssh-keygen -t rsa -b 4096 -N "" -f ~/.ssh/id_rsa

```

