'use client';
import { useState } from 'react';

export default function Home() {
  const [file, setFile] = useState<File | null>(null);
  const [isConverting, setIsConverting] = useState(false);
  const [progress, setProgress] = useState(0);
  const [downloadUrl, setDownloadUrl] = useState<string | null>(null);
  const [message, setMessage] = useState("");

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      setFile(e.target.files[0]);
      setDownloadUrl(null);
      setMessage("");
    }
  };

  const handleConvert = async () => {
    if (!file) {
      alert("Please select a video file");
      return;
    }

    setIsConverting(true);
    setProgress(0);
    setMessage("Uploading to Gateway...");

    const formData = new FormData();
    formData.append("video", file);

    try {
      const response = await fetch("http://localhost:8080/api/convert", {
        method: "POST",
        body: formData,
      });

      const data = await response.json();

      if (response.ok) {
        setMessage("Conversion started! Processing...");
        
        // Simulate progress
        for (let i = 10; i <= 100; i += 20) {
          await new Promise(r => setTimeout(r, 600));
          setProgress(i);
        }

        setDownloadUrl(data.download_link || "#");
        setMessage("✅ Conversion Complete!");
      } else {
        setMessage("Error: " + (data.error || "Something went wrong"));
      }
    } catch (error) {
      setMessage("Failed to connect to Gateway. Is Gateway running?");
      console.error(error);
    } finally {
      setIsConverting(false);
    }
  };

  return (
    <>
      <nav className="navbar">
        <h1>🎵 MP3 Converter</h1>
        <div>Secure • Fast • Via Gateway</div>
      </nav>

      <div className="main-container">
        <div className="card">
          <h2>Convert Video to MP3</h2>
          <p style={{ color: '#a5b4fc', marginBottom: '2rem' }}>
            Upload → Gateway → Converter → Notification
          </p>

          <div
            style={{
              border: '3px dashed #555',
              borderRadius: '16px',
              padding: '60px 20px',
              margin: '20px 0',
              cursor: 'pointer',
              background: 'rgba(0,0,0,0.3)'
            }}
            onClick={() => document.getElementById('fileInput')?.click()}
          >
            <input
              id="fileInput"
              type="file"
              accept="video/*"
              onChange={handleFileChange}
              hidden
            />

            {file ? (
              <div>
                <p>✅ {file.name}</p>
                <small>{(file.size / (1024 * 1024)).toFixed(2)} MB</small>
              </div>
            ) : (
              <div>
                <h3>📤 Drop Video Here</h3>
                <p>or click to select</p>
              </div>
            )}
          </div>

          {file && !isConverting && !downloadUrl && (
            <button
              onClick={handleConvert}
              style={{
                width: '100%',
                padding: '16px',
                fontSize: '1.1rem',
                background: 'linear-gradient(90deg, #6366f1, #a855f7)',
                border: 'none',
                borderRadius: '12px',
                color: 'white',
                cursor: 'pointer',
                marginTop: '10px'
              }}
            >
              🚀 Convert via Gateway
            </button>
          )}

          {isConverting && (
            <div style={{ marginTop: '20px' }}>
              <p>{message}</p>
              <div style={{ height: '12px', background: '#333', borderRadius: '10px', marginTop: '10px', overflow: 'hidden' }}>
                <div style={{ width: `${progress}%`, height: '100%', background: 'linear-gradient(90deg, #6366f1, #22d3ee)' }}></div>
              </div>
            </div>
          )}

          {downloadUrl && (
            <div style={{ marginTop: '25px' }}>
              <p style={{ color: '#4ade80', fontSize: '1.2rem' }}>✅ {message}</p>
              <a href={downloadUrl} download>
                <button style={{
                  padding: '14px 40px',
                  background: '#22c55e',
                  border: 'none',
                  borderRadius: '12px',
                  color: 'white',
                  fontSize: '1.1rem',
                  marginTop: '15px',
                  cursor: 'pointer'
                }}>
                  ⬇️ Download MP3
                </button>
              </a>
            </div>
          )}
        </div>
      </div>
    </>
  );
}