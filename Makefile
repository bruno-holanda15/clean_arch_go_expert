COMPOSEV2 := $(shell docker compose version 2> /dev/null)
CONTAINER_NAME := go_servers

ifdef COMPOSEV2
    COMMAND_DOCKER=docker compose
else
    COMMAND_DOCKER=docker-compose
endif

up:
	$(COMMAND_DOCKER) up -d
	@echo "Container $(CONTAINER_NAME) está iniciando..." && sleep 30;

down:
	$(COMMAND_DOCKER) down

server-logs:
	docker logs --tail 100 -f go_servers