# Requisitos 

Pagina inical com barra de pesquisa
- Opção de pesquisa avançada: a api do google fornece varios parametros que um usuario mais avançado pode querer usar

customização de themas, cores, dashboard (talvez), etc
keymaps customization (maybe with tui, maybe via config file)

extensibility: allow people to create their own plugins to render specifics pages, or to add new search engines
- Mainly, the idea that the user may be capable of creating is own config file for displaing the structure of stackoverflow, for example



# DESING

The tool should be responsive, adapting both for bigger and small screens
In fact, even if the screen is bigger, usually the search tools mantain the same size of a mobile screen

## Home page
```
           ██╗ █████╗  ██████╗ ██╗   ██╗ █████╗ ████████╗██╗██████╗ ██╗ ██████╗ █████╗  
           ██║██╔══██╗██╔════╝ ██║   ██║██╔══██╗╚══██╔══╝██║██╔══██╗██║██╔════╝██╔══██╗ 
           ██║███████║██║  ███╗██║   ██║███████║   ██║   ██║██████╔╝██║██║     ███████║ 
      ██   ██║██╔══██║██║   ██║██║   ██║██╔══██║   ██║   ██║██╔══██╗██║██║     ██╔══██║ 
      ╚█████╔╝██║  ██║╚██████╔╝╚██████╔╝██║  ██║   ██║   ██║██║  ██║██║╚██████╗██║  ██║ 
       ╚════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝╚═╝  ╚═╝╚═╝ ╚═════╝╚═╝  ╚═╝ 
      
╭──────────────────────────────────────── Search ────────────────────────────────────────────╮
│                                                                                           │
╰────────────────────────────────────────────────────────────────────────────────────────────╯
```

## Results page
```
╭──────────────────────────────────────── Search ────────────────────────────────────────────╮
│ Batata                                                                      │ Jaguatirica │
╰────────────────────────────────────────────────────────────────────────────────────────────╯

# Batata. A origem da batata ────────────────────────────────────────────────── Brasil Escola
│ 
│ A batata é um tubérculo pertencente à família das Solanaceae. Surgiu nos Andes e nas Ilhas 
│ Chilenas, foi levada para a Europa no século XVI e tornou-se base da ...
╰─────────────────────────────────────────────────── https://brasilescola.uol.com.br › saud...


# esse é o melhor jeito de fazer batata, derrete na boca e fica ... ───────────── TudoGostoso
│
│ Ingredientes · 6 batatas médias · 2 colheres (sopa) de azeite de oliva · Pimenta-do-reino
│ gosto · 4 dentes de alho · 6 colheres (sopa) de manteiga ...
╰─────────────────────────────────────────────────── https://www.tudogostoso.com.br › notic...


# Batata - Mundo Educação ───────────────────────────────────────────────────── Mundo Educação
│ 
│ A batata é pobre em gordura e rica em carboidratos, ou seja, é uma ótima fonte de energia. 
│ O tubérculo também possui quantidades consideráveis de vitamina
╰─────────────────────────────────────────────────── https://mundoeducacao.uol.com.br › sau...
```


## Advanced search

