'use client';
import { useState } from 'react';
import Link from 'next/link';

export default function LoginPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    
    // Baad mein real API call lagega
    setTimeout(() => {
      alert("✅ Login Successful! (Demo)");
      window.location.href = "/";
      setLoading(false);
    }, 1500);
  };

  return (
    <div className="main-container">
      <div className="card">
        <h2>Login</h2>
        <p style={{color: '#a5b4fc', marginBottom: '2rem'}}>Welcome back!</p>

        <form onSubmit={handleLogin}>
          <input
            type="email"
            placeholder="Email Address"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            style={{width: '100%', padding: '14px', margin: '10px 0', borderRadius: '10px', border: 'none', background: '#1f2937', color: 'white'}}
            required
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            style={{width: '100%', padding: '14px', margin: '10px 0', borderRadius: '10px', border: 'none', background: '#1f2937', color: 'white'}}
            required
          />
          
          <button 
            type="submit"
            disabled={loading}
            style={{width: '100%', padding: '16px', marginTop: '15px', background: '#6366f1', border: 'none', borderRadius: '12px', color: 'white', fontSize: '1.1rem', cursor: 'pointer'}}
          >
            {loading ? "Logging in..." : "Login"}
          </button>
        </form>

        <p style={{marginTop: '20px'}}>
          Don't have an account? <Link href="/register" style={{color: '#a5b4fc'}}>Register here</Link>
        </p>
      </div>
    </div>
  );
}