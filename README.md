# chronicler-back-end

Back-ednd do projeto `chronicler`, realizado como trabalho para a matéria de Engenharia de Software II

## Instruções

### Tecnologias e Versões
- Neste projeto está sendo utilizado `Go` na versão `1.21.7`.
- Como framework, está sendo utilizado `Go Fiber` na versão `2.52.0`. 

### Rodando o projeto na sua máquina local
- Para configurar o projeto em sua máquina, após clonar o repositório, basta utilizar o seguinte comando no terminal:
```sh
go mod tidy
```
- Para rodar o projeto em sua máquina, basta utilizar o seguinte comando:
```sh
air
```

## Padrões de projeto 

### Criação de entidades
- Para cada entidade, deve ser criada uma pasta seguindo o padrão `/src/[NOME_ENTIDADE]`.
- Para cada pasta de uma entidade criada, devem ser criadas as pastas:
    -`src/[NOME_ENTIDADE]/model` - onde serão armazenados os arquivos relativos aos modelos daquela entidade.
    -`src/[NOME_ENTIDADE]/controller` - onde serão armazenados os arquivos realtivos aos controllers daquela entidade.
    -`src/[NOME_ENTIDADE]/router` - onde serão armazenados os arquivos relativos às rotas gerenciadas pelo controller.

### Criação de branches
- As branches criadas devem seguir o seguinte padrão:
    -`feat/[NOME_DA_FEATURE]` - para criação de novas funcionalidades.
    -`fix/[NOME_DO_FIX]` - para correções de código.
    -`refactor/[NOME_DA_REFACTOR]` - para refatorações de código.

### Mensagens de commit
- As mensagens de commit devem ser escritas em `INGLÊS`.
- As mensagens de commit devem seguir o seguinte padrão:
    -`feature: [DESCRICAO_DA_FEATURE]` - para funcionalidades.
    -`fix: [DESCRICAO_DO_FIX]` - para correção de código.
    -`refactor: [DESCRICAO_DO_REFACTOR]` - para refatoração de código.

### Regras gerais
- O código deve ser implementado em `INGLÊS`.

