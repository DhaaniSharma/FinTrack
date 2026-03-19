import LoginForm from "./LoginForm"
import ParticleUniverse from "../../components/animations/ParticleUniverse"
import "./login.css"
import FinTrackLogo from "../../assets/logo/FinTrackLogo.png"

function LoginPage() {
  return (
    <div className="login-container">

      <div className="left-panel">
        <ParticleUniverse />

        <div className="branding">
          <img
            src={FinTrackLogo}
            alt="FinTrack Logo"
            className="brand-logo"
          />
          
          <h1 className="brand-name">FinTrack</h1>
          <p className="tagline">Smart. Secure. Financial Management</p>
          <p className="subtext">
            AI-driven personal finance intelligence
          </p>
        </div>
      </div>

      <div className="right-panel">
        <LoginForm />
      </div>

    </div>
  )
}

export default LoginPage
