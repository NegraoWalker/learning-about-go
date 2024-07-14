## Características da linguagem:
* Compilada;
* Rápido processamento;
* Não é uma linguagem orientada a objetos;

## Links importantes de ajuda:
* https://pkg.go.dev/ - Opção de pesquisa sobre packages disponíveis para serem usados
* https://pkg.go.dev/std - Packages Standart (Padrão)

## Rever ou pesquisar sobre:
- [ ] Ponteiros
- [ ] Structs
- [ ] Receiver



## Informações importantes da linguagem:
* Todo programa em go deve ter pelo menos uma função main;
* Package fmt significa format (formato);
* Se declaramos uma variável dentro de uma função temos que usá-la - Linguagem reclama de erro caso não faça;
* Em uma função podemos ter dois retornos;
* **Ponteiro** é uma variável que armazena um endereço de memória;
* A notação *&variavel* indica o endereço de memória apontado pelo ponteiro;
* A notação **variavel* aponta para o valor armazenado no endereço de memória que o ponteiro aponta;
* Melhor explicação sobre ponteiro: https://www.youtube.com/watch?v=Ip1VpLxNOvQ&ab_channel=Filhodanuvem
* Variáveis declaradas no escopo global estão disponíveis para todas as funções. Já variáveis declaradas no escopo local estão disponíveis apenas dentro da função;
* **struct(estrutura)** é um recurso que agrupa variáveis de tipos de dados diferentes. Cada variável dentro de uma struct é chamada de field(campo). Esse recurso é muito utilizado para agrupar dados relacionados;
* Se declararmos uma variável ou função com a primeira letra do seu nome em maiúscula, informamos ao compilador do Go que a mesma pode ser eventualmente usada em outro pacote. Já o contrário, a variável ou função só ficará disponível dentro do pacote que foi declarada;
* **receiver(receptor)** é um parâmetro especial que permite que uma função seja associada a um tipo específico, transformando-a em um método. Receivers são usados para definir métodos em tipos de structs (ou outros tipos) e são declarados na definição da função antes do nome da função;
* **map(mapa)** é estrutura de dados que associa chaves a valores, permitindo a busca rápida de valores baseados em suas chaves;
* **slice(fatia)** é uma estrutura de dados que fornece uma interface poderosa e flexível para sequências de elementos. Um slice é composto por três componentes principais:
    
    *Ponteiro (pointer)*: aponta para o primeiro elemento da array subjacente que é acessível pelo slice;

    *Comprimento (length)*: o número de elementos no slice;

    *Capacidade (capacity)*: o número de elementos entre o primeiro elemento do slice e o final da array subjacente;
* **strings** são imutáveis;
* **interface** é um tipo que especifica um conjunto de métodos, mas não implementa esses métodos. Interfaces são usadas para definir comportamentos comuns que diferentes tipos podem implementar, promovendo a flexibilidade e a reutilização do código;
* **package** é uma coleção de arquivos de código-fonte que são compilados juntos. Cada arquivo em um pacote começa com uma declaração de pacote, que indica o nome do pacote ao qual o arquivo pertence. Os pacotes permitem organizar o código em módulos reutilizáveis e gerenciáveis;
* **module** é maneira de gerenciar dependências e pacotes em projetos de software. Os módulos permitem que os desenvolvedores especifiquem quais versões de pacotes (também conhecidos como dependências) seu código depende, o que facilita a reprodutibilidade e a gestão de versões. Para iniciar um módulo abrir o terminal dentro do diretório do projeto e digitar o seguinte comando: 
*go mod init github.com/NegraoWalker/name-module*, com a execução desse programa será criado um arquivo go.mod, onde apresenta as dependências do projeto e suas versões;

