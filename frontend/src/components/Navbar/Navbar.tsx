import { Link } from "react-router-dom";
import "./Navbar.scss";

const Navbar = () => {
  return (
    <nav className="navbar">
      <Link className="navbar__link" to={"live"}>
        Live
      </Link>
      <Link className="navbar__link" to={"traffic"}>
        Traffic
      </Link>
      <Link className="navbar__link" to={"stats"}>
        Stats
      </Link>
    </nav>
  );
};

export default Navbar;
