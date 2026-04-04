Запуск проекта

```bash
docker-compose up -d --build
```
Запуск деплоя
```bash
ansible-playbook -i ansible/inventory.ini ansible/deploy.yml -K 
```