.PHONY: init
init: install

# Установка зависимостей
.PHONY: install
install:
	@npm install --ignore-scripts

# Обновление зависимостей
.PHONY: update
update:
	@npm update

# Аудит безопасности зависимостей
.PHONY: audit
audit:
	@npx cve-lite-cli ./
	@npm audit signatures

# Проверка устаревших пакетов
.PHONY: outdated
outdated:
	@npm outdated

# Форматирование кода
.PHONY: format
format:
	@npm run format

# Проверка кода
.PHONY: lint
lint:
	@npm run lint

# Тестирование
.PHONY: test
test:
	@npm run test
