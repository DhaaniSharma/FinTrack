function DashboardForm() {
  return (
    <div className="dashboard-container">

      {/* SIDEBAR */}
      <div className="sidebar">
        <div className="sidebar-top">
          <h2 className="logo">FinTrack</h2>
          <ul className="menu">
            <li className="active">Dashboard</li>
            <li>Transactions</li>
            <li>Investments</li>
            <li>Activity</li>
          </ul>
        </div>
        
        {/* BOTTOM ANCHORED LINKS WITH CUSTOM SVGS */}
        <div className="sidebar-bottom">
          <ul className="menu">
            <li className="sidebar-action">
              <svg className="menu-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                <circle cx="12" cy="12" r="3"></circle>
                <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"></path>
              </svg>
              Settings
            </li>
            <li className="sidebar-action">
              <svg className="menu-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"></path>
                <polyline points="16 17 21 12 16 7"></polyline>
                <line x1="21" y1="12" x2="9" y2="12"></line>
              </svg>
              Logout
            </li>
          </ul>
        </div>
      </div>

      {/* MAIN WRAPPER */}
      <div className="main-wrapper">
        
        {/* HEADER */}
        <div className="header">
          <input
            type="text"
            placeholder="Search..."
            className="search-bar"
          />
          <div className="profile">👤 User</div>
        </div>

        {/* DASHBOARD BODY */}
        <div className="dashboard-body">
          
          {/* LEFT COLUMN: MAIN CONTENT */}
          <div className="main-content">
            
            <div className="cards">
              <div className="card"><h4>Balance</h4></div>
              <div className="card"><h4>Savings</h4></div>
              <div className="card"><h4>Investment</h4></div>
            </div>

            <div className="charts">
              <div className="chart-box"><h4>Expense Breakdown</h4></div>
              <div className="chart-box"><h4>Saving Goal</h4></div>
            </div>

            <div className="cashflow">
              <h4>Cashflow</h4>
            </div>

          </div>

          {/* RIGHT COLUMN: SIDE PANEL */}
          <div className="right-panel">
            
            <div className="side-card ml-card">
              <h4>Spending Behavior</h4>
            </div>

            <div className="side-card">
              <h4>Recent Transactions</h4>
            </div>

            <div className="side-card">
              <h4>Upcoming Payments</h4>
            </div>

          </div>

        </div>
      </div>

    </div>
  )
}

export default DashboardForm