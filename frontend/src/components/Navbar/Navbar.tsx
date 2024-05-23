import { Link, useLocation } from "react-router-dom";
import "./Navbar.scss";

const NAVBAR_LINKS = [
  { to: "config", label: "Config" },
  { to: "traffic", label: "Traffic" },
  { to: "stats", label: "Stats" },
];

const Navbar = () => {
  const currentPage = useLocation().pathname;
  return (
    <nav className="navbar">
      {NAVBAR_LINKS.map(({ to, label }) => (
        <Link
          key={to}
          to={to}
          className={`navbar__link ${
            currentPage.includes(to) ? "navbar__link--active" : ""
          }`}
        >
          {label}
        </Link>
      ))}
    </nav>
  );
};

export default Navbar;
