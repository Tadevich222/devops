# kubectl-kubeplugin

Плагін для kubectl, який відображає статистику використання ресурсів Kubernetes (CPU та Memory).

## Опис

Цей плагін дозволяє швидко отримати інформацію про використання CPU та пам'яті для ресурсів Kubernetes у зручному CSV форматі.

## Передумови

- Встановлений `kubectl`
- Доступ до Kubernetes кластера
- Встановлений `metrics-server` у кластері

## Встановлення

### Крок 1: Створення директорії для плагінів

```bash
mkdir -p ~/.kube/plugins
```

### Крок 2: Копіювання скрипта

```bash
cp kubectl-kubeplugin ~/.kube/plugins/kubectl-kubeplugin
```

### Крок 3: Надання прав на виконання

```bash
chmod +x ~/.kube/plugins/kubectl-kubeplugin
```

### Крок 4: Додавання до PATH

Додайте наступний рядок до `~/.bashrc` або `~/.zshrc`:

```bash
export PATH="${PATH}:${HOME}/.kube/plugins"
```

Застосуйте зміни:

```bash
source ~/.bashrc
# або
source ~/.zshrc
```

## Використання

### Синтаксис

```bash
kubectl kubeplugin <resource_type> <namespace>
```

### Параметри

- `resource_type` - тип ресурсу Kubernetes (pod, node, deployment, тощо)
- `namespace` - namespace у якому шукати ресурси

### Приклади використання

#### Отримати статистику для подів у namespace kube-system

```bash
kubectl kubeplugin pod kube-system
```

#### Отримати статистику для подів у namespace default

```bash
kubectl kubeplugin pod default
```

#### Отримати статистику для нод

```bash
kubectl kubeplugin node kube-system
```

## Формат виводу

Плагін виводить дані у CSV форматі:

```
Resource, Namespace, Name, CPU, Memory
```

### Приклад виводу

```
tadevich@DESKTOP-K5102K1:/mnt/d/courses/devops$ kubectl kubeplugin pod kube-system
Resource, Namespace, Name, CPU, Memory
pod, kube-system, coredns-ccb96694c-g9cc8, 5m, 23Mi
pod, kube-system, local-path-provisioner-5cf85fd84d-r8bpb, 1m, 11Mi
pod, kube-system, metrics-server-5985cbc9d7-jkwm5, 10m, 28Mi
pod, kube-system, svclb-traefik-f6a92544-m4vpp, 0m, 2Mi
pod, kube-system, svclb-traefik-f6a92544-vsbjq, 0m, 2Mi
pod, kube-system, traefik-5d45fc8cc9-vs4lq, 1m, 48Mi
```

## Перевірка встановлення

Перевірте, що kubectl бачить плагін:

```bash
kubectl plugin list
```

Ви повинні побачити у списку:

```
/home/user/.kube/plugins/kubectl-kubeplugin
```

## Усунення несправностей

### Помилка: "command not found"

Перевірте, що:
1. Скрипт має права на виконання (`chmod +x`)
2. Директорія `~/.kube/plugins` додана до PATH
3. Ви перезавантажили термінал або виконали `source ~/.bashrc`

### Плагін не відображається у `kubectl plugin list`

Перевірте:
1. Ім'я файла починається з `kubectl-` (наприклад, `kubectl-kubeplugin`)
2. Файл має права на виконання
3. Файл знаходиться у директорії, яка є в PATH