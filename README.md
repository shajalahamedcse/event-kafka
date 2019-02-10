# event-kafka

##Setting Up Kafka

In this section, we will see how to create a topic in Kafka.

First, start `ZooKeeper`:

    $ bin/zookeeper-server-start.sh config/zookeeper.properties
    
We will run a single Kafka broker for the time being:
    
    $ bin/kafka-server-start.sh config/server.properties
    
Then, we will create our first topic:

    $ bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partition 1 --topic bkash-transactions
    

Running multiple Kafka instances is very easy, just `cp` the `server.properties` file and make the necessary changes:

    config/server-1.properties:
        broker.id=1
        listeners=PLAINTEXT://:9093
        log.dir=/tmp/kafka-logs-1
    
    config/server-2.properties:
        broker.id=2
        listeners=PLAINTEXT://:9094
        log.dir=/tmp/kafka-logs-2
        
To start the brokers, we must specify the file properly, as follows:
    
    $ bin/kafka-server-start.sh config/server-1.properties
    $ bin/kafka-server-start.sh config/server-2.properties