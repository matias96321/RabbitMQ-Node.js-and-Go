import rabbit from 'amqplib'  
    
const message = "Hello World"
const exchange = "hello_world_exchange"
const queue = "hello_world_queue"
const routingKey = "hello_world_key"

rabbit.connect("amqp://admin:123456@127.0.0.1:5672/").then((connection)=>{
    
    if(!connection)
    throw new Error('Failed to connect to RabbitMQ')
    
    connection.createChannel().then((channel)=>{

        if(!channel)
        throw new Error('Failed to open a channel')

        channel.assertExchange(exchange, 'direct', { durable: true })

        channel.assertQueue(queue, { durable: true })

        channel.bindQueue(queue, exchange, routingKey)

        channel.publish(exchange, routingKey, Buffer.from(message))
        
        console.log("Message sent");

        setTimeout( function() { channel.close(); connection.close();}, 500 );
    })
})



