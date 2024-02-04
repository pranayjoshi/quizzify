import React from "react";
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";

const Navbar = () => {
  const navigate = useNavigate();
  return (
    <nav className="flex items-center justify-between p-2 px-6 bg-gray-950 border-b">
      <Link
        to="/"
        className="bg-clip-text text-transparent bg-gradient-to-r from-green-400 to-blue-500 text-2xl font-bold"
      >
        Quizzify
      </Link>
      <div className="flex space-x-4">
        <button
          className="text-white bg-gradient-to-r from-green-400 to-blue-500 rounded-full px-6 py-4 opacity-80"
          onClick={() => navigate("/login")}
        >
          Sign In
        </button>

        <button
          className="text-white bg-gradient-to-r from-green-400 to-blue-500 rounded-full px-6 py-2 opacity-80"
          onClick={() => navigate("/register")}
        >
          Sign Up
        </button>
      </div>
    </nav>
  );
};

export default Navbar;
