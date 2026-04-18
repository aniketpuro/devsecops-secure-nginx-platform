from flask import Flask, request, jsonify
import pika
import json
import os
import time
from moviepy.editor import VideoFileClip

app = Flask(__name__)

RABBITMQ_HOST = os.getenv("RABBITMQ_HOST", "rabbitmq")
UPLOAD_FOLDER = "/app/uploads"
os.makedirs(UPLOAD_FOLDER, exist_ok=True)

@app.route('/health')
def health():
    return jsonify({"status": "healthy", "service": "converter-service"})

@app.route('/convert', methods=['POST'])
def convert_video():
    if 'video' not in request.files:
        return jsonify({"error": "No video file"}), 400
    
    file = request.files['video']
    if file.filename == '':
        return jsonify({"error": "No selected file"}), 400

    filename = file.filename
    video_path = os.path.join(UPLOAD_FOLDER, filename)
    file.save(video_path)

    try:
        # Video to MP3 conversion
        video = VideoFileClip(video_path)
        mp3_filename = filename.rsplit('.', 1)[0] + ".mp3"
        mp3_path = os.path.join(UPLOAD_FOLDER, mp3_filename)
        
        video.audio.write_audiofile(mp3_path)
        video.close()

        # RabbitMQ pe notification bhejo
        connection = pika.BlockingConnection(pika.ConnectionParameters(host=RABBITMQ_HOST))
        channel = connection.channel()
        channel.queue_declare(queue='conversion_complete', durable=True)
        
        message = {
            "email": "user@example.com",   # baad mein real user email aayega
            "filename": mp3_filename,
            "download_link": f"http://localhost:3000/download/{mp3_filename}"
        }
        channel.basic_publish(exchange='', routing_key='conversion_complete', body=json.dumps(message))
        connection.close()

        return jsonify({
            "message": "Conversion started",
            "mp3_filename": mp3_filename
        })

    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5000)