import React from "react";
import "./styles/Header.css";
import Link from "next/link";
function Header() {
  return (
    <header className="header">
      <Link href="/">
        <h3 className="title">Kadrion</h3>
      </Link>
      <nav className="navbar">
        <ul>
          <li className="navitem">
            <Link href="/roadmap">Roadmap</Link>
          </li>
          <li className="navitem">
            <Link href="/documentation">Documentation</Link>
          </li>
          <li className="navitem">
            <button className="github">
              <Link
                href="https://github.com/kadriontestops-tech"
                target="_blank"
              >
                Github
              </Link>
            </button>
          </li>
        </ul>
      </nav>
    </header>
  );
}

export default Header;
