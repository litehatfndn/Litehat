# 🧠 Litehat (LTH)

Decentralized MCP Server Deployment with Ethereum + Secure Containers + On-Chain Governance

---

## 📘 Introduction

Litehat is a decentralized platform for deploying and managing Model Context Protocol (MCP) servers. It combines Ethereum smart contracts with secure container execution and CLI tooling for scalable, auditable AI infrastructure.

---

## ✨ Features

- 🛠 Smart Contracts for on-chain governance  
- 📦 Secure, isolated container-based server deployment  
- 🔐 Secret encryption and access control  
- 🖥 CLI tools for deploy, monitor, governance  
- 📈 Auto-scaling based on demand  
- 🔗 Token-based governance with $LTH  
- 🔌 SDK and API support  
- ☁️ Cloud and on-prem ready  

---

## 📐 Architecture

> NOTE: GitHub Mobile or Web does not support live Mermaid rendering. This is normal.

<details>
<summary>Click to view diagram (GitHub preview may not show this)</summary>

```mermaid
graph TD
  Dev[Developer]
  CLI[Litehat CLI]
  Chain[Ethereum Smart Contracts]
  MCP[MCP Container Node]
  Gov[Governance Layer]

  Dev --> CLI
  CLI --> Chain
  CLI --> MCP
  Chain --> Gov
  MCP --> Chain
