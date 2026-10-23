# nav-game-go

Jogo de nave 2D em Go, feito com a engine [Ebiten](https://ebitengine.org).

## O que faz

Você controla uma nave na parte de baixo da tela e atira lasers nos meteoros que caem, com
estrelas passando ao fundo. Cada meteoro destruído vale 1 ponto (placar "Pontos" no canto). Se um
meteoro bate na nave, o jogo recomeça do zero.

| Tecla        | Ação                  |
| ------------ | --------------------- |
| `←` / `→`    | mover a nave          |
| `Espaço`     | atirar (com recarga)  |

As imagens e fontes ficam em `assets/` e vão embutidas no binário (`go:embed`).

## Por que existe

É um projeto de estudo (maio de 2025) para fazer um jogo em Go: o laço `Update`/`Draw` do Ebiten,
sprites, colisão por retângulos e temporizadores para o tiro e para o surgimento de meteoros e
estrelas.

## Instalação

Pré-requisitos: Go 1.24.3+ e, no Linux, as bibliotecas de desenvolvimento gráfico que o Ebiten
usa. No Fedora:

```sh
sudo dnf install mesa-libGL-devel libXrandr-devel libXcursor-devel libXinerama-devel \
  libXi-devel libXxf86vm-devel alsa-lib-devel pkgconf-pkg-config
```

Para outras distros, veja a [lista do Ebiten](https://ebitengine.org/en/documents/install.html).
Sem essas bibliotecas, a compilação falha no link, por exemplo com `cannot find -lXxf86vm`.

```sh
git clone https://github.com/rillmind/nav-game-go.git
cd nav-game-go
go build -o navgame .
./navgame
```

## Como testar

Não há testes automatizados. `go vet ./...` confere o código, mesmo sem as bibliotecas gráficas.
Para testar o jogo, rode `./navgame` (ou `go run .`) e confira que:

- a janela abre com a nave, as estrelas ao fundo e o placar `Pontos: 0`;
- as setas movem a nave e o espaço dispara lasers, com um intervalo entre os tiros;
- acertar um meteoro com um laser soma 1 ponto;
- deixar um meteoro bater na nave zera o placar e recomeça.
