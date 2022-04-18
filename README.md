# Mitnick Event Bot
> _Bot para agendamentos de eventos no Discord._

## Requisitos:
Caso tenha o interesse em utilizar o Bot, é necessário que você possua **um servidor de chat no Discord**. Além disso, temos os seguintes requisitos:

- Golang: [https://golang.org/](https://golang.org/)
- Como Criar um Bot no Discord: [https://discord.com/developers/docs/intro](https://discord.com/developers/docs/intro)
- Ter um Canal de Texto no Servidor Discord criado, para a utilização do Bot.

## O que Seria o Bot?
O _Mitnick Chat Bot_, é um bot que tem como objetivo auxiliar as pessoas a criarem eventos, como por exemplo, eventos de estudos, sem necessitar ser um administrador do Servidor para isso.

## Como ele Funciona?

### Criação do Evento
- O Usuário irá enviar uma mensagem como por exemplo: ```/mitnick criar "Nome do Evento" "Data do Evento" "Horário do Evento" "Local do Evento" "Descrição do Evento"```
- Será lançado um código do evento, para o usuário que fez a criação, como por exemplo: ```Código: 123456789```
- O Evento será lançado para criação e será enviado para o grupo de admin. 
- O admin irá aprovar o evento e o bot irá criar o evento no Discord.

### Cancelamento do Evento
- O Usuário irá enviar uma mensagem como por exemplo: ```/mitnick cancelar "Código do Evento"```
- O bot irá verificar se o evento existe e se o usuário que fez a criação do evento é o mesmo que está tentando cancelar.
- Se o evento existir e o usuário que fez a criação do evento for o mesmo que está tentando cancelar, o evento será cancelado e o bot irá deletar o evento no Discord.

## Como Contribuir
O primeiro passo para contribuir, é abrir uma *issue* no GitHub do projeto. Após isso, você pode fazer as mudanças e enviá-las por *pull request*. 

Se já tiver a *issue* criada, não crie uma nova, apenas envie o *pull request* referenciando a *issue* já existente.

Com a *issue* criada, siga o passo a passo abaixo:

1. Faça o fork do projeto no Github.
2. Crie uma branch para sua modificação (git checkout -b feature/fooBar).
3. Faça o commit (git commit -am 'Add some fooBar').
4. Faça o Push do Repositório Local, para o Repositório Remoto (git push origin feature/fooBar).
5. Crie um novo Pull Request.