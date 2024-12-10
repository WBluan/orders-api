# 📦 **Orders API**  

API REST desenvolvida em **Go** usando o framework **Fiber** e o banco de dados **SQLite3**. Esta API permite gerenciar **Clientes**, **Itens** e **Pedidos**, com operações de **CRUD (Create, Read, Update, Delete)** para cada uma dessas entidades.

---

## 🚀 **Funcionalidades**
- **Clientes**: Criar, consultar, editar e excluir clientes.
- **Itens**: Gerenciar os produtos disponíveis para pedidos.
- **Pedidos**: Criar pedidos, consultar histórico, atualizar e excluir pedidos.

---

## 📂 **Estrutura de Pastas**
```
📁 orders-api
├── 📁 models        # Estruturas de dados (structs) e tipos utilizados nas entidades
├── 📁 repository    # Conexão com o banco de dados e operações de persistência
├── 📁 routes        # Definição de rotas de cada entidade (clientes, pedidos e itens)
├── 📁 controllers   # Funções que controlam as rotas das API's
├── 📄 main.go       # Arquivo principal para inicializar o servidor e registrar as rotas
├── 📄 go.mod        # Arquivo de dependências do Go
├── 📄 orders.db     # Arquivo do banco de dados SQLite (não deve ser versionado)
└── 📄 README.md     # Documentação do projeto
```

## ⚙️ **Configuração e Instalação**

### 🛠️ **Pré-requisitos**
- **Go** (versão 1.20 ou superior) → [Baixar Go](https://go.dev/doc/install)
- **Postman** (para testar a API) → [Baixar Postman](https://www.postman.com/downloads/)
