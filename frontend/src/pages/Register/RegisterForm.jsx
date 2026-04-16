import { useState } from "react"
import { useNavigate } from "react-router-dom"
import "./register.css"
import { FaEye, FaEyeSlash } from "react-icons/fa"

function RegisterForm(){
const navigate = useNavigate();

const [showPassword,setShowPassword] = useState(false)
const [password, setPassword] = useState("")
const [touched, setTouched] = useState(false)
const [email, setEmail] = useState("")
const [otp, setOtp] = useState(["","","","","",""])
const [otpSent, setOtpSent] = useState(false)
const [otpVerified, setOtpVerified] = useState(false)

// Added state for full Registration tracking
const [fullName, setFullName] = useState("")
const [username, setUsername] = useState("")
const [phone, setPhone] = useState("")
const [address, setAddress] = useState("")
const [currency, setCurrency] = useState("USD")

const handleSendOtp = async () => {
  try {
    const res = await fetch("http://localhost:5050/send-otp", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email })
    })

    const data = await res.json()

    if (data.success) {
      setOtpSent(true)
      alert("OTP sent to your personal email! ✨")
    } else {
      alert("Failed to send OTP ❌ Make sure your backend node server is running and .env is configured.")
    }
  } catch(e) {
    alert("Connection error trying to reach OTP Server (Port 5050).")
  }
}
const handleOtpChange = (index, value) => {
if (!/^[0-9]?$/.test(value)) return

  const newOtp = [...otp]
  newOtp[index] = value
  setOtp(newOtp)

  // move to next box
  if (value && index < 5) {
    document.getElementById(`otp-${index + 1}`).focus()
  }
}

const handleKeyDown = (e, index) => {
  if (e.key === "Backspace" && !otp[index] && index > 0) {
    document.getElementById(`otp-${index - 1}`).focus()
  }
}
    
const handleVerifyOtp = async () => {
  const enteredOtp = otp.join("")
  try {
    const res = await fetch("http://localhost:5050/verify-otp", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, otp: enteredOtp })
    })

    const data = await res.json()

    if (data.verified) {
      setOtpVerified(true)
      alert("Email completely verified! ✅")
    } else {
      alert("Invalid OTP ❌ Ensure you typed exactly what was mailed.")
    }
  } catch(e) {
    alert("Connection error trying to verify OTP (Port 5050).")
  }
}

const handleRegister = async () => {
  if (!otpVerified) {
    alert("Please verify your email with the OTP before registering!");
    return;
  }
  
  if (!isValid) {
    alert("Please fix the password requirements.");
    return;
  }

  try {
    const payload = {
      user_name: username,
      email: email,
      phone: phone,
      address: address,
      password: password,
      currency: currency
    };

    const res = await fetch("http://localhost:8080/auth/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload)
    });

    if (res.ok) {
      alert("Registration Successful!");
      navigate("/");
    } else {
      const errText = await res.text();
      alert("Registration failed: " + errText);
    }
  } catch(e) {
    alert("System network error while hitting main Go backend.");
  }
}
const checks = {
  length: password.length >= 8 && password.length <= 12,
  uppercase: /[A-Z]/.test(password),
  lowercase: /[a-z]/.test(password),
  number: /[0-9]/.test(password),
  special: /[!@#$%^&*]/.test(password)
}
const isValid =
  checks.length &&
  checks.uppercase &&
  checks.lowercase &&
  checks.number &&
  checks.special

return(

<div className="register-form">

<h2>Get Started with FinTrack</h2>

<input 
  type="text" 
  placeholder="Full Name" 
  value={fullName}
  onChange={(e) => setFullName(e.target.value)}
/>

<input 
  type="text" 
  placeholder="Username" 
  value={username}
  onChange={(e) => setUsername(e.target.value)}
/>

<div className="email-otp">
  <input     type="email"
    placeholder="Email Address"
    value={email}
    onChange={(e) => setEmail(e.target.value)} />
    <button className= "send-btn" onClick ={handleSendOtp}>
        Send OTP
    </button>
</div>

{otpSent && (
  <div className="otp-section">

<div className="otp-box-container">

  {otp.map((digit, index) => (
    <input
      key={index}
      id={`otp-${index}`}
      type="text"
      maxLength="1"
      value={digit}
      onChange={(e) => handleOtpChange(index, e.target.value)}
      onKeyDown={(e) => handleKeyDown(e, index)}
      className="otp-box"
    />
  ))}

</div>

    <button className= "verify-btn" onClick={handleVerifyOtp}>
      Verify
    </button>

  </div>
)}

<input 
  type="text" 
  placeholder="Mobile Number"
  value={phone}
  onChange={(e) => setPhone(e.target.value)}
/>

<div className="password-wrapper">

<input
  type={showPassword ? "text" : "password"}
  placeholder="Password"
  className={`input-field ${touched ? (isValid ? "valid-border" : "invalid-border") : ""}`}
  value={password}
  onChange={(e) => {
    setPassword(e.target.value)
    setTouched(true)
  }}
/>

  <span
    className="eye-icon"
    onClick={() => setShowPassword(!showPassword)}
  >
    {showPassword ? <FaEye /> : <FaEyeSlash />}
  </span>

</div>
{touched && (
  <div className="password-rules">

    <p className={checks.length ? "valid" : ""}>
      8–12 characters
    </p>

    <p className={checks.uppercase ? "valid" : ""}>
      One uppercase letter
    </p>

    <p className={checks.lowercase ? "valid" : ""}>
      One lowercase letter
    </p>

    <p className={checks.number ? "valid" : ""}>
      One number
    </p>

    <p className={checks.special ? "valid" : ""}>
      One special character
    </p>

  </div>
)}



<input 
  type="text" 
  placeholder="Address"
  value={address}
  onChange={(e) => setAddress(e.target.value)}
/>

<select value={currency} onChange={(e) => setCurrency(e.target.value)}>
  <option value="USD">USD</option>
  <option value="INR">INR</option>
  <option value="EUR">EUR</option>
</select>

<button className="register-btn" onClick={handleRegister}>
  Register Securely
</button>

<p className="login-link">
Already have an account? <span onClick={() => navigate("/")} style={{ cursor: "pointer", color: "#3b82f6", fontWeight: "bold" }}>Log in</span>
</p>

</div>


)

}

export default RegisterForm