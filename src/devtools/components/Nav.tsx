import React from "react";
import { Link } from "react-router-dom";

const Nav: React.FC = () => (
  <div className="nav">
    <Link to="/">Home</Link>
    <Link to="/test">Test</Link>
    <Link to="/emotes">Emotes</Link>
    <Link to="/bridge">Bridge</Link>
    <Link to="/layout">Layout</Link>
    <Link to="/directory">Directory</Link>
    <Link to="/storybook">Storybook</Link>
  </div>
);

export default Nav;
