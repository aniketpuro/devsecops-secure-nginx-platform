'use client';
import { useState } from 'react';
import Link from 'next/link';

export default function RegisterPage() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    
    setTimeout(() => {
      alert("✅ Account Created Successfully! (Demo)");
      window.location.href = "/login";
      setLoading(false);
    }, 1500);
  };

  return (
    <div className="main-container">
      <div className="card">
        <h2>Create Account</h2>
        <p style={{color: '#a5b4fc', marginBottom: '2rem'}}>Join us today</p>

        <form onSubmit={handleRegister}>
          <input
            type="text"
            placeholder="Full Name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            style={{width: '100%', padding: '14px', margin: '10px 0', borderRadius: '10px', border: 'none', background: '#1f2937', color: 'white'}}
            required
          />
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
            style={{width: '100%', padding: '16px', marginTop: '15px', background: '#22c55e', border: 'none', borderRadius: '12px', color: 'white', fontSize: '1.1rem', cursor: 'pointer'}}
          >
            {loading ? "Creating Account..." : "Register"}
          </button>
        </form>

        <p style={{marginTop: '20px'}}>
          Already have an account? <Link href="/login" style={{color: '#a5b4fc'}}>Login here</Link>
        </p>
      </div>
    </div>
  );
}