import { useState } from "react"
import { FaEye, FaEyeSlash } from "react-icons/fa"
import "./login.css"
import FinTrackLogo from "../../assets/logo/FinTrackLogo.png"

function LoginForm() {

  const [showPassword, setShowPassword] = useState(false)

  return (
    <div className="login-form">

      <h2 className="login-title">Log into FinShield</h2>

      {/* EMAIL FIELD */}
      <input
        type="text"
        placeholder="Email or Username"
        className="input-field"
      />

      {/* PASSWORD FIELD */}
      <div className="password-wrapper">

        <input
          type={showPassword ? "text" : "password"}
          placeholder="Password"
          className="input-field"
        />

        <span
          className="eye-icon"
          onClick={() => setShowPassword(!showPassword)}
        >
          {showPassword ? <FaEye /> : <FaEyeSlash />}
        </span>

      </div>

      <button className="login-button">
        Log in
      </button>
      {/* FORGOT PASSWORD LINK */}
      <p className="forgot-password">
        Forgot password?
      </p>
      
      {/* DIVIDER */}
      <div className="divider">
      <span>OR</span>
      </div>

      {/*login with google*/}
        <button className="google-login">
          Continue with Google
        </button>

      {/*Create Account*/}
      <button className="create-account">
        Create New Account
      </button>

      <div className="login-FOOTER">
        <img
          src={FinTrackLogo}
          alt="FinTrack"
          className="footer-logo"
        />
        <p className="footer-text">© 2024 FinTrack.</p>
      </div>

    </div>
  )
}

export default LoginForm