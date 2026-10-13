# Wdrożenie od zera:  make all
# Usunięcie wszystkiego z AWS:  make destroy
#
# Wymagane zmienne środowiskowe (patrz README / .env.example):
#   DOCKERHUB_IMAGE, DISCORD_WEBHOOK_URL, GRAFANA_ADMIN_PASSWORD (opcjonalnie)

-include .env
export

.PHONY: all infra configure destroy output test local

all: infra configure

infra:
	cd terraform && terraform init -input=false && terraform apply -auto-approve

configure:
	cd ansible && ansible-galaxy collection install -r requirements.yml >/dev/null && \
	  ansible-playbook playbook.yml

output:
	cd terraform && terraform output

destroy:
	cd terraform && terraform destroy -auto-approve

test:
	cd app && go test -v ./...

local:
	docker build -t devops-projekt:local . && docker run --rm -p 8080:8080 devops-projekt:local
