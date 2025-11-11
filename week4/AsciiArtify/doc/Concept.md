# Concept — Local Kubernetes options for AsciiArtify (minikube | kind | k3d)

**Дата:** 2025-11-10  
**Автор:** Євдокімов Святослав

## Мета документа
Допомогти команді AsciiArtify вибрати інструмент для локальної розробки / PoC Kubernetes-кластера. Порівняння охоплює три варіанти: **minikube**, **kind** (Kubernetes IN Docker) та **k3d** (k3s-in-docker). Також розглядаємо ризики ліцензування Docker Desktop і альтернативу — **Podman**.

---

## Джерела (важливі)
- [Minikube](https://minikube.sigs.k8s.io/docs/) — офіційна документація. 
- [Kind]([Minikube](https://minikube.sigs.k8s.io/docs/) — офіційна документація.
- [k3d](https://k3d.io/stable/) — офіційний сайт.
- [Docker Desktop](https://www.docker.com/products/docker-desktop/) — сторінка ліцензування / pricing. (умови використання для бізнесу). 
- [Podman](https://podman.io/) — порівняння та rootless концепція (альтернатива Docker Desktop).

---

## Вступ — що це за інструменти

- **minikube** — запускає одноко́мп’ютерний Kubernetes (VM або контейнерний драйвер). Підходить для локального навчання, відлагодження і простих PoC. Підтримує різні драйвери (docker, podman, virtualbox тощо).

- **kind** — запускає Kubernetes-кластери, де кожна нода — Docker-контейнер. Часто використовується для тестування (CI) і швидких локальних кластерів. Підтримує мульти-ноди та інтеграцію з різними контейнерними рантаймами (docker, podman, nerdctl).

- **k3d** — обгортка для запуску **k3s** (легкий дистрибутив Kubernetes від Rancher) всередині Docker-контейнерів. Дуже швидко піднімає кластери, економний по ресурсах, зручний для локальної розробки та CI.

---

## Таблиця: швидке порівняння

| Критерій | minikube | kind | k3d |
|---|---:|---:|---:|
| Призначення | Local dev & learning | Test/CI & local dev | Fast local dev & PoC |
| Підтримувані ОС | Linux / macOS / Windows. Підтримка x86_64, ARM64 та ін.  | Linux / macOS / Windows (потребує контейнер runtime) | Linux / macOS / Windows з Docker (k3d вимагає Docker / container runtime)  |
| Архітектури | x86_64, ARM64, ARMv7 (більше варіантів бінарних релізів)  | залежить від образів ноди (docker images)  | залежить від Docker/хост-архітектури (k3s оптимізований для низьких ресурсів)  |
| Ресурси / швидкість підняття | середній (VMs або container driver) | швидкий (контейнери) | дуже швидкий (k3s оптимізований)  |
| Багато-ноди | є (але на локалці обмежено) | легко (контейнери як ноди) | легко (multi-node)  |
| CI-підтримка | підходить (менше в CI) | відмінно підходить для CI (часто використовується у GitHub Actions)  | гарно підходить для локального CI / Actions (швидко стартує)  |
| Простота використання | дуже добрий UX для початківців | чудово для інженерів/CI | дуже простий і швидкий для PoC |
| Розширені функції | dashboard, вбудований addon manager | фокус на Kubernetes тестуванні | інтеграції k3s, швидке створення кластерів |
| Документація / спільнота | велика (Kubernetes SIG)  | сильна (SIGs, популярний у CI)  | активна (Rancher / k3d community)  |

---

## Переваги та недоліки (розгорнуто)

### Minikube
**Переваги**
- Простий старт для новачків та зручний UX.  
- Підтримує кілька драйверів (docker, podman, virtualbox, hyperv), дозволяє тестувати різні CRI (containerd, CRI-O). 

**Недоліки**
- Менш швидкий, ніж k3d/kind при створенні мульти-нодових середовищ.  
- Для деяких сценаріїв масштабування і CI — менш зручний, ніж kind/k3d.

### Kind
**Переваги**
- Швидкий старт і чудово підходить для CI (Docker контейни як ноди). 
- Добре підходить для тестування Kubernetes manifests, operator-ів, інтеграційних тестів.

**Недоліки**
- Іноді потрібно додатково налаштовувати доступ до образів (image building and loading) та мережу (особливо на Windows/macOS з Docker Desktop).

### K3d
**Переваги**
- Дуже швидко стартує, низькі вимоги до ресурсів (k3s). Ідеально для PoC та локальної розробки ML-сервісів (легко підняти multi-node). 
- Простий CLI для створення/видалення кластерів, зручний для демонстрацій і повторюваних PoC.

**Недоліки**
- При крайніх production-реплікаціях k3s може мати відмінності від "повного" upstream Kubernetes; але для PoC це не заважає.

---

## Ризики ліцензування Docker Desktop & альтернативи

### Docker Desktop ліцензія
Docker Desktop **безкоштовний** для персонального використання, для невеликих бізнесів (менше 250 співробітників та < $10M річного обороту), освітніх проєктів і відкритого ПО. У великих організаціях комерційне використання може вимагати підписки (Pro/Team/Business). Перед використанням у комерційному середовищі — перевірте умови ліцензії.

**Рекомендація для AsciiArtify:** якщо команда буде використовувати Docker Desktop на робочих машинах і стартап росте — закласти витрати на підписку або використовувати альтернативи (Podman / nerdctl) у CI і на dev-станціях.

### Podman як альтернатива
- Podman — daemonless, підтримує rootless контейнери, краще безпекове підґрунтя для локальної розробки. Є сумісність з багатьма workflow-ами Docker (podman CLI сумісний у багатьох випадках). Kind і інші інструменти все більше підтримують podman/nerdctl як runtime.

**Практичний підхід:** для розробки — можна використовувати Podman (особливо на Linux) або nerdctl; у CI (GitHub Actions) — використовувати images + runner-агенти. Для macOS/Windows, де Podman має нюанси, можна або використовувати Docker Desktop (якщо ліцензійно ок) або запустити Docker-in-VM.

---

## Рекомендація (що обрати для PoC AsciiArtify)

**Рекомендовано: `k3d`** для PoC з таких причин:
- Дуже швидке підняття кластеру (малий стартовий час) і низьке споживання ресурсів (k3s). 
- Легко створювати multi-node для тестування горизонтального масштабування (важливо для ML-сервісів, коли потрібно тестувати CPU/GPU/replica behaviour).  
- Проста інтеграція з локальним Docker (або з containerd/nerdctl) та з CI (GitHub Actions має готові екшени для k3d). 

**Альтернатива / коли обрати іншу:**
- `kind` — якщо пріоритет — **точні тести** Kubernetes manifests і інтеграція в CI, особливо якщо вже звичні workflow з kind. 
- `minikube` — якщо мета навчитися Kubernetes, або потрібні специфічні драйвери/досвід з VM-драйверами.

---

## Демонстрація (швидка) — **k3d** (Hello World)

> Передумови: встановлено `docker` (або compatible runtime), `k3d` та `kubectl`.

[![asciicast](https://asciinema.org/a/RyQhbMEqsgbQAbbSjzrfzQuD7.svg)](https://asciinema.org/a/RyQhbMEqsgbQAbbSjzrfzQuD7)

### Команди (термінал)
```bash
# 1) Встановити k3d (якщо ще не встановлено)
# Linux / macOS (інструкція): https://k3d.io
# curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash

# 2) Створити кластер для PoC (1 control-plane + 2 агенти)
k3d cluster create asciiartify --agents 2 --wait

# 3) Підключитись kubectl (k3d автоматично налаштовує kubeconfig)
kubectl get nodes

# 4) Розгорнути простий Hello World (nginx-демонстрація)
kubectl create deployment hello --image=nginxdemos/hello
kubectl expose deployment hello --type=NodePort --port=80

# 5) Подивитись, який NodePort призначено
kubectl get svc

# 6) Доступ з хоста: у більшості налаштувань k3d пробросить порт на localhost: (перегляньте NodePort, напр., 30080)
# або скористайтесь k3d kubeconfig та port-forward:
kubectl port-forward svc/hello 8080:80
# Відкрийте у браузері: http://localhost:8080