from flask import Flask, request, jsonify
import pika
import os
import json
import time

app = Flask(__name__)

RABBITMQ_HOST = os.getenv("RABBITMQ_HOST", "rabbitmq")
QUEUE_NAME = "conversion_complete"

def send_notification(user_email: str, filename: str, download_link: str):
    """Simple console + future email notification"""
    print(f"\n{'='*60}")
    print(f"📧 NOTIFICATION SERVICE")
    print(f"{'='*60}")
    print(f"✅ Conversion Completed!")
    print(f"📁 File     : {filename}")
    print(f"📧 To       : {user_email}")
    print(f"🔗 Download : {download_link}")
    print(f"Time       : {time.strftime('%Y-%m-%d %H:%M:%S')}")
    print(f"{'='*60}\n")
    
    # Baad mein real email code add kar sakte hain
    # send_email(user_email, filename, download_link)

@app.route('/health')
def health():
    return jsonify({"status": "healthy", "service": "notification-service"})

def callback(ch, method, properties, body):
    try:
        message = json.loads(body.decode())
        user_email = message.get("email")
        filename = message.get("filename")
        download_link = message.get("download_link", "http://localhost/download")

        send_notification(user_email, filename, download_link)
        ch.basic_ack(delivery_tag=method.delivery_tag)
    except Exception as e:
        print("Error processing message:", e)
        ch.basic_nack(delivery_tag=method.delivery_tag)

def start_rabbitmq_consumer():
    print("Notification Service waiting for messages...")
    while True:
        try:
            connection = pika.BlockingConnection(pika.ConnectionParameters(host=RABBITMQ_HOST))
            channel = connection.channel()
            channel.queue_declare(queue=QUEUE_NAME, durable=True)
            channel.basic_consume(queue=QUEUE_NAME, on_message_callback=callback)
            channel.start_consuming()
        except Exception as e:
            print("RabbitMQ connection failed, retrying in 5s...", e)
            time.sleep(5)

if __name__ == "__main__":
    start_rabbitmq_consumer()