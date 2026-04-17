import pika
import time
import os
from pymongo import MongoClient

print("Starting Converter Service...")
time.sleep(15) # Wait for RabbitMQ and MongoDB to fully start

# 1. Connect to MongoDB
MONGO_URI = os.getenv("MONGO_URI", "mongodb://admin:securepass123@mongodb:27017/")
client = MongoClient(MONGO_URI)
db = client.mp3_converter
tasks_collection = db.tasks
print("Connected to MongoDB!")

# 2. Connect to RabbitMQ
RABBITMQ_HOST = os.getenv("RABBITMQ_HOST", "rabbitmq")
credentials = pika.PlainCredentials('admin', 'securepass123')
connection = pika.BlockingConnection(pika.ConnectionParameters(host=RABBITMQ_HOST, credentials=credentials))
channel = connection.channel()

# Queue banate hain agar nahi hai toh
channel.queue_declare(queue='mp3_tasks', durable=True)

# 3. Message aane par kya karna hai (Callback Function)
def callback(ch, method, properties, body):
    task_id = body.decode()
    print(f" [x] Received MP3 Conversion Task: {task_id}")
    
    # DB me status 'processing' kardo
    tasks_collection.update_one({"task_id": task_id}, {"$set": {"status": "processing"}}, upsert=True)
    
    # Yahan baad me asli MP3 conversion code aayega (FFmpeg wala)
    print("     -> Converting file... (Simulating 5 seconds work)")
    time.sleep(5) 
    
    # DB me status 'completed' kardo
    tasks_collection.update_one({"task_id": task_id}, {"$set": {"status": "completed"}})
    print(f" [x] Finished Task: {task_id}")
    
    # RabbitMQ ko bata do ki task ho gaya
    ch.basic_ack(delivery_tag=method.delivery_tag)

# 4. Sunna shuru karo (Listening mode)
channel.basic_qos(prefetch_count=1)
channel.basic_consume(queue='mp3_tasks', on_message_callback=callback)

print(' [*] Waiting for tasks. To exit press CTRL+C')
channel.start_consuming()