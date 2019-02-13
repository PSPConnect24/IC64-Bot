#----------------------------------------------------------------------
# Go settings
#----------------------------------------------------------------------
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

#----------------------------------------------------------------------
# Standard option to run
#----------------------------------------------------------------------
all: build

#----------------------------------------------------------------------
# General building procedure
#----------------------------------------------------------------------
build: deps
	@echo "Building..."
	@$(GOBUILD)

#----------------------------------------------------------------------
# Cleanup
#----------------------------------------------------------------------
clean:
	@echo "Cleaning..."
	@$(GOCLEAN)
	@rm -f $(BINARY)

#----------------------------------------------------------------------
# Dependency retrieval.
#----------------------------------------------------------------------
deps:
	@echo "Fetching dependencies..."
	@$(GOGET) -u github.com/bwmarrin/discordgo
	@$(GOGET) -u github.com/MikeModder/anpan
	@$(GOGET) -u github.com/apaxa-go/eval
	@$(GOGET) -u github.com/Apfel/GOtag