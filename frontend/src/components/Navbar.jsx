// src/components/Navbar.jsx
import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <nav style={{ padding: "15px", backgroundColor: "#f3f4f6", display: "flex", gap: "20px" }}>
      <Link to="/">🏠 Главная</Link>
      <Link to="/grants">📌 Гранты</Link>
      <Link to="/programs">🎓 Программы</Link>
    </nav>
  );
};

export default Navbar;