Talvez abrir como um painel lateral?
Talvez abrir como um folding em baixo da barra?
All the possible filters from google api are [here](https://developers.google.com/custom-search/v1/reference/rest/v1/cse/list?apix=true&hl=pt-br)
```
           ██╗ █████╗  ██████╗ ██╗   ██╗ █████╗ ████████╗██╗██████╗ ██╗ ██████╗ █████╗  
           ██║██╔══██╗██╔════╝ ██║   ██║██╔══██╗╚══██╔══╝██║██╔══██╗██║██╔════╝██╔══██╗ 
           ██║███████║██║  ███╗██║   ██║███████║   ██║   ██║██████╔╝██║██║     ███████║ 
      ██   ██║██╔══██║██║   ██║██║   ██║██╔══██║   ██║   ██║██╔══██╗██║██║     ██╔══██║ 
      ╚█████╔╝██║  ██║╚██████╔╝╚██████╔╝██║  ██║   ██║   ██║██║  ██║██║╚██████╗██║  ██║ 
       ╚════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝╚═╝  ╚═╝╚═╝ ╚═════╝╚═╝  ╚═╝ 
      
╭──────────────────────────────────────── Search ────────────────────────────────────────────╮
│                                                                                           │
╰────────────────────────────────────────────────────────────────────────────────────────────╯
╭──────────────────────────────── ADVANCED SEARCH OPTIONS ───────────────────────────────────╮
│─────────────────────────────────────── Language ───────────────────────────────────────────│
│ Allow Simplified Chinese                                                               [x] │
│                                                                                            │
│                                                                                            │
│ Allow Languages                                         All [x]     Manually Selection [ ] │
│                                                                                            │
│─────────────────────────────────── Date and Location ──────────────────────────────────────│
│ Answer newers than:                                                                        │
│   days [number]          weeks [number]          months [number]            years [number] │
│                                                                                            │
│                                                                                            │
│ Allow Searchs From Location                        All [x]     Manually Selection [ ]      │
│                                                                                            │
│─────────────────────────────────── Advanced Options ───────────────────────────────────────│
│ .....                                                                                      │
│                                                                                            │
│─────────────────────────────────── Advanced Options ───────────────────────────────────────│
│ .....                                                                                      │
╰────────────────────────────────────────────────────────────────────────────────────────────╯

```
Some options may unfold
```
╭──────────────────────────────── ADVANCED SEARCH OPTIONS ───────────────────────────────────╮
│─────────────────────────────────────── Language ───────────────────────────────────────────│
│ Allow Simplified Chinese                                                               [x] │
│                                                                                            │
│                                                                                            │
│ Allow Languages                                         All [ ]     Manually Selection [x] │
│        │                                                                                   │
│        │─────────────────────────────────────────────────────── Afeganistão [x]            │
│        │────────────────────────────────────────────────────────────Albânia [x]            │
│        │────────────────────────────────────────────────────────────Argélia [x]            │
│        │────────────────────────────────────────────────────────────Argélia [x]            │
│        │─────────────────────────────────────────────────────────────...                   │
│        │──────────────────────────────────────────────────────────────Iêmen [x]            │
│        │─────────────────────────────────────────────────────────Iugoslávia [x]            │
│        │─────────────────────────────────────────────────────────────Zâmbia [x]            │
│        ╰───────────────────────────────────────────────────────────Zimbábue [x]            │
│                                                                                            │
│                                                                                            │
│─────────────────────────────────── Date and Location ──────────────────────────────────────│
│ Answer newers than:                                                                        │
│   days [number]        weeks [number]        months [number]            years [number]     │
│                                                                                            │
│                                                                                            │
│ Allow Searchs From                                      All [ ]     Manually Selection [x] │
│        │                                                                                   │
│        │─────────────────────────────────────────────────────── Afeganistão [x]            │
│        │────────────────────────────────────────────────────────────Albânia [x]            │
│        │────────────────────────────────────────────────────────────Argélia [x]            │
│        │────────────────────────────────────────────────────────────Argélia [x]            │
│        │─────────────────────────────────────────────────────────────...                   │
│        │──────────────────────────────────────────────────────────────Iêmen [x]            │
│        │─────────────────────────────────────────────────────────Iugoslávia [x]            │
│        │─────────────────────────────────────────────────────────────Zâmbia [x]            │
│        ╰───────────────────────────────────────────────────────────Zimbábue [x]            │
│                                                                                            │
│─────────────────────────────────── Advanced Options ───────────────────────────────────────│
│ .....                                                                                      │
│                                                                                            │
│─────────────────────────────────── Advanced Options ───────────────────────────────────────│
│ .....                                                                                      │
╰────────────────────────────────────────────────────────────────────────────────────────────╯
```
