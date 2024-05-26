import { Link, useLocation } from "react-router-dom";
import "./Navbar.scss";
import { logout } from "../../api/proxyApi";

const NAVBAR_LINKS = [
  { to: "config", label: "Config" },
  { to: "traffic", label: "Traffic" },
  { to: "stats", label: "Stats" },
];

const Navbar = () => {
  const currentPage = useLocation().pathname;
  return (
    <nav className="navbar">
      <div className="navbar__links">
        {NAVBAR_LINKS.map(({ to, label }) => (
          <Link
            key={to}
            to={to}
            className={`navbar__link ${currentPage.includes(to) ? "navbar__link--active" : ""
              }`}
          >
            {label}
          </Link>
        ))}
      </div>
      <Link className="navbar__link" to={"/"} onClick={logout}>
        Logout
      </Link>
    </nav>
  );
};

export default Navbar;
