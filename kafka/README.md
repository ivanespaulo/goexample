## Kafka e Go

Este final de semana fiz uma Poc do Kafka para testar as libs feitas na lignuagem Go, e antes de iniciar os producers e consumers tive que fazer um resumo e destalhamento nos pontos mais importantes do Kafka.

Kafka foi Desenvolvido pelo LinkedIn e torno-se open source no início de 2011. Em novembro de 2014, Jun Rao, Jay Kreps, e Neha Narkhede, que trabalharam com o Kafka no LinkedIn, criaram uma nova empresa chamada Confluent com foco em Kafka.

O Apache Kafka é uma plataforma de streaming de eventos capaz de lidar com trilhões de eventos. Inicialmente concebido como uma fila de mensagens, o Kafka é baseado na abstração de um log de confirmação distribuído. 

<h2 align="center">
  <br/>
  <img src="https://github.com/jeffotoni/goexample/blob/master/kafka/img/kafka1.png" alt="logo" width="1200" title="Referencia da imagem: confluent.io" />
  <br />
  <br />
  <br />
</h2>

Existe várias formas de subirmos o Kafka, instalando local baixando tgz, utilizando Cloud ou usando Docker ou Docker Compose.

Poderiamos utilizar:

 - [kafka.apache quickstart](https://kafka.apache.org/quickstart)
 - [confluent.io docker](https://docs.confluent.io/current/quickstart/ce-docker-quickstart.html#ce-docker-quickstart)

A [confluent.io](https://confluent.io) é sem dúvidas uma das versões que particularmente gosto muito pela flexibilidade e vamos utilizar exatamente ela.


Aqui todo o serviço que encontra-se em nosso docker-compose.yaml irá baixar as imagens e fazer seu start.

```bash
$ git clone https://github.com/jeffotoni/goexample.git

$ cd kafka/cp-all-in-one

$ docker-compose up --build

$ docker-compose ps

```

<h2 align="center">
  <br/>
  <img src="https://github.com/jeffotoni/goexample/blob/master/kafka/img/docker-compose-ps.png" alt="logo" width="1200" />
  <br />
  <br />
  <br />
</h2>


Agora que o seriço está rodando podemos ir no browser e acessar [localhost:9021](http://localhost:9021)
Neste ambiente poderemos visualizar todo arsenal que o kafka disponibiliza de forma visual.

<h2 align="center">
  <br/>
  <img src="https://github.com/jeffotoni/goexample/blob/master/kafka/img/confluent-browser.png" alt="logo" width="1200" />
  <br />
  <br />
  <br />
</h2>


## Usando kafka-shell local

Caso queira usar o kafka bash basta instalar usando o comando abaixo.

Poderá encontrar aqui: [kafka-shell](https://github.com/devshawn/kafka-shell).

```bash

$ pip3 install kafka-shell

```

<h2 align="center">
  <br/>
  <img src="https://github.com/jeffotoni/goexample/blob/master/kafka/img/kafka-shell.png" alt="logo" width="1200" />
  <br />
  <br />
  <br />
</h2>

Em nosso exemplo estamos usando a plataforma da confluent, e em seus brokers e zookeeper já possuem o kafka-sell, vou apresentar abaixo como executa-los.
Como estamos usando Docker Compose para subir todo serviço do kafka iremos usar Docker Compose exec ou docker exec.


### Usando docker-compose exec para executar kafka-shell

#### Criando Topico
```bash

$ docker-compose exec broker kafka-topics --create --topic my-topic-golang-test1 \
--partitions 1 --replication-factor 1 --if-not-exists --zookeeper zookeeper:2181

```

#### Listando todos meus topicos
```bash

$ docker-compose exec broker kafka-topics --list --zookeeper zookeeper:2181

```

#### Describe um tópico 
```bash

$ docker-compose exec broker kafka-topics --describe my-topic-golang-test1 --zookeeper zookeeper:2181

```

#### Producer mensanges

```bash

$ docker-compose exec broker  \
  bash -c "seq 100 | kafka-console-producer --request-required-acks 1 \
  --broker-list localhost:9092 --topic my-topic-golang-test1 && echo 'Produced 100 messages.'"

```

#### Consumer mensagens

```bash

$ docker-compose exec broker  \
  kafka-console-consumer --bootstrap-server localhost:9092 \
  --topic my-topic-golang-test1 --from-beginning --max-messages 100

```

## ZOOKEEPER

O Zookeeper é um sistema centralizador e de gerenciamento para qualquer tipo de sistema distribuído. Sistema distribuído são diferentes módulos de software executando em diferentes nós / clusters (podem estar em locais geograficamente distantes), mas executando como um sistema. O Zookeeper facilita a comunicação entre os nós, compartilhando configurações entre os nós, mantém o controle de qual nó é líder, qual nó se junta / sai etc. O Zookeeper é quem mantém os sistemas distribuídos sãos e mantém a consistência. O Zookeeper é basicamente uma plataforma de orquestração.

O Zookeeper em si é um sistema distribuído que consiste em vários nós em um conjunto. O Zookeeper é um serviço centralizado para manter esses metadados.

O Zookeeper também desempenha um papel vital para servir a muitos outros propósitos, como detecção de líder, gerenciamento de configuração, sincronização, detecção de quando um novo nó entra ou sai do cluster, etc.


Kafka usa o Zookeeper para o seguinte:

### Elegendo um controlador
O controlador é um dos intermediários e é responsável por manter o relacionamento líder / seguidor para todas as partições. Quando um nó é desligado, é o controlador que instrui outras réplicas a se tornarem líderes de partição para substituir os líderes de partição no nó que está desaparecendo. O Zookeeper é usado para eleger um controlador, verifique se há apenas um e escolha um novo se ele travar.

### Associação ao cluster
Quais corretores estão ativos e fazem parte do cluster? isso também é gerenciado através do ZooKeeper.

### Configuração de tópico
Quais tópicos existem, quantas partições cada um possui, onde estão as réplicas, quem é o líder preferencial, quais substituições de configuração são definidas para cada tópico

Cotas - quantos dados cada cliente tem permissão para ler e gravar

ACLs - quem tem permissão para ler e gravar em qual tópico (consumidor antigo de alto nível) - Quais grupos de consumidores existem, quem são seus membros e qual é o último deslocamento que cada grupo obteve de cada partição.


## Broker

Um broker é o componente responsável por receber as requisições de producers e consumers, armazenar as mensagens e executar a replicação das mesmas.
Os brokers são gerenciados por outro componente o zookeeper. Este componente é bastante utilizado para controlar os diferentes integrantes de um cluster.
Além das tarefas descritas acima, os brokers também realizam outras tarefas, como gerenciar os líderes de cada partição, realizar a limpeza de dados ou a compactação das mensagens.
Pretendo escrever em detalhes cada um destes tópicos avançados.


## Log

Um log pode ser descrito como uma sequência temporal de mensagens, onde as novas mensagens sempre são adicionadas no final do log. Desta forma, uma mensagem enviada em t0 sempre estará posicionada antes de uma mensagem enviada em t1.

Cada mensagem dentro do log possui algumas informações:

1. Timestamp: data-hora da inserção
2. Offset: índice da mensagem na partição
3. Key: chave da mensagem
4. Value: a mensagem propriamente dita chamado de payload

Todas as mensagens dentro de uma partição serão um conjunto chave/valor.


## Partições

A primeira coisa a entender é que uma partição de tópico é a unidade de paralelismo em Kafka.

A unidade de armazenamento de Kafka é uma partição. Uma partição é uma sequência imutável e ordenada de mensagens anexadas. Uma partição não pode ser dividida em vários intermediários ou mesmo em vários discos.

Você especifica quantos dados ou por quanto tempo os dados devem ser retidos, após os quais o Kafka limpa as mensagens em ordem - independentemente de a mensagem ter sido consumida.

Partições são divididas em segmentos, portanto, o Kafka precisa encontrar regularmente as mensagens no disco que precisam ser removidas. Com um único arquivo muito longo das mensagens de uma partição, essa operação é lenta e propensa a erros. Para corrigir isso (e outros problemas que veremos), a partição é dividida em segmentos.

Kafka sempre fornece os dados de uma única partição para um thread do consumidor. Assim, o grau de paralelismo no consumidor (dentro de um grupo de consumidores) é limitado pelo número de partições sendo consumidas. Portanto, em geral, quanto mais partições houver em um cluster Kafka, maior será a taxa de transferência possível.

#### Resumo: 

 - Partições são a unidade de armazenamento da Kafka
 - Partições são divididas em segmentos
 - Segmentos são dois arquivos: seu log e índice
 - Os índices mapeiam cada deslocamento para a posição de suas mensagens no log, são usados ​​para procurar mensagens
 - Os índices armazenam compensações em relação à compensação base do seu segmento
 - Os lotes de mensagens compactadas são agrupados como carga útil de uma mensagem do wrapper
 - Os dados armazenados no disco são os mesmos que o broker recebe do produtor pela rede e envia aos seus consumidores

## Qual é o numero de Partições que deveriamos criar para nosso cenário?

Uma fórmula aproximada para escolher o número de partições é baseada na taxa de transferência. Você mede o tempo todo que pode obter em uma única partição para produção (chame de p ) e consumo (chame de c ). 

MAX(t/p, t/c)

t: taxa de transferência desejada
p: taxa de transferência do producer
c: taxa de transferência do consumer

Embora seja possível aumentar o número de partições ao longo do tempo, é preciso ter cuidado se as mensagens forem produzidas com chaves. Ao publicar uma mensagem com chave, o Kafka mapeia deterministicamente a mensagem para uma partição com base no hash da chave. Isso garante que as mensagens com a mesma chave sejam sempre roteadas para a mesma partição.

Em geral, mais partições em um cluster Kafka levam a uma taxa de transferência mais alta. No entanto, é preciso estar ciente do impacto potencial de ter muitas partições no total ou por broker em coisas como disponibilidade e latência. 


## PRODUCERS

#### Acks = 0
o produtor não aguarda nenhum tipo de resposta do cluster. É o modo com throughput mais elevado. É importante levar em conta que nesse modo a perda de dados é possível uma vez que o produtor não aguarda nenhum tipo de sinal do cluster.

#### Acks = 1
, o produtor aguarda por um ok do líder da partição. Sendo assim sabemos que ao menos 1 broker recebeu a mensagem. Já é uma configuração bem mais segura que Acks=0, mas não é 100% segura uma vez que o broker líder pode cair antes que a replicação seja realizada, e o produtor não seria notificado nesse cenário.

#### Acks = -1
O produtor aguarda o retorno até que o líder e todas as réplicas recebam a mensagem. É o modo mais seguro, 


## CONSUMER

Não adianta ter mais consumidores do que partições. Caso o grupo 1 possuísse 5 consumidores, 3 deles ficariam ociosos pois o Kafka não conseguiria mandar mensagens de uma mesma partição à mais de um consumidor do mesmo grupo.


## Estratégias de commit de offsets

#### No máximo uma vez
Neste modo, o consumidor commita o offset para o Kafka assim que recebe a mensagem.
Mensagens podem ser perdidas, mas nunca processadas com duplicação.

#### Pelo menos uma vez
O offset é commitado após o processamento da mensagem
Mensagens nunca serão perdidas, mas podem ser processadas com duplicação.

#### Exatamente uma vez
Uma mensagem tem a garantia de ser enviada uma única vez para um determinado consumidor.


## Bash curl rest-proxy


#### List Info Topics
```bash

$ curl "http://localhost:8082/topics" | jq

```

#### List Info Topic específico
```bash

$ curl "http://localhost:8082/topics/topicgo1" | jq

```

#### List Info Partitions Topic
```bash

$ curl "http://localhost:8082/topics/topicgo1/partitions" | jq

```

#### Produce JSON Menssage

```bash
$ curl -X POST -H "Content-Type: application/vnd.kafka.json.v2+json" \
      -H "Accept: application/vnd.kafka.v2+json" \
      --data '{"records":[{"value":{"msg":"success 4"}}]}' "http://localhost:8082/topics/topicgo1"

```

#### Create a Consumer

```bash
$ curl -X POST -H "Content-Type: application/vnd.kafka.v2+json" \
      --data '{"name": "my_consumer_instance", "format": "json", "auto.offset.reset": "earliest"}' \
      http://localhost:8082/consumers/go_json_consumer

```

#### Out
```json
  {
  	"instance_id":"my_consumer_instance",
	"base_uri":"http://rest-proxy:8082/consumers/go_json_consumer/instances/my_consumer_instance"
  } 
```

#### Subscription Consumer
```bash
$ curl -X POST -H "Content-Type: application/vnd.kafka.v2+json" --data '{"topics":["topicgo1"]}' \
 http://localhost:8082/consumers/go_json_consumer/instances/my_consumer_instance/subscription
 ```

#### Consume JSON Menssage

```bash
$ curl -X GET -H "Accept: application/vnd.kafka.json.v2+json" \
  http://localhost:8082/consumers/go_json_consumer/instances/my_consumer_instance/records

```

#### Delete
```bash
curl -X DELETE -H "Content-Type: application/vnd.kafka.v2+json" \
      http://localhost:8082/consumers/go_json_consumer/instances/my_consumer_instance

```

### Producer feito em Go

Para comunicarmos com Kafka via Go temos varias formas e algumas libs disponíveis.

#### Algumas das opções disponíveis são:

 - [sarama](https://github.com/Shopify/sarama) A API expõe conceitos de baixo nível do protocolo Kafka mas não suporta recursos recentes do Go, como contextos. Ele também passa todos os valores como ponteiros, o que causa um grande número de alocações de memória dinâmica, coletas de lixo mais frequentes e maior uso de memória.

 - [confluent-kafka-go](https://docs.confluent.io/current/clients/confluent-kafka-go/index.html#pkg-overview) É um Package que fornece aos produtores e consumidores Apache Kafka de alto nível, usando ligações na parte superior da biblioteca librdkafka C. Bem completa e muito rápida.

 - [goka](https://github.com/lovoo/goka) É um cliente Kafka mais recente do Go que se concentra em um padrão de uso específico. Ele fornece abstrações para usar o Kafka como uma mensagem que passa o barramento entre serviços, em vez de um log de eventos ordenado. O pacote também depende do sarama para todas as interações com o Kafka.

 - [kafka-go](https://github.com/segmentio/kafka-go/) Fornece APIs de nível baixo e alto para interagir com o Kafka, espelhar conceitos e implementar interfaces da biblioteca padrão Go para facilitar o uso e a integração.


Temos exemplos utilizando a lib confluent-kafka e kafka-go

#### Producer / lib kafka-go
```go

func main() {

  flagTopic := flag.String("topic", "topicgo1", "string")
  flag.Parse()

  kafkaURL := "localhost:9092"
  topic := *flagTopic

  fmt.Println("Url: ", kafkaURL)
  fmt.Println("Topic: ", topic)

  writer := newKafkaWriter(kafkaURL, topic)
  defer writer.Close()
  fmt.Println("Go Start Producing ... !!")
  for i := 0; ; i++ {
    uuid := fmt.Sprint(uuid.New())
    msgJson := `{"uuid":"` + uuid + `", "key":` + strconv.Itoa(i) + `,"msg":success", "event":"kafka test"}`
    msg := kafka.Message{
      Key:   []byte(uuid),
      Value: []byte(msgJson),
    }
    err := writer.WriteMessages(context.Background(), msg)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println("Key-", i)
  }
}

```

#### Consumer / lib kafka-go

```go

func main() {

  flagTopic := flag.String("topic", "topicgo1", "string")
  flagGroup := flag.String("group", "logger-group1", "string")
  flag.Parse()

  kafkaURL := "localhost:9092"
  topic := *flagTopic         
  groupID := *flagGroup     

  fmt.Println("Url: ", kafkaURL)
  fmt.Println("Topic: ", topic)
  fmt.Println("Group: ", groupID)

  reader := getKafkaReader(kafkaURL, topic, groupID)

  defer reader.Close()
  fmt.Println("start consuming ... !!")
  for {
      fmt.Println("consumer: ", t.Format("2006-01-02 15:04:05"))
      m, err := reader.ReadMessage(context.Background())
      if err != nil {
        log.Fatalln(err)
      }

      fmt.Printf("message at topic:%v partition:%v offset:%v  %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
  }

  fmt.Println("Ticker stopped")
}

```







