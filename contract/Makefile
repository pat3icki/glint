# ============================================================
# Glint Smart Contracts — Makefile
# ============================================================
# Targets:
#   make init        – install all npm dependencies
#   make compile     – compile the Solidity contracts
#   make test        – run the Hardhat test suite
#   make clean       – remove build artefacts
#   make clean-all   – remove build artefacts + node_modules
# ============================================================

.PHONY: init compile test clean clean-abi clean-bin clean-bindings clean-all

CONTRACT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

# Writable cache dir for Hardhat (avoids EROFS errors on /home/.local in some envs)
HARDHAT_HOME := $(CONTRACT_DIR)/.hardhat-home

# Common env vars for all Hardhat invocations
HARDHAT_ENV := HARDHAT_DISABLE_TELEMETRY_PROMPT=true \
               XDG_DATA_HOME=$(HARDHAT_HOME) \
               XDG_CONFIG_HOME=$(HARDHAT_HOME)

# ── Dependency installation ──────────────────────────────────

init: 
	@echo "→ Installing npm dependencies..."
	cd $(CONTRACT_DIR) && npm install --legacy-peer-deps
	@echo "→ Setting up contracts directory..."
	mkdir -p $(CONTRACT_DIR)/contracts
	cp $(CONTRACT_DIR)/src/glint/glint.sol $(CONTRACT_DIR)/contracts/GiftCardNFT.sol
	cp $(CONTRACT_DIR)/src/paymaster/paymaster.sol $(CONTRACT_DIR)/contracts/GiftCardPaymaster.sol
	@echo "✓ Init complete. Run 'make compile' to build."

# ── Build ────────────────────────────────────────────────────

compile:
	@echo "→ Compiling Solidity contracts..."
	mkdir -p $(HARDHAT_HOME)
	cd $(CONTRACT_DIR) && $(HARDHAT_ENV) npx hardhat compile
	@echo "✓ Compilation done. Artefacts in contract/artifacts/"

# ── Test ─────────────────────────────────────────────────────

test:
	@echo "→ Running Hardhat tests..."
	mkdir -p $(HARDHAT_HOME)
	cd $(CONTRACT_DIR) && $(HARDHAT_ENV) npx hardhat test

# ── Clean ────────────────────────────────────────────────────

clean:
	@echo "→ Removing build artefacts..."
	mkdir -p $(HARDHAT_HOME)
	cd $(CONTRACT_DIR) && $(HARDHAT_ENV) npx hardhat clean

	rm -rf $(HARDHAT_HOME) $(CONTRACT_DIR)/abi $(CONTRACT_DIR)/bin $(CONTRACT_DIR)/bindings
	@echo "✓ Clean done."

clean-abi:
	rm -rf $(CONTRACT_DIR)/abi

clean-bin:
	rm -rf $(CONTRACT_DIR)/bin

clean-bindings:
	rm -rf $(CONTRACT_DIR)/bindings

clean-all: clean
	@echo "→ Removing node_modules..."
	rm -rf $(CONTRACT_DIR)/node_modules $(CONTRACT_DIR)/contracts
	@echo "✓ Full clean done."

