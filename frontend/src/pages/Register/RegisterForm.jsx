import { useState } from "react"
import "./register.css"
import { FaEye, FaEyeSlash } from "react-icons/fa"


function RegisterForm(){

const [showPassword,setShowPassword] = useState(false)
const [password, setPassword] = useState("")
const [touched, setTouched] = useState(false)

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

<input type="text" placeholder="Full Name"/>

<input type="text" placeholder="Username"/>

<input type="email" placeholder="Email Address"/>

<input type="text" placeholder="Mobile Number"/>

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



<input type="text" placeholder="Address"/>

<select>
<option>Select Currency</option>
<option>INR</option>
<option>USD</option>
<option>EUR</option>
</select>

<button className="register-btn">
Register
</button>

<p className="login-link">
Already have an account? <span>Log in</span>
</p>

</div>

)

}

export default RegisterForm